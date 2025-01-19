package main

import (
	"fmt"

	"example.com/hello/internal/greeter"
	"moul.io/banner"
)

func Hello() string {
	return banner.Inline(greeter.MakeHelloGreeting())
}

func main() {
	fmt.Println(Hello())
}
