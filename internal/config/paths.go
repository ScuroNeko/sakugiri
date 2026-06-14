package config

import (
	"os"
	"path/filepath"
)

func GetConfigDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appConfigDir := filepath.Join(dir, "sakugiri")
	return appConfigDir, nil
}
