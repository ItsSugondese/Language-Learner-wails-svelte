package file_path_enums

var FilePathEnums = newFilePath()

func newFilePath() *filePaths {
	return &filePaths{
		AssetsFolderAbsolutePath:            "AssetsFolderAbsolutePath",
		StorageFolderAbsolutePath:           "StorageFolderAbsolutePath",
		GermanTextStorageFolderAbsolutePath: "GermanTextStorageFolderAbsolutePath",
		AudioFolderAbsolutePath:             "AudioFolderAbsolutePath",
		GermanAudioFolderAbsolutePath:       "GermanAudioFolderAbsolutePath",
		EnglishAudioFolderAbsolutePath:      "EnglishAudioFolderAbsolutePath",

		GermanSampleFolderPath:               "GermanSampleFolderPath",
		GermanRandomAudioFolderAbsolutePath:  "GermanRandomAudioFolderAbsolutePath",
		GermanAllWordFolderPath:              "GermanAllWordFolderPath",
		GermanAllWordAudioFolderAbsolutePath: "GermanAllWordAudioFolderAbsolutePath",
	}
}

type filePaths struct {
	AssetsFolderAbsolutePath            string
	StorageFolderAbsolutePath           string
	GermanTextStorageFolderAbsolutePath string
	AudioFolderAbsolutePath             string
	GermanAudioFolderAbsolutePath       string
	EnglishAudioFolderAbsolutePath      string

	GermanSampleFolderPath               string
	GermanRandomAudioFolderAbsolutePath  string
	GermanAllWordFolderPath              string
	GermanAllWordAudioFolderAbsolutePath string
}
