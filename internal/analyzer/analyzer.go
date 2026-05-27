package analyzer

import (
	"github.com/777genius/claude-notifications/internal/config"
	"github.com/777genius/claude-notifications/pkg/jsonl"
)

// Tool categories for state machine classification
//
// TODO: Future improvement - detect passive Bash commands
// Currently all Bash commands are treated as "active" (code-changing).
// Could be improved by parsing command strings to differentiate:
//   - Passive: ls, cd, pwd, git status, git log, git diff, find, grep
//   - Active: mkdir, rm, mv, cp, git commit, npm install, etc.
//
// This requires:
//  1. Storing tool Input in ToolUse struct (pkg/jsonl)
//  2. Parsing command string from Input["command"]
//  3. Handling complex cases: pipes (|), redirects (>), chains (&&)
//
// Complexity: Medium-High. Edge cases are tricky (e.g. "cat file > output").
var (
	ActiveTools   = []string{"Write", "Edit", "Bash", "NotebookEdit", "SlashCommand", "KillShell"}
	QuestionTools = []string{"AskUserQuestion"}
	PlanningTools = []string{"ExitPlanMode", "TodoWrite"}
	PassiveTools  = []string{"Read", "Grep", "Glob", "WebFetch", "WebSearch", "Search", "Fetch", "Task"}
)

// Status represents the current task status
type Status string

const (
	StatusTaskComplete        Status = "task_complete"
	StatusReviewComplete      Status = "review_complete"
	StatusQuestion            Status = "question"
	StatusPlanReady           Status = "plan_ready"
	StatusSessionLimitReached Status = "session_limit_reached"
	StatusAPIError            Status = "api_error"
	StatusAPIErrorOverloaded  Status = "api_error_overloaded"
	StatusUnknown             Status = "unknown"
)

// AnalyzeTranscript analyzes a transcript file and determines the current status
func AnalyzeTranscript(transcriptPath string, cfg *config.Config) (Status, error) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}

// AnalyzeTranscriptWithMessages analyzes a transcript and also returns the parsed messages.
// This allows callers to reuse the messages (e.g., for summary generation) without re-reading the file.
func AnalyzeTranscriptWithMessages(transcriptPath string, cfg *config.Config) (Status, []jsonl.Message, error) {
	_ = "STUB: not implemented"
	// Parse JSONL file
	return *new(Status), nil, nil
}

// PRIORITY CHECK 1: Session limit reached
// This takes precedence over all other status detection

// PRIORITY CHECK 2: API errors (uses isApiErrorMessage flag from JSONL)

// Find last user message timestamp
// This ensures we only analyze tools from the CURRENT response,
// not from previous user requests (avoids "ghost" ExitPlanMode problem)

// Filter assistant messages AFTER last user message

// Take last 15 messages (temporal window) from filtered set

// Extract tools with positions

// STATE MACHINE LOGIC - tool-based detection only

// 1. If we have tools, analyze them

// 1a. Last tool is ExitPlanMode → plan just created

// 1b. Last tool is AskUserQuestion → waiting for user

// 1c. ExitPlanMode exists AND tools after it → plan executed

// 1d. Review detection: only read-like tools + long text response
// Read-like tools: Read, Grep, Glob (searching/analyzing code)
// No active tools: no Write, Edit, Bash, etc.
// Long text: >200 chars (indicates substantial analysis/review)

// Extract recent text to check length

// 1e. Last tool is active (Write/Edit/Bash) → work completed

// 1f. Any tool usage at all → likely task completed
// (matches bash version: toolCount >= 1 → task_complete)

// 2. No tools found
// If notifyOnTextResponse is enabled (default: true), treat as task_complete
// This handles cases like extended thinking where Claude responds with text only

// contains checks if a slice contains a string
func contains(slice []string, str string) bool { _ = "STUB: not implemented"; return false }

// GetStatusForPreToolUse determines status for PreToolUse hook
// This is called BEFORE tool execution, so we only have the tool name
func GetStatusForPreToolUse(toolName string) Status { _ = "STUB: not implemented"; return *new(Status) }

// detectSessionLimitReached checks if the last assistant messages contain "Session limit reached"
func detectSessionLimitReached(messages []jsonl.Message) bool {
	_ = "STUB: not implemented"
	// Check last 3 assistant messages for the session limit text
	return false
}

// Extract text from recent messages

// Check each text for the session limit phrase

// detectAPIErrors checks for API errors using the isApiErrorMessage flag in JSONL.
// Claude Code sets isApiErrorMessage=true on synthetic assistant messages
// when the API returns an error (400, 401, 429, 500, 529, etc).
// Returns the specific error status or StatusUnknown if no API error found.
func detectAPIErrors(messages []jsonl.Message) Status {
	_ = "STUB: not implemented"
	// Check if there are any recent API error messages (after last user message)
	return *new(Status)
}

// Get the last API error messages to determine the type

// Check the last error message to determine the type

// Primary check: use the structured "error" field from JSONL
// Real values: "authentication_failed" for 401, "unknown" for others

// Fallback: check text content for 401 indicators
// (in case Claude Code changes the error field format)

// Any other API error (400, 429, 500, 529, connection error, etc.)

// containsIgnoreCase checks if string contains substring (case insensitive)
func containsIgnoreCase(s, substr string) bool { _ = "STUB: not implemented"; return false }
