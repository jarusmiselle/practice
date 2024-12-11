package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Success!")
		return
	}
	for _, filename := range os.Args[1:] {

		bs, err := os.ReadFile(filename)

		if err != nil {
			panic(err)
		}

		os.Stdout.Write(bs)
	}

}
