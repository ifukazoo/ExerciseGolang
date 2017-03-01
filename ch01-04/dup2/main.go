package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		// stdin
		countLines(os.Stdin, "stdin", counts, names)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
				continue
			}
			countLines(f, filename, counts, names)
			err = f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			}
		}
	}
	for line, n := range counts {
		if 1 < n {
			fmt.Printf("%d\t%s\t%s\n", n, names[line], line)
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]int, names map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		names[input.Text()] = append(names[input.Text()], filename)
	}
}
