package config

// Config represents the plugin configuration
type Config struct {
	Notifications NotificationsConfig   `json:"notifications"`
	Statuses      map[string]StatusInfo `json:"statuses"`
	Debug         DebugConfig           `json:"debug,omitempty"`
}

// DebugConfig represents debug/diagnostic settings
type DebugConfig struct {
	Benchmark bool `json:"benchmark"` // Enable benchmark timing output to log file
}

// NotificationsConfig represents notification settings
type NotificationsConfig struct {
	Desktop                                     DesktopConfig    `json:"desktop"`
	Webhook                                     WebhookConfig    `json:"webhook"`
	SuppressQuestionAfterTaskCompleteSeconds    *int             `json:"suppressQuestionAfterTaskCompleteSeconds"`
	SuppressQuestionAfterAnyNotificationSeconds *int             `json:"suppressQuestionAfterAnyNotificationSeconds"`
	NotifyOnSubagentStop                        bool             `json:"notifyOnSubagentStop"`      // Send notifications when subagents (Task tool) complete, default: false
	SuppressForSubagents                        *bool            `json:"suppressForSubagents"`      // Suppress notifications when transcript_path contains /subagents/, default: true
	NotifyOnTextResponse                        *bool            `json:"notifyOnTextResponse"`      // Send notifications for text-only responses (no tools), default: true
	RespectJudgeMode                            *bool            `json:"respectJudgeMode"`          // Honor CLAUDE_HOOK_JUDGE_MODE=true env var to suppress notifications, default: true
	SuppressFilters                             []SuppressFilter `json:"suppressFilters,omitempty"` // Rules for suppressing notifications by status/branch/folder
	TeamMode                                    string           `json:"teamMode,omitempty"`        // Team mode: "always" (no suppression, default), "wait-all" (suppress lead, notify when all idle), "never" (silent in team mode)
}

// DesktopConfig represents desktop notification settings
type DesktopConfig struct {
	Enabled          bool    `json:"enabled"`
	Sound            bool    `json:"sound"`
	TerminalBell     *bool   `json:"terminalBell"`     // Send BEL to /dev/tty for terminal tab indicators (default: true)
	Volume           float64 `json:"volume"`           // Volume level 0.0-1.0, default 1.0 (full volume)
	AudioDevice      string  `json:"audioDevice"`      // Audio output device name (empty = system default)
	AppIcon          string  `json:"appIcon"`          // Path to app icon
	ClickToFocus     bool    `json:"clickToFocus"`     // macOS: activate terminal on notification click (default: true)
	TerminalBundleID string  `json:"terminalBundleId"` // macOS: override auto-detected terminal bundle ID (empty = auto)
}

// WebhookConfig represents webhook settings
type WebhookConfig struct {
	Enabled        bool                   `json:"enabled"`
	Preset         string                 `json:"preset"`
	URL            string                 `json:"url"`
	ChatID         string                 `json:"chat_id"`
	Format         string                 `json:"format"`
	Headers        map[string]string      `json:"headers"`
	PayloadFields  map[string]interface{} `json:"payloadFields,omitempty"`
	Retry          RetryConfig            `json:"retry"`
	CircuitBreaker CircuitBreakerConfig   `json:"circuitBreaker"`
	RateLimit      RateLimitConfig        `json:"rateLimit"`
}

// RetryConfig represents retry settings
type RetryConfig struct {
	Enabled        bool   `json:"enabled"`
	MaxAttempts    int    `json:"maxAttempts"`
	InitialBackoff string `json:"initialBackoff"` // e.g. "1s"
	MaxBackoff     string `json:"maxBackoff"`     // e.g. "10s"
}

// CircuitBreakerConfig represents circuit breaker settings
type CircuitBreakerConfig struct {
	Enabled          bool   `json:"enabled"`
	FailureThreshold int    `json:"failureThreshold"` // failures before opening
	Timeout          string `json:"timeout"`          // time to wait in open state, e.g. "30s"
	SuccessThreshold int    `json:"successThreshold"` // successes needed in half-open
}

// RateLimitConfig represents rate limiting settings
type RateLimitConfig struct {
	Enabled           bool `json:"enabled"`
	RequestsPerMinute int  `json:"requestsPerMinute"`
}

