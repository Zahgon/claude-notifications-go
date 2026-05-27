package summary

import (
	"regexp"
	"time"

	"github.com/777genius/claude-notifications/internal/analyzer"
	"github.com/777genius/claude-notifications/internal/config"
	"github.com/777genius/claude-notifications/pkg/jsonl"
)

const (
	// Message window sizes for different notification types
	// These determine how many recent assistant messages to analyze
	QuestionMessagesWindow = 8 // Based on bash version, good balance for question detection
	ReviewMessagesWindow   = 5 // Smaller window for focused review summaries
	TaskMessagesWindow     = 5 // Smaller window for task completion summaries
)

var (
	// Regex patterns for markdown cleanup
	headerPattern     = regexp.MustCompile(`^#+\s*`)
	bulletPattern     = regexp.MustCompile(`^[-*•]\s*`)
	backtickPattern   = regexp.MustCompile("`")
	multiSpacePattern = regexp.MustCompile(`\s+`)
	emojiPattern      = regexp.MustCompile(`^[\p{So}\p{Sk}]+\s*`)

	// Extended markdown patterns for full cleanup
	codeBlockPattern     = regexp.MustCompile("```[\\s\\S]*?```")        // Code blocks
	linkPattern          = regexp.MustCompile(`\[([^\]]+)\]\([^\)]+\)`)  // [text](url) -> text
	imagePattern         = regexp.MustCompile(`!\[([^\]]*)\]\([^\)]+\)`) // ![alt](url) -> alt
	boldPattern          = regexp.MustCompile(`(\*\*|__)(.+?)(\*\*|__)`) // **text** or __text__
	italicPattern        = regexp.MustCompile(`(\*|_)([^*_]+)(\*|_)`)    // *text* or _text_
	strikethroughPattern = regexp.MustCompile(`~~(.+?)~~`)               // ~~text~~
	blockquotePattern    = regexp.MustCompile(`^>\s*`)                   // > quote
)

// getRecentAssistantMessages safely extracts recent assistant messages from current response
// Filters by last user timestamp to ensure we only get messages from the CURRENT response,
// not from previous user requests. Falls back to last N messages if filtering fails.
func getRecentAssistantMessages(messages []jsonl.Message, limit int) []jsonl.Message {
	_ = "STUB: not implemented"
	// Filter by user timestamp (current response only)
	return nil
}

// If filtered result is not empty, use it (limited to window size)

// Fallback: last N messages (for backward compatibility and edge cases)

// GenerateFromMessages generates a status-specific summary from already-parsed messages.
// This avoids re-reading the transcript file when messages are already available.
//
// Returns the joined "<body> <actions>" string for backward compatibility. New
// callers that want body and actions separately should use GenerateFromMessagesStructured.
func GenerateFromMessages(messages []jsonl.Message, status analyzer.Status, cfg *config.Config) string {
	_ = "STUB: not implemented"
	return ""
}

// GenerateFromMessagesStructured returns the summary body and the action summary
// (e.g. "📝 1 new  ▶ 2 cmds  ⏱ 41s") as separate strings. Either may be empty.
//
// Webhook formatters that render structured layouts (Discord embed fields) use
// this to avoid re-parsing the joined output.
func GenerateFromMessagesStructured(messages []jsonl.Message, status analyzer.Status, cfg *config.Config) (body, actions string) {
	_ = "STUB: not implemented"
	return "", ""
}

// GenerateFromTranscript generates a status-specific summary from transcript
func GenerateFromTranscript(transcriptPath string, status analyzer.Status, cfg *config.Config) string {
	_ = "STUB: not implemented"
	return ""
}

// generateQuestionBody generates the body text for question status (no actions appended).
// Improved logic: extracts meaningful question text with markdown cleanup.
func generateQuestionBody(messages []jsonl.Message) string {
	_ = "STUB: not implemented"
	// 1) Try to extract AskUserQuestion tool (with recency check)
	return ""
}

// 2) Get recent messages from current response using helper

// Strategy A: Find texts with "?" and prioritize short ones

// If we found questions, pick the shortest one (likely most direct)

// Strategy B: No "?" found, take first sentence from last assistant message

// 3) Final fallback: generic prompt

// generatePlanBody generates the body text for plan_ready status (no actions appended).
// Matches bash: lib/summarizer.sh lines 471-492.
func generatePlanBody(messages []jsonl.Message) string { _ = "STUB: not implemented"; return "" }

// Get first non-empty line, clean markdown

// generateReviewBody generates the body text for review_complete status (no actions appended).
// Matches bash: lib/summarizer.sh lines 494-521.
func generateReviewBody(messages []jsonl.Message) string { _ = "STUB: not implemented"; return "" }

// Count Read tool usage

// generateTaskBody generates the body text for task_complete status (no actions appended).
// Matches bash: lib/summarizer.sh lines 523-653.
func generateTaskBody(messages []jsonl.Message, cfg *config.Config) string {
	_ = "STUB: not implemented"
	return ""
}

