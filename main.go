package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

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

func dirTree(out io.Writer, root string, printFiles bool) error {
	indent := ""
	var buffer bytes.Buffer
	err := treeWalk(&buffer, root, indent, printFiles)
	if err != nil {
		return fmt.Errorf("could not stat %s: %v", root, err)
	}
	fmt.Fprint(out, buffer.String())
	return nil
}

func treeWalk(out io.Writer, root, indent string, printFiles bool) error {

	fi, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("could not stat %s: %v", root, err)
	}

	if !fi.IsDir() {
		return nil
	}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("could not read dir %s: %v", root, err)
	}

	var filesList []os.FileInfo

	if !printFiles {
		for _, file := range files {
			if file.IsDir() {
				filesList = append(filesList, file)
			}
		}
	} else {
		filesList = files
	}

	for i, fi := range filesList {

		var add string
		if i == len(filesList)-1 {
			fmt.Fprint(out, indent+"└───")
			add = "\t"
		} else {
			fmt.Fprint(out, indent+"├───")
			add = "│\t"
		}

		fileSize := ""
		if printFiles && !fi.IsDir() {
			fileSize = " (empty)"
			if size := fi.Size(); size > 0 {
				fileSize = fmt.Sprintf(" (%vb)", size)
			}
		}
		line := fmt.Sprintf("%s%s", fi.Name(), fileSize)
		fmt.Fprintln(out, line)
		if err := treeWalk(out, filepath.Join(root, fi.Name()), indent+add, printFiles); err != nil {
			return err
		}
	}
	return nil
}
