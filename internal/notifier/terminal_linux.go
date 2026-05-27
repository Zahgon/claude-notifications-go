//go:build linux

// ABOUTME: Linux-specific notification handling with click-to-focus support.
// ABOUTME: Uses background daemon for persistent D-Bus connection when click-to-focus is enabled.
package notifier

import (
	"github.com/777genius/claude-notifications/internal/config"
)

// macOS stub functions - these are not used on Linux but required for compilation

// GetTerminalBundleID returns empty string on Linux
// as terminal bundle IDs are a macOS-specific concept.
func GetTerminalBundleID(configOverride string) string {
	_ = "STUB: not implemented"

	// GetTerminalNotifierPath returns an error on Linux
	// as terminal-notifier is macOS-only.
	return ""
}

func GetTerminalNotifierPath() (string, error) { _ = "STUB: not implemented"; return "", nil }

// IsTerminalNotifierAvailable returns false on Linux.
func IsTerminalNotifierAvailable() bool {
	_ = "STUB: not implemented"

	// EnsureClaudeNotificationsApp is a no-op on Linux.
	return false
}

func EnsureClaudeNotificationsApp() error {
	_ = "STUB: not implemented"

	// sendLinuxNotification sends a notification on Linux.
	// When clickToFocus is enabled, uses the daemon for click-to-focus support.
	// Falls back to beeep when daemon is unavailable.
	// cwd is the working directory of the project; used for window-specific focus. May be empty.
	return nil
}

func sendLinuxNotification(title, body, appIcon string, cfg *config.Config, cwd string) error {
	_ = "STUB: not implemented"
	// If click-to-focus is disabled, use beeep directly
	return nil
}

// Try to use daemon for click-to-focus

// Fallback to beeep (no click-to-focus)

// sendViaDaemon sends a notification via the background daemon.
// Returns an error if daemon is not available or fails.
// cwd is used to extract the project folder name for window-specific focus.
func sendViaDaemon(title, body, cwd string) error {
	_ = "STUB: not implemented"
	// Start daemon on-demand (no-op if already running)
	return nil
}

// Create client and send notification

// Extract folder name from cwd for title-based window focus

// Send notification with 30 second timeout.
// Detect focus target in the hook process (not the daemon), since the daemon may
// have been started from a different environment.

// IsDaemonAvailable checks if the notification daemon is available and running.
// Exported for testing and status checks.
func IsDaemonAvailable() bool { _ = "STUB: not implemented"; return false }

// StartDaemon starts the notification daemon on-demand.
// Returns true if daemon started successfully or was already running.
func StartDaemon() bool { _ = "STUB: not implemented"; return false }

// StopDaemon stops the running notification daemon.
func StopDaemon() error { _ = "STUB: not implemented"; return nil }
