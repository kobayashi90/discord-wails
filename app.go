package main

import (
    "github.com/wailsapp/wails/v2"
    "github.com/wailsapp/wails/v2/pkg/options"
    "github.com/wailsapp/wails/v2/pkg/options/windows"
)

func main() {
    appURL := "https://discord.com/app"

    err := wails.Run(&options.App{
        Title:  "Discord",
        Width:  1024,
        Height: 768,
        Frameless: false,
        StartHidden: false,
        DisableResize: false,
        Windows: &windows.Options{
            WebviewIsTransparent: false,
            WindowIsTranslucent:  false,
        },
        Assets: nil, // We don't use a local frontend here
        OnStartup: func(ctx *wails.Context) {
            ctx.Runtime.Browser().OpenURL(appURL)
        },
    })

    if err != nil {
        panic(err)
    }
}
