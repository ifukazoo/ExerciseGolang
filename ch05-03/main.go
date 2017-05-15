package main

import (
	"fmt"
	"log"
	"os"

	"./collectnode"
)

func main() {
	data, err := collectnode.CollectText(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("%v\n", v)
	}
}
