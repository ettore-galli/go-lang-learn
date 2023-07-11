package executor

import (
	"fmt"
	"sync"
)

type ParallelExecutorConfig struct {
	Buffer  int
	Workers int
}

type ParallelExecutor[P any, M any] struct {
	Config    ParallelExecutorConfig
	Producer  func() []P
	Processor func(item P) M
	Consumer  func(item M)
}

func (exe *ParallelExecutor[P, M]) Perform() {
	var wg sync.WaitGroup
	wg.Add(1)

	toBeProcessed := make(chan P, exe.Config.Buffer)

	workThread := func(ch chan P, w *sync.WaitGroup) {
		for {
			produced, ok := <-ch
			if !ok {
				break
			}

			intermediate := exe.Processor(produced)
			exe.Consumer(intermediate)
			w.Done()
		}

	}

	workersRange := make([]int, exe.Config.Workers)

	for range workersRange {
		go workThread(toBeProcessed, &wg)
	}

	go func() {
		for _, produced := range exe.Producer() {
			fmt.Println(produced)
			toBeProcessed <- produced
			wg.Add(1)
		}
		wg.Done()
	}()

	wg.Wait()

}
