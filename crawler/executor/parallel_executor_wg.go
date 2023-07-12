package executor

import (
	"sync"
)

type ParallelExecutorWgConfig struct {
	Buffer  int
	Workers int
}

type ParallelExecutorWg[P any, M any] struct {
	Config    ParallelExecutorWgConfig
	Producer  func() []P
	Processor func(item P) M
	Consumer  func(item M)
}

func (exe *ParallelExecutorWg[P, M]) Perform() {
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
