package config

import (
	"cloudflare-ddns/pkg/slogger"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig(name string, path []string, defaults map[string]any) {
	slogger.Info("Config Init", "config-file", name, "config-path", path)
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")

	for _, value := range path {
		viper.AddConfigPath(value)
	}

	for key, value := range defaults {
		if value == "" {
			envKey := strings.ToUpper(strings.Replace(key, "-", "_", -1))
			viper.SetDefault(key, os.Getenv(envKey))
			continue
		}
		viper.SetDefault(key, value)
	}

	if err := viper.ReadInConfig(); err != nil {
		slogger.Error("Config error", "error", "file not found")
		return
	}
	slogger.Info("Config Init", "loaded", true, "total", len(viper.AllKeys()))

}
