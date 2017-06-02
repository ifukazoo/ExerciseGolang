package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		word, img, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Printf("count words and images error:%s\n", err)
			os.Exit(1)
		}
		fmt.Printf("%v\t%v\n", word, img)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	} else if n.Type == html.TextNode {
		words += countWords(n.Data)
	}
	w, i := countWordsAndImages(n.FirstChild) // 親子
	words += w
	images += i
	w, i = countWordsAndImages(n.NextSibling) // 兄弟
	words += w
	images += i
	return
}

func scanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading non letter.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if unicode.IsLetter(r) {
			break
		}
	}
	// Scan until letter, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if !unicode.IsLetter(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func countWords(text string) (words int) {
	in := bufio.NewScanner(strings.NewReader(text))
	in.Split(scanWords)
	for in.Scan() {
		words++
	}
	return
}
