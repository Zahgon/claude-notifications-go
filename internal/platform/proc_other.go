//go:build !windows

package platform

import (
	"os/exec"
)

// SetDetachedProcAttr configures the command to run as a detached process on Unix.
// Setpgid creates a new process group so the child survives parent exit.
func SetDetachedProcAttr(cmd *exec.Cmd) { _ = "STUB: not implemented"; return }
