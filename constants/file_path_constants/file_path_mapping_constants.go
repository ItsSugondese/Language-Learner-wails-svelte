package filepathconstants

import (
	"lang-learner-wails/enums/struct-enums/file-path-enums"
)

// FilePathMapping represents the file path mapping struct
type FilePathMapping struct {
	Path string
}

// Define the file path mappings
var (
	AssetsFolder = FilePathMapping{
		Path: AssetsFolderAbsolutePath,
	}
	StorageFolder = FilePathMapping{
		Path: StorageFolderAbsolutePath,
	}
	GermanTextStorageFolder = FilePathMapping{
		Path: GermanTextStorageFolderAbsolutePath,
	}
	AudioFolder = FilePathMapping{
		Path: AudioFolderAbsolutePath,
	}
	GermanAudioFolder = FilePathMapping{
		Path: GermanAudioFolderAbsolutePath,
	}
	EnglishAudioFolder = FilePathMapping{
		Path: EnglishAudioFolderAbsolutePath,
	}

	GermanSampleFolder = FilePathMapping{
		Path: SampleFolderPath,
	}
	GermanRandomAudioFolder = FilePathMapping{
		Path: RandomAudioFolderAbsolutePath,
	}
	GermanAllWordFolder = FilePathMapping{
		Path: AllWordFolderPath,
	}
	GermanAllWordAudioFolder = FilePathMapping{
		Path: AllWordAudioFolderAbsolutePath,
	}
)

// FilePathMappings map for easy lookup
var FilePathMappings = map[string]FilePathMapping{
	file_path_enums.FilePathEnums.AssetsFolderAbsolutePath:            AssetsFolder,
	file_path_enums.FilePathEnums.StorageFolderAbsolutePath:           StorageFolder,
	file_path_enums.FilePathEnums.GermanTextStorageFolderAbsolutePath: GermanTextStorageFolder,
	file_path_enums.FilePathEnums.AudioFolderAbsolutePath:             AudioFolder,
	file_path_enums.FilePathEnums.GermanAudioFolderAbsolutePath:       GermanAudioFolder,
	file_path_enums.FilePathEnums.EnglishAudioFolderAbsolutePath:      EnglishAudioFolder,

	file_path_enums.FilePathEnums.GermanSampleFolderPath:               GermanSampleFolder,
	file_path_enums.FilePathEnums.GermanRandomAudioFolderAbsolutePath:  GermanRandomAudioFolder,
	file_path_enums.FilePathEnums.GermanAllWordFolderPath:              GermanAllWordFolder,
	file_path_enums.FilePathEnums.GermanAllWordAudioFolderAbsolutePath: GermanAllWordAudioFolder,
}
