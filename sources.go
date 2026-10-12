package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Release struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func getLatestClash() (string, error) {
	resp, err := http.Get("https://api.github.com/repos/clash-verge-rev/clash-verge-rev/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var rel Release
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return "", err
	}
	fmt.Printf("Latest Clash: %s\n", rel.TagName)
	// Return appropriate asset URL
	return "https://github.com/clash-verge-rev/clash-verge-rev/releases/latest", nil
}
