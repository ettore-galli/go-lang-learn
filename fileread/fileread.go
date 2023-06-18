package main

import (
	"github.com/ettore-galli/golang-fileread/csvutils"

	"fmt"
)

func main() {
	mx1, _ := csvutils.ReadCsvMap("./data/input.csv", true)
	fmt.Println(mx1)

	mx2, _ := csvutils.ReadCsvMap("./data/input.csv", false)
	fmt.Println(mx2)
}
