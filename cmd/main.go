package main

import (
	"cloudflare-ddns/pkg/config"
	"cloudflare-ddns/pkg/slogger"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		slogger.Fatal("Failed to get current working directory: %v", err)
	}

	slogger.Info("Starting Cloudflare DDNS...")
	config.InitConfig("config", []string{"/config/", "config", dir + "/config"}, map[string]any{})
}
