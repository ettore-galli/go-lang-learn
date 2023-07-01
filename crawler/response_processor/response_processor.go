package response_processor

import (
	"fmt"
)

type ResponseProcessor struct{}

func (rp *ResponseProcessor) ProcessResponse(url string, response string) {
	fmt.Printf("\n===== %v =====\n", url)
	fmt.Printf("\n=====ini=====\n%v\n=====end=====\n", response)
}

func NewResponseProcessor() *ResponseProcessor {
	return &ResponseProcessor{}
}
