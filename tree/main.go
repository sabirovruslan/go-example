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

func readRecursivelyDir(out io.Writer, path string, printFiles bool, prefix string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return
	}

	if printFiles != true {
		files = removeFiles(files)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for i, file := range files {
		nameFmt := fmt.Sprintf("%v├───%v", prefix, file.Name())
		prefixNext := prefix + "│\t"
		if len(files)-1 == i {
			prefixNext = prefix + "\t"
			nameFmt = fmt.Sprintf("%v└───%v", prefix, file.Name())
		}
		if file.IsDir() {
			_, _ = fmt.Fprintln(out, nameFmt)
			err = readRecursivelyDir(out, path+"/"+file.Name(), printFiles, prefixNext)
			if err != nil {
				return
			}
		} else if printFiles {
			size := "empty"
			if file.Size() > 0 {
				size = fmt.Sprintf("%vb", file.Size())
			}
			nameFmt = fmt.Sprintf("%v (%v)", nameFmt, size)
			_, _ = fmt.Fprintln(out, nameFmt)
		}
	}
	return
}

func removeFiles(list []os.FileInfo) (res []os.FileInfo) {
	for _, f := range list {
		if f.IsDir() {
			res = append(res, f)
		}
	}
	return
}
