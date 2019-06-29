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
	return readRecursivelyDir(out, path, printFiles, "")
}

func readRecursivelyDir(out io.Writer, path string, printFiles bool, level string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	files, err := f.Readdir(-1)
	if err != nil {
		return err
	}

	if printFiles != true {
		files = removeFiles(files)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for i, file := range files {
		nameFmt := file.Name()
		levelFmt := level + "│\t"
		if len(files)-1 == i {
			levelFmt = level + "\t"
			nameFmt = fmt.Sprintf("%v└───%v", level, nameFmt)
		} else {
			nameFmt = fmt.Sprintf("%v├───%v", level, nameFmt)
		}
		if file.IsDir() {
			fmt.Println(nameFmt)
			err := readRecursivelyDir(out, path+"/"+file.Name(), printFiles, levelFmt)
			if err != nil {
				return err
			}
		} else if printFiles {
			size := "empty"
			if file.Size() > 0 {
				size = fmt.Sprintf("%vb", file.Size())
			}
			fmt.Printf("%v (%v)\n", nameFmt, size)
		}
	}
	return err
}

func removeFiles(list []os.FileInfo) (res []os.FileInfo) {
	for _, f := range list {
		if f.IsDir() {
			res = append(res, f)
		}
	}
	return
}
