//go:build !linux

package main

// runDaemon is a stub for non-Linux platforms
func runDaemon() { _ = "STUB: not implemented"; return }
