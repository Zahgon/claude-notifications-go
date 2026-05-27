//go:build !windows

package teamstate

// withFileLock executes fn while holding an exclusive file lock (flock).
// This ensures cross-process safety between Stop and TeammateIdle hooks.
func withFileLock(teamName string, fn func() error) error { _ = "STUB: not implemented"; return nil }
