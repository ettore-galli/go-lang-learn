package executor

import (
	"fmt"
	"testing"
)

func TestLinearExecutor(t *testing.T) {

	results := []string{}

	producer := func() []int {
		return []int{1, 2, 3}
	}

	processor := func(item int) string {
		return fmt.Sprintf("*%v*", item)
	}

	consumer := func(item string) {
		results = append(results, item)
	}

	linearExecutor := LinearExecutor[int, string]{
		Producer:  producer,
		Processor: processor,
		Consumer:  consumer,
	}

	linearExecutor.Perform()

	want := []string{"*1*", "*2*", "*3*"}

	for idx, g := range results {
		w := want[idx]
		if g != w {
			t.Errorf("\nGot.: \n\t%v \nWant:\n\t %v", g, w)
		}
	}

}
