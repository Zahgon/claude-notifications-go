//go:build darwin

package notifier

import (
	"github.com/777genius/claude-notifications/internal/config"
)

// terminalBundleIDMap maps TERM_PROGRAM values to macOS bundle identifiers
var terminalBundleIDMap = map[string]string{
	"Apple_Terminal": "com.apple.Terminal",
	"iTerm.app":      "com.googlecode.iterm2",
	"WarpTerminal":   "dev.warp.Warp-Stable",
	"kitty":          "net.kovidgoyal.kitty",
	"ghostty":        "com.mitchellh.ghostty",
	"WezTerm":        "com.github.wez.wezterm",
	"Alacritty":      "org.alacritty",
	"Hyper":          "co.zeit.hyper",
	"vscode":         "com.microsoft.VSCode",
}

// GetTerminalBundleID determines the bundle ID of the current terminal.
// Priority:
// 1. configOverride (if provided)
// 2. __CFBundleIdentifier env var (set by some terminals like Warp)
// 3. TERM_PROGRAM env var mapped to known bundle IDs
// 4. Inside tmux: TERM_PROGRAM from tmux session environment
// 5. Fallback to com.apple.Terminal
func GetTerminalBundleID(configOverride string) string {
	_ = "STUB: not implemented"
	// 1. Use config override if provided
	return ""
}

// 2. Check __CFBundleIdentifier (directly contains bundle ID)

// 3. Map TERM_PROGRAM to bundle ID

// 4. Inside tmux: check TERM_PROGRAM from tmux session environment

// 5. Fallback to standard Terminal.app

// getBundleIDFromTmuxEnv retrieves TERM_PROGRAM from the tmux environment.
// Inside tmux, TERM_PROGRAM is overwritten to "tmux", but the original value
// is preserved in tmux's global environment (set by the terminal that started tmux).
func getBundleIDFromTmuxEnv() string {
	_ = "STUB: not implemented"
	// Try session environment first, then global
	return ""
}

// Output format: "TERM_PROGRAM=WarpTerminal\n"

// GetTerminalNotifierPath returns the path to terminal-notifier binary.
// Priority:
// 1. terminal-notifier-modern (embedded in plugin): uses UNUserNotificationCenter, works on macOS 10.14+
// 2. terminal-notifier (embedded in plugin): legacy NSUserNotificationCenter
// 3. System-installed (via brew): $(which terminal-notifier)
func GetTerminalNotifierPath() (string, error) { _ = "STUB: not implemented"; return "", nil }

// 1. Check ClaudeNotifier (preferred — modern UNUserNotificationCenter with Claude icon)

// Development checkout fallback: make build-notifier writes the bundle to
// swift-notifier/ClaudeNotifier.app, while plugin-dir runs set
// CLAUDE_PLUGIN_ROOT to the repo root.

// 2. Check legacy terminal-notifier

// 3. Check system installation (brew install terminal-notifier)

// IsTerminalNotifierAvailable checks if terminal-notifier is available
func IsTerminalNotifierAvailable() bool { _ = "STUB: not implemented"; return false }

// EnsureClaudeNotificationsApp creates ClaudeNotifications.app if it doesn't exist.
// This allows the notification icon to work even when users update the plugin
// without running /claude-notifications-go:notifications-init.
func EnsureClaudeNotificationsApp() error { _ = "STUB: not implemented"; return nil }

// Already exists

// Create app structure

// Create iconset and convert to icns

// Generate icon sizes using sips

// Ignore errors, some sizes may fail

// Copy original as 512x512

// Convert to icns

// Create Info.plist

// Create minimal executable

// Register with Launch Services

// sendLinuxNotification is a stub for macOS.
// On macOS, click-to-focus is handled via terminal-notifier.
func sendLinuxNotification(title, body, appIcon string, cfg *config.Config, cwd string) error {
	_ = "STUB: not implemented"
	return nil
}

// IsDaemonAvailable returns false on macOS (Linux daemon is not applicable).
func IsDaemonAvailable() bool {
	_ = "STUB: not implemented"

	// StartDaemon is a no-op on macOS.
	return false
}

func StartDaemon() bool {
	_ = "STUB: not implemented"

	// StopDaemon is a no-op on macOS.
	return false
}

func StopDaemon() error { _ = "STUB: not implemented"; return nil }
