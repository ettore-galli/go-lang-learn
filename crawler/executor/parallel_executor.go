package executor

import (
	"fmt"
	"sync"
)

type ParallelExecutorConfig struct {
	Buffer  int
	Workers int
}

type MonitorUpdate struct {
	Message string
	JobMap  map[int]string
}

type WorkerStatus struct {
	Worker int
	Status string
}

type ParallelExecutor[P any, M any] struct {
	Config    ParallelExecutorConfig
	Producer  func() []P
	Processor func(item P) (M, error)
	Consumer  func(item M) error
	Monitor   func(update MonitorUpdate)
}

func (executor *ParallelExecutor[P, M]) Perform() {
	processingQueue := make(chan P, executor.Config.Buffer)
	var workersWaitGroup sync.WaitGroup
	var jobMonitorMap = make(map[int]string)
	jobMonitoringQueue := make(chan WorkerStatus, executor.Config.Workers)

	monitorEventsThread := func(jobMap map[int]string, jobMonQ chan WorkerStatus) {
		for {
			event := <-jobMonQ
			fmt.Printf("event %v\n", event)
			jobMap[event.Worker] = event.Status
			executor.Monitor(MonitorUpdate{Message: "", JobMap: jobMap})
		}
	}

	workThread := func(ch chan P, w *sync.WaitGroup, wid int, jobMap map[int]string) {
		fmt.Printf("Start worker %v\n", wid)
		jobMonitoringQueue <- WorkerStatus{Worker: wid, Status: "start"}
		for {
			produced, ok := <-ch
			if !ok {
				break
			}

			intermediate, procErr := executor.Processor(produced)
			if procErr != nil {
				continue
			}

			consErr := executor.Consumer(intermediate)
			if consErr != nil {
				continue
			}

		}
		w.Done()
		jobMonitoringQueue <- WorkerStatus{Worker: wid, Status: "end"}
		fmt.Printf("End worker %v\n", wid)
	}

	startWorkers := func(workers int, w *sync.WaitGroup, jobMap map[int]string) {
		workersRange := make([]int, workers)
		for wid := range workersRange {
			go workThread(processingQueue, &workersWaitGroup, wid, jobMap)
		}
	}

	produceData := func(procQ chan P) {
		for _, produced := range executor.Producer() {
			procQ <- produced
		}
		close(procQ)
	}

	executorMain := func() {
		go monitorEventsThread(jobMonitorMap, jobMonitoringQueue)

		workersWaitGroup.Add(executor.Config.Workers)

		go startWorkers(executor.Config.Workers, &workersWaitGroup, jobMonitorMap)

		go produceData(processingQueue)

		workersWaitGroup.Wait()
	}

	executorMain()

}
