package executor

import (
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

type parallelExecutorInternals[P any] struct {
	processingQueue    chan P
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
	internals parallelExecutorInternals[P]
}

func (executor *ParallelExecutor[P, M]) monitorEventsThread() {
	for {
		event := <-executor.internals.jobMonitoringQueue
		executor.internals.jobMonitorMap[event.Worker] = event.Status
		executor.Monitor(MonitorUpdate{Message: "", JobMap: executor.internals.jobMonitorMap})
	}
}

func (executor *ParallelExecutor[P, M]) mainWorkThread(ch chan P, wid int) {
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
}

func (executor *ParallelExecutor[P, M]) startWorkers() {
	workersRange := make([]int, executor.Config.Workers)
	for wid := range workersRange {
		go executor.mainWorkThread(executor.internals.processingQueue, wid)
	}
}

func (executor *ParallelExecutor[P, M]) produceData() {
	for _, produced := range executor.Producer() {
		executor.internals.processingQueue <- produced
	}
	close(executor.internals.processingQueue)
}

func (executor *ParallelExecutor[P, M]) Perform() {
	executor.internals.processingQueue = make(chan P, executor.Config.Buffer)
	executor.internals.jobMonitorMap = make(map[int]string)
	executor.internals.jobMonitoringQueue = make(chan WorkerStatus, executor.Config.Workers)

	executorMain := func() {
		go executor.monitorEventsThread()

		executor.internals.workersWaitGroup.Add(executor.Config.Workers)

		go executor.startWorkers()

		go executor.produceData()

		executor.internals.workersWaitGroup.Wait()
	}

	executorMain()

}
