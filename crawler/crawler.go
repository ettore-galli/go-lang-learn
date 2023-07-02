package main

import (
	"ettore-galli/crawler/executor"
	"ettore-galli/crawler/http_tools"
	"ettore-galli/crawler/response_writer"
)

func Crawler() {
	producer := func() []string {
		return []string{"https://www.google.it", "https://www.ibm.com"}
	}

	processor := func(url string) string {
		response, err := http_tools.NewHttpContentGetter().PerformGet(url)
		if err != nil {
			return string(err.Error())
		}
		return response
	}

	consumer := func(item string) {
		response_writer.NewResponseWriter().WriteResponse("./out", item)

	}

	linearExecutor := executor.LinearExecutor[string, string]{
		Producer:  producer,
		Processor: processor,
		Consumer:  consumer,
	}
	linearExecutor.Perform()

}

func main() {
	Crawler()
}
