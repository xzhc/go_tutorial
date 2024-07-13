package main

import "fmt"

//Declare a type constraint
type Number interface {
	int64 | float64
}

func main() {
	//Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	//Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))
	fmt.Printf("Generic sums: %v and %v\n", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))
	fmt.Printf("type parameters infered: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	fmt.Printf("Generic sums with type constraint: %v and %v\n", SumNumbers(ints), SumNumbers(floats))

}

// SumInts adds together the values of m
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}

//SumFloats adds together the values of m

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}

	return s
}

//SumIntsOrFloats sums the value of map m .It supports both int64 and float64
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

//SumNumbers sums the values of map m. It supports both int64 and float64
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}
