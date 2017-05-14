package main

import (
	"fmt"
	"log"
	"os"

	"./countdata"
)

func main() {
	data, err := countdata.Count(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", data)
}
