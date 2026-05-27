package notifier

import (
	"os/exec"
	"sync"

	"github.com/777genius/claude-notifications/internal/analyzer"
	"github.com/777genius/claude-notifications/internal/audio"
	"github.com/777genius/claude-notifications/internal/config"
)

const macOSPermissionDeniedMessage = "Notification permission denied. Enable in System Settings > Notifications."

var execCommand = exec.Command

// NotificationPermissionDeniedError indicates macOS rejected the native
// ClaudeNotifier path because notification permission is denied for the app.
type NotificationPermissionDeniedError struct {
	Details string
}

func (e *NotificationPermissionDeniedError) Error() string { _ = "STUB: not implemented"; return "" }

// Notifier sends desktop notifications
type Notifier struct {
	cfg         *config.Config
	audioPlayer *audio.Player
	playerInit  sync.Once
	playerErr   error
	mu          sync.Mutex
	wg          sync.WaitGroup
	closing     bool // Prevents new sounds from being enqueued after Close() is called
}

// New creates a new notifier
func New(cfg *config.Config) *Notifier { _ = "STUB: not implemented"; return nil }

// isTimeSensitiveStatus returns true for statuses that should break through Focus Mode
func isTimeSensitiveStatus(status analyzer.Status) bool { _ = "STUB: not implemented"; return false }

// SendDesktop sends a desktop notification.
// On macOS, it always prefers ClaudeNotifier/terminal-notifier to avoid
// Script Editor attribution and optionally enables click-to-focus.
// On Linux with clickToFocus enabled, it uses the background daemon.
// cwd is the working directory of the project; used for window-specific focus. May be empty.
func (n *Notifier) SendDesktop(status analyzer.Status, message, sessionID, cwd string) error {
	_ = "STUB: not implemented"
	// Send terminal bell for terminal tab indicators (e.g. Ghostty, tmux)
	return nil
}

// Extract session name, git branch and folder name from message
// Format: "[session-name|branch folder] actual message" or "[session-name folder] actual message"

// Build clean title (status only + session name)
// Format: "✅ Completed [peak]" or "✅ Completed"

// Build subtitle from branch and folder name
// Format: "main · notification_plugin_go" or just folder name

// gitBranch may contain "branch folder" (space-separated from hooks.go format)

// Get app icon path if configured

// macOS: prefer ClaudeNotifier/terminal-notifier so the common path keeps the
// native app attribution. If that fails, fall back to beeep as a delivery
// safety net rather than dropping the notification entirely.

// Linux: Try daemon for click-to-focus support

// Fall through to beeep

// Standard path: beeep (Windows, Linux fallback)

