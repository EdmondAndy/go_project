package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	paths := os.Args[1:]
	if len(paths) == 0 {
		fmt.Println("Please provide at least one directory path.")
		return
	}

	var dirs []byte

	for _, dir := range paths {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Printf("Error reading directory %s: %v\n", dir, err)
			return
		}

		dirs = append(dirs, dir...)
		dirs = append(dirs, '\n')

		for _, file := range files {
			if file.IsDir() {
				dirs = append(dirs, '\t')
				dirs = append(dirs, file.Name()...)
				dirs = append(dirs, '/', '\n')
			}
		}

		dirs = append(dirs, '\n')
	}

	err := ioutil.WriteFile("directories.txt", dirs, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}