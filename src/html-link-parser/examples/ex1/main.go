package main

import (
	"fmt"
	"log"
	"strings"

	htmllinkparser "github.com/edmondanndy/go_project/src/html-link-parser"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link 
to another page</a>
  <a href="/page-with-query">A link with a query</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := htmllinkparser.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", links)
	// for _, link := range links {
	// 	fmt.Printf("Found link: %s (%s)\n", link.Text, link.Href)
	// }
}