package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func logConfig() {
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("start loging...")
}

func leftTreePanel() *tview.TreeView {
	tree := tview.NewTreeView().
		SetRoot(tview.NewTreeNode("/").
			AddChild(tview.NewTreeNode("/bin").
				AddChild(tview.NewTreeNode("/bin/ls")).
				AddChild(tview.NewTreeNode("/bin/cat")).
				AddChild(tview.NewTreeNode("/bin/rm")).
				AddChild(tview.NewTreeNode("/bin/mkdir")).
				AddChild(tview.NewTreeNode("/bin/mv")).
				AddChild(tview.NewTreeNode("/bin/cp")).
				AddChild(tview.NewTreeNode("/bin/grep")).
				AddChild(tview.NewTreeNode("/bin/sed")).
				AddChild(tview.NewTreeNode("/bin/awk")).
				AddChild(tview.NewTreeNode("/bin/find"))))
	tree.SetBorder(true).SetTitle("Files")
	return tree
}

func fileListPanel() *tview.Table {
	table := tview.NewTable().
		SetBorders(true).
		SetSelectable(true, false).
		SetFixed(1, 1)
	feedTable(table)
	return table
}

func feedTable(table *tview.Table) {
	contents := filesProvicer()
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

func filesProvicer() [][]string {
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

func main() {
	logConfig()
	app := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(leftTreePanel(), 0, 1, false).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Command"), 0, 1, false).
				AddItem(fileListPanel(), 0, 5, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Status"), 2, 1, false),
			0, 3, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
