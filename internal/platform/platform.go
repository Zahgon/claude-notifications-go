package platform

// OS returns the current operating system
func OS() string { _ = "STUB: not implemented"; return "" }

// TempDir returns the platform-specific temporary directory (without trailing slash)
func TempDir() string { _ = "STUB: not implemented"; return "" }

// Remove trailing slash if present (macOS $TMPDIR ends with /)

// FileMTime returns the modification time of a file as Unix timestamp
// Returns 0 if the file doesn't exist or on error
func FileMTime(path string) int64 { _ = "STUB: not implemented"; return 0 }

// CurrentTimestamp returns the current Unix timestamp
func CurrentTimestamp() int64 { _ = "STUB: not implemented"; return 0 }

// FileAge returns the age of a file in seconds
// Returns -1 if the file doesn't exist
func FileAge(path string) int64 { _ = "STUB: not implemented"; return 0 }

// FileExists checks if a file exists
func FileExists(path string) bool { _ = "STUB: not implemented"; return false }

// CleanupOldFiles removes files older than maxAge seconds matching a pattern
func CleanupOldFiles(dir, pattern string, maxAge int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore errors

// AtomicCreateFile creates a file atomically using O_EXCL flag
// Returns true if file was created, false if it already exists
func AtomicCreateFile(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NormalizePath normalizes a file path (removes double slashes, etc.)
func NormalizePath(path string) string { _ = "STUB: not implemented"; return "" }

// ExpandEnv expands environment variables in a string (like ${VAR} or $VAR)
func ExpandEnv(s string) string { _ = "STUB: not implemented"; return "" }

// IsWindows returns true if running on Windows
func IsWindows() bool { _ = "STUB: not implemented"; return false }

// IsMacOS returns true if running on macOS
func IsMacOS() bool { _ = "STUB: not implemented"; return false }

// IsLinux returns true if running on Linux
func IsLinux() bool { _ = "STUB: not implemented"; return false }
