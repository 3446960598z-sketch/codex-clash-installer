.PHONY: all win mac clean

all: win mac

win:
	GOOS=windows GOARCH=amd64 go build -o dist/codex-clash-installer-win-amd64.exe .

mac:
	GOOS=darwin GOARCH=arm64 go build -o dist/codex-clash-installer-mac-arm64 .
	GOOS=darwin GOARCH=amd64 go build -o dist/codex-clash-installer-mac-amd64 .

clean:
	rm -rf dist
