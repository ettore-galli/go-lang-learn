package concur

import (
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
