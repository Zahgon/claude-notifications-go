package hooks

import (
	"io"
	"time"

	"github.com/777genius/claude-notifications/internal/analyzer"
	"github.com/777genius/claude-notifications/internal/config"
	"github.com/777genius/claude-notifications/internal/dedup"
	"github.com/777genius/claude-notifications/internal/state"
	"github.com/777genius/claude-notifications/internal/teamstate"
	"github.com/777genius/claude-notifications/internal/webhook"
	"github.com/777genius/claude-notifications/pkg/jsonl"
)

// HookData represents the data received from Claude Code hooks
type HookData struct {
	TranscriptPath string `json:"transcript_path"`
	SessionID      string `json:"session_id"`
	CWD            string `json:"cwd"`
	ToolName       string `json:"tool_name,omitempty"`
	HookEventName  string `json:"hook_event_name,omitempty"`
	// Team-related fields (present in TeammateIdle, TaskCreated, TaskCompleted hooks)
	TeamName     string `json:"team_name,omitempty"`
	TeammateName string `json:"teammate_name,omitempty"`
}

// notifierInterface defines the interface for sending desktop notifications
type notifierInterface interface {
	SendDesktop(status analyzer.Status, message, sessionID, cwd string) error
	Close() error
}

// webhookInterface defines the interface for sending webhook notifications
type webhookInterface interface {
	SendAsyncWithContext(sendCtx webhook.SendContext)
	Shutdown(timeout time.Duration) error
}

// Handler handles hook events
type Handler struct {
	cfg          *config.Config
	dedupMgr     *dedup.Manager
	stateMgr     *state.Manager
	teamStateMgr *teamstate.Manager
	notifierSvc  notifierInterface
	webhookSvc   webhookInterface
	pluginRoot   string
}

// NewHandler creates a new hook handler
func NewHandler(pluginRoot string) (*Handler, error) {
	_ = "STUB: not implemented"
	// Load config
	return nil, nil
}

// Validate config

// HandleHook handles a hook event
func (h *Handler) HandleHook(hookEvent string, input io.Reader) error {
	_ = "STUB: not implemented"
	// Benchmark instrumentation (enabled via config debug.benchmark)
	return nil
}

// Add panic recovery for robustness

// Skip notifications when running in background judge mode (e.g., double-shot-latte plugin)
// The CLAUDE_HOOK_JUDGE_MODE env var is set by plugins that spawn background Claude instances
// to evaluate context/decide on continuation - we don't want notifications from these
// Can be disabled via config: "respectJudgeMode": false

// Ensure notifier resources are cleaned up when function exits

// Ensure webhook sender waits for in-flight requests before exit

// Parse hook data

// Validate session ID

// Phase 1: Early duplicate check (per hook event type)

// Check if any notification method is enabled

// Determine status based on hook type

// reused by generateMessage to avoid double I/O

// Check session state first (60s TTL) to suppress duplicates after PreToolUse

// Check if this is a subagent transcript and should be suppressed

// Team mode: check if this session is a team lead and suppress if needed

// Record that the lead has stopped

// Check if all teammates are already idle

// Not all teammates idle yet — suppress notification, wait for TeammateIdle

// All teammates are idle — proceed with notification and mark as notified

// teamMode "always" or not a team lead: fall through to normal processing

// Analyze the transcript to determine status

// Note: We don't delete session state here to preserve cooldown info
// State files have TTL and will be cleaned up automatically

// Check config: should we suppress subagent notifications?
// First check path-based suppression (covers subagents and teammates)

// Then check the legacy notifyOnSubagentStop flag

// If enabled, handle like Stop

// If status is unknown, skip

// Check suppress-filters before any state mutations (dedup lock, cooldowns)

// Phase 2: Acquire lock before sending (per hook event type)

// Note: Lock is NOT released - it ages out naturally after 2s to prevent rapid duplicates

// Check cooldown for question status BEFORE updating notification time

// Load state to log its contents

// First, check if we should suppress question after ANY notification (not just task_complete)

// Lock will be released by defer

// Also check legacy cooldown after task_complete

// Lock will be released by defer

// Update state (only for task_complete, PreToolUse already updated state)

