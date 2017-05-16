package main

import (
	"fmt"
	"log"
	"os"

	"./collectlink"
)

func main() {
	data, err := collectlink.CollectLink(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("%v\n", v)
	}
}
