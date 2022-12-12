package main

import (
	"embed"
	"github.com/Nik77x/Blender_Manager/backend"
	"github.com/Nik77x/Blender_Manager/backend/Data"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := backend.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Blender_Manager",
		Width:  1175,
		Height: 680,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
			&Data.BlendInfo{},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
