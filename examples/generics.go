package main

import "fmt"

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

type GeneralNumber interface {
	int |
		int8 |
		int16 |
		int32 |
		int64 |
		float32 |
		float64
}

func SumMap[T GeneralNumber](m map[string]T) T {
	var s T
	for _, v := range m {
		s += v
	}
	return s
}

func limap(lista []int, fu func(int) int) []int {
	transformed := make([]int, len(lista))
	for i, v := range lista {
		transformed[i] = fu(v)
	}
	return transformed
}

func transform(x int) int {
	return 2*x + 1
}

func transformfloat(x float64) float64 {
	return (3*x + 1) / 2
}

func limapgen[T any](lista []T, fu func(T) T) []T {
	transformed := make([]T, len(lista))
	for i, v := range lista {
		transformed[i] = fu(v)
	}
	return transformed
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v, %v\nGeneric Sums ...: %v, %v\n",
		SumInts(ints),
		SumFloats(floats),
		SumMap(ints),
		SumMap(floats),
	)

	fmt.Println(limap([]int{1, 2, 3, 4}, transform))
	fmt.Println(limapgen([]float64{1, 2, 2.1, 3, 4}, transformfloat))
}
