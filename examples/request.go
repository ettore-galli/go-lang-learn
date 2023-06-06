package main

import (
	"fmt"
	"io"
	"net/http"
)

func makeRequest(url string) (http.Response, error) {
	resp, err := http.Get(url)
	if err == nil {
		return *resp, nil
	}

	return http.Response{}, err
}

func getContent(url string) string {
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		content, _ := io.ReadAll(resp.Body)
		return string(content)
	}

	return err.Error()
}

func getContentType(url string) string {
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		return resp.Header.Get("Content-type")
	}
	return err.Error()
}

func main() {
	var url string = "https://example.com"

	fmt.Println(getContent(url))
	fmt.Println(getContentType(url))

}
