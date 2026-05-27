package notifier

// multiplexerHandler describes a terminal multiplexer integration.
type multiplexerHandler struct {
	name      string
	detect    func() bool
	buildArgs func(title, message, bundleID string) ([]string, error)
}

// multiplexerHandlers is the ordered list of supported multiplexers.
// First detected wins.
var multiplexerHandlers = []multiplexerHandler{
	{"tmux", IsTmux, buildTmuxClickArgs},
	{"zellij", IsZellij, buildZellijClickArgs},
	{"wezterm", IsWezTerm, buildWezTermClickArgs},
	{"kitty", IsKitty, buildKittyClickArgs},
}

// detectMultiplexerArgs tries each registered multiplexer.
// Returns (args, name) if detected and target obtained,
// (nil, name) if detected but target failed,
// (nil, "") if no multiplexer detected.
func detectMultiplexerArgs(title, message, bundleID string) ([]string, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// buildTmuxClickArgs captures tmux target and builds notifier args.
// For iTerm2 tmux -CC (control mode), uses the iTerm2 Python API helper
// instead of standard tmux select-window (which doesn't switch iTerm2 tabs).
func buildTmuxClickArgs(title, message, bundleID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fallback to standard tmux select-window

// buildZellijClickArgs captures zellij tab target and builds notifier args.
func buildZellijClickArgs(title, message, bundleID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildWezTermClickArgs captures WezTerm pane target and builds notifier args.
func buildWezTermClickArgs(title, message, bundleID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildKittyClickArgs captures Kitty window target and builds notifier args.
func buildKittyClickArgs(title, message, bundleID string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
