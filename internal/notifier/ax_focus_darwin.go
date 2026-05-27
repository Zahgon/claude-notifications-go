//go:build darwin

package notifier

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework ApplicationServices -framework AppKit -framework CoreGraphics
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#import <AppKit/AppKit.h>
#import <ApplicationServices/ApplicationServices.h>
#import <CoreGraphics/CoreGraphics.h>

// Private CGS API declarations (stable, used by Moom/Magnet/Raycast et al.)
typedef int CGSConnectionID;
typedef uint64_t CGSSpaceID;
#define CGSAllSpacesMask 7
extern CGSConnectionID CGSMainConnectionID(void);
extern CFArrayRef CGSCopySpacesForWindows(CGSConnectionID cid, int selector, CFArrayRef windowIDs);
extern CGError CGSManagedDisplaySetCurrentSpace(CGSConnectionID cid, CFStringRef displayID, CGSSpaceID spaceID);
extern CFStringRef CGSCopyBestManagedDisplayForRect(CGSConnectionID cid, CGRect rect);

static int findPID(const char *bundleID) {
	@autoreleasepool {
		NSString *bid = [NSString stringWithUTF8String:bundleID];
		NSArray *apps = [NSRunningApplication runningApplicationsWithBundleIdentifier:bid];
		if (!apps || apps.count == 0) return -1;
		return (int)((NSRunningApplication *)apps[0]).processIdentifier;
	}
}

static void activateByPID(int pid) {
	@autoreleasepool {
		NSRunningApplication *app = [NSRunningApplication runningApplicationWithProcessIdentifier:(pid_t)pid];
		if (!app) return;
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdeprecated-declarations"
		[app activateWithOptions:NSApplicationActivateIgnoringOtherApps];
#pragma clang diagnostic pop
	}
}

// titleMatchesFolder checks if a window title contains folderName as a
// distinct component. Different apps use different dash separators:
//   VS Code:   "file.go \u2014 my-project \u2014 Visual Studio Code"  (em dash U+2014)
//   JetBrains: "file.go \u2013 my-project \u2013 PhpStorm"            (en dash U+2013)
//   Others:    "file.go - my-project - Terminal"                       (hyphen U+002D)
// Tries each separator in turn and returns YES on the first exact component match.
static BOOL titleMatchesFolder(NSString *title, NSString *folder) {
	NSArray *separators = @[@" \u2014 ", @" \u2013 ", @" - "];
	for (NSString *sep in separators) {
		NSArray *components = [title componentsSeparatedByString:sep];
		for (NSString *comp in components) {
			NSString *trimmed = [comp stringByTrimmingCharactersInSet:
				[NSCharacterSet whitespaceCharacterSet]];
			if ([trimmed isEqualToString:folder]) return YES;
		}
	}
	return NO;
}

// findWindowID returns the CGWindowID of the first window owned by pid whose
// title contains folderName as a distinct component, searching across all Spaces.
// Requires Screen Recording permission; caller must check CGPreflightScreenCaptureAccess first.
static CGWindowID findWindowID(int pid, const char *folderName, CGRect *outBounds) {
	@autoreleasepool {
		*outBounds = CGRectZero;
		CFArrayRef allInfo = CGWindowListCopyWindowInfo(
			kCGWindowListOptionAll | kCGWindowListExcludeDesktopElements,
			kCGNullWindowID
		);
		if (!allInfo) return 0;

		NSString *folder = [NSString stringWithUTF8String:folderName];
		CGWindowID targetWID = 0;

		for (NSDictionary *info in (__bridge NSArray *)allInfo) {
			NSNumber *pidNum = info[(__bridge NSString *)kCGWindowOwnerPID];
			if (!pidNum || pidNum.intValue != pid) continue;
			NSString *name = info[(__bridge NSString *)kCGWindowName];
			if (!name || !titleMatchesFolder(name, folder)) continue;
			NSNumber *wid = info[(__bridge NSString *)kCGWindowNumber];
			if (!wid) continue;
			targetWID = (CGWindowID)wid.unsignedIntValue;
			CFDictionaryRef boundsDict = (__bridge CFDictionaryRef)info[(__bridge NSString *)kCGWindowBounds];
			if (boundsDict) CGRectMakeWithDictionaryRepresentation(boundsDict, outBounds);
			break;
		}
		CFRelease(allInfo);
		return targetWID;
	}
}

