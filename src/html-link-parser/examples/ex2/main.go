package main

import (
	"fmt"
	"log"
	"strings"

	htmllinkparser "github.com/edmondanndy/go_project/src/html-link-parser"
)

var exampleHtml = `
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
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