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
	workersWaitGroup   sync.WaitGroup
	jobMonitorMap      map[int]string
	jobMonitoringQueue chan WorkerStatus
}

type ParallelExecutor[P any, M any] struct {
	Config    ParallelExecutorConfig
	Producer  func() []P
	Processor func(item P) (M, error)
	Consumer  func(item M) error
	Monitor   func(update MonitorUpdate)
	internals parallelExecutorInternals
}

func (executor *ParallelExecutor[P, M]) monitorEventsThread() {
	for {
		event := <-executor.internals.jobMonitoringQueue
		fmt.Printf("event %v\n", event)
		executor.internals.jobMonitorMap[event.Worker] = event.Status
		executor.Monitor(MonitorUpdate{Message: "", JobMap: executor.internals.jobMonitorMap})
	}
}

func (executor *ParallelExecutor[P, M]) mainWorkThread(ch chan P, wid int) {
	fmt.Printf("Start worker %v\n", wid)
	executor.internals.jobMonitoringQueue <- WorkerStatus{Worker: wid, Status: "start"}
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
	executor.internals.jobMonitoringQueue <- WorkerStatus{Worker: wid, Status: "end"}
	fmt.Printf("End worker %v\n", wid)
}

func (executor *ParallelExecutor[P, M]) startWorkers(workers int, processingQueue chan P) {
	workersRange := make([]int, workers)
	for wid := range workersRange {
		go executor.mainWorkThread(processingQueue, wid)
	}
}

func (executor *ParallelExecutor[P, M]) Perform() {
	processingQueue := make(chan P, executor.Config.Buffer)

	executor.internals.jobMonitorMap = make(map[int]string)
	executor.internals.jobMonitoringQueue = make(chan WorkerStatus, executor.Config.Workers)

	produceData := func(procQ chan P) {
		for _, produced := range executor.Producer() {
			procQ <- produced
		}
		close(procQ)
	}

	executorMain := func() {
		go executor.monitorEventsThread()

		executor.internals.workersWaitGroup.Add(executor.Config.Workers)

		go executor.startWorkers(executor.Config.Workers, processingQueue)

		go produceData(processingQueue)

		executor.internals.workersWaitGroup.Wait()
	}

	executorMain()

}
