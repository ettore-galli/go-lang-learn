package http_tools

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestPerformGet(t *testing.T) {

	sampleHtml := "<http><body><h1>Hello, world!</h1></body></http>"
	expectedHtml := "<http><body><h1>Hello, world!</h1></body></http>"

	ht := HttpContentGetter{http_get: func(url string) (resp *http.Response, err error) {
		return &http.Response{Body: io.NopCloser(strings.NewReader(sampleHtml))}, nil
	}}

	response, _ := ht.PerformGet("any-url")

	if response != expectedHtml {
		t.Errorf("Got.: \n%v \nWant: %v", response, expectedHtml)
	}

}
