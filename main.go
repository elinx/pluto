package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func logConfig() {
	f, err := os.OpenFile("pluto.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("start loging...")
}

func leftTreePanel() *tview.TreeView {
	root := tview.NewTreeNode("/").
		SetColor(tcell.ColorRed)
	addTreeNode(root, ".")
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)
	tree.SetBorder(true).SetTitle("Files")
	tree.SetSelectedFunc(treeSelectionFunc)
	return tree
}

func treeSelectionFunc(root *tview.TreeNode) {
	reference := root.GetReference()
	if reference == nil {
		return
	}
	path := reference.(string)
	log.Println("tree selection func: ", path)
	if len(root.GetChildren()) == 0 {
		addTreeNode(root, path)
	} else {
		root.SetExpanded(!root.IsExpanded())
	}
}

func addTreeNode(root *tview.TreeNode, rootPath string) {
	log.Println("add tree node: ", rootPath)
	files, err := os.ReadDir(rootPath)
	if err != nil {
		log.Fatalf("read dir error: %v", err)
	}
	for _, file := range files {
		node := tview.NewTreeNode(file.Name()).
			SetReference(filepath.Join(rootPath, file.Name())).
			SetSelectable(file.IsDir()).
			SetColor(tcell.ColorGreen)
		if file.IsDir() {
			node.SetColor(tcell.ColorRed)
		}
		root.AddChild(node)
		// addTreeNode(node, file.Name())
	}
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

func addLayout(app *tview.Application) *tview.Flex {
	leftTreeView := leftTreePanel()
	flex := tview.NewFlex().
		AddItem(leftTreeView, 0, 1, true).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Command"), 0, 1, false).
				AddItem(fileListPanel(), 0, 5, true).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Status"), 2, 1, false),
			0, 3, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	return flex
}

func main() {
	logConfig()
	app := tview.NewApplication()
	layout := addLayout(app)
	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}
