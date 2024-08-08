package tests

import (
	"path/filepath"
)

// findMainGo searches for main.go in the parent directory
func findMainGo() (string, error) {
	dir, err := filepath.Abs("..")
	if err != nil {
		return "", err
	}
	mainPath := filepath.Join(dir, "main.go")
	return mainPath, nil
}
