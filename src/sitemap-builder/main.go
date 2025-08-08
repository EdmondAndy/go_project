package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
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

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "http://gophercises.com", "The URL to fetch the sitemap from")
	maxDepth := flag.Int("depth", 10, "The maximum depth to traverse the sitemap")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)

	toXml := urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		urlStr: struct{}{},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		}
		for url, _ := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}

			for _, link := range get(url) {
				if _, ok := seen[link]; !ok {
					nq[link] = struct{}{}
				}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlstr string) []string {
	resp, err:= http.Get(urlstr)
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
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(resp.Body, base), withPrefix(base)) // filter out links that do not start with base
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
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
	return hrefs
}

func filter(links []string, keepFn func(string) bool) []string {
	var filtered []string
	for _, link := range links {
		if keepFn(link) {
			filtered = append(filtered, link)
		}
	}
	return filtered
}

func withPrefix(pfx string) func(string) bool {
	return func(s string) bool {
		return strings.HasPrefix(s, pfx)
	}
}