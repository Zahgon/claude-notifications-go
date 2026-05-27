//go:build linux

// ABOUTME: IPC protocol types for communication between daemon client and server.
// ABOUTME: Uses JSON-over-Unix-socket for simple, reliable inter-process communication.
package daemon

import (
	"errors"
)

// Common errors
var (
	ErrDaemonNotAvailable = errors.New("daemon not available")
	ErrDaemonNotRunning   = errors.New("daemon not running")
)

// Protocol version for compatibility checking
const ProtocolVersion = "1.0"

// MessageType identifies the type of IPC message
type MessageType string

const (
	MessageTypeNotify MessageType = "notify"
	MessageTypePing   MessageType = "ping"
	MessageTypeStop   MessageType = "stop"
)

// Request is the wrapper for all IPC requests
type Request struct {
	Type    MessageType    `json:"type"`
	Notify  *NotifyRequest `json:"notify,omitempty"`
	Version string         `json:"version"`
}

// Response is the wrapper for all IPC responses
type Response struct {
	Type   MessageType     `json:"type"`
	Notify *NotifyResponse `json:"notify,omitempty"`
	Ping   *PingResponse   `json:"ping,omitempty"`
	Error  string          `json:"error,omitempty"`
}

// NotifyRequest contains notification details sent to the daemon
type NotifyRequest struct {
	Title            string `json:"title"`
	Body             string `json:"body"`
	FocusTarget      string `json:"focus_target"`                 // Terminal identifier (empty = auto-detect)
	FocusFolder      string `json:"focus_folder,omitempty"`       // Project folder name for window-specific focus
	FocusWindowID    string `json:"focus_window_id,omitempty"`    // Exact X11 window ID captured in the hook process
	FocusWindowTitle string `json:"focus_window_title,omitempty"` // Exact window title captured in the hook process when available
	Timeout          int    `json:"timeout"`                      // Notification timeout in seconds
}

// NotifyResponse contains the result of a notification request
type NotifyResponse struct {
	Success        bool   `json:"success"`
	NotificationID uint32 `json:"notification_id"`
	Error          string `json:"error,omitempty"`
}

// PingResponse contains daemon status information
type PingResponse struct {
	Version string `json:"version"`
	Uptime  int64  `json:"uptime"` // Seconds since daemon started
}

// GetSocketPath returns the Unix socket path for the daemon.
// Uses XDG_RUNTIME_DIR if available, falls back to /tmp with UID suffix.
func GetSocketPath() string {
	_ = "STUB: not implemented"
	// Prefer XDG_RUNTIME_DIR (usually /run/user/1000)
	return ""
}

// Fallback to /tmp with UID for isolation

// GetPidFilePath returns the path to the daemon's PID file.
func GetPidFilePath() string { _ = "STUB: not implemented"; return "" }
