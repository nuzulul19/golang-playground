package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// run command `go run main.go <filename>` e.g `go run main.go myfile.txt`
	file := os.Args[1]

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
