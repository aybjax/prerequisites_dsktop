package main

import (
	"embed"
	"fmt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	appMenu := menu.NewMenu()

	aboutMenu := appMenu.AddSubmenu("Help")
	aboutMenu.AddSeparator()
	aboutMenu.AddText("About", keys.CmdOrCtrl("h"), app.HelpMenu)
	// if runtime.GOOS == "darwin" {
	// 	appMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	// }

	// Create application with options
	err := wails.Run(&options.App{
		Title:  fmt.Sprintf("Discover prerequisites with aybjax. Version: %s", AppVersion),
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Menu:       appMenu,
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
