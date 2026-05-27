package platform

import (
	"os/exec"
)

// SetDetachedProcAttr configures the command to run as a detached process on Windows.
// CREATE_NEW_PROCESS_GROUP (0x200) detaches from parent's console group.
// CREATE_NO_WINDOW (0x08000000) prevents a visible console window.
func SetDetachedProcAttr(cmd *exec.Cmd) { _ = "STUB: not implemented"; return }