// Generate message

// Acquire content lock to prevent race between different hooks (Stop vs Notification)
// This ensures only one process can check and update duplicate state at a time

// Error (not "lock busy") - continue without lock as fallback

// Lock is held by another process - it's already handling this notification

// Release lock on exit if acquired

// Check for duplicate message content (3 minutes = 180 seconds window)

// Update last notification time and message

// Send notifications

// handlePreToolUse handles PreToolUse hook
func (h *Handler) handlePreToolUse(hookData *HookData) analyzer.Status {
	_ = "STUB: not implemented"
	return *new(analyzer.Status)
}

// Write session state BEFORE returning (prevents race with Notification hook)
// This matches bash version behavior: state is written BEFORE notification is sent

// handleNotificationEvent handles Notification hook
// Always returns StatusQuestion as per design: Notification hook is triggered
// when Claude needs user input (e.g., permission dialogs, questions)
func (h *Handler) handleNotificationEvent(hookData *HookData) (analyzer.Status, error) {
	_ = "STUB: not implemented"
	return *new(analyzer.Status), nil
}

// handleTeammateIdle handles the TeammateIdle hook event.
// Records the teammate as idle, checks if all teammates are idle + lead stopped,
// and sends a notification when both conditions are met.
func (h *Handler) handleTeammateIdle(hookData *HookData) error {
	_ = "STUB: not implemented"
	return nil
}

// Dedup: prevent rapid duplicate TeammateIdle events for the same teammate

// Get team info to know all expected members

// Record this teammate as idle

// Check if all conditions are met: lead stopped + all teammates idle

// All conditions met — send notification

func skipUTF8BOM(input io.Reader) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// handleStopEvent handles Stop/SubagentStop hooks.
// Returns the parsed messages alongside the status so callers can reuse them
// (e.g., for summary generation) without re-reading the transcript file.
func (h *Handler) handleStopEvent(hookData *HookData) (analyzer.Status, []jsonl.Message, error) {
	_ = "STUB: not implemented"
	return *new(analyzer.Status), nil, nil
}

// generateMessage generates a notification body and action summary.
// If messages are provided (from handleStopEvent), uses them directly to avoid re-reading the transcript.
func (h *Handler) generateMessage(hookData *HookData, status analyzer.Status, messages []jsonl.Message) (body, actions string) {
	_ = "STUB: not implemented"
	// Use pre-parsed messages if available (eliminates ~234ms double I/O)
	return "", ""
}

// Fallback: read transcript from file (for non-Stop hooks)

// joinMessageParts mirrors summary.appendActions: joins body and actions with a
// single space when actions is non-empty.
func joinMessageParts(body, actions string) string { _ = "STUB: not implemented"; return "" }

// sendNotifications sends desktop and webhook notifications.
//
// body is the summary text (no metadata prefix, no action segments).
// actions is the formatted action summary (e.g. "📝 1 new  ▶ 2 cmds  ⏱ 41s") or "".
func (h *Handler) sendNotifications(status analyzer.Status, body, actions, sessionID, cwd string) {
	_ = "STUB: not implemented"
	// Add panic recovery to prevent notification failures from crashing the plugin
	return
}

// Format: "[sessionname|branch folder] message" or "[sessionname folder] message"

// Send desktop notification (check per-status enabled)

// Send webhook notification (async, check per-status enabled)

// isSubagentTranscript checks if the transcript path indicates a subagent session.
// Claude Code stores subagent transcripts in paths containing /subagents/ segment.
func isSubagentTranscript(transcriptPath string) bool {
	_ = "STUB: not implemented"
	// Normalize path separators for cross-platform compatibility
	return false
}

// cleanupOldLocks cleans up old lock and state files but preserves session state for cooldown
func (h *Handler) cleanupOldLocks() {
	_ = "STUB: not implemented"
	// Cleanup old locks (older than 60 seconds)
	return
}

// Cleanup old state files (older than 60 seconds)

func (h *Handler) maybeEmitDesktopPermissionGuidance(err error) { _ = "STUB: not implemented"; return }

func (h *Handler) shouldEmitPermissionGuidance() bool { _ = "STUB: not implemented"; return false }
