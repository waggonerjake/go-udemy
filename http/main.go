package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		panic(err)
	}

	lw := logWriter{}

	//Take something that implements Reader interface (resp.Body) and
	//copy it to something that implements the Writer interface (os.Stdout)
	io.Copy(lw, resp.Body)
}

//Implementing the writer interface (https://pkg.go.dev/io@go1.16.6#Writer)
func (logWriter) Write(bs []byte) (int, error) {
	_, err := fmt.Println(string(bs))
	return len(bs), err
}
