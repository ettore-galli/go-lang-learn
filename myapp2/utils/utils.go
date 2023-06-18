package utils

import (
	"fmt"
)

func Somma100(a int, b int) int {
	return 100 * (a + b)
}

func FormatSomma100(a int, b int) string {
	return fmt.Sprintf("Somma100(%v, %v) => %v", a, b, Somma100(a, b))
}

func StampaSomma100(a int, b int) {
	fmt.Printf("Somma100(%v, %v) => %v", a, b, Somma100(a, b))
}
