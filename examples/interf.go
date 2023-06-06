package main

import (
	"fmt"
	"os"
	"strings"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Capper struct {
	InnnerWriter Writer
}

func byteUpperCase(p []byte) []byte {
	return []byte(strings.ToUpper(string(p)))
}

func (k Capper) Write(p []byte) (n int, err error) {
	return k.InnnerWriter.Write(byteUpperCase(p))
}

func main() {
	c := Capper{os.Stdout}
	fmt.Fprintln(c, "!Dlrow Olleh")
}
