package main

import (
	"fmt"
	"io"
	"os"
	"sort"
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

func dirTree(out io.Writer, path string, printFiles bool) error {
	list, err := readDir(path)
	if err != nil {
		return err
	}

	for _, file := range list {
		if file.IsDir() {
			fmt.Println(file.Name())
			err := dirTree(out, path+"/"+file.Name(), printFiles)
			if err != nil {
				return err
			}

		} else if printFiles {
			fmt.Println(file.Name())
		}
	}

	return nil
}

func readDir(path string) ([]os.FileInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	return files, nil
}