// StatusChannelConfig represents per-channel status overrides.
type StatusChannelConfig struct {
	Enabled *bool `json:"enabled,omitempty"` // nil = inherit default enabled behavior
}

// StatusInfo represents configuration for a specific status
type StatusInfo struct {
	Enabled *bool                `json:"enabled,omitempty"` // nil = true (default for backward compatibility)
	Desktop *StatusChannelConfig `json:"desktop,omitempty"`
	Webhook *StatusChannelConfig `json:"webhook,omitempty"`
	Title   string               `json:"title"`
	Sound   string               `json:"sound"`
}

// SuppressFilter defines conditions for suppressing notifications.
// All specified (non-nil) fields must match for the filter to suppress.
// Omitted fields act as wildcards (match any value).
type SuppressFilter struct {
	Name      string  `json:"name,omitempty"`
	Status    *string `json:"status,omitempty"`
	GitBranch *string `json:"gitBranch"` // no omitempty — nil means "any", "" means "no branch"
	Folder    *string `json:"folder,omitempty"`
}

// Matches returns true if all specified fields match the given values.
func (f *SuppressFilter) Matches(status, gitBranch, folder string) bool {
	_ = "STUB: not implemented"
	return false
}

// HasConditions returns true if the filter has at least one condition field set.
func (f *SuppressFilter) HasConditions() bool { _ = "STUB: not implemented"; return false }

// intPtr returns a pointer to the given int value
func intPtr(v int) *int {
	_ = "STUB: not implemented"

	// stringPtr returns a pointer to the given string value
	return nil
}

func stringPtr(v string) *string { _ = "STUB: not implemented"; return nil }

const defaultSuppressQuestionAfterAnyNotificationSeconds = 7

// DefaultConfig returns a config with sensible defaults
func DefaultConfig() *Config {
	_ = "STUB: not implemented"
	// Get plugin root from environment, fallback to current directory
	return nil
}

// Full volume by default

// macOS: activate terminal on click (default: enabled)
// TerminalBundleID: "" - empty means auto-detect

// Load loads configuration from a file
// If the file doesn't exist, returns default config
func Load(path string) (*Config, error) {
	_ = "STUB: not implemented"
	// If path doesn't exist, use default config
	return nil, nil
}

// Expand environment variables in paths

// Expand environment variables in sound paths

// Apply defaults for missing fields

// GetStableConfigDir returns the stable config directory outside the plugin cache.
// This directory survives plugin updates (bootstrap.sh rm -rf of cache).
func GetStableConfigDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetStableConfigPath returns the stable config file path outside the plugin cache.
func GetStableConfigPath() (string, error) { _ = "STUB: not implemented"; return "", nil }

// LoadFromPluginRoot loads configuration with a resilient fallback chain:
// 1. Stable path (~/.claude/claude-notifications-go/config.json) — preferred
// 2. Old path (pluginRoot/config/config.json) — fallback, auto-migrates to stable
// 3. Default config — if neither path has valid config
//
// Corrupted config files are non-fatal: a warning is printed to stderr and
// logged, then the next source in the chain is tried.
func LoadFromPluginRoot(pluginRoot string) (*Config, error) {
	_ = "STUB: not implemented"
	// 1. Try stable path
	return nil, nil
}

// Corrupted stable config — warn and fall through to old path

// 2. Try old path (pluginRoot/config/config.json)

// Corrupted old config — warn, return defaults (non-fatal)

// Migrate to stable path (best-effort)

// 3. Neither path has config — return defaults

// migrateConfig copies config from oldPath to stablePath atomically.
// Uses temp file + rename in the same directory for safe atomic write.
func migrateConfig(oldPath, stablePath string) error { _ = "STUB: not implemented"; return nil }

// Create temp file in same dir — guarantees same filesystem for safe os.Rename

// cleanup on any error path

// ApplyDefaults fills in missing fields with default values
func (c *Config) ApplyDefaults() {
	_ = "STUB: not implemented"
	// Desktop defaults
	return
}

// Default to full volume

// AppIcon: Keep empty if not set (no default)

// Webhook defaults

// Cooldown defaults (nil = not set in config, apply defaults)

// Status defaults

// Fill in missing statuses

// Validate validates the configuration
func (c *Config) Validate() error {
	_ = "STUB: not implemented"
	// Validate volume
	return nil
}

