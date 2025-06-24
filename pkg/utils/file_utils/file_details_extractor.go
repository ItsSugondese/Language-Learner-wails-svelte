package file_utils

import (
	"bufio"
	"fmt"
	filepathconstants "lang-learner-wails/constants/file_path_constants"
	"os"
	"strings"
	"sync"
)

func GetFileNamesInPathFromDirectory(dir string, withDir bool) ([]string, error) {
	var fileNames []string

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			if withDir {
				fileNames = append(fileNames, dir+filepathconstants.FileSeparator+file.Name())
			} else {
				fileNames = append(fileNames, file.Name())
			}
		}
	}

	return fileNames, nil
}

func GetAllFromFileAsSlices(path string) ([]string, error) {
	var lines []string

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return lines, nil
}

func GetCombinedLinesFromFilesParallel(filePaths []string) ([]string, error) {
	var wg sync.WaitGroup
	lineChan := make(chan []string, len(filePaths))
	errChan := make(chan error, len(filePaths))

	for _, path := range filePaths {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			lines, err := GetAllFromFileAsSlices(p)
			if err != nil {
				errChan <- err
				return
			}
			lineChan <- lines
		}(path)
	}

	wg.Wait()
	close(lineChan)
	close(errChan)

	// Check if any errors occurred
	if len(errChan) > 0 {
		return nil, <-errChan // return the first error
	}

	var allLines []string
	for lines := range lineChan {
		allLines = append(allLines, lines...)
	}

	return allLines, nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
