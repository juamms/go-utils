package utils

import (
	"encoding/json"
	"io/ioutil"
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

// LoadJSONFromFile loads a JSON file from `path` and populated the given `model`.
func LoadJSONFromFile(path string, model interface{}) error {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &model)

	if err != nil {
		return err
	}

	return nil
}

// ParseConfig reads a `config.json` file on the same directory as the executable and parses its JSON to the given model.
func ParseConfig(model interface{}) error {
	executablePath, err := GetExecutablePath()

	if err != nil {
		return err
	}

	path := SafeJoinPaths(executablePath, "config.json")

	return LoadJSONFromFile(path, model)
}
