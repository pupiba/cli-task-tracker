package utils

import "os"

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	return false
}
