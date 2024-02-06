package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	openFile(os.Args[1])
}

func openFile(file string) {
	f, err := os.Open(file)

	if err != nil {
		fmt.Println("An error ocurred: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
