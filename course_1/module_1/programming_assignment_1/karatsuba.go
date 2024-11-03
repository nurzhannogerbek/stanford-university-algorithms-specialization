package main

import (
	"fmt"
	"math"
)

// KaratsubaMultiplier - struct containing static methods to perform multiplication using the Karatsuba algorithm.
type KaratsubaMultiplier struct{}

// KaratsubaMultiply - static method to initiate multiplication using the Karatsuba algorithm.
func (KaratsubaMultiplier) KaratsubaMultiply(x, y int) int {
	return karatsuba(x, y)
}

// karatsuba - recursive function implementing the Karatsuba algorithm.
func karatsuba(x, y int) int {
	// Base case: if either number is small enough, perform simple multiplication.
	if x < 100 || y < 100 { // Threshold increased to reduce recursive depth
		return x * y
	}

	// Determine the number of digits in the larger number using a logarithmic approach.
	n := int(math.Max(math.Log10(float64(x))+1, math.Log10(float64(y))+1))
	m := n / 2

	// Calculate the power of ten once to reduce redundant calculations.
	powerOfTenM := int(math.Pow(10, float64(m)))

	// Split x and y into high and low parts.
	highX, lowX := x/powerOfTenM, x%powerOfTenM
	highY, lowY := y/powerOfTenM, y%powerOfTenM

	// Recursively calculate three products.
	z0 := karatsuba(lowX, lowY)
	z2 := karatsuba(highX, highY)
	z1 := karatsuba(lowX+highX, lowY+highY)

	// Combine the results using the stored value `powerOfTenM`.
	return (z2 * powerOfTenM * powerOfTenM) + ((z1 - z2 - z0) * powerOfTenM) + z0
}

func main() {
	x := 3141592653589793238462643383279502884197169399375105820974944592
	y := 2718281828459045235360287471352662497757247093699959574966967627
	result := KaratsubaMultiplier{}.KaratsubaMultiply(x, y)
	fmt.Println("Product:", result)
}
