package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

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
	urlFlag := flag.String("url", "http://gophercises.com", "The URL to fetch the sitemap from")
	flag.Parse()

	fmt.Println(*urlFlag)
	resp, err:= http.Get(*urlFlag)
	if err != nil {	
		log.Fatal(err)
	}
	defer resp.Body.Close()

/*
   /some-path - add domain to include full url
   https://gophercises.com/some-path - use it as is
   #fragment - ignore it
   mailto:jon@calhoun.io - ignore it
*/
	reqUrl := resp.Request.URL
	fmt.Println(reqUrl.String())
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	links, _ := link.Parse(resp.Body)
	var hrefs []string
	for _, l := range links {
		switch {
			case strings.HasPrefix(l.Href, "/"):
				hrefs = append(hrefs, base+l.Href)
			case strings.HasPrefix(l.Href, "http"):
				hrefs = append(hrefs, l.Href)
			// default:
			// 	fmt.Printf("we are skipping this link: %s\n", l.Href)
		}
	}
	for _, href := range hrefs {
		fmt.Println(href)
	}
}