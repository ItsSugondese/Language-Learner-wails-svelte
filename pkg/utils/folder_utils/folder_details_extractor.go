package folder_utils

import (
	"os"
)

func ListFoldersInDir(dirPath string) ([]string, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name()) // or entry.Name() for just folder name
		}
	}
	return dirs, nil
}
