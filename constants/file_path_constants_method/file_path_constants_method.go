package file_path_constants_method

import (
	filepathconstants "lang-learner-wails/constants/file_path_constants"
	"os"
	"path/filepath"
)

func GetAllWordPath() string {
	return filepath.Join(GetJavaResPath(), filepathconstants.FileSeparator, "german", filepathconstants.FileSeparator, "all-word")

}

func GetJavaResPath() string {
	return os.Getenv("JAVA_PROJECT_RESOURCE_PATH")
}
