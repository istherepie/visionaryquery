package config

import "os"

// FileExists checks for the existance of a config file
func FileExists(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return false
	}

	return true
}
