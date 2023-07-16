package executor

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestParallelExecutor(t *testing.T) {
	var resultMux sync.Mutex

	results := []string{}

	producer := func() []int {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	processor := func(item int) (string, error) {
		return fmt.Sprintf("*%v*", item), nil
	}

	consumer := func(item string) error {
		resultMux.Lock()
		results = append(results, item)
		resultMux.Unlock()
		return nil
	}

	monitor := func(update MonitorUpdate) {
		fmt.Printf("Monitor: %v\n", update.JobMap)
	}

	parallelExecutor := ParallelExecutor[int, string]{
		Config:    ParallelExecutorConfig{Buffer: 3, Workers: 2},
		Producer:  producer,
		Processor: processor,
		Consumer:  consumer,
		Monitor:   monitor,
	}

	parallelExecutor.Perform()

	fmt.Printf("Results: <<<%v>>>\n", results)

	wantLen := 10

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

func TestParallelExecutorDemo(t *testing.T) {

	results := []string{}

	producer := func() []int {
		return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	}

	processor := func(item int) (string, error) {
		time.Sleep(110 * time.Millisecond)
		return fmt.Sprintf("*%v*", item), nil
	}

	consumer := func(item string) error {
		time.Sleep(170 * time.Millisecond)
		results = append(results, item)
		return nil
	}

	monitor := func(update MonitorUpdate) {
		fmt.Printf("Monitor: %v\n", update.JobMap)
	}

	parallelExecutor := ParallelExecutor[int, string]{
		Config:    ParallelExecutorConfig{Buffer: 3, Workers: 2},
		Producer:  producer,
		Processor: processor,
		Consumer:  consumer,
		Monitor:   monitor,
	}

	parallelExecutor.Perform()

	fmt.Printf("Results: <<<%v>>>\n", results)

}
