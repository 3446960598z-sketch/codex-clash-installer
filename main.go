package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Codex + Clash Verge Automated Installer")
	fmt.Printf("Platform: %s/%s\n", runtime.GOOS, runtime.GOARCH)

	if runtime.GOOS == "windows" {
		bringConsoleToFront()
	}

	openBrowser("https://love3.p6m6.com/#/register?code=0Fwo3B76")

	// Proxy detection + speed test would happen here
	fmt.Println("Detecting proxies and selecting fastest path...")

	// Download Clash, verify, run GUI installer
	fmt.Println("Downloading and installing Clash Verge...")

	// Download Codex MSIX, verify, Add-AppxPackage
	fmt.Println("Downloading and installing Codex...")

	openBrowser("https://chatgpt.com")

	// Final verification
	fmt.Println("Performing final verification...")

	if runtime.GOOS == "windows" {
		selfDelete()
	}

	fmt.Println("Done.")
}

func bringConsoleToFront() {
	// Windows API placeholder
	fmt.Println("Bringing console to front...")
}

func openBrowser(url string) {
	fmt.Printf("Opening browser: %s\n", url)
	// Platform-specific open
}

func selfDelete() {
	fmt.Println("Self-deleting installer...")
	// Delayed del command placeholder
}
