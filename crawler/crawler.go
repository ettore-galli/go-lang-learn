package main

import (
	"ettore-galli/crawler/http_tools"
	"ettore-galli/crawler/response_processor"
)

func main() {
	url := "https://www.google.it"
	cg := http_tools.NewHttpContentGetter()
	rp := response_processor.NewResponseProcessor()
	response, _ := cg.PerformGet(url)
	rp.ProcessResponse(url, response)

}
