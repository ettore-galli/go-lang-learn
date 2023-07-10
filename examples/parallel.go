package main

import "fmt"

func ExecuteParallel[T any, U any]() {

}

func main() {

	ch := make(chan int)
	ctl := make(chan int)

	go func(cx chan<- int) {
		for v := range [10]int{} {
			cx <- v
			fmt.Printf("After send %v\n", v)
		}
		close(cx)
	}(ch)

	go func(cx <-chan int, ctl chan<- int) {
		for data := range cx {
			fmt.Printf("Received %v\n", data)
		}
		ctl <- 0
	}(ch, ctl)

	<-ctl
	close(ctl)

}
