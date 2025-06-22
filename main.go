package main

import (
	"context"
	"embed"
	"lang-learner-wails/services/file_services"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	fileService := file_services.NewFileService()
	filePickerService := file_services.NewFilePicker()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "lang-learner-wails",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			// Initialize all services that need context
			app.startup(ctx)
			filePickerService.WailsInit(ctx)
			// If fileService also needs context:
			// fileService.SetContext(ctx)
		},
		Bind: []interface{}{
			app,
			fileService,
			filePickerService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
