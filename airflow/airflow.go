package airflow

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/astronomerio/astro-cli/utils"
)

func initDirs(root string, dirs []string) bool {
	// Any inputs exist
	exists := false

	// Create the dirs
	for _, dir := range dirs {
		// Create full path to directory
		fullpath := filepath.Join(root, dir)

		// Move on if already exists
		if utils.Exists(fullpath) {
			exists = true
			continue
		}

		// Create directory
		if err := os.MkdirAll(dir, 0777); err != nil {
			fmt.Println(err)
		}
	}

	return exists
}

func initFiles(root string, files map[string]string) bool {
	// Any inputs exist
	exists := false

	// Create the files
	for file, content := range files {
		// Create full path to file
		fullpath := filepath.Join(root, file)

		// Move on if already exiss
		if utils.Exists(fullpath) {
			exists = true
			continue
		}

		// Write files out
		if err := utils.WriteStringToFile(fullpath, content); err != nil {
			fmt.Println(err)
		}
	}

	return exists
}

// Init will scaffold out a new airflow project
func Init(path string) error {
	// List of directories to create
	dirs := []string{"dags", "plugins", "include"}

	// Map of files to create
	files := map[string]string{
		".dockerignore":    dockerignore,
		"Dockerfile":       dockerfile,
		"packages.txt":     "",
		"requirements.txt": "",
	}

	// Initailize directories
	initDirs(path, dirs)

	// Initialize files
	initFiles(path, files)

	return nil
}

// Create new airflow deployment
func Create() error {
	return nil
}
