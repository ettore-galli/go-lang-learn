package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

func main() {
	webDemo()
	printDemo()
}
