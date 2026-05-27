package webhook

import (
	"github.com/777genius/claude-notifications/internal/analyzer"
	"github.com/777genius/claude-notifications/internal/config"
)

const (
	discordEmbedAuthorLimit = 256
	discordEmbedFooterLimit = 2048
)

// Formatter renders a SendContext into a preset-specific webhook payload.
//
// Implementations should treat ctx.Message as the pre-joined notification text
// and use the structured fields (RawBody, ActionSummary, SessionName, Folder,
// GitBranch) only when the underlying transport benefits from richer layouts.
type Formatter interface {
	Format(ctx SendContext, statusInfo config.StatusInfo) (interface{}, error)
}

// SlackFormatter formats messages for Slack
type SlackFormatter struct{}

func (f *SlackFormatter) Format(ctx SendContext, statusInfo config.StatusInfo) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DiscordFormatter formats messages for Discord using native embed structure:
// author / title / description / inline fields / footer.
type DiscordFormatter struct{}

func (f *DiscordFormatter) Format(ctx SendContext, statusInfo config.StatusInfo) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fallback for callers that haven't populated structured fields yet.

// buildDiscordAuthor returns the embed author line, e.g.:
//
//	"phoenix 439d1884 · claude-utils (main)"
//	"phoenix 439d1884 · claude-utils"      // no git branch
//	"phoenix 439d1884"                      // no folder either
//
// Returns "" when no session-derived label is available.
func buildDiscordAuthor(ctx SendContext) string { _ = "STUB: not implemented"; return "" }

// buildDiscordFooter returns the embed footer text.
// Uses the raw session UUID so the footer is not redundant with the friendly
// label that already appears in the author line.
func buildDiscordFooter(ctx SendContext) string { _ = "STUB: not implemented"; return "" }

// truncateMiddle keeps both the start and end of a string visible while
// enforcing a hard character limit for Discord embed fields.
func truncateMiddle(s string, limit int) string { _ = "STUB: not implemented"; return "" }

// actionEmojiField maps the leading emoji of an action segment to a field name.
// Keep in sync with summary.buildActionsString.
var actionEmojiField = []struct {
	prefix string
	name   string
}{
	{"📝", "New"},
	{"✏️", "Edited"},
	{"▶", "Commands"},
	{"⏱", "Duration"},
}

// parseActionSummary splits an action summary string (e.g.
// "📝 1 new  ▶ 2 cmds  ⏱ 41s") into Discord embed fields. Unknown segments are
// collected into a single "Details" field so future emoji additions never lose
// information.
func parseActionSummary(s string) []map[string]interface{} { _ = "STUB: not implemented"; return nil }

// Discord rejects fields with empty value (HTTP 400). Skip
// segments that have no payload after their emoji prefix.

// TelegramFormatter formats messages for Telegram with HTML
type TelegramFormatter struct {
	ChatID string
}

func (f *TelegramFormatter) Format(ctx SendContext, statusInfo config.StatusInfo) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getColorForStatus returns color hex code for status (Slack)
func getColorForStatus(status analyzer.Status) string { _ = "STUB: not implemented"; return "" }

// Green

// Teal

// Yellow/Orange

// Blue

// Gray

// getDiscordColorInt returns Discord color integer for status
func getDiscordColorInt(status analyzer.Status) int { _ = "STUB: not implemented"; return 0 }

// Green

// Teal

// Yellow

// Blue

// Gray

// getEmojiForStatus returns emoji for status (Telegram)
func getEmojiForStatus(status analyzer.Status) string { _ = "STUB: not implemented"; return "" }

// LarkFormatter formats messages for Feishu/Lark with interactive cards
type LarkFormatter struct{}

func (f *LarkFormatter) Format(ctx SendContext, statusInfo config.StatusInfo) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getLarkColorTemplate returns Lark color template for status
func getLarkColorTemplate(status analyzer.Status) string { _ = "STUB: not implemented"; return "" }
