package util

import (
	"io"
	"net/http"
	"strings"
)

func GetPublicIP() string {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "N/A"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "N/A"
	}

	ip := strings.TrimSpace(string(body))
	return ip
}
