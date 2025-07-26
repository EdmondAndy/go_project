package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func run() ([]string, error) {
	searchDir := "./"

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})
	
	if e != nil {
		panic(e)
	}

	// for _, file := range fileList {
	// 	fmt.Println(file)
	// }

	return fileList, nil
}

func main() {
	lists, err := run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fileName := "lists.txt"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", fileName, err)
		return
	}
	defer file.Close()

	for _, list := range lists {
		_, err := fmt.Fprintln(file, list)
		if err != nil {
			fmt.Printf("Error writing to file %s: %v\n", fileName, err)
			return
		}
	}
	fmt.Printf("Successfully wrote string slice to %s\n", fileName)
}