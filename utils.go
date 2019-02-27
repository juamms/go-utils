package utils

import (
	"os"
	"path/filepath"
	"strconv"
)

// GetExecutablePath gets the path of the current executable.
func GetExecutablePath() (string, error) {
	ex, err := os.Executable()

	if err != nil {
		return "", err
	}

	return filepath.Dir(ex), nil
}

// SafeJoinPaths joins `path1` and `path2` and applies the correct separator.
func SafeJoinPaths(path1, path2 string) string {
	return filepath.FromSlash(filepath.Join(path1, path2))
}

// StringToInt converts a string to an int, and if the conversion fails, it uses the `defaultValue`.
func StringToInt(str string, defaultValue int) int {
	result, err := strconv.Atoi(str)

	if err != nil {
		return defaultValue
	}

	return result
}
