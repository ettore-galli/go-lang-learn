package main

import (
	"fmt"
	"strings"
)

func words(sentence string) []string {
	return strings.Fields(sentence)
}

func lowered(strsIn []string) []string {
	lowStrs := make([]string, len(strsIn))
	for idx, element := range strsIn {
		lowStrs[idx] = strings.ToLower(element)
	}

	return lowStrs

}

func mapped(strsIn []string) map[string]int {
	wm := map[string]int{}
	for _, element := range strsIn {
		_, ok := wm[element]
		if !ok {
			wm[element] = 0
		}
		wm[element] = wm[element] + 1

	}
	return wm
}

func main() {
	var sentence string = "The quick Brown Fok jumps over the lazy fox and dog and quick"
	fmt.Println(sentence)
	fmt.Println(&sentence)

	fmt.Println(mapped(lowered(words(sentence))))

}
