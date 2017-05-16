package collectlink

import (
	"io"

	"golang.org/x/net/html"
)

func visit(n *html.Node, result []string) []string {
	if n == nil {
		return result
	}
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "src" {
				result = append(result, a.Val)
			}
		}
	}
	result = visit(n.FirstChild, result)  // 親子
	result = visit(n.NextSibling, result) // 兄弟
	return result
}

func CollectLink(r io.Reader) ([]string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var result []string
	result = visit(doc, result)
	return result, nil
}
