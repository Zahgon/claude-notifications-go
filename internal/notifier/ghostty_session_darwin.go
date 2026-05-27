//go:build darwin

package notifier

import (
	"time"
)

const ghosttyFrontmostTerminalInfoTimeout = 1500 * time.Millisecond

type ghosttyFrontmostTerminalInfo struct {
	ID               string
	WorkingDirectory string
	Name             string
}

type ghosttyWorktreeContext struct {
	RepoRoot     string
	WorktreeName string
}

var ghosttyFrontmostTerminalInfoRunner = readGhosttyFrontmostTerminalInfo

// MaybeCaptureGhosttyTerminalID stores the current Ghostty terminal ID for the
// Claude session when we can confidently prove the frontmost Ghostty terminal is
// the active Claude session that emitted this hook.
func MaybeCaptureGhosttyTerminalID(configOverride, sessionID, cwd string) {
	_ = "STUB: not implemented"
	return
}

func loadStoredGhosttyTerminalID(sessionID string) string { _ = "STUB: not implemented"; return "" }

func ghosttyFrontmostTerminalMatchesSession(info ghosttyFrontmostTerminalInfo, cwd string) bool {
	_ = "STUB: not implemented"
	return false
}

func deriveGhosttyWorktreeContext(cwd string) (ghosttyWorktreeContext, bool) {
	_ = "STUB: not implemented"
	return *new(ghosttyWorktreeContext), false
}

func ghosttyTerminalNameMatchesWorktree(lowerName, worktreeName string) bool {
	_ = "STUB: not implemented"
	return false
}

func ghosttyTerminalNameContainsWorktree(name, worktreeName string) bool {
	_ = "STUB: not implemented"
	return false
}

func readGhosttyFrontmostTerminalInfo() (ghosttyFrontmostTerminalInfo, error) {
	_ = "STUB: not implemented"
	return *new(ghosttyFrontmostTerminalInfo), nil
}

const ghosttyFrontmostTerminalInfoAppleScript = `
on normalizePath(thePath)
	if thePath is "/" then
		return "/"
	end if
	if thePath ends with "/" then
		return text 1 thru -2 of thePath
	end if
	return thePath
end normalizePath

on run
	tell application "Ghostty"
		if not frontmost then
			error "Ghostty not frontmost" number 1003
		end if
		set term to focused terminal of selected tab of front window
		set delim to ASCII character 31
		return ((id of term as string) & delim & (my normalizePath(working directory of term)) & delim & (name of term as string))
	end tell
end run
`
