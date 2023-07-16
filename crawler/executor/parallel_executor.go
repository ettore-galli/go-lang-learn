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

type parallelExecutorInternals struct {
	workersWaitGroup sync.WaitGroup
}

type ParallelExecutor[P any, M any] struct {
	Config    ParallelExecutorConfig
	Producer  func() []P
	Processor func(item P) (M, error)
	Consumer  func(item M) error
	Monitor   func(update MonitorUpdate)
	internals parallelExecutorInternals
}

func (executor *ParallelExecutor[P, M]) monitorEventsThread(jobMap map[int]string, jobMonQ chan WorkerStatus) {
	for {
		event := <-jobMonQ
		fmt.Printf("event %v\n", event)
		jobMap[event.Worker] = event.Status
		executor.Monitor(MonitorUpdate{Message: "", JobMap: jobMap})
	}
}

func (executor *ParallelExecutor[P, M]) mainWorkThread(ch chan P, wid int, jobMonitoringQueue chan WorkerStatus, jobMap map[int]string) {
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
	executor.internals.workersWaitGroup.Done()
	jobMonitoringQueue <- WorkerStatus{Worker: wid, Status: "end"}
	fmt.Printf("End worker %v\n", wid)
}

func (executor *ParallelExecutor[P, M]) Perform() {
	processingQueue := make(chan P, executor.Config.Buffer)

	var jobMonitorMap = make(map[int]string)
	jobMonitoringQueue := make(chan WorkerStatus, executor.Config.Workers)

	startWorkers := func(workers int, jobMap map[int]string) {
		workersRange := make([]int, workers)
		for wid := range workersRange {
			go executor.mainWorkThread(processingQueue, wid, jobMonitoringQueue, jobMap)
		}
	}

	produceData := func(procQ chan P) {
		for _, produced := range executor.Producer() {
			procQ <- produced
		}
		close(procQ)
	}

	executorMain := func() {
		go executor.monitorEventsThread(jobMonitorMap, jobMonitoringQueue)

		executor.internals.workersWaitGroup.Add(executor.Config.Workers)

		go startWorkers(executor.Config.Workers, jobMonitorMap)

		go produceData(processingQueue)

		executor.internals.workersWaitGroup.Wait()
	}

	executorMain()

}
