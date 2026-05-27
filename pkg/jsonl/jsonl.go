package jsonl

import (
	"io"
)

// Message represents a Claude Code transcript message
type Message struct {
	ParentUUID        string         `json:"parentUuid"`
	Type              string         `json:"type"`
	Message           MessageContent `json:"message"`
	Timestamp         string         `json:"timestamp"`
	IsApiErrorMessage bool           `json:"isApiErrorMessage,omitempty"`
	Error             string         `json:"error,omitempty"`
}

// MessageContent represents the content of a message
// Content can be either a string (user text messages) or an array (tool results, assistant messages)
type MessageContent struct {
	Role          string    `json:"role"`
	Content       []Content `json:"-"` // Array content (tool_result, assistant messages)
	ContentString string    `json:"-"` // String content (user text messages)
}

// Content represents a content block in a message
type Content struct {
	Type  string                 `json:"type"`
	Name  string                 `json:"name,omitempty"`
	Text  string                 `json:"text,omitempty"`
	Input map[string]interface{} `json:"input,omitempty"`
}

// UnmarshalJSON implements custom JSON unmarshaling for MessageContent
// Handles both string content (user text messages) and array content (tool results, assistant messages)
func (m *MessageContent) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	// Create an alias to avoid recursion
	return nil
}

// Unmarshal everything except content

// Try to unmarshal content as a string (user text messages)

// Try to unmarshal content as an array (tool results, assistant messages)

// Content is neither string nor array (or is null/empty), that's okay

// MarshalJSON implements custom JSON marshaling for MessageContent
// Outputs content as string if ContentString is set, otherwise as array
func (m MessageContent) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// Create auxiliary struct with content as interface{}
	return nil, nil
}

// Choose content format based on which field is set

// ParseFile parses a JSONL file and returns all messages
func ParseFile(path string) ([]Message, error) { _ = "STUB: not implemented"; return nil, nil }

// Parse parses JSONL from a reader and returns all messages.
// Uses bufio.Reader instead of bufio.Scanner to handle arbitrarily long lines
// (e.g. base64-encoded images, large code diffs in Claude Code transcripts).
func Parse(r io.Reader) ([]Message, error) { _ = "STUB: not implemented"; return nil, nil }

// 64KB initial buffer

// Process line even if err != nil (last line may lack trailing newline)

// Skip invalid JSON lines instead of failing

// GetLastApiErrorMessages returns the last N messages with isApiErrorMessage=true
func GetLastApiErrorMessages(messages []Message, count int) []Message {
	_ = "STUB: not implemented"
	return nil
}

// HasRecentApiError checks if there are API error messages after the last user message
func HasRecentApiError(messages []Message) bool { _ = "STUB: not implemented"; return false }

// If no user timestamp, any API error counts

// Check if error is after last user message

// GetLastAssistantMessages returns the last N assistant messages
func GetLastAssistantMessages(messages []Message, count int) []Message {
	_ = "STUB: not implemented"
	return nil
}

// Return last N messages

// ExtractTools extracts all tools from messages with their positions
func ExtractTools(messages []Message) []ToolUse { _ = "STUB: not implemented"; return nil }

// ToolUse represents a tool use with its position
type ToolUse struct {
	Position int
	Name     string
}

// GetLastTool returns the last tool used, or empty string if none
func GetLastTool(tools []ToolUse) string { _ = "STUB: not implemented"; return "" }

// CountToolsAfterPosition counts how many tools were used after a given position
func CountToolsAfterPosition(tools []ToolUse, position int) int {
	_ = "STUB: not implemented"
	return 0
}

// FindToolPosition finds the position of a tool by name (last occurrence)
// Returns -1 if not found
func FindToolPosition(tools []ToolUse, name string) int { _ = "STUB: not implemented"; return 0 }

// ExtractTextFromMessages extracts all text content from messages
func ExtractTextFromMessages(messages []Message) []string { _ = "STUB: not implemented"; return nil }

// FindLastToolUse finds the last occurrence of a specific tool use in messages
// Returns nil if not found
func FindLastToolUse(messages []Message, toolName string) *Content {
	_ = "STUB: not implemented"
	return nil
}

// ExtractToolInput extracts the input parameters from a specific tool use
// Returns empty map if tool not found
func ExtractToolInput(messages []Message, toolName string) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// GetLastUserTimestamp returns the timestamp of the last user message with text content
// Includes both string content (normal user messages) and array content with type="text" (interrupted tool use)
// Excludes tool_result messages
func GetLastUserTimestamp(messages []Message) string { _ = "STUB: not implemented"; return "" }

// Check for string content (normal user text messages)

// Check for array content with type="text" (interrupted tool use: "[Request interrupted by user for tool use]")

// GetLastAssistantTimestamp returns the timestamp of the last assistant message
func GetLastAssistantTimestamp(messages []Message) string { _ = "STUB: not implemented"; return "" }

// FilterMessagesAfterTimestamp filters messages that occurred after given timestamp
// Returns only assistant messages after the timestamp
// This is used to filter messages to only those in the current response (after last user message)
func FilterMessagesAfterTimestamp(messages []Message, afterTimestamp string) []Message {
	_ = "STUB: not implemented"
	return nil

	// No user message - return all assistant messages
}

// Parse the timestamp

// Invalid timestamp - return all assistant messages

// Include only messages AFTER user message

// filterAssistantMessages returns only assistant messages from the list
func filterAssistantMessages(messages []Message) []Message { _ = "STUB: not implemented"; return nil }

// CountToolsByNames counts tools matching any of the given names
func CountToolsByNames(tools []ToolUse, names []string) int { _ = "STUB: not implemented"; return 0 }

// HasAnyActiveTool checks if any active tool was used
func HasAnyActiveTool(tools []ToolUse, activeTools []string) bool {
	_ = "STUB: not implemented"
	return false
}

// ExtractRecentText extracts concatenated text from last N assistant messages
func ExtractRecentText(messages []Message, count int) string { _ = "STUB: not implemented"; return "" }

// Join all texts with spaces
