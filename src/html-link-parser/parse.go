package htmllinkparser

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	// dfs(doc, "")
	// 1. Find <a> nodes in document
	// 2. for each link node...
	//   2.a build a Link
	// 3. return the Links
	return links, nil
}

func buildLink(n *html.Node) Link {
	var link Link
	for _, attr := range n.Attr{
		if attr.Key == "href" {
			link.Href = attr.Val
			break
		}
	}
	link.Text = text(n)
	return link
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	var links []*html.Node
	if n.Type == html.ElementNode && n.Data == "a" {
		links = append(links, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, linkNodes(c)...)
	}
	return links
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	for c:= n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}