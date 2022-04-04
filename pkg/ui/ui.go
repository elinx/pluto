package ui

import (
	"github.com/rivo/tview"
)

type UI struct {
	*tview.Flex
	leftPanel   *LeftPanel
	middlePanel *MiddlePanel
	rightPanel  *tview.Box
}

func NewUI(rootPath string) *UI {
	leftTreeView := NewLeftPanel(rootPath)
	middlePanel := NewMiddlePanel()
	rightPanel := tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)")
	flex := tview.NewFlex().
		AddItem(leftTreeView, 0, 1, true).
		AddItem(middlePanel, 0, 3, false).
		AddItem(rightPanel, 3, 1, false)
	return &UI{
		Flex:        flex,
		leftPanel:   leftTreeView,
		middlePanel: middlePanel,
		rightPanel:  rightPanel,
	}
}
