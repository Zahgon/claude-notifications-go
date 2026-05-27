package platform

// GitMetadata contains commonly requested git values for a working tree.
type GitMetadata struct {
	Branch            string
	CommitHash        string
	CommitShortHash   string
	CommitAuthorName  string
	CommitAuthorEmail string
	UserName          string
	UserEmail         string
}

// GetGitBranch returns the current git branch name for the given directory.
// Returns empty string if not in a git repository or on error.
func GetGitBranch(cwd string) string { _ = "STUB: not implemented"; return "" }

// GetGitMetadata returns commonly used git metadata for the given directory.
// Empty fields indicate the directory is not a git repo or the value is unavailable.
func GetGitMetadata(cwd string) GitMetadata { _ = "STUB: not implemented"; return *new(GitMetadata) }

func getGitOutput(cwd string, args ...string) string { _ = "STUB: not implemented"; return "" }

// "HEAD" is returned when in detached HEAD state
