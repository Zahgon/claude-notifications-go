package webhook

import (
	"regexp"
	"time"

	"github.com/777genius/claude-notifications/internal/analyzer"
	"github.com/777genius/claude-notifications/internal/config"
	"github.com/777genius/claude-notifications/internal/platform"
)

var templatePattern = regexp.MustCompile(`\$\{\{\s*([^{}]+?)\s*\}\}`)

// SendContext carries per-notification metadata used by webhook templates.
//
// Message remains the pre-joined "[session|branch folder] body actions" string
// so existing ${{message}} templates and downstream consumers keep working. The
// structured fields below let formatters that render rich layouts (e.g. Discord
// embeds) avoid re-parsing the joined output.
type SendContext struct {
	Status    analyzer.Status
	Message   string
	SessionID string
	CWD       string

	SessionName   string // friendly session label (e.g. "phoenix 439d1884")
	GitBranch     string // empty when CWD is not a git working tree
	Folder        string // filepath.Base(CWD); empty when CWD is empty
	RawBody       string // summary body without prefix/actions
	ActionSummary string // action segment only (e.g. "📝 1 new  ▶ 2 cmds  ⏱ 41s")
}

type runtimeContext struct {
	sendCtx    SendContext
	statusInfo config.StatusInfo
	now        time.Time

	gitLoaded bool
	gitMeta   platform.GitMetadata
}

func newRuntimeContext(sendCtx SendContext, statusInfo config.StatusInfo) *runtimeContext {
	_ = "STUB: not implemented"
	return nil
}

func (c *runtimeContext) resolveHeaders(headers map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *runtimeContext) resolvePayloadFields(fields map[string]interface{}) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *runtimeContext) resolveValue(path string, value interface{}) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (c *runtimeContext) resolveString(input string) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (c *runtimeContext) lookupTemplateValue(token string) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (c *runtimeContext) lookupGitTemplateValue(token string) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (c *runtimeContext) gitMetadata() platform.GitMetadata {
	_ = "STUB: not implemented"
	return *new(platform.GitMetadata)
}

func stringifyTemplateValue(value interface{}) string { _ = "STUB: not implemented"; return "" }

func mergePayloadMaps(base, overrides map[string]interface{}) { _ = "STUB: not implemented"; return }
