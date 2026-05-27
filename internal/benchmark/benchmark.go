package benchmark

import (
	"sync"
	"time"
)

// Context provides lightweight timing instrumentation for hook execution.
// Enabled via config.json: { "debug": { "benchmark": true } }.
// Output goes to the plugin log file via the provided logFunc.
type Context struct {
	mu      sync.Mutex
	events  []TimingEvent
	starts  map[string]time.Time
	enabled bool
	logFunc func(format string, args ...interface{})
}

// TimingEvent represents a single timing measurement.
type TimingEvent struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"` // "interval", "checkpoint", "async"
	Duration float64 `json:"duration_ms"`
	Success  bool    `json:"success"`
}

// New creates a new benchmark context.
// When enabled, timing data is written to the log file via logFunc.
func New(enabled bool, logFunc func(string, ...interface{})) *Context {
	_ = "STUB: not implemented"
	return nil
}

// Enabled returns whether benchmarking is active.
func (c *Context) Enabled() bool {
	_ = "STUB: not implemented"

	// Start begins timing a named operation.
	return false
}

func (c *Context) Start(name string) { _ = "STUB: not implemented"; return }

// Elapsed records the duration since the matching Start call.
func (c *Context) Elapsed(name string) { _ = "STUB: not implemented"; return }

// Checkpoint records a point-in-time event with a success/failure flag.
func (c *Context) Checkpoint(name string, success bool) { _ = "STUB: not implemented"; return }

// Async marks a non-blocking operation that runs in the background.
func (c *Context) Async(name string) { _ = "STUB: not implemented"; return }

// Report outputs all collected timing data to the plugin log file.
func (c *Context) Report() { _ = "STUB: not implemented"; return }
