package app

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/pluto/pkg/ui"
	"github.com/rivo/tview"
)

type App struct {
	*tview.Application
	ui     *ui.UI
	config *Configuration
}

func NewApp(config *Configuration) *App {
	return &App{
		config:      config,
		Application: tview.NewApplication(),
		ui:          ui.NewUI(config.resolveStartupDir()),
	}
}

func (app *App) Run() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlC {
			log.Println("Ctrl-C pressed. Exiting...")
			rootPath := app.ui.GetCurrentRootPath()
			log.Println("current root path:", rootPath)
			app.config.SetLeftOffDir(rootPath)
			app.config.Serialize()
			app.Stop()
		}
		return event
	})
	app.ui.SetupKeyboard(app.Application)
	if err := app.SetRoot(app.ui, true).Run(); err != nil {
		panic(err)
	}
}
