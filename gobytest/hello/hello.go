package main

import (
	"fmt"

	"example.com/hello/internal/calc"
	"example.com/hello/internal/greeter"
	"moul.io/banner"
)

func Hello() string {
	return banner.Inline(greeter.MakeHelloGreeting())
}

func main() {
	fmt.Println(Hello())
	shape := calc.Triangle{Base: 7.0, Height: 8.0}
	fmt.Println(shape, shape.Area())
}
