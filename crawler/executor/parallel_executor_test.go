package executor

import (
	"fmt"
	"testing"
)

func Contains[T comparable](slice []T, element T) bool {
	for _, g := range slice {
		if g == element {
			return true
		}
	}
	return false
}

func TestParallelExecutor(t *testing.T) {
	results := []string{}

	producer := func() []int {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	processor := func(item int) string {
		return fmt.Sprintf("*%v*", item)
	}

	consumer := func(item string) {
		fmt.Println(item)
		results = append(results, item)
	}

	parallelExecutor := ParallelExecutor[int, string]{
		Config:    ParallelExecutorConfig{Buffer: 3, Workers: 2},
		Producer:  producer,
		Processor: processor,
		Consumer:  consumer,
	}

	parallelExecutor.Perform()

	wantLen := 10

	fmt.Printf("<<%v>>\n", results)

	if len(results) != wantLen {
		t.Errorf("Result expected to have length %v, got %v", wantLen, len(results))
	}

	wantItems := []string{"*0*", "*1*", "*2*", "*3*", "*4*", "*5*", "*6*", "*7*", "*8*", "*9*"}

	for _, g := range wantItems {
		isPresent := Contains(results, g)
		if !isPresent {
			t.Errorf("\n%v not in results", g)
		}
	}

}