// sendWithTerminalNotifier sends notification via terminal-notifier on macOS
// with click-to-focus support (clicking notification activates the terminal)
func (n *Notifier) sendWithTerminalNotifier(title, message, subtitle, sessionID string, timeSensitive bool, cwd string, clickToFocus bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Append shared options: subtitle, threadID, timeSensitive, nosound

// Always suppress sound in Swift — Go manages sound via audio player

func runClaudeNotifierApp(appPath string, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// buildNotifierCommand builds the execution command for a notifier binary.
// ClaudeNotifier.app must be launched via LaunchServices so
// UNUserNotificationCenter gets valid bundle metadata under hardened runtime.
func buildNotifierCommand(notifierPath string, args []string) *exec.Cmd {
	_ = "STUB: not implemented"
	return nil
}

// claudeNotifierAppPath extracts ClaudeNotifier.app from the embedded
// terminal-notifier-modern executable path.
func claudeNotifierAppPath(notifierPath string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// buildTerminalNotifierArgs constructs command-line arguments for terminal-notifier.
// When cwd is provided, uses -execute with a focus script instead of -activate.
// Exported for testing purposes.
func buildTerminalNotifierArgs(title, message, bundleID, cwd string, clickToFocus bool) []string {
	_ = "STUB: not implemented"
	return nil
}

func buildTerminalNotifierArgsWithOptions(title, message, bundleID, cwd, ghosttyTerminalID string, clickToFocus bool) []string {
	_ = "STUB: not implemented"
	return nil
}

// Note: -sender option removed because it conflicts with -activate on macOS Sequoia (15.x)
// Using -sender causes click-to-focus to stop working.

// Add group ID to prevent notification stacking issues

// buildFocusScript returns the shell command for -execute in terminal-notifier.
// For Ghostty: uses AXDocument attribute (OSC 7 CWD) via Accessibility API,
// falling back to plain app activation.
// For all apps (including Electron editors and regular terminals): invokes the
// binary's focus-window subcommand which uses CGS + AXTitle APIs to find and
// raise the correct window across Spaces.
// Returns "" when cwd is empty or unusable (caller should use -activate instead).
func buildFocusScript(bundleID, cwd string) string { _ = "STUB: not implemented"; return "" }

func buildFocusScriptWithOptions(bundleID, cwd, ghosttyTerminalID string) string {
	_ = "STUB: not implemented"
	return ""
}

// All other terminals: use focus-window subcommand (AXTitle matching + CGS Space switching).
// Previously used AppleScript (-execute osascript), but macOS Tahoe (26.x) broke
// Automation permission prompts for notification click handlers — osascript fails silently.
// The focus-window approach uses Accessibility + Screen Recording instead of Automation,
// with graceful fallback to app-level activation when permissions are not granted.
// See: https://github.com/777genius/claude-notifications-go/issues/47

// isElectronEditorBundleID reports whether bundleID belongs to an Electron-based
// editor (VS Code, Cursor, etc.). These apps don't support AppleScript window
// enumeration (-1708) and require the binary focus-window subcommand instead.
func isElectronEditorBundleID(bundleID string) bool { _ = "STUB: not implemented"; return false }

// Cursor

// isGhosttyBundleID reports whether bundleID is Ghostty.
func isGhosttyBundleID(bundleID string) bool { _ = "STUB: not implemented"; return false }

// shellQuote wraps s in single quotes, escaping internal single quotes
// using the '\" technique (end quote, literal apostrophe, resume quote).
func shellQuote(s string) string { _ = "STUB: not implemented"; return "" }

// buildBinaryFocusScript builds the -execute script for apps that use the
// binary's focus-window subcommand (all macOS terminals including Electron editors and Ghostty).
// Returns "" (causing -activate fallback) if os.Executable() fails.
func buildBinaryFocusScript(bundleID, cwd, ghosttyTerminalID string) string {
	_ = "STUB: not implemented"
	return ""
}

// buildElectronEditorFocusScript builds the -execute script for Electron-based
// editors (VS Code, Cursor). Invokes the binary's focus-window subcommand which
// activates the app, waits for AXWindows to populate, then raises the window
// matching cwd.
func buildElectronEditorFocusScript(bundleID, cwd string) string {
	_ = "STUB: not implemented"
	return ""
}

// buildGhosttyFocusScript builds the -execute script for Ghostty.
// Invokes the binary's focus-window subcommand which activates Ghostty,
// waits for AXWindows to populate, then raises the window matching cwd via
// AXDocument (OSC 7 file:// URL). AXDocument is window-level only; tabs and
// split panes within a window are not individually addressable.
func buildGhosttyFocusScript(bundleID, cwd, ghosttyTerminalID string) string {
	_ = "STUB: not implemented"
	return ""
}

// cwdToFileURL converts an absolute path to a file:// URL. Ghostty exposes the
// window CWD (set via OSC 7) as a file:// URL in the AXDocument attribute.
// Uses net/url for RFC-3986-compliant percent-encoding.
func cwdToFileURL(cwd string) string { _ = "STUB: not implemented"; return "" }

// SendQuickNotification sends a one-off notification without requiring a
// Notifier instance.
// executeCmd is the shell command run when the user clicks the notification (may be empty).
func SendQuickNotification(title, message, executeCmd string) error {
	_ = "STUB: not implemented"
	return nil
}

// Fallback: osascript (no click action, just informational)

// sendWithBeeep sends notification via beeep (cross-platform)
func (n *Notifier) sendWithBeeep(title, message, appIcon, sound string) error {
	_ = "STUB: not implemented"
	// Platform-specific AppName handling:
	// - Windows: Use fixed AppName to prevent registry pollution. Each unique AppName
	//   creates a persistent entry in HKEY_CURRENT_USER\SOFTWARE\Microsoft\Windows\
	//   CurrentVersion\Notifications\Settings\ that is never cleaned up.
	//   See: https://github.com/777genius/claude-notifications-go/issues/4
	// - macOS/Linux: Use unique AppName to prevent notification grouping/replacement,
	//   allowing multiple notifications to be displayed simultaneously.
	return nil
}

// Send notification using beeep with proper title and clean message

// playSoundDetached spawns a detached child process to play the sound.
// The parent hook process does not wait for audio to finish, eliminating
// the ~3.6s delay from notifier.Close() wg.Wait().
// Falls back to playSoundAsync (inline playback) if spawn fails.
func (n *Notifier) playSoundDetached(sound string) { _ = "STUB: not implemented"; return }

// Do NOT call cmd.Wait() — child process runs independently

// playSoundAsync plays sound asynchronously if enabled (inline fallback)
func (n *Notifier) playSoundAsync(sound string) { _ = "STUB: not implemented"; return }

// Check if notifier is closing to prevent WaitGroup race

// Use SafeGo to protect against panics in sound playback goroutine

// initPlayer initializes the audio player once
func (n *Notifier) initPlayer() error { _ = "STUB: not implemented"; return nil }

// playSound plays a sound file using the audio module
func (n *Notifier) playSound(soundPath string) { _ = "STUB: not implemented"; return }

// Initialize player once

// Play sound

// Close waits for all sounds to finish playing and cleans up resources
func (n *Notifier) Close() error {
	_ = "STUB: not implemented"
	// Set closing flag to prevent new sounds from being enqueued
	return nil
}

// Wait for all sounds to finish

// Close audio player if it was initialized

// sendTerminalBell writes a BEL character to /dev/tty to trigger terminal
// tab indicators (e.g. Ghostty tab highlight, tmux window bell flag).
//
// When the hook subprocess has no controlling tty (notably Claude Code hooks,
// which detach from the parent terminal), /dev/tty open fails with ENXIO.
// In that case we fall back to writing BEL into the tmux pane's tty directly,
// using $TMUX_PANE to locate the pane and `tmux display-message` to resolve
// its tty path. Tmux reads the BEL from the pty and sets the window's bell
// flag, so tab indicators (window-status-bell-style) still light up.
func sendTerminalBell() { _ = "STUB: not implemented"; return }

// sendTmuxPaneBell writes a BEL byte to the current tmux pane's tty as a
// fallback path for environments without a controlling tty (e.g. Claude Code
// hook subprocesses). No-op when not running under tmux.
func sendTmuxPaneBell() { _ = "STUB: not implemented"; return }

// extractSessionInfo extracts session name and git branch from message
// Format: "[session-name|branch] message" or "[session-name] message"
// Returns session name, git branch (may be empty), and clean message
func extractSessionInfo(message string) (sessionName, gitBranch, cleanMessage string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// Check if message starts with [

// Find closing bracket

// Extract content inside brackets

// Check if there's a pipe separator for git branch

// Extract clean message (everything after "] ")
