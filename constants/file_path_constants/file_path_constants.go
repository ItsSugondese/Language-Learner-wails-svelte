package filepathconstants

import (
	"os"
	"path/filepath"
	"runtime"
)

// FilePathConstants defines constants related to file paths.
var (
	_, b, _, _           = runtime.Caller(0)
	FileSeparator string = string(filepath.Separator)
	ProjectPath   string = filepath.Dir(filepath.Join(filepath.Dir(b), ".."))
	ProjectName   string = os.Getenv("PROJECT_NAME")
	PresentDir    string = filepath.Join(filepath.Dir(ProjectPath), "")
	UploadDir     string = filepath.Join(PresentDir, "?same-document", "?same", FileSeparator)

	AssetsFolderAbsolutePath            = filepath.Join(ProjectPath, "frontend", "src", "assets")
	StorageFolderAbsolutePath           = filepath.Join(AssetsFolderAbsolutePath, "storage")
	GermanTextStorageFolderAbsolutePath = filepath.Join(StorageFolderAbsolutePath, "german")

	AudioFolderAbsolutePath        = filepath.Join(StorageFolderAbsolutePath, "audio")
	GermanAudioFolderAbsolutePath  = filepath.Join(AudioFolderAbsolutePath, "german")
	EnglishAudioFolderAbsolutePath = filepath.Join(AudioFolderAbsolutePath, "english")
)
