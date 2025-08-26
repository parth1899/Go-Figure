package main

import (
	"fmt"
)

// adds the values of map m
// two non generic functions
func SumInts(m map[string]int64) int64 {
	var sum int64

	for _, v := range m {
		sum += v
	}

	return sum
}

func SumFloats(m map[string]float64) float64 {
	var sum float64

	for _, v := range m {
		sum += v
	}

	return sum
}

// one generic function to avoid writing two functions
// we need type parameters + ordinary function parameters
// we call the function with type arguments + ordinary function arguments
// each type parameter has a "type constraint"
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	// Go requires that map keys be comparable. So declaring K as comparable is necessary so you can use K as the key in the map variable

	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

// declaring "AllowedFormats" interface type to use as a type Constraint
type AllowedFormats interface {
	int64 | float64
}

// use: func SumIntsOrFloats[K comparable, V AllowedFormats](m map[K]V) V

func main() {
	// initialise the maps
	ints := map[string]int64{
		"first":  12,
		"second": 13,
	}

	floats := map[string]float64{
		"first":  12.12,
		"second": 13.09,
	}

	fmt.Printf("Non Generic Sums: %v and %v\n", SumInts(ints), SumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n", SumIntsOrFloats[string, int64](ints), SumIntsOrFloats[string, float64](floats))
	// we can also omit the type arguments, The compiler infers type arguments from the types of function arguments
	fmt.Printf("Generic Sums: %v and %v (type parameters are inferred)\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
	// Note that this isn’t always possible. For example, if you needed to call a generic function that had no arguments, you would need to include the type arguments in the function call.

	// Declaring Type Constraints
	// declare a type constraint as an interface

}
