//go:build !darwin && !linux

package notifier

import (
	"github.com/777genius/claude-notifications/internal/config"
)

// GetTerminalBundleID returns empty string on non-macOS platforms
// as terminal bundle IDs are a macOS-specific concept.
func GetTerminalBundleID(configOverride string) string {
	_ = "STUB: not implemented"

	// GetTerminalNotifierPath returns an error on non-macOS platforms
	// as terminal-notifier is macOS-only.
	return ""
}

func GetTerminalNotifierPath() (string, error) { _ = "STUB: not implemented"; return "", nil }

// IsTerminalNotifierAvailable returns false on non-macOS platforms.
func IsTerminalNotifierAvailable() bool {
	_ = "STUB: not implemented"

	// EnsureClaudeNotificationsApp is a no-op on non-macOS platforms.
	return false
}

func EnsureClaudeNotificationsApp() error {
	_ = "STUB: not implemented"

	// sendLinuxNotification is a stub for non-Linux platforms.
	// On Windows, this falls back to beeep directly.
	return nil
}

func sendLinuxNotification(title, body, appIcon string, cfg *config.Config, cwd string) error {
	_ = "STUB: not implemented"
	return nil
}

// IsDaemonAvailable returns false on non-Linux platforms.
func IsDaemonAvailable() bool {
	_ = "STUB: not implemented"

	// StartDaemon is a no-op on non-Linux platforms.
	return false
}

func StartDaemon() bool {
	_ = "STUB: not implemented"

	// StopDaemon is a no-op on non-Linux platforms.
	return false
}

func StopDaemon() error { _ = "STUB: not implemented"; return nil }
