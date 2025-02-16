package concur

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func mockWsChecker(url string) bool {
	return url != "www.notworking.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{"www.pippo.it", "www.pluto.it", "www.notworking.com"}
	want := map[string]bool{"www.pippo.it": true, "www.pluto.it": true, "www.notworking.com": false}
	got := CheckWebsites(mockWsChecker, websites)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}

}

func slowStubWsChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {

	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWsChecker, urls)
	}

}

func makeTestServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	slowServer := makeTestServer(10 * time.Millisecond)
	fastServer := makeTestServer(1 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
