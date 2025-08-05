package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/edmondandy/go_project/src/cyoa"
)

func main() {
	file := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()

	d := json.NewDecoder(f)
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}
}