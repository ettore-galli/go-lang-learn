package greeter

import "example.com/hello/internal/iter"

func MakeHelloGreeting() string {
	return iter.RepeatChar("Hello, ", 3) + "world!"
}
