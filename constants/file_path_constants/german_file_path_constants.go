package filepathconstants

import "path/filepath"

var (
	SampleFolderPath = filepath.Join(GermanTextStorageFolderAbsolutePath, "sample")

	RandomAudioFolderAbsolutePath = filepath.Join(GermanAudioFolderAbsolutePath, "random")

	AllWordFolderPath              = filepath.Join(GermanTextStorageFolderAbsolutePath, "all-word")
	AllWordAudioFolderAbsolutePath = GermanAudioFolderAbsolutePath
)
