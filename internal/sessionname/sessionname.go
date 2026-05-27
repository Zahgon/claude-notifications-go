package sessionname

// Lists for friendly name generation (same as bash version)
var adjectives = []string{
	"bold", "brave", "bright", "calm", "clever",
	"cool", "cosmic", "crisp", "daring", "eager",
	"fair", "fancy", "fast", "gentle", "glad",
	"grand", "happy", "kind", "lively", "lucky",
	"merry", "noble", "proud", "quick", "quiet",
	"rapid", "smart", "solid", "swift", "warm",
	"wise", "witty", "zesty", "agile", "alert",
}

var nouns = []string{
	"bear", "bird", "cat", "deer", "eagle",
	"fish", "fox", "hawk", "lion", "owl",
	"star", "moon", "sun", "wind", "wave",
	"tree", "river", "mountain", "ocean", "cloud",
	"tiger", "wolf", "dragon", "phoenix", "falcon",
	"comet", "galaxy", "planet", "nova", "meteor",
	"forest", "canyon", "valley", "peak", "storm",
}

// GenerateSessionName generates a friendly single-word name from a session ID (UUID).
// Returns a deterministic name like "cat" or "eagle".
//
// Args:
//   - sessionID: UUID string (e.g., "73b5e210-ec1a-4294-96e4-c2aecb2e1063")
//
// Returns:
//   - Friendly name string (e.g., "peak")
func GenerateSessionName(sessionID string) string {
	_ = "STUB: not implemented"
	// Return "unknown" if no session ID
	return ""
}

// Remove dashes and convert to lowercase

// Use first 8 hex chars as seed for word selection

// Fallback for short IDs

// Combine adjectives and nouns into a single pool for more variety

// Convert hex to decimal for array indexing

// GenerateSessionLabel generates a friendly name with session ID prefix.
// Returns a string like "bold 06ddb8f7" for better session identification.
func GenerateSessionLabel(sessionID string) string { _ = "STUB: not implemented"; return "" }

// Extract first 8 chars of UUID (before first dash)

// hexToInt converts hex string to int (takes first 6 characters for safety)
func hexToInt(hex string) int { _ = "STUB: not implemented"; return 0 }

// Return 0 on parse error
