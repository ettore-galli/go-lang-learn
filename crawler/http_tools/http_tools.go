package http_tools

import (
	"io"
	"net/http"
)

type HttpContentGetter struct {
	http_get func(url string) (resp *http.Response, err error)
}

func (hcg *HttpContentGetter) PerformGet(url string) (string, error) {
	resp, err := hcg.http_get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return string(body), nil
}

func NewHttpContentGetter() *HttpContentGetter {
	return &HttpContentGetter{http.Get}
}
