package main

import (
	"context"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// OnStartup is called when the app starts up
func (a *App) OnStartup(ctx context.Context) {
	a.ctx = ctx
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Discord",
		Width:            1200,
		Height:           800,
		BackgroundColour: &options.RGBA{R: 54, G: 57, B: 63, A: 1}, // Discord's dark theme color
		OnStartup:        app.OnStartup,
		WebviewIsTransparent: false,
		WindowStartState: options.Normal,
	})

	if err != nil {
		log.Fatal("Error:", err)
	}
}
