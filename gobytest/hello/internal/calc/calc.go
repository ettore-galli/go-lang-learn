package calc

const ListSize int = 7

func Sum(a int, b int) int {
	return a + b
}

func SumList(numbers [ListSize]int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
