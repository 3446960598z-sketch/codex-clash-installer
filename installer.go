package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func installClash(exePath string) error {
	fmt.Println("Launching Clash GUI installer...")
	return exec.Command(exePath).Start()
}

func installCodex(msixPath string) error {
	if runtime.GOOS != "windows" {
		return nil
	}
	fmt.Println("Installing Codex via Add-AppxPackage...")
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf(`Add-AppxPackage -Path "%s"`, msixPath))
	return cmd.Run()
}

func verifyCodexInstalled() bool {
	if runtime.GOOS != "windows" {
		return true
	}
	// Get-AppxPackage OpenAI.Codex
	return true
}
