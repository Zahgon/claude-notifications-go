package notifier

import (
	"time"
)

const iTerm2SessionIDEnv = "ITERM_SESSION_ID"

const (
	iTerm2HealthcheckFlag           = "--healthcheck"
	iTerm2HealthcheckExitDisabled   = 11
	iTerm2HealthcheckExitModuleMiss = 12
	iTerm2HealthcheckExitOther      = 13
)

var sendQuickNotification = SendQuickNotification

var (
	iTerm2HealthcheckSuccessTTL = 10 * time.Minute
	iTerm2PythonAPIPromptTTL    = 24 * time.Hour
	iTerm2PythonAPIHealthcheck  = checkIterm2PythonAPIHealth
)

type iTerm2HelperHealth int

const (
	iTerm2HelperUnavailable iTerm2HelperHealth = iota
	iTerm2HelperReady
	iTerm2HelperDisabled
)

// buildIterm2FocusScript prefers iTerm2's exact session reveal URL when the
// current shell exported ITERM_SESSION_ID. This targets the precise tab/pane
// via the iTerm2 Python API helper. If the helper is unavailable or the exact
// session can no longer be resolved, it falls back to app-level activation
// instead of focus-window to avoid confusing Screen Recording prompts on iTerm2.
func buildIterm2FocusScript(cwd string) string { _ = "STUB: not implemented"; return "" }

// If the exact helper fails at click time, keep the fallback at simple
// app activation instead of switching to focus-window, which would ask
// for Screen Recording even though the underlying iTerm2 issue is the
// Python API helper.

func checkIterm2PythonAPIHealth(pythonPath, scriptPath string) iTerm2HelperHealth {
	_ = "STUB: not implemented"
	return *new(iTerm2HelperHealth)
}

func shouldPromptIterm2PythonAPI(output string) bool { _ = "STUB: not implemented"; return false }

func isUsableFocusCWD(cwd string) bool { _ = "STUB: not implemented"; return false }

func iTerm2HealthcheckSuccessMarkerPath() string { _ = "STUB: not implemented"; return "" }

func iTerm2PythonAPIPromptMarkerPath() string { _ = "STUB: not implemented"; return "" }

func isRecentMarker(path string, ttl time.Duration) bool { _ = "STUB: not implemented"; return false }

func touchMarker(path string) { _ = "STUB: not implemented"; return }

func promptIterm2PythonAPIDisabled() { _ = "STUB: not implemented"; return }

func stringsTrimSpace(s string) string { _ = "STUB: not implemented"; return "" }
