package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/xackery/quail-wails/app"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	//a := NewApp()
	a := app.New()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "quai-gui",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		StartHidden:      true,
		BackgroundColour: &options.RGBA{R: 21, G: 23, B: 24, A: 255},
		OnStartup:        a.OnStartup,
		OnDomReady:       a.OnDomReady,
		OnShutdown:       a.OnShutdown,
		Bind: []interface{}{
			a,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