// switchToWindowSpace switches the current visible Space to the one containing
// windowID, using bounds to select the correct display.
static void switchToWindowSpace(CGWindowID windowID, CGRect bounds) {
	@autoreleasepool {
		CGSConnectionID conn = CGSMainConnectionID();
		CFArrayRef spaces = CGSCopySpacesForWindows(conn, CGSAllSpacesMask,
			(__bridge CFArrayRef)@[@(windowID)]);
		if (!spaces) return;
		if (CFArrayGetCount(spaces) > 0) {
			CGSSpaceID spaceID = [(NSNumber *)CFArrayGetValueAtIndex(spaces, 0) unsignedLongLongValue];
			CFStringRef displayID = CGSCopyBestManagedDisplayForRect(conn, bounds);
			if (displayID) {
				CGSManagedDisplaySetCurrentSpace(conn, displayID, spaceID);
				CFRelease(displayID);
			}
		}
		CFRelease(spaces);
	}
}

static int hasScreenRecordingAccess(void) {
	return CGPreflightScreenCaptureAccess() ? 1 : 0;
}

static int requestScreenRecordingAccess(void) {
	return CGRequestScreenCaptureAccess() ? 1 : 0;
}

// restoreAndRaiseWindow restores a minimized window before raising it.
// Returns 2 when the window was successfully un-minimized so the caller can
// retry after Dock animation, or 1 after attempting the normal raise path.
static int restoreAndRaiseWindow(AXUIElementRef appEl, AXUIElementRef window) {
	CFTypeRef minRef = NULL;
	if (AXUIElementCopyAttributeValue(window, CFSTR("AXMinimized"), &minRef) == kAXErrorSuccess && minRef) {
		if (CFGetTypeID(minRef) == CFBooleanGetTypeID() && CFBooleanGetValue((CFBooleanRef)minRef)) {
			AXError err = AXUIElementSetAttributeValue(window, CFSTR("AXMinimized"), kCFBooleanFalse);
			CFRelease(minRef);
			if (err == kAXErrorSuccess) {
				return 2;
			}
		} else {
			CFRelease(minRef);
		}
	}

	AXUIElementPerformAction(window, CFSTR("AXRaise"));
	AXUIElementSetAttributeValue(appEl, CFSTR("AXFrontmost"), kCFBooleanTrue);
	return 1;
}

// raiseWindowByAXDocument enumerates AXWindows for the given PID and raises
// the first window whose AXDocument attribute exactly matches fileURL. Ghostty
// sets AXDocument to the shell CWD (via OSC 7) as a file:// URL.
// Returns 1 after raising, 2 when a minimized window was restored and should be
// retried, 0 if not found, -1 if Accessibility permission is missing.
// NOTE: AXWindows only populates after the app has been activated; callers
// must call activateByPID and wait before calling this function.
static int raiseWindowByAXDocument(int pid, const char *fileURL) {
	if (!AXIsProcessTrusted()) {
		return -1;
	}

	AXUIElementRef appEl = AXUIElementCreateApplication((pid_t)pid);
	if (!appEl) return 0;

	CFTypeRef windowsRef = NULL;
	if (AXUIElementCopyAttributeValue(appEl, CFSTR("AXWindows"), &windowsRef) != kAXErrorSuccess || !windowsRef) {
		CFRelease(appEl);
		return 0;
	}

	CFArrayRef windows = (CFArrayRef)windowsRef;
	CFIndex count = CFArrayGetCount(windows);
	int found = 0;

	for (CFIndex i = 0; i < count; i++) {
		AXUIElementRef w = (AXUIElementRef)CFArrayGetValueAtIndex(windows, i);
		CFTypeRef docRef = NULL;
		if (AXUIElementCopyAttributeValue(w, CFSTR("AXDocument"), &docRef) != kAXErrorSuccess) continue;

		CFIndex len = CFStringGetMaximumSizeForEncoding(
			CFStringGetLength((CFStringRef)docRef), kCFStringEncodingUTF8) + 1;
		char *buf = (char *)malloc(len);
		BOOL ok = buf && CFStringGetCString((CFStringRef)docRef, buf, len, kCFStringEncodingUTF8);
		CFRelease(docRef);

		if (ok && strcmp(buf, fileURL) == 0) {
			found = restoreAndRaiseWindow(appEl, w);
			free(buf);
			break;
		}
		free(buf);
	}

	CFRelease(windowsRef);
	CFRelease(appEl);
	return found;
}

