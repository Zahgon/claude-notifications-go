package teamstate

// teamConfig represents the relevant fields from ~/.claude/teams/{name}/config.json
type teamConfig struct {
	Name          string       `json:"name"`
	LeadSessionID string       `json:"leadSessionId"`
	Members       []teamMember `json:"members"`
}

// teamMember represents a member entry in team config
type teamMember struct {
	AgentID   string `json:"agentId"`
	Name      string `json:"name"`
	AgentType string `json:"agentType"`
}

// TeamInfo holds detected team information for the current session
type TeamInfo struct {
	TeamName   string
	Members    []string // non-lead member names
	ConfigPath string
}

// State tracks team notification state (persisted to /tmp)
type State struct {
	TeamName    string           `json:"team_name"`
	LeadStopped bool             `json:"lead_stopped"`
	LeadStopAt  int64            `json:"lead_stop_at,omitempty"`
	IdleMembers map[string]int64 `json:"idle_members"` // member name → unix timestamp
	NotifiedAt  int64            `json:"notified_at,omitempty"`
}

// Manager handles team detection and state tracking.
// All state mutations use file-level locking (flock) for cross-process safety,
// since Stop and TeammateIdle hooks run as separate OS processes.
type Manager struct {
	claudeDir string // defaults to ~/.claude
}

// NewManager creates a new team state manager.
// claudeDir can be empty to use the default (~/.claude).
func NewManager(claudeDir string) *Manager { _ = "STUB: not implemented"; return nil }

// DetectTeamLead checks if the given session ID is a team lead.
// Returns team info if found, nil otherwise.
func (m *Manager) DetectTeamLead(sessionID string) *TeamInfo { _ = "STUB: not implemented"; return nil }

// Found our team — collect non-lead member names

// Team with no non-lead members — not a real team scenario

// DetectTeamByName finds team info by team name.
// Returns team info if found, nil otherwise.
func (m *Manager) DetectTeamByName(teamName string) *TeamInfo {
	_ = "STUB: not implemented"
	return nil
}

// statePath returns the path to the state file for a team
func statePath(teamName string) string { _ = "STUB: not implemented"; return "" }

// lockPath returns the path to the flock file for a team
func lockPath(teamName string) string { _ = "STUB: not implemented"; return "" }

// loadStateUnlocked reads team state from disk without locking.
// Caller must hold the file lock.
func loadStateUnlocked(teamName string) (*State, error) { _ = "STUB: not implemented"; return nil, nil }

// Corrupted state — return fresh

// saveStateUnlocked persists team state to disk atomically without locking.
// Caller must hold the file lock.
func saveStateUnlocked(s *State) error { _ = "STUB: not implemented"; return nil }

// Atomic write via temp file + rename

// LoadState loads team state from disk (for read-only access / tests).
func (m *Manager) LoadState(teamName string) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SaveState persists team state to disk (for tests).
func (m *Manager) SaveState(s *State) error { _ = "STUB: not implemented"; return nil }

// RecordLeadStopped marks the team lead as stopped and persists state.
// Uses file locking to prevent races with concurrent TeammateIdle hooks.
func (m *Manager) RecordLeadStopped(teamName string) error { _ = "STUB: not implemented"; return nil }

// RecordTeammateIdle marks a teammate as idle and persists state.
// Uses file locking to prevent races with concurrent Stop hooks.
func (m *Manager) RecordTeammateIdle(teamName, teammateName string) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckAllIdle checks if the lead has stopped AND all expected members are idle.
// Returns true if a notification should be sent.
// Uses file locking for consistent reads.
func (m *Manager) CheckAllIdle(teamName string, expectedMembers []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Prevent duplicate notifications: check if we already notified

// MarkNotified records that a notification was sent and resets state for next cycle.
// Uses file locking for atomic read-modify-write.
func (m *Manager) MarkNotified(teamName string) error { _ = "STUB: not implemented"; return nil }

// Reset state for next cycle: lead will stop again, teammates will go idle again

// Cleanup removes state and lock files older than maxAge seconds.
func (m *Manager) Cleanup(maxAgeSec int64) { _ = "STUB: not implemented"; return }
