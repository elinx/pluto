package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UI struct {
	*tview.Flex
	leftPanel   *LeftPanel
	middlePanel *MiddlePanel
	rightPanel  *tview.Box
}

func NewUI(rootPath string) *UI {
	leftPanel := NewLeftPanel(rootPath)
	middlePanel := NewMiddlePanel()
	rightPanel := tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)")
	flex := tview.NewFlex().
		AddItem(leftPanel, 0, 1, true).
		AddItem(middlePanel, 0, 3, true).
		AddItem(rightPanel, 20, 1, false)
	return &UI{
		Flex:        flex,
		leftPanel:   leftPanel,
		middlePanel: middlePanel,
		rightPanel:  rightPanel,
	}
}

func (ui *UI) SetupKeyboard(app *tview.Application) {
	ui.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			for i := 0; i < ui.GetItemCount(); i++ {
				if ui.GetItem(i).HasFocus() {
					j := (i + 1) % ui.GetItemCount()
					app.SetFocus(ui.GetItem(j))
					break
				}
			}
			return nil
		}
		return event
	})
}