// findSwitchAndActivate locates a window by title across Spaces, switches to
// its Space and activates the app. The AX raise step is handled separately by
// raiseWindowByAXTitle so that Go can retry it with backoff.
// Returns 1 ok, 0 window not found, -1 no Screen Recording permission.
static int findSwitchAndActivate(int pid, const char *folderName) {
	if (!CGPreflightScreenCaptureAccess()) {
		return -1;
	}

	CGRect bounds;
	CGWindowID targetWID = findWindowID(pid, folderName, &bounds);
	if (!targetWID) return 0;

	switchToWindowSpace(targetWID, bounds);
	usleep(300000); // wait for Space transition animation

	activateByPID(pid);
	return 1;
}

// raiseWindowByAXTitle enumerates AXWindows for the given PID and raises the
// first window whose AXTitle contains folderName as a distinct component.
// Returns 1 after raising, 2 when a minimized window was restored and should be
// retried, 0 if not found, -1 if Accessibility permission is missing.
static int raiseWindowByAXTitle(int pid, const char *folderName) {
	if (!AXIsProcessTrusted()) {
		return -1;
	}

	AXUIElementRef appEl = AXUIElementCreateApplication((pid_t)pid);
	if (!appEl) return 0;

	CFTypeRef windowsRef = NULL;
	if (AXUIElementCopyAttributeValue(appEl, CFSTR("AXWindows"), &windowsRef) != kAXErrorSuccess || !windowsRef) {
		CFRelease(appEl);
		return 0;
	}

	CFArrayRef windows = (CFArrayRef)windowsRef;
	CFIndex count = CFArrayGetCount(windows);
	int found = 0;

	NSString *folder = [NSString stringWithUTF8String:folderName];
	for (CFIndex i = 0; i < count; i++) {
		AXUIElementRef w = (AXUIElementRef)CFArrayGetValueAtIndex(windows, i);
		CFTypeRef titleRef = NULL;
		if (AXUIElementCopyAttributeValue(w, CFSTR("AXTitle"), &titleRef) != kAXErrorSuccess) continue;

		NSString *title = (__bridge NSString *)titleRef;
		BOOL matched = titleMatchesFolder(title, folder);
		CFRelease(titleRef);
		if (matched) {
			found = restoreAndRaiseWindow(appEl, w);
			break;
		}
	}

	CFRelease(windowsRef);
	CFRelease(appEl);
	return found;
}
*/
import "C"

import (
	"time"
)

const windowFocusRetryAfterRestore = 2

const ghosttyAppleScriptFocusTimeout = 1500 * time.Millisecond

var ghosttyAppleScriptRunner = runGhosttyAppleScriptFocus
var ghosttyAppleScriptIDRunner = runGhosttyAppleScriptFocusByID

type FocusWindowOptions struct {
	GhosttyTerminalID string
}

// retryWindowFocus calls fn with increasing delays until a non-zero result.
// Returns 1 (found), -1 (no permission), or 0 (not found after all attempts).
// Worst case: 150+250+400 = 800ms. Best case: 150ms.
func retryWindowFocus(fn func() C.int) C.int { _ = "STUB: not implemented"; return *new(C.int) }

