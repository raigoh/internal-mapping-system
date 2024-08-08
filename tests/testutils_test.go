package tests

import (
	"fmt"
	"os"
	"path/filepath"
)

// findMainGo searches for main.go in the parent directory and its subdirectories
func findMainGo() (string, error) {
	// Start from the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %v", err)
	}

	// Go up one level to the parent directory
	parentDir := filepath.Dir(currentDir)

	// Walk through the parent directory and its subdirectories
	var mainPath string
	err = filepath.Walk(parentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "main.go" {
			mainPath = path
			return filepath.SkipAll // Stop walking once main.go is found
		}
		return nil
	})

	if err != nil {
		return "", fmt.Errorf("error while searching for main.go: %v", err)
	}

	if mainPath == "" {
		return "", fmt.Errorf("main.go not found in parent directory or its subdirectories")
	}

	fmt.Printf("Debug: Found main.go at: %s\n", mainPath) // Add this line for debugging

	return mainPath, nil
}
