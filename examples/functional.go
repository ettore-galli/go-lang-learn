package main

import (
	"fmt"
)

func fmap(input string) string {
	return input
}

func main() {
	var input string = "https://example.com"

	fmt.Println(fmap(input))

}
