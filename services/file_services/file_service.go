package file_services

import (
	"encoding/base64"
	filepathconstants "lang-learner-wails/constants/file_path_constants"
	file_utils "lang-learner-wails/pkg/utils/file_utils"
	"lang-learner-wails/pkg/utils/folder_utils"
	"os"
	"path/filepath"
	"strings"
)

type FileService struct{}

func NewFileService() *FileService {
	return &FileService{}
}

func (f *FileService) ListFiles(pathEnum string, withDir bool) ([]string, error) {
	filePathSlices, err := file_utils.GetFileNamesInPathFromDirectory(filepathconstants.FilePathMappings[pathEnum].Path, withDir)

	if err != nil {
		return nil, err
	}

	return filePathSlices, nil
}

func (f *FileService) ListDataFromFilePathEnum(pathEnum string, fileName string) ([]string, error) {
	filePathSlices, err := file_utils.GetAllFromFileAsSlices(filepath.Join(filepathconstants.FilePathMappings[pathEnum].Path, fileName))

	if err != nil {
		return nil, err
	}

	return filePathSlices, nil
}

func (f *FileService) ListDataFromFilePath(path string) ([]string, error) {
	filePathSlices, err := file_utils.GetAllFromFileAsSlices(path)

	if err != nil {
		return nil, err
	}

	return filePathSlices, nil
}

func (f *FileService) FindFileFromFilePathEnumIfExistsAndReturnEncoded(pathEnum string, subDirName, fileName string, shouldSearchAll bool) (string, error) {
	var dirToSearch []string
	//var err error

	mainPath := filepathconstants.FilePathMappings[pathEnum].Path
	if shouldSearchAll {
		foldersToSearch, err := folder_utils.ListFoldersInDir(mainPath)

		for _, subFolder := range foldersToSearch {
			dirToSearch = append(dirToSearch, filepath.Join(mainPath, subFolder))
		}
		if err != nil {
			return "", err
		}
	} else {
		dirToSearch = []string{filepath.Join(mainPath, strings.ToLower(subDirName))}
	}

	for _, dir := range dirToSearch {
		fileToSearch := filepath.Join(dir, strings.ToLower(fileName))
		if file_utils.FileExists(fileToSearch) {
			data, err := os.ReadFile(fileToSearch)

			if err != nil {
				return "", err
			}
			return "data:audio/mpeg;base64," + base64.StdEncoding.EncodeToString(data), nil
		}

	}

	return "", nil
}
