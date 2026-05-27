package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/777genius/claude-notifications/internal/errorhandler"
	"github.com/777genius/claude-notifications/internal/notifier"
)

const version = "1.39.1"
const windowsLazyUpdateRetryAfter = time.Hour

var (
	currentGOOS               = runtime.GOOS
	scheduleWindowsLazyUpdate = scheduleWindowsLazyUpdateImpl
)

func main() {
	// Initialize global error handler with panic recovery
	// logToConsole=true: errors will be shown in console
	// exitOnCritical=false: don't exit on critical errors (let caller decide)
	// recoveryEnabled=true: recover from panics
	errorhandler.Init(true, false, true)

	// Add global panic recovery
	defer errorhandler.HandlePanic()

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "handle-hook":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Error: hook event name required\n")
			printUsage()
			os.Exit(1)
		}
		handleHook(os.Args[2])
	case "focus-window":
		if len(os.Args) < 4 {
			fmt.Fprintf(os.Stderr, "Error: focus-window requires bundleID and cwd arguments\n")
			os.Exit(1)
		}
		opts, err := parseFocusWindowOptions(os.Args[4:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "focus-window: %v\n", err)
			os.Exit(1)
		}
		if err := notifier.FocusAppWindowWithOptions(os.Args[2], os.Args[3], opts); err != nil {
			fmt.Fprintf(os.Stderr, "focus-window: %v\n", err)
			os.Exit(1)
		}
	case "play-sound":
		runPlaySound(os.Args[2:])
	case "daemon", "--daemon":
		runDaemon()
	case "windows-hooks":
		runWindowsHooks(os.Args[2:])
	case "version", "--version", "-v":
		fmt.Printf("claude-notifications v%s\n", version)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

type hookSettings struct {
	Hooks map[string][]hookMatcherGroup `json:"hooks"`
}

type hookMatcherGroup struct {
	Matcher string        `json:"matcher,omitempty"`
	Hooks   []hookCommand `json:"hooks"`
}

type hookCommand struct {
	Type    string `json:"type"`
	Command string `json:"command"`
	Timeout int    `json:"timeout"`
	Shell   string `json:"shell"`
}

func runWindowsHooks(args []string) { _ = "STUB: not implemented"; return }

func parseWindowsHooksExecutable(args []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func newPowerShellHook(exePath, hookName string) hookCommand {
	_ = "STUB: not implemented"
	return *new(hookCommand)
}

func powershellDoubleQuoted(value string) string { _ = "STUB: not implemented"; return "" }

func handleHook(hookEvent string) {
	_ = "STUB: not implemented"
	// Add panic recovery for this function
	return
}

// Determine plugin root

// Initialize logger

// Create handler

// Handle hook

type pluginManifest struct {
	Version string `json:"version"`
}

func maybeScheduleWindowsLazyUpdate(pluginRoot string) { _ = "STUB: not implemented"; return }

func readPluginManifestVersion(pluginRoot string) string { _ = "STUB: not implemented"; return "" }

func windowsLazyUpdateStampPath(pluginRoot string) string { _ = "STUB: not implemented"; return "" }

func windowsLazyUpdateRecentlyScheduled(stampPath, stampKey string) bool {
	_ = "STUB: not implemented"
	return false
}

func writeWindowsLazyUpdateStamp(stampPath, stampKey string) error {
	_ = "STUB: not implemented"
	return nil
}

func scheduleWindowsLazyUpdateImpl(pluginRoot string) error { _ = "STUB: not implemented"; return nil }

func findWindowsPowerShell() (string, error) { _ = "STUB: not implemented"; return "", nil }

func findWindowsBash() (string, error) { _ = "STUB: not implemented"; return "", nil }

func windowsBashCandidates() []string { _ = "STUB: not implemented"; return nil }

func shellSingleQuoted(value string) string { _ = "STUB: not implemented"; return "" }

func powershellSingleQuoted(value string) string { _ = "STUB: not implemented"; return "" }

func getPluginRoot() string {
	_ = "STUB: not implemented"
	// Try CLAUDE_PLUGIN_ROOT environment variable first
	return ""
}

// Try to find plugin root relative to executable

// Executable is in bin/, so plugin root is parent directory

// Otherwise, try parent of executable dir

// Fallback to current directory

// runPlaySound plays a sound file and exits. Designed to be spawned as a detached
// child process so the parent hook process does not wait for audio to finish.
// Usage: play-sound <path> [--volume <0.0-1.0>] [--device <name>]
func runPlaySound(args []string) { _ = "STUB: not implemented"; return }

// Parse optional flags

func parseFocusWindowOptions(args []string) (notifier.FocusWindowOptions, error) {
	_ = "STUB: not implemented"
	return *new(notifier.FocusWindowOptions), nil
}

func printUsage() { _ = "STUB: not implemented"; return }
