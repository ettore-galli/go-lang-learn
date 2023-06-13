package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
)

type GetFromUrlResponse struct {
	url         string
	success     bool
	message     string
	contentType string
	body        string
}

func GetFromUrl(url string) GetFromUrlResponse {
	resp, err := http.Get(url)
	if err != nil {
		return GetFromUrlResponse{
			url:     url,
			success: false,
			message: fmt.Sprintf("ERROR: %v", err),
		}
	}
	defer resp.Body.Close()
	cType := resp.Header.Get("content-type")
	body, _ := io.ReadAll(resp.Body)

	return GetFromUrlResponse{
		url:         url,
		success:     true,
		message:     "OK",
		contentType: cType,
		body:        string(body),
	}

}

func (r GetFromUrlResponse) GetShortRepr() string {
	return fmt.Sprintf("%v: %v %v -> %v\n", r.url, r.success, r.message, r.body[:int(math.Min(30, float64(len(r.body))))])
}

func producer(url string) string {
	result := GetFromUrl(url)
	return result.GetShortRepr()
}

func consumer(result string) {
	fmt.Printf("Result %v\n", result)
}

func multiUrl(urls []string) {
	input_queue := make(chan string, 100)
	work_queue := make(chan string, 100)

	for _, url := range urls {
		input_queue <- url
	}

	close(input_queue)

	for url := range input_queue {
		go func(url string) {
			fmt.Printf("\u2756Sending %v\n", url)
			work_queue <- producer(url)
		}(url)
	}

	for range urls {
		fmt.Printf("\u2728Done with %v", <-work_queue)
	}
	close(work_queue)

}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://libero.it",
	}

	multiUrl(urls)

}
