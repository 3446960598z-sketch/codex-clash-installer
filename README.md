# Codex + Clash Verge Automated Installer

Cross-platform CLI tool (Windows & macOS) that automates downloading and installing the latest Clash Verge Rev and Codex (ChatGPT desktop app).

## Features
- Dynamic latest version retrieval (GitHub API / CDN)
- Multi-source speed test + automatic proxy detection (system proxy, Clash ports, env vars)
- Multi-threaded download with total + segmented progress bars
- Download to the directory of the running executable
- SHA256 verification
- Guided Windows flow: topmost console → open registration link → download Clash → GUI install → download Codex MSIX → Add-AppxPackage → open chatgpt.com → final verify → self-delete
- macOS support via .command launcher with admin password popup

## Build
```bash
make
# or
go build -o dist/codex-clash-installer-win-amd64.exe .
```

## Usage
Run the platform binary. On Windows it follows the guided sequence and cleans itself up after successful verification.
