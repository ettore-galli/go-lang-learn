package main

import "fmt"

type FibCalcFunction func(int, FibCalcFunction) int

func fib(n int, recur FibCalcFunction) int {
	fmt.Printf("fib(%v)\n", n)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return recur(n-1, recur) + recur(n-2, recur)
}

func memo(f func(int, FibCalcFunction) int) func(int, FibCalcFunction) int {

	m := make(map[int]int)

	return func(x int, recur FibCalcFunction) int {
		if v, ok := m[x]; ok {
			fmt.Printf("Found in cache: %v\n", x)
			return v
		}
		result := f(x, recur)
		m[x] = result
		return result
	}
}

func main() {
	mfib := memo(fib)
	fmt.Println(memo(fib)(9, mfib))
}
