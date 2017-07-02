package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func startElement(n *html.Node, hasChild bool) {
	switch n.Type {
	case html.ElementNode:
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, v := range n.Attr {
			fmt.Printf(" %s=%s", v.Key, v.Val)
		}
		if hasChild {
			fmt.Printf(">\n")
		}
		depth++
	case html.CommentNode:
		fmt.Printf("%*s//%s\n", depth*2, "", n.Data)
	case html.TextNode:
		if len(strings.TrimSpace(n.Data)) != 0 {
			fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		}
	}
}
func endElement(n *html.Node, hasChild bool) {
	switch n.Type {
	case html.ElementNode:
		depth--
		if hasChild {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		} else {
			fmt.Printf("/>\n")
		}
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, hasChild bool)) {
	c := n.FirstChild
	if pre != nil {
		pre(n, c != nil)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n, c != nil)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline2: %v\n", err)
		os.Exit(1)
	}

	forEachNode(doc, startElement, endElement)
}
