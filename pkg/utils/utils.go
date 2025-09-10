package utils

import (
	"os"
	"time"
)

const FILENAME string = "tasks.json"

func InitStorageFile(file string) {
	if !FileExists(file) {
		new_file, _ := os.Create(file)
		defer new_file.Close()
	}
}

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}

func FormattedTime(t time.Time) string {
	return t.Format("15:04 02.01.2006")
}

func UnormattedTime(str_t string) (time.Time, error) {
	return time.Parse("15:04 02.01.2006", str_t)
}
