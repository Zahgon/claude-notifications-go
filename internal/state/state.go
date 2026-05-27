package state

import (
	"github.com/777genius/claude-notifications/internal/analyzer"
)

// SessionState represents per-session state
type SessionState struct {
	SessionID               string `json:"session_id"`
	LastInteractiveTool     string `json:"last_interactive_tool"`
	LastTimestamp           int64  `json:"last_ts"`
	LastTaskCompleteTime    int64  `json:"last_task_complete_ts,omitempty"`
	LastNotificationTime    int64  `json:"last_notification_ts,omitempty"`
	LastNotificationStatus  string `json:"last_notification_status,omitempty"`
	LastNotificationMessage string `json:"last_notification_message,omitempty"`
	GhosttyTerminalID       string `json:"ghostty_terminal_id,omitempty"`
	CWD                     string `json:"cwd"`
}

// Manager manages session state
type Manager struct {
	tempDir string
}

// NewManager creates a new state manager
func NewManager() *Manager { _ = "STUB: not implemented"; return nil }

// getStatePath returns the path to the state file for a session
func (m *Manager) getStatePath(sessionID string) string { _ = "STUB: not implemented"; return "" }

// Load loads session state from disk
// Returns nil if state file doesn't exist
func (m *Manager) Load(sessionID string) (*SessionState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Save saves session state to disk
func (m *Manager) Save(state *SessionState) error { _ = "STUB: not implemented"; return nil }

// Delete deletes session state
func (m *Manager) Delete(sessionID string) error { _ = "STUB: not implemented"; return nil }

// UpdateInteractiveTool updates the last interactive tool and timestamp
func (m *Manager) UpdateInteractiveTool(sessionID, toolName, cwd string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateGhosttyTerminalID stores the exact Ghostty terminal ID associated with a
// Claude session so future notification clicks can target the correct tab.
func (m *Manager) UpdateGhosttyTerminalID(sessionID, terminalID string) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateTaskComplete updates the last task complete timestamp
func (m *Manager) UpdateTaskComplete(sessionID string) error { _ = "STUB: not implemented"; return nil }

// ShouldSuppressQuestion checks if a question notification should be suppressed
// due to being within the cooldown window after a task completion
func (m *Manager) ShouldSuppressQuestion(sessionID string, cooldownSeconds int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Check if we're within the cooldown window

// UpdateState updates state based on the detected status
func (m *Manager) UpdateState(sessionID string, status analyzer.Status, toolName, cwd string) error {
	_ = "STUB: not implemented"
	return nil
}

// Cleanup cleans up old state files (older than maxAge seconds)
func (m *Manager) Cleanup(maxAge int64) error { _ = "STUB: not implemented"; return nil }

// UpdateLastNotification updates the last notification timestamp, status, and message
func (m *Manager) UpdateLastNotification(sessionID string, status analyzer.Status, message string) error {
	_ = "STUB: not implemented"
	return nil
}

// ShouldSuppressQuestionAfterAnyNotification checks if a question notification should be suppressed
// due to being within the cooldown window after ANY notification
func (m *Manager) ShouldSuppressQuestionAfterAnyNotification(sessionID string, cooldownSeconds int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Check if we're within the cooldown window

// Import logging to add debug output
// Note: This creates a circular dependency, so we'll skip logging here
// and rely on the caller to log the result

// normalizeMessage normalizes a message for comparison by:
// - Trimming whitespace
// - Removing trailing dots
// - Converting to lowercase
func normalizeMessage(msg string) string { _ = "STUB: not implemented"; return "" }

// IsDuplicateMessage checks if the given message is a duplicate of a recent notification
// within the specified time window (in seconds)
func (m *Manager) IsDuplicateMessage(sessionID string, message string, windowSeconds int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Check if we're within the time window

// Compare normalized messages
