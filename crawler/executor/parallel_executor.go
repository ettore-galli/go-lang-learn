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

func (executor *ParallelExecutor[P, M]) Perform() {
	var wg sync.WaitGroup

	processingQueue := make(chan P, executor.Config.Buffer)

	workThread := func(ch chan P, w *sync.WaitGroup, wid int) {
		for {
			produced, ok := <-ch
			if !ok {
				break
			}
			intermediate := executor.Processor(produced)
			executor.Consumer(intermediate)
		}
		w.Done()
	}

	startWorkers := func(workers int, w *sync.WaitGroup) {
		workersRange := make([]int, workers)
		for wid := range workersRange {
			go workThread(processingQueue, &wg, wid)
		}
	}

	produceData := func(procQ chan P) {
		for _, produced := range executor.Producer() {
			procQ <- produced
		}
		close(procQ)
	}

	executorMain := func() {
		wg.Add(executor.Config.Workers)
		go startWorkers(executor.Config.Workers, &wg)
		go produceData(processingQueue)
		wg.Wait()
	}

	executorMain()

}
