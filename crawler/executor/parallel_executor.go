package executor

import (
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

	toBeProcessed := make(chan P, exe.Config.Buffer)

	workThread := func(ch chan P, w *sync.WaitGroup, wid int) {
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

	for wid := range workersRange {
		go workThread(toBeProcessed, &wg, wid)
	}

	data := exe.Producer()

	wg.Add(len(data))
	for _, produced := range data {
		toBeProcessed <- produced
	}
	wg.Wait()

}
