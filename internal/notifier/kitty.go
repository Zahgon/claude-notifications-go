package notifier

// IsKitty returns true if the current process is running inside Kitty
// with remote control enabled. Checks both $KITTY_WINDOW_ID (always set
// in Kitty) and $KITTY_LISTEN_ON (only set when remote control is configured).
func IsKitty() bool { _ = "STUB: not implemented"; return false }

// getKittyPath returns the absolute path to the kitten binary.
// ClaudeNotifier.app runs without the user's PATH, so we need the full path.
func getKittyPath() string { _ = "STUB: not implemented"; return "" }

// GetKittyWindowTarget returns the window ID and listen socket path from environment variables.
func GetKittyWindowTarget() (windowID, listenOn string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// buildKittyNotifierArgs constructs command-line arguments for terminal-notifier
// when running inside Kitty. Uses -activate (to focus the terminal app)
// and -execute (to switch to the correct Kitty window) on click.
func buildKittyNotifierArgs(title, message, windowID, listenOn, bundleID string) []string {
	_ = "STUB: not implemented"
	return nil
}
