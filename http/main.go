package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error occurred:", err)
		return
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	n := len(bs)
	fmt.Println(string(bs))
	fmt.Printf("wrote %v bytes\n", n)
	return n, nil
}
