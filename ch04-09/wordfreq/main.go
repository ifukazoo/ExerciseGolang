package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
		// これだと記号でつながった単語も1単語となってしまう.
		// 単語を分割するなら,はじめからbufio.ScanWordsをセットする
		// 必要がなくなるので,これはこのままとする.

	}
	if err := in.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	for w, n := range counts {
		fmt.Printf("%s\t%d\n", w, n)
	}
}
