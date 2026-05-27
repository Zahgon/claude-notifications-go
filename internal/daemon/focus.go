//go:build linux

// ABOUTME: Window focus methods for Linux desktop environments.
// ABOUTME: Implements a fallback chain to focus windows on GNOME, KDE, Sway, and other compositors.
package daemon

// FocusMethod represents a method for focusing a window
type FocusMethod struct {
	Name string
	Fn   func(terminalName, folderName string) error
}

// GetFocusMethods returns the ordered list of focus methods to try
func GetFocusMethods() []FocusMethod { _ = "STUB: not implemented"; return nil }

// TryFocus attempts to focus a window using available tools.
// folderName is the project folder name used for title-based window search (may be empty).
// It tries each method in order until one succeeds.
func TryFocus(terminalName, folderName string) error { _ = "STUB: not implemented"; return nil }

// TryFocusWithWindowID preserves the previous API for callers that only have an exact X11 window ID.
func TryFocusWithWindowID(terminalName, folderName, windowID string) error {
	_ = "STUB: not implemented"
	return nil
}

// TryFocusWithHints attempts exact focus using hook-time hints first, then falls back to
// compositor-specific methods.
func TryFocusWithHints(terminalName, folderName, windowID, windowTitle string) error {
	_ = "STUB: not implemented"
	return nil
}

func tryX11WindowID(windowID string) error { _ = "STUB: not implemented"; return nil }

func activateWindowIDWithXdotool(windowID string) error { _ = "STUB: not implemented"; return nil }

func activateWindowIDWithWmctrl(windowID string) error { _ = "STUB: not implemented"; return nil }

func normalizeX11WindowID(windowID string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func tryWindowTitle(windowTitle string) error { _ = "STUB: not implemented"; return nil }

func activateWindowTitleWithWmctrl(windowTitle string) error { _ = "STUB: not implemented"; return nil }

func activateWindowTitleWithXdotool(windowTitle string) error {
	_ = "STUB: not implemented"
	return nil
}

// TryActivateWindowByTitle uses the activate-window-by-title GNOME extension.
// https://extensions.gnome.org/extension/5021/activate-window-by-title/
// This method does NOT require unsafe_mode and works on GNOME 42+.
func TryActivateWindowByTitle(terminalName, folderName string) error {
	_ = "STUB: not implemented"
	return nil
}

// busctl can succeed (exit code 0) even when no window was activated.
// The extension returns a boolean; ensure we only treat "true" as success.

// TryGnomeShellEvalByTitle uses GNOME Shell's Eval to find and focus window by title.
// Requires unsafe_mode or development-tools enabled.
func TryGnomeShellEvalByTitle(terminalName, folderName string) error {
	_ = "STUB: not implemented"
	return nil
}

// JavaScript to find window by title and activate it

// TryGnomeShellEval uses GNOME Shell's Eval method to activate an app.
// Requires unsafe_mode or development-tools enabled.
func TryGnomeShellEval(terminalName, folderName string) error {
	_ = "STUB: not implemented"
	return nil
}

// JavaScript to find and activate the app's windows

// TryGnomeFocusApp uses GNOME Shell's FocusApp method (available since GNOME 45).
func TryGnomeFocusApp(terminalName, folderName string) error { _ = "STUB: not implemented"; return nil }

// TryWlrctl uses wlrctl for wlroots-based compositors (Sway, etc.).
func TryWlrctl(terminalName, folderName string) error { _ = "STUB: not implemented"; return nil }

// Try app_id first (more reliable)

// Fallback to title

// TryKdotool uses kdotool for KDE Plasma.
func TryKdotool(terminalName, folderName string) error { _ = "STUB: not implemented"; return nil }

// Search by class

// TryXdotool uses xdotool for X11-based desktop environments
// (XFCE, MATE, Cinnamon, i3, bspwm, and X11 sessions of GNOME/KDE).
func TryXdotool(terminalName, folderName string) error { _ = "STUB: not implemented"; return nil }

// xdotool returns bottom-most windows first; prefer the top-most candidate.

type xdotoolSearch struct {
	label string
	args  []string
}

func buildXdotoolSearches(terminalName, folderName string) []xdotoolSearch {
	_ = "STUB: not implemented"
	return nil
}

func runXdotoolSearch(args ...string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func splitWindowIDs(output string) []string { _ = "STUB: not implemented"; return nil }

func prioritizeXdotoolCandidates(windowIDs []string, searchLabel, folderName string) []string {
	_ = "STUB: not implemented"
	return nil
}

func getXdotoolWindowName(windowID string) string { _ = "STUB: not implemented"; return "" }

// DetectFocusTools returns a map of available focus tools.
func DetectFocusTools() map[string]bool { _ = "STUB: not implemented"; return nil }

// Check command-line tools

// Check GNOME activate-window-by-title extension
