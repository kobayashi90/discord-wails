package main

import (
    "embed"

    "github.com/wailsapp/wails/v2"
    "github.com/wailsapp/wails/v2/pkg/options"
    "github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
    err := wails.Run(&options.App{
        Title:            "Discord",
        Width:            1024,
        Height:           768,
        Assets:           assets,
        WindowStartState: options.Normal,
        Windows: &windows.Options{
            WebviewIsTransparent: false,
            WindowIsTranslucent:  false,
        },
    })

    if err != nil {
        panic(err)
    }
}
