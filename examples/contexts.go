package main

import (
	"context"
	"fmt"
	"time"
)

func upToInfinite(ctx context.Context) <-chan int {

	outChan := make(chan int, 100)

	go func(c chan int, cnt context.Context) {
		n := 0

		for {
			select {
			case c <- n:
				n++
			case k := <-cnt.Done():
				fmt.Printf("k=%v %T\n", k, k)
				close(c)
				return
			}
		}

	}(outChan, ctx)

	return outChan
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	for v := range upToInfinite(ctx) {
		fmt.Println(v)
	}

}
