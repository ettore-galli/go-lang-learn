package main

import (
	"fmt"

	"example.com/hello/greeter"
)

func Hello() string {
	return greeter.MakeHelloGreeting()
}

func main() {
	fmt.Println(Hello())
}
