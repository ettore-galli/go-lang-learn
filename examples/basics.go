package main

import (
	"fmt"
)

func PrintType(object any) {
	fmt.Printf("%v, %T\n", object, object)
}

func FizzBuzz(maxn int) {
	for i := 0; i < maxn; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizz buzz")

		case i%3 == 0:
			fmt.Println("fizz")

		case i%5 == 0:
			fmt.Println("buzz")

		default:
			fmt.Println(i)

		}

	}
}

func CharSet() {
	base := 100000
	for i := base; i < base+1000; i++ {
		if i%78 == 0 {
			fmt.Print("\n")
		}
		fmt.Print(fmt.Sprintf("%s", string(rune(i))))

	}
}

func EvenEnded() {
	startNum := 1000
	endNum := 9999
	evenEndeds := 0
	for alfa := startNum; alfa <= endNum; alfa++ {
		for beta := alfa; beta <= endNum; beta++ {
			strnum := fmt.Sprintf("%d", alfa*beta)
			ending := strnum[len(strnum)-1:]
			if ending == "0" || ending == "2" || ending == "4" || ending == "6" || ending == "8" {
				evenEndeds++
			}

		}
	}
	fmt.Println(evenEndeds)
}

func ternary(condition bool, alfa any, beta any) any {
	if condition {
		return alfa
	}
	return beta
}

func MaxVal(items []int) int {
	if len(items) == 0 {
		return 0
	}
	if len(items) == 1 {
		return items[0]
	}
	remMax := MaxVal(items[1:])
	return ternary(items[0] > remMax, items[0], remMax).(int)
}

func SlicesDemo() {
	var list []string

	list = append(list, "a", "cicci", "kj")

	fmt.Println(list)

	for i, s := range list {
		fmt.Printf("%v %v\n", i, s)
	}
}

func main() {
	// var hello string

	// hello = "Hello, world! \u5344"

	// PrintType(hello)

	// FizzBuzz(20)

	// CharSet()

	// EvenEnded()

	// SlicesDemo()
	values := []int{1, 4, 3, 9, 1, 4}
	fmt.Println(MaxVal(values))
}
