package main

import (
	"fmt"
	"os"
	"strconv"
)

func readdata(file string) (int64, error) {
	dat, err := os.ReadFile(file)

	if err != nil {
		fmt.Printf("errore file: %v\n", err)
		return 0, err
	}
	id, err := strconv.ParseInt(string(dat), 10, 64)
	if err != nil {
		fmt.Printf("errore conv: %v\n", err)
		return 0, err
	}
	return id, nil
}

func main() {
	fmt.Println(readdata("examples/data.txt"))
}
