package notifier

// IsTmux returns true if the current process is running inside a tmux session.
func IsTmux() bool { _ = "STUB: not implemented"; return false }

// getTmuxSocketPath extracts the tmux socket path from the TMUX env var.
// TMUX format: "/private/tmp/tmux-501/default,12345,0"
func getTmuxSocketPath() string { _ = "STUB: not implemented"; return "" }

// Socket path is everything before the first comma

// getTmuxPath returns the absolute path to the tmux binary.
// ClaudeNotifier.app runs without the user's PATH, so we need the full path.
func getTmuxPath() string { _ = "STUB: not implemented"; return "" }

// GetTmuxPaneTarget returns the tmux pane ID (e.g. "%42") of the pane where
// Claude Code is running, for use with tmux select-pane / select-window commands.
//
// Prefers $TMUX_PANE (set by tmux per-pane at creation, always points to the
// process's own pane) over "tmux display-message" (which returns the currently
// active pane and may be wrong if the user switched tabs).
func GetTmuxPaneTarget() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Best-effort fallback for environments where TMUX_PANE is not available.
// Use the resolved tmux binary and explicit socket so this still works when
// the notifier runs outside the user's shell PATH.

// IsTmuxControlMode returns true if any tmux client attached to the current
// server is running in control mode (-CC), typically used by iTerm2.
// In control mode, standard tmux select-window doesn't cause iTerm2 to
// switch tabs; the iTerm2 Python API must be used instead.
//
// Uses list-clients (not display-message) because display-message evaluates
// #{client_control_mode} for the temporary command-line client we spawn,
// which is never in control mode. list-clients enumerates all persistent
// (attached) clients, so we can check if ANY is in control mode.
func IsTmuxControlMode() bool { _ = "STUB: not implemented"; return false }

// buildTmuxNotifierArgs constructs command-line arguments for terminal-notifier
// when running inside tmux. Uses both -activate (to focus the terminal app)
// and -execute (to switch to the correct tmux session/window/pane) on click.
func buildTmuxNotifierArgs(title, message, paneTarget, bundleID string) []string {
	_ = "STUB: not implemented"
	// Use absolute path to tmux and explicit socket — ClaudeNotifier.app
	// runs without the user's shell PATH, so bare "tmux" won't be found.
	return nil
}

// switch-client is a separate command so its failure (e.g. no attached
// clients) doesn't abort the select-window/select-pane chain.

// Add group ID to prevent notification stacking issues
