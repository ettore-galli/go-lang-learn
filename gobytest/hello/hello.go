package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/hello/internal/calc"
	"example.com/hello/internal/greeter"
	"moul.io/banner"
)

func Hello() string {
	return banner.Inline(greeter.MakeHelloGreeting())
}

func Greet(b io.Writer, name string) {
	fmt.Fprintf(b, "Hello, %s!", name)
}

func printDemo() {
	fmt.Println(Hello())
	shape := calc.Triangle{Base: 7.0, Height: 8.0}
	fmt.Println(shape, shape.Area())
	Greet(os.Stdout, "Ettore")
}

func webDemo() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

type Delayer interface {
	Sleep()
}

type BasicDelayer struct {
}

func (BasicDelayer) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableDelayer struct {
	sleeper      func(time.Duration)
	delaySeconds int64
}

func (ConfigurableDelayer) delayDuration(seconds int64) time.Duration {
	return time.Duration(seconds) * time.Second
}

func (delayer ConfigurableDelayer) Sleep() {
	delayer.sleeper(delayer.delayDuration(delayer.delaySeconds))
}

func Countdown(b io.Writer, delayer Delayer) {
	const finalWord = "Go!"
	const countdownStart = 3
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(b, "%d\n", i)
		delayer.Sleep()
	}
	fmt.Fprint(b, finalWord)
}

func main() {
	if false {
		webDemo()
	}

	if false {
		printDemo()
	}
	if true {
		Countdown(os.Stdout, &BasicDelayer{})
	}
	// if true {
	// 	Countdown(os.Stdout, &ConfigurableDelayer{delaySeconds: 2, sleeper: time.Sleep})
	// }
}
