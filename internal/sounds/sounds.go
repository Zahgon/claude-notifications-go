// ABOUTME: Sound discovery package for listing available notification sounds.
// ABOUTME: Pure filesystem scanning with no audio dependencies (CGO-free).

package sounds

// SoundInfo represents a discovered sound file.
type SoundInfo struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Format      string `json:"format"`
	Source      string `json:"source"`
	Description string `json:"description"`
}

// DiscoverOptions controls which sound sources to scan.
type DiscoverOptions struct {
	PluginRoot     string // Root directory of the plugin (for built-in sounds)
	IncludeBuiltIn bool
	IncludeSystem  bool
	MaxSystemDepth int // Max directory depth for Linux system sounds (default 5)
}

// descriptions maps sound names to human-readable descriptions.
var descriptions = map[string]string{
	// Built-in sounds
	"task-complete":   "Triumphant completion chime",
	"review-complete": "Gentle notification tone",
	"question":        "Attention-grabbing sound",
	"plan-ready":      "Professional planning tone",
	"error":           "Error alert sound",
	// macOS system sounds
	"Glass":     "Crisp, clean chime",
	"Hero":      "Triumphant fanfare",
	"Ping":      "Subtle ping sound",
	"Pop":       "Quick pop sound",
	"Purr":      "Gentle purr",
	"Funk":      "Distinctive funk groove",
	"Sosumi":    "Pleasant notification",
	"Basso":     "Deep bass sound",
	"Blow":      "Breeze-like whoosh",
	"Frog":      "Unique ribbit sound",
	"Submarine": "Sonar-like ping",
	"Bottle":    "Cork pop sound",
	"Morse":     "Morse code beeps",
	"Tink":      "Light metallic sound",
}

// Discover scans for available sounds and returns them grouped by source.
// Built-in sounds are listed first, then system sounds.
func Discover(opts DiscoverOptions) []SoundInfo { _ = "STUB: not implemented"; return nil }

// Sort within each source group for stable output

// built-in first, then system

// FindByName searches for a sound by name with 3-level matching:
// 1. Exact match
// 2. Case-insensitive match
// 3. Prefix match (case-insensitive)
// Built-in sounds are prioritized over system sounds at every level.
func FindByName(name string, available []SoundInfo) (SoundInfo, bool) {
	_ = "STUB: not implemented"
	return *new(SoundInfo), false
}

// Level 1: exact match (prefer built-in)

// Level 2: case-insensitive match (prefer built-in)

// Level 3: prefix match (prefer built-in)

// findPreferBuiltIn finds the first match, preferring built-in over system sources.
func findPreferBuiltIn(available []SoundInfo, match func(SoundInfo) bool) (SoundInfo, bool) {
	_ = "STUB: not implemented"
	return *new(SoundInfo), false
}

// findSoundsDirectory locates the plugin's sounds/ directory.
func findSoundsDirectory(pluginRoot string) string { _ = "STUB: not implemented"; return "" }

// Try relative to this source file (development mode)

// internal/sounds/sounds.go -> project root

// Try CLAUDE_PLUGIN_ROOT env

// discoverBuiltIn scans the plugin's sounds/ directory for built-in MP3 files.
func discoverBuiltIn(pluginRoot string) []SoundInfo { _ = "STUB: not implemented"; return nil }

// discoverSystem scans platform-specific system sound directories.
func discoverSystem(maxDepth int) []SoundInfo { _ = "STUB: not implemented"; return nil }

// discoverMacOSSounds finds AIFF files in /System/Library/Sounds/.
func discoverMacOSSounds() []SoundInfo { _ = "STUB: not implemented"; return nil }

// discoverLinuxSounds walks /usr/share/sounds/ for OGG and WAV files.
func discoverLinuxSounds(maxDepth int) []SoundInfo { _ = "STUB: not implemented"; return nil }

// skip errors silently

// Limit depth

// remove leading dot

// discoverWindowsSounds finds WAV files in %SYSTEMROOT%/Media/.
func discoverWindowsSounds() []SoundInfo { _ = "STUB: not implemented"; return nil }
