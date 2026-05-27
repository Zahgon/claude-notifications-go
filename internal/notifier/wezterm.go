package notifier

// IsWezTerm returns true if the current process is running inside WezTerm.
func IsWezTerm() bool { _ = "STUB: not implemented"; return false }

// getWezTermPath returns the absolute path to the wezterm binary.
// ClaudeNotifier.app runs without the user's PATH, so we need the full path.
func getWezTermPath() string { _ = "STUB: not implemented"; return "" }

// GetWezTermPaneTarget returns the pane ID and unix socket path from environment variables.
// No external commands needed — the pane ID is already available in $WEZTERM_PANE.
func GetWezTermPaneTarget() (paneID, socketPath string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// buildWezTermNotifierArgs constructs command-line arguments for terminal-notifier
// when running inside WezTerm. Uses -activate (to focus the terminal app)
// and -execute (to switch to the correct WezTerm pane) on click.
func buildWezTermNotifierArgs(title, message, paneID, socketPath, bundleID string) []string {
	_ = "STUB: not implemented"
	return nil
}
