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

// func multiUrl(urls []string) {
// 	barrier := make(chan int)
// 	q := make(chan string)

// 	go func(urls []string) {
// 		for _, url := range urls {
// 			fmt.Printf("Sending %v\n", url)
// 			result := GetFromUrl(url)
// 			q <- result.GetShortRepr()
// 		}
// 		close(q)
// 	}(urls)

// 	go func() {
// 		for {
// 			res, ok := <-q
// 			fmt.Printf("Read from q %v\n", res)
// 			if !ok {
// 				fmt.Println("End of queue")
// 				break
// 			}
// 			fmt.Printf("Result %v\n", res)
// 		}
// 		fmt.Println("END LOOP")
// 		barrier <- 0
// 	}()

// 	<-barrier
// }

func multiUrl(urls []string) {
	q := make(chan string, 100)

	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Sending %v\n", url)
			result := GetFromUrl(url)
			q <- result.GetShortRepr()
		}(url)
	}

	for range urls {
		fmt.Printf("Done with %v", <-q)
	}

}

func multiUrlBarrier(urls []string) {
	barrier := make(chan int, 100)
	q := make(chan string)

	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Sending %v\n", url)
			result := GetFromUrl(url)
			q <- result.GetShortRepr()
		}(url)
	}

	for range urls {
		fmt.Printf("Done with %v", <-q)
	}

	<-barrier
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://libero.it",
	}

	multiUrl(urls)

	// fmt.Println(GetFromUrl("https://www.libero.it").GetShortRepr())

}
