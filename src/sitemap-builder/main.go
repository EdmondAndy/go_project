package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	link "github.com/edmondandy/go_project/src/html-link-parser"
)

/*
  1. GET the webpage
  2. parse all the links on the page
  3. build proper urls with our links
  4. filter out any links w/ a diff domain
  5. Find all pages (BFS)
  6. print out XML
*/

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "The URL to fetch the sitemap from")
	flag.Parse()

	fmt.Println(*urlFlag)
	resp, err:= http.Get(*urlFlag)
	if err != nil {	
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	links, _ := link.Parse(resp.Body)
	for _, l := range links {
		fmt.Printf("Found link: %s (%s)\n", l.Text, l.Href)
	}
}

