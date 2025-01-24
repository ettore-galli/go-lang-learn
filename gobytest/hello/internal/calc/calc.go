package calc

import "math"

const ListSize int = 7

func Sum(a int, b int) int {
	return a + b
}

func SumArray(numbers [ListSize]int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumSlice(numbers []int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumSlices(slices ...[]int) []int {
	var sums []int
	for _, slice := range slices {
		sums = append(sums, SumSlice(slice))
	}
	return sums
}

func SumTails(slices ...[]int) []int {
	var sums []int
	for _, slice := range slices {
		var sum int
		if len(slice) > 1 {
			sum = SumSlice(slice[1:])
		} else {
			sum = 0
		}
		sums = append(sums, sum)
	}
	return sums
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 2 * 3.14 * math.Pow(c.radius, 2.0)
}

func Perimeter(rectangle Rectangle) float64 {
	return 2.0 * (rectangle.width + rectangle.height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
