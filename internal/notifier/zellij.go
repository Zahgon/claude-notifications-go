package notifier

// IsZellij returns true if the current process is running inside a zellij session.
func IsZellij() bool { _ = "STUB: not implemented"; return false }

// getZellijPath returns the absolute path to the zellij binary.
// ClaudeNotifier.app runs without the user's PATH, so we need the full path.
func getZellijPath() string { _ = "STUB: not implemented"; return "" }

// GetZellijTabTarget returns the active tab name and session name for the current zellij session.
func GetZellijTabTarget() (tabName, sessionName string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// parseActiveTabName extracts the name of the focused tab from zellij dump-layout output (KDL format).
// It looks for lines matching: tab ... name="..." ... focus=true
func parseActiveTabName(layout string) string { _ = "STUB: not implemented"; return "" }

// Must be a top-level "tab" line (not "pane" or nested content)

// Must have focus=true

// Extract name="..."

// extractKDLStringAttr extracts the value of a key="value" attribute from a KDL line.
func extractKDLStringAttr(line, key string) string {
	_ = "STUB: not implemented"
	// Search for key="
	return ""
}

// Find closing quote

// buildZellijNotifierArgs constructs command-line arguments for terminal-notifier
// when running inside zellij. Uses -activate (to focus the terminal app)
// and -execute (to switch to the correct zellij tab) on click.
func buildZellijNotifierArgs(title, message, tabName, sessionName, bundleID string) []string {
	_ = "STUB: not implemented"
	return nil
}

// Add group ID to prevent notification stacking issues
