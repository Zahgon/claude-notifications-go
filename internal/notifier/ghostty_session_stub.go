//go:build !darwin

package notifier

func MaybeCaptureGhosttyTerminalID(configOverride, sessionID, cwd string) {
	_ = "STUB: not implemented"
	return
}

func loadStoredGhosttyTerminalID(sessionID string) string { _ = "STUB: not implemented"; return "" }