// Validate webhook preset (only if webhooks are enabled)

// Validate webhook format (only if webhooks are enabled)

// Validate webhook URL if enabled

// Validate Telegram chat_id if Telegram preset is used

// Validate cooldowns (both fields, if explicitly set)

// Validate teamMode

// Validate suppress-filters

// GetStatusInfo returns status information for a given status
func (c *Config) GetStatusInfo(status string) (StatusInfo, bool) {
	_ = "STUB: not implemented"
	return *new(StatusInfo), false
}

// IsDesktopEnabled returns true if desktop notifications are enabled
func (c *Config) IsDesktopEnabled() bool { _ = "STUB: not implemented"; return false }

// IsWebhookEnabled returns true if webhook notifications are enabled
func (c *Config) IsWebhookEnabled() bool { _ = "STUB: not implemented"; return false }

// IsAnyNotificationEnabled returns true if at least one notification method is enabled
func (c *Config) IsAnyNotificationEnabled() bool { _ = "STUB: not implemented"; return false }

// GetSuppressQuestionAfterTaskCompleteSeconds returns the cooldown in seconds
// after task completion before question notifications are allowed (default: 12)
func (c *Config) GetSuppressQuestionAfterTaskCompleteSeconds() int {
	_ = "STUB: not implemented"
	return 0
}

// GetSuppressQuestionAfterAnyNotificationSeconds returns the cooldown in seconds
// after any notification before question notifications are allowed (default: 7).
func (c *Config) GetSuppressQuestionAfterAnyNotificationSeconds() int {
	_ = "STUB: not implemented"
	return 0
}

// ShouldNotifyOnTextResponse returns true if notifications should be sent for text-only responses (default: true)
func (c *Config) ShouldNotifyOnTextResponse() bool { _ = "STUB: not implemented"; return false }

// Default: notify on text responses

// ShouldSuppressForSubagents returns true if notifications should be suppressed
// when transcript_path contains /subagents/ (default: true)
func (c *Config) ShouldSuppressForSubagents() bool { _ = "STUB: not implemented"; return false }

// Default: suppress subagent notifications

// IsBenchmarkEnabled returns true if benchmark timing is enabled via config
func (c *Config) IsBenchmarkEnabled() bool { _ = "STUB: not implemented"; return false }

// ShouldRespectJudgeMode returns true if CLAUDE_HOOK_JUDGE_MODE=true env var should suppress notifications (default: true)
func (c *Config) ShouldRespectJudgeMode() bool { _ = "STUB: not implemented"; return false }

// Default: respect judge mode

// GetTeamMode returns the team notification mode: "always" (default), "wait-all", or "never"
func (c *Config) GetTeamMode() string { _ = "STUB: not implemented"; return "" }

// IsStatusEnabled returns true if notifications for this status are enabled
// Returns true by default (if Enabled is nil or not specified) for backward compatibility
func (c *Config) IsStatusEnabled(status string) bool { _ = "STUB: not implemented"; return false }

// unknown statuses are enabled by default

// nil means enabled (backward compatibility)

func isStatusChannelEnabled(channel *StatusChannelConfig) bool {
	_ = "STUB: not implemented"
	return false
}

// IsTerminalBellEnabled returns true if terminal bell (BEL) should be sent (default: true)
func (c *Config) IsTerminalBellEnabled() bool { _ = "STUB: not implemented"; return false }

// Default: enabled

// IsStatusDesktopEnabled returns true if desktop notifications for this status are enabled
// Considers global desktop.enabled, per-status enabled, and per-channel desktop override.
func (c *Config) IsStatusDesktopEnabled(status string) bool {
	_ = "STUB: not implemented"
	return false
}

// IsStatusWebhookEnabled returns true if webhook notifications for this status are enabled
// Considers global webhook.enabled, per-status enabled, and per-channel webhook override.
func (c *Config) IsStatusWebhookEnabled(status string) bool {
	_ = "STUB: not implemented"
	return false
}

// ShouldFilter returns true if any suppress-filter rule matches the given context.
// When true, the notification should be suppressed entirely (both desktop and webhook).
func (c *Config) ShouldFilter(status, gitBranch, folder string) bool {
	_ = "STUB: not implemented"
	return false
}
