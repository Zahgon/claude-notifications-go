package notifier

// getiTerm2PythonEnv returns the absolute paths to the Python interpreter
// inside the iTerm2 venv and the tab-switch helper script.
// Returns ("", "", false) if either is not found.
func getiTerm2PythonEnv() (pythonPath string, scriptPath string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

const iTerm2BundleID = "com.googlecode.iterm2"

func isIterm2BundleID(bundleID string) bool { _ = "STUB: not implemented"; return false }

// buildIterm2TmuxNotifierArgs constructs terminal-notifier arguments for
// iTerm2 + tmux. The Python helper handles both tmux -CC (via tmuxWindowPane)
// and plain tmux (via tmux client tty fallback) to avoid mutating the wrong
// tmux client in multi-tab setups.
func buildIterm2TmuxNotifierArgs(title, message, paneTarget, bundleID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildTmuxCCNotifierArgs is kept as a compatibility wrapper for the existing
// tmux -CC tests; the helper now supports both control mode and plain tmux.
func buildTmuxCCNotifierArgs(title, message, paneTarget, bundleID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
