package ui

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/pluto/pkg/util"
	"github.com/rivo/tview"
)

type LeftPanel struct {
	*tview.TreeView
}

func NewLeftPanel(rootPath string) *LeftPanel {
	return &LeftPanel{
		TreeView: leftTreePanel(rootPath),
	}
}

func leftTreePanel(rootPath string) *tview.TreeView {
	root := tview.NewTreeNode(rootPath).
		SetColor(tcell.ColorRed)
	addTreeNode(root, rootPath)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)
	tree.SetBorder(true).SetTitle("Files")
	tree.SetSelectedFunc(treeSelectionFunc)
	return tree
}

func addTreeNode(root *tview.TreeNode, rootPath string) {
	if !util.IsDir(rootPath) {
		return
	}
	files, err := os.ReadDir(rootPath)
	if err != nil {
		log.Fatalf("read dir error: %v", err)
	}
	for _, file := range files {
		node := tview.NewTreeNode(file.Name()).
			SetReference(filepath.Join(rootPath, file.Name())).
			SetColor(tcell.ColorGreen)
		if file.IsDir() {
			node.SetColor(tcell.ColorRed)
		}
		root.AddChild(node)
	}
}

var onSelectionCallbacks = make(map[string]func(string))

func RegisterOnSelectionCallback(path string, callback func(string)) {
	onSelectionCallbacks[path] = callback
}

func treeSelectionFunc(root *tview.TreeNode) {
	reference := root.GetReference()
	if reference == nil {
		return
	}
	path := reference.(string)
	for _, callback := range onSelectionCallbacks {
		callback(path)
	}
	log.Println("tree selection func: ", path)
	if len(root.GetChildren()) == 0 {
		addTreeNode(root, path)
	} else {
		root.SetExpanded(!root.IsExpanded())
	}
}
