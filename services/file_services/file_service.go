package file_services

import (
	"encoding/base64"
	"fmt"
	delimiter_constants "lang-learner-wails/constants/delimiter-constants"
	filepathconstants "lang-learner-wails/constants/file_path_constants"
	file_utils "lang-learner-wails/pkg/utils/file_utils"
	"lang-learner-wails/pkg/utils/folder_utils"
	"lang-learner-wails/services/api_services"
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

func (f *FileService) ListDataFromAllFiles(pathEnum string) (files []string, err error) {
	allFilesPath, err := f.ListFiles(pathEnum, true)

	if err != nil {
		return nil, err
	}

	for _, path := range allFilesPath {
		dataFromFile, err := f.ListDataFromFilePath(path)
		if err != nil {
			return nil, err
		}
		files = append(files, dataFromFile...)
	}
	return files, nil
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
		fileToSearch := filepath.Join(dir, strings.ToLower(fileName)+".mp3")
		if file_utils.FileExists(fileToSearch) {
			data, err := os.ReadFile(fileToSearch)

			if err != nil {
				return "", err
			}
			return "data:audio/mpeg;base64," + base64.StdEncoding.EncodeToString(data), nil
		}
	}
	return f.callApiAndSaveAudio(fileName, shouldSearchAll, mainPath, subDirName)
}

func (f *FileService) LoadAllAudioInDirectory(textPathEnum string, audioPathEnum string, topicName string) error {
	textPath := filepathconstants.FilePathMappings[textPathEnum].Path
	fileData, err := f.ListDataFromFilePath(filepath.Join(textPath, topicName))

	if err != nil {
		return err
	}

	for index, fileName := range fileData {
		fmt.Printf("%d of %d\n", index+1, len(fileData))

		parts := strings.Split(fileName, delimiter_constants.PipeDelimiter)
		if len(parts) > 0 {
			word := strings.ToLower(strings.TrimSpace(parts[0]))
			fileToSearch := filepath.Join(filepathconstants.FilePathMappings[audioPathEnum].Path, strings.ToLower(topicName), word+".mp3")

			if !file_utils.FileExists(fileToSearch) {
				_, err := f.callApiAndSaveAudio(word, false, filepathconstants.FilePathMappings[audioPathEnum].Path, topicName)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

func (f *FileService) callApiAndSaveAudio(fileName string, isRandomToStore bool, mainPath string, subDirName string) (string, error) {
	data, err := api_services.ElevenLabsTextToSpeechAudioBytes(fileName)

	if err != nil {
		println(err)
		return "", err
	}

	if data != nil {

		var folderToSave string

		if isRandomToStore {
			folderToSave = filepathconstants.RandomAudioFolderAbsolutePath
		} else {
			folderToSave = filepath.Join(mainPath, strings.ToLower(subDirName))
		}

		err := os.MkdirAll(folderToSave, os.ModePerm)
		if err != nil {
			return "", err
		}

		err = os.WriteFile(filepath.Join(folderToSave, strings.ToLower(fileName)+".mp3"), data, 0644)
		if err != nil {
			return "", err
		}

		return "data:audio/mpeg;base64," + base64.StdEncoding.EncodeToString(data), nil
	}

	println("Erorr with api")
	return "", nil
}
