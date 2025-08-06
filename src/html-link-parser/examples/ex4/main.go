package main

import (
	"fmt"
	"log"
	"strings"

	htmllinkparser "github.com/edmondandy/go_project/src/html-link-parser"
)

var exampleHtml = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
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