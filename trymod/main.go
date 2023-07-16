package main

import (
	"fmt"

	"github.com/haukened/emojify"
)

func main() {
	fmt.Println("External modules")
	fmt.Println(emojify.Render(":beer:"))
}
