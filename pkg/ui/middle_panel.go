package ui

import (
	"log"
	"os"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/pluto/pkg/util"
	"github.com/rivo/tview"
)

const (
	DirectoryView = "Directory"
	FileView      = "File"
)

type MiddlePanel struct {
	*tview.Flex
	commandPanel *tview.Box
	mainPanel    *tview.Pages
	statusPanel  *tview.Box
}

func NewMiddlePanel() *MiddlePanel {
	commandPanel := tview.NewBox().SetBorder(true).SetTitle("Command")
	directoryView := fileListPanel()
	fileContentView := tview.NewTextView()
	fileContentView.SetBorder(true)
	mainPanel := tview.NewPages().
		AddPage(DirectoryView, directoryView, true, true).
		AddPage(FileView, fileContentView, true, true)
	statusPanel := tview.NewBox().SetBorder(true).SetTitle("Status")
	middlePanel := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(commandPanel, 0, 1, false).
		AddItem(mainPanel, 0, 5, true).
		AddItem(statusPanel, 2, 1, false)
	RegisterOnSelectionCallback("middle", func(rootPath string) {
		if util.IsDir(rootPath) {
			directoryView.Clear()
			feedTable(directoryView, rootPath)
			mainPanel.SwitchToPage(DirectoryView)
		} else {
			fileContentView.Clear()
			feedContent(fileContentView, rootPath)
			mainPanel.SwitchToPage(FileView)
		}
	})
	return &MiddlePanel{
		Flex:         middlePanel,
		commandPanel: commandPanel,
		mainPanel:    mainPanel,
		statusPanel:  statusPanel,
	}
}

func fileListPanel() *tview.Table {
	style := tcell.Style{}.
		Background(tcell.ColorDimGray).
		Foreground(tcell.ColorPurple).
		Attributes(0)

	table := tview.NewTable().
		SetSelectable(true, false).
		SetFixed(1, 1).
		SetSelectedStyle(style)
	table.SetBorder(true)
	return table
}

func feedTable(table *tview.Table, rootPath string) {
	contents := filesProvider(rootPath)
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

func filesProvider(rootPath string) [][]string {
	title := []string{
		"Name", "Szie", "Kind", "Date Added",
	}
	content := [][]string{title}
	files, err := os.ReadDir(rootPath)
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

func feedContent(contentView *tview.TextView, rootPath string) {
	content, err := os.ReadFile(rootPath)
	if err != nil {
		log.Fatal(err)
	}
	contentView.SetText(string(content))
}
