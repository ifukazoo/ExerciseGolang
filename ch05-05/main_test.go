package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func Test_countWords(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"word", 1},
		{"word word", 2},

		{"!", 0},
		{"!word", 1},
		{"word!word", 2},
		{"word!word(word)", 3},

		{"!表", 1},
		{"表!図表", 2},
		{"図表!図表(図表)", 3},
	}
	for _, test := range tests {
		if got := countWords(test.input); got != test.want {
			t.Errorf("test(%v) return[%v], want[%v]", test.input, got, test.want)
		}
	}
}
func Test_countWordsAndImages(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{`<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`, 3},
		{`<ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`, 2},
	}
	for _, test := range tests {
		doc, _ := html.Parse(strings.NewReader(test.input))
		if got, _ := countWordsAndImages(doc); got != test.want {
			t.Errorf("test(%v) return[%v], want[%v]", test.input, got, test.want)
		}
	}
}
