package file_services

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	filepathconstants "lang-learner-wails/constants/file_path_constants"
)

type FilePicker struct {
	ctx context.Context
}

func NewFilePicker() *FilePicker {
	return &FilePicker{}
}

func (f *FilePicker) WailsInit(ctx context.Context) {
	f.ctx = ctx
}

func (f *FilePicker) SelectFile() (string, error) {
	if f.ctx == nil {
		return "", fmt.Errorf("context not initialized")
	}
	return runtime.OpenFileDialog(f.ctx, runtime.OpenDialogOptions{
		Title:            "Pick a file",
		DefaultDirectory: filepathconstants.GermanSampleFolder.Path,
	})
}