func retryWindowFocusWithDelays(fn func() int, delays []time.Duration, sleep func(time.Duration)) int {
	_ = "STUB: not implemented"
	return 0
}

// FocusAppWindow raises the window matching cwd for the given bundleID app.
// For Ghostty: first tries an exact terminal ID match when available, then
// falls back to exact AppleScript cwd matching, then finally AXDocument
// (OSC 7 file:// URL) if Automation is unavailable or no exact match is found.
// For other apps: uses CGS to find the window across Spaces then raises via AXTitle. macOS only.
func FocusAppWindow(bundleID, cwd string) error { _ = "STUB: not implemented"; return nil }

func FocusAppWindowWithOptions(bundleID, cwd string, opts FocusWindowOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Ask macOS for permission first. If access is still unavailable after
// the system prompt flow, show our own explanation with a Settings link.

// Window not found by title, but still activate the app so the user gets
// at least app-level focus. This matches the previous AppleScript behavior
// which always called "activate" before searching for windows.

func focusGhosttyWindow(
	pid int,
	bundleID, cwd string,
	exactFocus func(string) error,
	fallback func(int, string, string) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func focusGhosttyWindowWithOptions(
	pid int,
	bundleID, cwd string,
	opts FocusWindowOptions,
	exactFocus func(string, FocusWindowOptions) error,
	fallback func(int, string, string) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func tryGhosttyExactFocus(cwd string, opts FocusWindowOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func tryGhosttyAppleScriptFocus(cwd string) error { _ = "STUB: not implemented"; return nil }

func runGhosttyAppleScriptFocusByID(terminalID string) error { _ = "STUB: not implemented"; return nil }

func runGhosttyAppleScriptFocus(candidates []string) error { _ = "STUB: not implemented"; return nil }

func ghosttyFocusCandidates(cwd string) []string { _ = "STUB: not implemented"; return nil }

func normalizeGhosttyWorkingDir(path string) string { _ = "STUB: not implemented"; return "" }

func focusGhosttyWindowByAXDocument(pid int, bundleID, cwd string) error {
	_ = "STUB: not implemented"
	return nil
}

const ghosttyFocusAppleScript = `
on normalizePath(thePath)
	if thePath is "/" then
		return "/"
	end if
	if thePath ends with "/" then
		return text 1 thru -2 of thePath
	end if
	return thePath
end normalizePath

on focusMatchingTerminal(searchPaths)
	tell application "Ghostty"
		set allTerminals to terminals
		repeat with candidate in searchPaths
			set normalizedCandidate to my normalizePath(contents of candidate)
			repeat with t in allTerminals
				try
					set termDir to my normalizePath(working directory of t)
					if termDir is normalizedCandidate then
						focus t
						return true
					end if
				end try
			end repeat
		end repeat
	end tell
	return false
end focusMatchingTerminal

on run argv
	if my focusMatchingTerminal(argv) then
		return
	end if
	error "ghostty terminal not found" number 1001
end run
`

const ghosttyFocusByIDAppleScript = `
on run argv
	if (count of argv) is 0 then
		error "missing Ghostty terminal id" number 1002
	end if

	set wantedID to item 1 of argv

	tell application "Ghostty"
		repeat with t in terminals
			try
				if (id of t as string) is wantedID then
					focus t
					return
				end if
			end try
		end repeat
	end tell

	error "ghostty terminal id not found" number 1002
end run
`

// promptScreenRecordingOnce sends a one-time notification explaining why Screen
// Recording access is needed. Clicking the notification opens the settings pane.
// Uses the plugin's own notification system without an osascript fallback.
func promptScreenRecordingOnce() { _ = "STUB: not implemented"; return }

// already prompted

// Mark as prompted before sending (avoid duplicate prompts on error)

// promptAccessibilityOnce sends a one-time notification explaining why
// Accessibility access is needed for click-to-focus window selection.
func promptAccessibilityOnce() { _ = "STUB: not implemented"; return }

// already prompted

// Mark as prompted before sending (avoid duplicate prompts on error)
