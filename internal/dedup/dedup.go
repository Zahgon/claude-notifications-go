package dedup

// Manager handles deduplication using two-phase locking
type Manager struct {
	tempDir string
}

// NewManager creates a new deduplication manager
func NewManager() *Manager { _ = "STUB: not implemented"; return nil }

// getLockPath returns the path to the lock file for a session and hook event
// If hookEvent is empty, uses a global lock for the session (backward compatibility)
func (m *Manager) getLockPath(sessionID string, hookEvent ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// CheckEarlyDuplicate performs Phase 1 check for duplicates
// Returns true if this is a duplicate and should be skipped
// hookEvent parameter is optional - if provided, checks hook-specific lock file
func (m *Manager) CheckEarlyDuplicate(sessionID string, hookEvent ...string) bool {
	_ = "STUB: not implemented"
	return false
}

// Check lock age

// If mtime is unavailable (Windows issue) or lock is fresh (<2s), treat as duplicate

// AcquireLock performs Phase 2 lock acquisition
// Returns true if lock was successfully acquired
// hookEvent parameter is optional - if provided, uses hook-specific lock file
func (m *Manager) AcquireLock(sessionID string, hookEvent ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Try to create lock atomically

// Lock acquired successfully

// Lock exists - check if it's stale

// If lock is fresh (<2s), we're a duplicate

// Lock is stale - try to replace it
// Ignore error - someone else might have deleted it

// Try again

// ReleaseLock releases a lock (optional, locks are cleaned up automatically)
// hookEvent parameter is optional - if provided, releases hook-specific lock file
func (m *Manager) ReleaseLock(sessionID string, hookEvent ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// Cleanup cleans up old lock files (older than maxAge seconds)
func (m *Manager) Cleanup(maxAge int64) error { _ = "STUB: not implemented"; return nil }

// CleanupForSession cleans up lock file for a specific session
func (m *Manager) CleanupForSession(sessionID string) error { _ = "STUB: not implemented"; return nil }

// AcquireContentLock acquires a lock for content-based deduplication
// Uses a separate lock file with longer TTL (5s) to prevent race conditions
// between different hook types (Stop, Notification) with same content
func (m *Manager) AcquireContentLock(sessionID string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Try to create lock atomically

// Lock exists - check if it's stale (5s TTL for content lock)

// Lock is fresh - wait a bit and try again
// This gives the first process time to complete

// Lock is stale - try to replace it

// ReleaseContentLock releases the content-based deduplication lock
func (m *Manager) ReleaseContentLock(sessionID string) error { _ = "STUB: not implemented"; return nil }
