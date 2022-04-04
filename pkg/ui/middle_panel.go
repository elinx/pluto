package ui

import (
	"log"
	"os"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MiddlePanel struct {
	*tview.Flex
	commandPanel  *tview.Box
	fileListPanel *tview.Table
	statusPanel   *tview.Box
}

func NewMiddlePanel() *MiddlePanel {
	commandPanel := tview.NewBox().SetBorder(true).SetTitle("Command")
	fileList := fileListPanel()
	statusPanel := tview.NewBox().SetBorder(true).SetTitle("Status")
	middlePanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(commandPanel, 0, 1, false).
		AddItem(fileList, 0, 5, true).
		AddItem(statusPanel, 2, 1, false)
	return &MiddlePanel{
		Flex:          middlePanel,
		commandPanel:  commandPanel,
		fileListPanel: fileList,
		statusPanel:   statusPanel,
	}
}

func fileListPanel() *tview.Table {
	table := tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false).
		SetFixed(1, 1)
	return table
}

func feedTable(table *tview.Table) {
	contents := filesProvider()
	for i, v := range contents {
		for j, vv := range v {
			color := tcell.ColorDefault
			if i == 0 {
				color = tcell.ColorYellow
			}
			cell := tview.NewTableCell(vv).SetTextColor(color)
			cell.SetExpansion(1)
			table.SetCell(i, j, cell)
		}
	}
}

func filesProvider() [][]string {
	title := []string{
		"Name", "Szie", "Kind", "Date Added",
	}
	content := [][]string{title}
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatalf("read dir error: %v", err)
	}
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			log.Fatalf("get file info error: %v", err)
		}
		content = append(content, []string{
			file.Name(), strconv.Itoa(int(info.Size())),
			file.Type().String(), info.ModTime().String(),
		})
	}
	return content
}
