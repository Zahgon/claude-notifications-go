package daemon

const claudeNotificationsDesktopEntryID = "claude-notifications"

// escapeJS escapes a string for safe interpolation into JavaScript single-quoted strings.
// Prevents JS injection when values are passed to GNOME Shell.Eval.
func escapeJS(s string) string { _ = "STUB: not implemented"; return "" }

// GetAppID returns the .desktop app ID for a terminal name.
func GetAppID(terminalName string) string { _ = "STUB: not implemented"; return "" }

// GetDesktopEntryID returns the desktop entry ID (without .desktop suffix) for a terminal.
// This is the value expected by the freedesktop "desktop-entry" notification hint.
func GetDesktopEntryID(terminalName string) string { _ = "STUB: not implemented"; return "" }

// GetNotificationDesktopEntryID returns the desktop-entry hint value to use for
// notifications. GNOME on Wayland shows a long-running loading cursor when the
// clicked notification advertises the terminal/editor desktop entry and the app
// never consumes the generated activation token. A dedicated hidden desktop
// file with StartupNotify=false avoids that spinner while preserving click
// handling via our daemon.
func GetNotificationDesktopEntryID(terminalName string) string {
	_ = "STUB: not implemented"
	return ""
}

func isGnomeWaylandSession() bool { _ = "STUB: not implemented"; return false }

func hasClaudeNotificationsDesktopEntry() bool { _ = "STUB: not implemented"; return false }

func getClaudeNotificationsDesktopEntryPath() string { _ = "STUB: not implemented"; return "" }

// GetWlrctlAppID returns the wlroots app_id for a terminal name.
func GetWlrctlAppID(terminalName string) string { _ = "STUB: not implemented"; return "" }

// GetKdotoolClass returns the window class for kdotool search.
func GetKdotoolClass(terminalName string) string { _ = "STUB: not implemented"; return "" }

// GetXdotoolClass returns the X11 WM_CLASS for xdotool search.
func GetXdotoolClass(terminalName string) string { _ = "STUB: not implemented"; return "" }

// GetSearchTerm returns a window title search term for a terminal name.
func GetSearchTerm(terminalName string) string { _ = "STUB: not implemented"; return "" }

// GetSearchTermWithFolder returns the window title search term, using the project
// folder name for VS Code when available (more specific than "Visual Studio Code").
func GetSearchTermWithFolder(terminalName, folderName string) string {
	_ = "STUB: not implemented"
	return ""
}

// GetTerminalName detects the current terminal from environment variables.
func GetTerminalName() string {
	_ = "STUB: not implemented"
	// Try TERM_PROGRAM first (set by many terminals)
	return ""
}

// Check VS Code indicators

// Check GNOME Terminal indicators

// Check Terminator (does not set TERM_PROGRAM, but always sets TERMINATOR_UUID)

// Fallback to generic terminal

// GetX11WindowID returns the current terminal window's X11 window ID when available.
// It is captured in the hook process and later used by the daemon for exact focus on X11.
func GetX11WindowID() string { _ = "STUB: not implemented"; return "" }

// GetExactWindowTitle returns an exact top-level window title for terminals that expose
// a reliable per-terminal identifier. Currently Terminator can provide this via
// TERMINATOR_UUID + remotinator.
func GetExactWindowTitle(terminalName string) string { _ = "STUB: not implemented"; return "" }

func getTerminatorWindowTitle() string { _ = "STUB: not implemented"; return "" }
