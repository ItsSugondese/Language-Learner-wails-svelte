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

func (f *FileService) ListFiles(path string) ([]string, error) {
	fullPath := filepath.Join(filepathconstants.AssetsFolderAbsolutePath, path)
	filePathSlices, err := file_utils.GetFileNamesInPathFromDirectory(fullPath)

	if err != nil {
		return nil, err
	}

	return filePathSlices, nil
}
