package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	bs := make([]byte, 1024)
	file.Read(bs)
	fmt.Println(string(bs))
}
