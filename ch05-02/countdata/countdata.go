package countdata

import (
	"io"

	"golang.org/x/net/html"
)

func visit(n *html.Node, result map[string]int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		result[n.Data]++
	}
	visit(n.FirstChild, result)  // 親子
	visit(n.NextSibling, result) // 兄弟
}

func Count(r io.Reader) (map[string]int, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	result := make(map[string]int)
	visit(doc, result)
	return result, nil
}
