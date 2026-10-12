#!/bin/bash
cd "$(dirname "$0")"
osascript -e 'do shell script "'"$(pwd)"'/codex-clash-installer-mac-arm64" with administrator privileges'
