//go:build !darwin

package notifier

// FocusAppWindow is not supported on non-darwin platforms.
func FocusAppWindow(bundleID, cwd string) error { _ = "STUB: not implemented"; return nil }

// FocusAppWindowWithOptions is not supported on non-darwin platforms.
func FocusAppWindowWithOptions(bundleID, cwd string, opts FocusWindowOptions) error {
	_ = "STUB: not implemented"
	return nil
}

type FocusWindowOptions struct {
	GhosttyTerminalID string
}
