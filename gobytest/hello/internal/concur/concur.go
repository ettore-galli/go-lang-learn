package concur

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