// generateAPIErrorBody returns the body text for api_error (401 authentication) status.
func generateAPIErrorBody() string { _ = "STUB: not implemented"; return "" }

// generateAPIErrorOverloadedBody returns the body text for api_error_overloaded status.
// Extracts the actual error text from the API error message.
func generateAPIErrorOverloadedBody(messages []jsonl.Message) string {
	_ = "STUB: not implemented"
	return ""
}

// extractAskUserQuestion extracts the last AskUserQuestion with recency check
// Returns (question, isRecent)
func extractAskUserQuestion(messages []jsonl.Message) (string, bool) {
	_ = "STUB: not implemented"
	// Find last AskUserQuestion tool
	return "", false
}

// Extract question from input.questions[0].question

// Check recency (60s window)

// Check if question is within 60s of last assistant message

// extractExitPlanModePlan extracts the plan text from ExitPlanMode tool
func extractExitPlanModePlan(messages []jsonl.Message) string { _ = "STUB: not implemented"; return "" }

// calculateDuration calculates duration between last user and last assistant messages
func calculateDuration(messages []jsonl.Message) string { _ = "STUB: not implemented"; return "" }

// formatDuration formats duration into human-readable string
func formatDuration(d time.Duration) string { _ = "STUB: not implemented"; return "" }

// countToolsByType counts tools since last user message
func countToolsByType(messages []jsonl.Message) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

// Find last user timestamp

// Count tools after user message

// Check if this message is after user message

// getActionsString calculates duration, counts tools, and returns formatted actions string
func getActionsString(messages []jsonl.Message) string { _ = "STUB: not implemented"; return "" }

// appendActions appends actions suffix to message if non-empty
func appendActions(message, actions string) string { _ = "STUB: not implemented"; return "" }

// buildActionsString builds actions summary with tool counts and duration
func buildActionsString(toolCounts map[string]int, duration string) string {
	_ = "STUB: not implemented"

	// Write
	return ""
}

// Edit

// Bash

// Add duration at the end

// Helper functions

func extractFirstSentence(text string) string {
	_ = "STUB: not implemented"
	// Find first sentence (ending with . ! or ?)
	// If first sentence is too short (< 20 chars), try to include second sentence too
	return ""
}

// For dots, check if this is really end of sentence:
// - Must be followed by space + uppercase letter, or end of string
// - Should not be preceded by a digit (to avoid "v1.6.0")

// Check if preceded by digit (version numbers like v1.6.0)

// Check if followed by digit (decimal numbers like 1.5)

// Check if followed by letter without space (abbreviations, domains)

// Include punctuation in the sentence

// Calculate total length so far

// If we have at least one sentence and either:
// 1. Total length >= minSentenceLength, OR
// 2. Total length >= maxLength
// Then return what we have

// First sentence too short, continue to get second

// Too long, return what we had before last sentence

// Good length, return

// No sentence ending found

// Return first 100 chars if no punctuation found

func truncateText(text string, maxLen int) string { _ = "STUB: not implemented"; return "" }

// Step 1: Try to find sentence boundary (., !, ?) within maxLen
// Look for the last sentence-ending punctuation in the allowed range
// Use runes to avoid cutting in the middle of a multi-byte character

// Check for sentence enders: ". ", "! ", "? " (followed by space or newline)
// Also check for end of string within maxLen

// Try sentence endings with space/newline after
// Find the FIRST suitable sentence ending (to avoid partial next sentences)

// Check if this position is suitable: not too early

// Found a suitable sentence ending
// Only use it if we haven't found one yet, or this is a better one
// (we want the FIRST suitable one, not the last)

// Stop searching for this ender

// Found a suitable sentence, no need to check other enders

// Also try sentence ending at the very end of searchText (no space after)

// Found a sentence boundary, truncate there (including the punctuation)

// Step 2: No sentence boundary found, try word boundary
// Still use runes to be safe

// CleanMarkdown cleans markdown formatting from text
// Removes all markdown syntax while preserving the actual text content
func CleanMarkdown(text string) string {
	_ = "STUB: not implemented"
	// Step 1: Remove code blocks first (they can contain markdown-like syntax)
	return ""
}

// Step 2: Convert images to alt text (must be before links since images are ![](url))

// Step 3: Convert links to text only

// Step 4: Remove strikethrough

// Step 5: Remove bold (both ** and __)

// Step 6: Remove italic (both * and _)
// Need to be careful with edge cases

// Step 7: Remove backticks (inline code)

// Step 8: Process line by line for line-based patterns

// Remove headers (# text)

// Remove blockquotes (> text)

// Remove bullet points (- text, * text, • text)

// Trim again

// Step 9: Join lines and normalize whitespace

// GetDefaultMessage returns a default message for a status
func GetDefaultMessage(status analyzer.Status, cfg *config.Config) string {
	_ = "STUB: not implemented"
	return ""
}

// Remove emoji from title for message

// GenerateSimple generates a simple message based on status
func GenerateSimple(status analyzer.Status, cfg *config.Config) string {
	_ = "STUB: not implemented"
	return ""
}
