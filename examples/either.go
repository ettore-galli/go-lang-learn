package main

import (
	"fmt"
)

type MagicResult interface {
	GetNumbers() []int
	GetMagicWord() string
}

type MagicNumbers struct {
	magicWord string
	numbers   []int
}

func (mn MagicNumbers) GetNumbers() []int {
	return mn.numbers
}

func (mn MagicNumbers) GetMagicWord() string {
	return mn.magicWord
}

type NoMagic struct {
	magicWord string
}

func (mn NoMagic) GetNumbers() []int {
	return []int{}
}

func (mn NoMagic) GetMagicWord() string {
	return mn.magicWord
}

func isMagic(mr MagicResult) bool {
	_, ok := mr.(MagicNumbers)
	return ok
}

func getMagicNumbers(magicWord string) MagicResult {
	if magicWord == "abracadabra" {
		return MagicNumbers{magicWord, []int{4, 8, 15, 16, 23, 42}}
	} else {
		return NoMagic{magicWord}
	}

}

func processMagic(mr MagicResult) string {
	if isMagic(mr) {
		return fmt.Sprintf("%v: It's a kind of Magic %v \n", mr.GetMagicWord(), mr.GetNumbers())
	} else {
		return fmt.Sprintf("%v: No magic", mr.GetMagicWord())
	}
}

func main() {
	fmt.Println("EITHER")
	fmt.Println(processMagic(getMagicNumbers("abracadabra")))
	fmt.Println(processMagic(getMagicNumbers("boo")))
}
