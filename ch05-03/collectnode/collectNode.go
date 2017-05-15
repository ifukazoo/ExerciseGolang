package collectnode

import (
	"io"
	"regexp"

	"golang.org/x/net/html"
)

var emptynode *regexp.Regexp

func init() {
	emptynode = regexp.MustCompile(`(?m)^\s+$`)
}

func isUselessNode(n *html.Node) bool {
	return n.Parent != nil && (n.Parent.Data == "script" || n.Parent.Data == "style")
}

func visit(n *html.Node, result []string) []string {
	if n == nil {
		return result
	}
	if n.Type == html.TextNode && !isUselessNode(n) {
		//空白のみのTextNodeは除外する
		if !emptynode.MatchString(n.Data) {
			result = append(result, n.Data)
		}
	}
	result = visit(n.FirstChild, result)  // 親子
	result = visit(n.NextSibling, result) // 兄弟
	return result
}

func CollectText(r io.Reader) ([]string, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var result []string
	result = visit(doc, result)
	return result, nil
}
