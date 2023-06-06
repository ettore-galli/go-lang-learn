package main

import (
	"fmt"
)

func producer(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Printf("Produced: %d\n", i)

	}
	close(c)
}

func consumer(c chan int, ctl chan int) {
	for {
		item, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Consumed: %d\n", item)
	}
	ctl <- 1
}

func main() {
	c := make(chan int)
	ctl := make(chan int)
	go producer(c)
	go consumer(c, ctl)
	<-ctl
}
