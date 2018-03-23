package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type TreeNode struct {
	Name     string
	Size     int64
	SubNodes []TreeNode
}

func StringTree(object []TreeNode, printFiles bool) (result string) {
	var spaces []bool
	result += outputDirTree(object, spaces, printFiles)
	return
}

func outputDirTree(treeNodes []TreeNode, spaces []bool, printFiles bool) (result string) {
	for index, node := range treeNodes {
		lastNode := (index >= len(treeNodes)-1)
		result += stringLine(node.Name, node.Size, spaces, lastNode, printFiles)

		if len(node.SubNodes) > 0 {
			spacesChild := append(spaces, lastNode)
			result += outputDirTree(node.SubNodes, spacesChild, printFiles)
		}
	}
	return
}

func stringLine(name string, size int64, spaces []bool, last bool, printFiles bool) (result string) {
	for _, space := range spaces {
		if space {
			result += "\t"
		} else {
			result += "│\t"
		}
	}

	indicator := "├───"
	if last {
		indicator = "└───"
	}

	result += indicator + name

	if printFiles && size == 0 {
		result += " (empty)"
	}

	if printFiles && size > 0 {
		result += fmt.Sprintf(" (%vb)", size)
	}

	result += "\n"

	return
}

func walk(dirname string, printFiles bool) ([]TreeNode, error) {
	var items []TreeNode

	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		return nil, err
	}

	for _, file := range files {

		if !file.IsDir() && !printFiles {
			continue
		}

		if !file.IsDir() {
			items = append(items, TreeNode{
				Name: file.Name(),
				Size: file.Size()})
		} else {
			newNode := filepath.Join(dirname, file.Name())
			subNodes, err := walk(newNode, printFiles)
			if err != nil {
				return nil, err
			}

			items = append(items, TreeNode{
				Name:     file.Name(),
				Size:     -1,
				SubNodes: subNodes})
		}
	}
	return items, nil
}

func dirTree(output io.Writer, directory string, printFiles bool) error {
	items, err := walk(directory, printFiles)

	if err != nil {
		return err
	}
	fmt.Fprint(output, StringTree(items, printFiles))
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
