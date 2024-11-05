package main

import (
	"fmt"
	"math/big"
)

// KaratsubaMultiplier - struct containing static methods to perform multiplication using the Karatsuba algorithm.
type KaratsubaMultiplier struct{}

// KaratsubaMultiply - static method to initiate multiplication using the Karatsuba algorithm for big.Int values.
func (KaratsubaMultiplier) KaratsubaMultiply(x, y *big.Int) *big.Int {
	return karatsuba(x, y)
}

// karatsuba - recursive function implementing the Karatsuba algorithm for big.Int values.
func karatsuba(x, y *big.Int) *big.Int {
	// Set a threshold for switching to simple multiplication to avoid deep recursion.
	const threshold = 10

	// Base case: if either number has fewer than "threshold" bits, use simple multiplication.
	if x.BitLen() < threshold || y.BitLen() < threshold {
		result := new(big.Int)
		return result.Mul(x, y)
	}

	// Determine the number of digits in the larger number.
	n := x.BitLen()
	if y.BitLen() > n {
		n = y.BitLen()
	}
	m := n / 2

	// Calculate the power of ten once to reduce redundant calculations.
	powerOfTenM := new(big.Int).Lsh(big.NewInt(1), uint(m))

	// Split x and y into high and low parts.
	highX := new(big.Int).Div(x, powerOfTenM)
	lowX := new(big.Int).Mod(x, powerOfTenM)
	highY := new(big.Int).Div(y, powerOfTenM)
	lowY := new(big.Int).Mod(y, powerOfTenM)

	// Recursively calculate three products.
	z0 := karatsuba(lowX, lowY)
	z2 := karatsuba(highX, highY)

	// Calculate (lowX + highX) * (lowY + highY) and store the result in z1.
	sumX := new(big.Int).Add(lowX, highX)
	sumY := new(big.Int).Add(lowY, highY)
	z1 := karatsuba(sumX, sumY)

	// Combine the results to get the final product.
	// result = (z2 * powerOfTenM^2) + ((z1 - z2 - z0) * powerOfTenM) + z0
	result := new(big.Int).Mul(z2, new(big.Int).Mul(powerOfTenM, powerOfTenM))
	z1.Sub(z1, z2).Sub(z1, z0)
	result.Add(result, new(big.Int).Mul(z1, powerOfTenM)).Add(result, z0)

	return result
}

func main() {
	// Define large integers as strings and parse them into big.Int.
	x := new(big.Int)
	y := new(big.Int)
	x.SetString("3141592653589793238462643383279502884197169399375105820974944592", 10)
	y.SetString("2718281828459045235360287471352662497757247093699959574966967627", 10)

	// Perform Karatsuba multiplication.
	result := KaratsubaMultiplier{}.KaratsubaMultiply(x, y)

	fmt.Println("Product:", result)
}
