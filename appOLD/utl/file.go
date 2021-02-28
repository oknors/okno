package utl

import (
	"os"
	"path/filepath"
)

// FileExists returns whether a file exists
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return true
}

// EnsureDir checks a file could be written to a path, creates the directories
// as needed
func EnsureDir(fileName string) bool {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			return false
		}
	}
	return true
}
