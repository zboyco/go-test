package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8080/test", nil)
	// req, err := http.NewRequest("GET", "http://localhost:8080/test", bytes.NewBuffer([]byte("hello world")))
	if err != nil {
		panic(err)
	}

	resp, err := (&DigestAuth{
		Username: "john",
		Password: "hello",
		Client:   http.DefaultClient,
	}).Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
