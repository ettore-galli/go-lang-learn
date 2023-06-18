package clockutils

import (
	"fmt"
	"time"
)

type OutMap map[string]string

type ClockApp struct {
	sleep        func(d time.Duration)
	now          func() time.Time
	print        func(format string, a ...any) (n int, err error)
	loopExecutor func(functionToExecute func())
}

func NewClockApp() *ClockApp {
	return &ClockApp{
		sleep: time.Sleep,
		now:   time.Now,
		print: fmt.Printf,
		loopExecutor: func(functionToExecute func()) {
			for {
				functionToExecute()
			}
		},
	}
}

func (ca ClockApp) PerformClock() {
	fmt.Println(time.Now())
	fmt.Println(time.Date(2023, time.June, 18, 10, 10, 10, 0, time.UTC))
	fmt.Printf("\n")
	ca.loopExecutor(func() {
		currentTime := ca.now()
		ca.print("\r ~~~ ( %v ) ~~~ ", currentTime)
		ca.sleep(100 * time.Millisecond)
	})

}
