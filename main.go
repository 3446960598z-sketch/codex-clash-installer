package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Codex + Clash Verge Installer")
	fmt.Printf("OS: %s Arch: %s\n", runtime.GOOS, runtime.GOARCH)

	// TODO: implement full flow based on requirements
	// - Topmost window (Windows)
	// - Open registration URL
	// - Proxy detection + speed test
	// - Download to exe dir
	// - Verify SHA256
	// - Install (GUI for Clash, Add-AppxPackage for Codex)
	// - Open chatgpt.com
	// - Final verification
	// - Self-delete

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("Starting guided installation flow...")
	// Placeholder for the full implementation
	return nil
}
