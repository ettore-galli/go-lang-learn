package clockutils

import (
	"fmt"
	"testing"
	"time"
)

func TestPerformClock(t *testing.T) {
	var printCollector string

	ca := ClockApp{
		sleep: func(_ time.Duration) { time.Sleep(0) },
		now: func() time.Time {
			return time.Date(2023, time.June, 18, 10, 10, 10, 0, time.UTC)
		},
		print: func(format string, a ...any) (n int, err error) {
			printCollector = fmt.Sprintf(format, a)
			return 0, nil
		},
		loopExecutor: func(functionToExecute func()) {
			functionToExecute() // Dummy executes once
		},
	}

	ca.PerformClock()

	want := "\r ~~~ ( [2023-06-18 10:10:10 +0000 UTC] ) ~~~ "
	if printCollector != want {
		t.Errorf("Got.: \n%v (%T) \nWant: %v (%T)", printCollector, printCollector, want, want)
	}

}
