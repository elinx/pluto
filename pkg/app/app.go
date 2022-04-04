package app

import (
	"github.com/pluto/pkg/ui"
	"github.com/rivo/tview"
)

type App struct {
	*tview.Application
	ui *ui.UI
}

func NewApp(rootPath string) *App {
	return &App{
		Application: tview.NewApplication(),
		ui:          ui.NewUI(rootPath),
	}
}

func (app *App) Run() {
	app.ui.SetupKeyboard(app.Application)
	if err := app.SetRoot(app.ui, true).Run(); err != nil {
		panic(err)
	}
}
