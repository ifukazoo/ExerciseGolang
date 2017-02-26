package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	longString = "gggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"
	sepapator  = "#"
)

var manyStrings []string

func init() {
	manyStrings = strings.Split(longString, "")
}

func measureLostTime(f func() string) int64 {
	start := time.Now()
	f()
	return time.Since(start).Nanoseconds()
}

func inefficientWork() string {
	var joined string
	for _, c := range manyStrings {
		joined += c
		joined += sepapator
	}
	return joined
}
func useLibraryWork() string {
	return strings.Join(manyStrings, sepapator)
}

func main() {
	slow := measureLostTime(inefficientWork)
	fast := measureLostTime(useLibraryWork)
	fmt.Printf("slow version %v nsec\n", slow)
	fmt.Printf("fast version %v nsec\n", fast)
}
