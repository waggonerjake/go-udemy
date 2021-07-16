package main

import (
	"io"
	"os"
)

func main() {
	if file, err := os.Open(os.Args[1]); err != nil {
		panic(err)
	} else {
		io.Copy(os.Stdout, file)
	}
}
