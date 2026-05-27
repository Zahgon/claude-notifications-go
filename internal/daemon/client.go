//go:build linux

// ABOUTME: Client library for communicating with the notification daemon.
// ABOUTME: Provides functions to check daemon status, start on-demand, and send notifications.
package daemon

// Client communicates with the daemon via Unix socket
type Client struct {
	socketPath string
}

// NewClient creates a new daemon client
func NewClient() (*Client, error) {
	_ = "STUB: not implemented"
	return nil,

		// Check if socket exists
		nil
}

// SendNotification sends a notification request to the daemon.
// focusFolder is the project folder name for window-specific focus (may be empty).
// focusWindowID and focusWindowTitle are optional exact window hints captured in the hook process.
func (c *Client) SendNotification(
	title,
	body,
	focusTarget,
	focusFolder,
	focusWindowID,
	focusWindowTitle string,
	timeout int,
) (*NotifyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ping checks if the daemon is responding and returns status info
func (c *Client) Ping() (*PingResponse, error) { _ = "STUB: not implemented"; return nil, nil }

// Stop requests the daemon to shut down
func (c *Client) Stop() error { _ = "STUB: not implemented"; return nil }

// send sends a request to the daemon and returns the response
func (c *Client) send(req Request) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// Set deadlines

// Send request

// Read response

// IsDaemonRunning checks if the daemon is running and responsive
func IsDaemonRunning() bool { _ = "STUB: not implemented"; return false }

// StartDaemonOnDemand starts the daemon if it's not already running.
// Returns true if daemon is running (either started now or was already running).
func StartDaemonOnDemand() bool {
	_ = "STUB: not implemented"
	// Check if already running
	return false
}

// Find the daemon binary

// Start daemon in background

// Create new session (detach from terminal)

// Redirect stdout/stderr to /dev/null

// Wait for daemon to be ready (up to 5 seconds)

// StopDaemon stops the running daemon
func StopDaemon() error { _ = "STUB: not implemented"; return nil }

// findDaemonBinary locates the daemon binary
func findDaemonBinary() (string, error) {
	_ = "STUB: not implemented"
	// Check CLAUDE_PLUGIN_ROOT
	return "", nil
}

// Try different binary names

// Check if current executable supports daemon mode

// Try to find in PATH

// GetDaemonPID returns the PID of the running daemon, or 0 if not running
func GetDaemonPID() int { _ = "STUB: not implemented"; return 0 }

// Check if process exists

// On Unix, FindProcess always succeeds; check if process exists with signal 0
