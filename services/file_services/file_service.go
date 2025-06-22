package file_services

import (
	filepathconstants "lang-learner-wails/constants/file_path_constants"
	file_utils "lang-learner-wails/pkg/utils/file_utils"
	"path/filepath"
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

func (f *FileService) ListDataFromFiles(pathEnum string, fileName string) ([]string, error) {
	filePathSlices, err := file_utils.GetAllFromFileAsSlices(filepath.Join(filepathconstants.FilePathMappings[pathEnum].Path, fileName))

	if err != nil {
		return nil, err
	}

	return filePathSlices, nil
}
