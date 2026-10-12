# Codex + Clash Verge Automated Installer

Cross-platform (Windows / macOS) automated installer for the latest Clash Verge Rev and Codex (ChatGPT desktop app).

## Features
- Dynamic latest version detection
- Multi-source + automatic proxy speed selection (system proxy, Clash ports, environment variables)
- Multi-threaded download with progress bars
- Downloads installers next to the executable
- SHA256 verification
- Windows guided flow: topmost window → open registration link → Clash GUI install → Codex MSIX via Add-AppxPackage → open chatgpt.com → final verification → self-delete
- macOS support via launcher script

## Build
```bash
make
```

## Usage
Run the binary for your platform. On Windows the tool follows the guided sequence and removes itself after successful verification.

See also:
- [README.zh.md](README.zh.md) (中文)
- [README.ja.md](README.ja.md) (日本語)
