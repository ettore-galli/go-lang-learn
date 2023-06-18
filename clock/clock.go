package main

import (
	"github.com/ettore-galli/golang-clock/clockutils"
)

func main() {
	ca := clockutils.NewClockApp()
	ca.PerformClock()

}
