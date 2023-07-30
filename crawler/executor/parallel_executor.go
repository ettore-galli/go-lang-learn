package executor

import (
	"fmt"
	"sync"
)

type ParallelExecutorConfig struct {
	Buffer  int
	Workers int
}

type ProcessingLogEntry struct {
	itemIdentifier string
	success        bool
	message        string
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
	processingLog      []ProcessingLogEntry
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

func (executor *ParallelExecutor[P, M]) mainWorkThread(ch chan P, workerId int) {
	executor.internals.jobMonitoringQueue <- WorkerStatus{Worker: workerId, Status: "start"}
	for {

		produced, ok := <-ch

		if !ok {
			executor.internals.processingLog = append(
				executor.internals.processingLog,
				ProcessingLogEntry{
					success: true,
					message: "No more data in queue",
				},
			)
			break
		}

		intermediate, procErr := executor.Processor(produced)

		if procErr != nil {
			executor.internals.processingLog = append(
				executor.internals.processingLog,
				ProcessingLogEntry{
					itemIdentifier: string("produced"),
					success:        false,
					message:        fmt.Sprintf("Processing error: %v", procErr)},
			)
			continue
		}

		consErr := executor.Consumer(intermediate)

		if consErr != nil {
			executor.internals.processingLog = append(
				executor.internals.processingLog,
				ProcessingLogEntry{
					itemIdentifier: string("consumed"),
					success:        false,
					message:        fmt.Sprintf("Consumning error: %v", procErr)},
			)
			continue
		}

	}
	executor.internals.workersWaitGroup.Done()
	executor.internals.jobMonitoringQueue <- WorkerStatus{Worker: workerId, Status: "end"}
}

func (executor *ParallelExecutor[P, M]) startWorkers(numberOfWorkers int) {
	workersRange := make([]int, numberOfWorkers)
	for workerId := range workersRange {
		go executor.mainWorkThread(executor.internals.processingQueue, workerId)
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

	go executor.monitorEventsThread()
	executor.internals.workersWaitGroup.Add(executor.Config.Workers)
	go executor.startWorkers(executor.Config.Workers)
	go executor.produceData()
	executor.internals.workersWaitGroup.Wait()
	fmt.Println(executor.internals.processingLog) // TODO: Replace with something better

}
