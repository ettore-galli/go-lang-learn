package concur

import (
	"net/http"
	"time"
)

type WebsiteChecker func(string) bool

type result = struct {
	url    string
	exists bool
}

func CheckWebsites(wsCheck WebsiteChecker, urls []string) map[string]bool {
	var resultMap = make(map[string]bool)
	var resultChannel = make(chan (result))

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wsCheck(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		res := <-resultChannel
		resultMap[res.url] = res.exists
	}

	return resultMap
}

func tryMakeRequest(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func tryGetAnyResponse(url string) chan struct{} {
	response := make(chan struct{})
	go func() {
		http.Get(url)
		close(response)
	}()
	return response
}

func Racer(alfa string, beta string) string {
	select {
	case <-tryGetAnyResponse(alfa):
		return alfa
	case <-tryGetAnyResponse(beta):
		return beta
	}
}
