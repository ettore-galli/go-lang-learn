package main

import (
	"ettore-galli/crawler/executor"
	"ettore-galli/crawler/http_tools"
	"ettore-galli/crawler/response_writer"
	"fmt"
)

func Crawler() {
	producer := func() []string {
		return []string{"https://www.google.it", "https://www.ibm.com"}
	}

	processor := func(url string) (string, error) {
		response, err := http_tools.NewHttpContentGetter().PerformGet(url)
		if err != nil {
			return "", err
		}
		return response, nil
	}

	consumer := func(item string) error {
		response_writer.NewResponseWriter().WriteResponse("./out", item)
		return nil
	}

	monitor := func(update executor.MonitorUpdate) {
		fmt.Printf("Monitor: %v\n", update.JobMap)
	}

	config := executor.ParallelExecutorConfig{Buffer: 3, Workers: 2}

	executor := executor.ParallelExecutor[string, string]{
		Config:    config,
		Producer:  producer,
		Processor: processor,
		Consumer:  consumer,
		Monitor:   monitor,
	}
	executor.Perform()

}

func main() {
	Crawler()
}
