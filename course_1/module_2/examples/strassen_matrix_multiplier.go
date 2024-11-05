package main

import (
	"errors"
	"fmt"
)

// StrassenMatrixMultiplier is a struct for matrix multiplication using Strassen's algorithm.
type StrassenMatrixMultiplier struct {
	MatrixA, MatrixB [][]int
	AdjustedSize     int
	OriginalSize     int
}

// NewStrassenMatrixMultiplier is a constructor that initializes a StrassenMatrixMultiplier instance.
func NewStrassenMatrixMultiplier(matrixA, matrixB [][]int) (*StrassenMatrixMultiplier, error) {
	// Ensure matrices are square and of the same dimensions.
	if len(matrixA) != len(matrixB) || len(matrixA) != len(matrixA[0]) || len(matrixB) != len(matrixB[0]) {
		return nil, errors.New("both matrices must be square and have the same dimensions")
	}

	originalSize := len(matrixA)
	adjustedSize := getNextPowerOfTwo(originalSize)

	// Pad matrices to the nearest power of two if necessary.
	if adjustedSize != originalSize {
		matrixA = padMatrix(matrixA, adjustedSize)
		matrixB = padMatrix(matrixB, adjustedSize)
	}

	return &StrassenMatrixMultiplier{
		MatrixA:      matrixA,
		MatrixB:      matrixB,
		AdjustedSize: adjustedSize,
		OriginalSize: originalSize,
	}, nil
}

// Multiply performs Strassen's matrix multiplication and removes padding if added.
func (s *StrassenMatrixMultiplier) Multiply() [][]int {
	result := s.strassenMultiply(s.MatrixA, s.MatrixB)
	// Remove padding if it was added to the matrices
	return unPadMatrix(result, s.OriginalSize)
}

// strassenMultiply is a recursive method for performing Strassen's matrix multiplication.
func (s *StrassenMatrixMultiplier) strassenMultiply(matrixA, matrixB [][]int) [][]int {
	size := len(matrixA)
	// Base case: If the matrix is 1x1, return the product directly.
	if size == 1 {
		return [][]int{{matrixA[0][0] * matrixB[0][0]}}
	}

	// Split matrices into quadrants.
	a11, a12, a21, a22 := splitMatrix(matrixA)
	b11, b12, b21, b22 := splitMatrix(matrixB)

	// Compute the 7 products required by Strassen's algorithm.
	product1 := s.strassenMultiply(a11, subtractMatrix(b12, b22))
	product2 := s.strassenMultiply(addMatrix(a11, a12), b22)
	product3 := s.strassenMultiply(addMatrix(a21, a22), b11)
	product4 := s.strassenMultiply(a22, subtractMatrix(b21, b11))
	product5 := s.strassenMultiply(addMatrix(a11, a22), addMatrix(b11, b22))
	product6 := s.strassenMultiply(subtractMatrix(a12, a22), addMatrix(b21, b22))
	product7 := s.strassenMultiply(subtractMatrix(a11, a21), addMatrix(b11, b12))

	// Combine products to form the resulting quadrants.
	c11 := addMatrix(subtractMatrix(addMatrix(product5, product4), product2), product6)
	c12 := addMatrix(product1, product2)
	c21 := addMatrix(product3, product4)
	c22 := subtractMatrix(subtractMatrix(addMatrix(product5, product1), product3), product7)

	// Combine quadrants into a single matrix.
	return combineQuadrants(c11, c12, c21, c22)
}

// getNextPowerOfTwo finds the smallest power of two greater than or equal to the given integer.
func getNextPowerOfTwo(n int) int {
	power := 1
	for power < n {
		power *= 2
	}
	return power
}

// padMatrix expands the matrix to the specified size, filling in zeros for the new entries.
func padMatrix(matrix [][]int, size int) [][]int {
	padded := make([][]int, size)
	for i := range padded {
		padded[i] = make([]int, size)
		if i < len(matrix) {
			copy(padded[i], matrix[i])
		}
	}
	return padded
}

// unPadMatrix removes extra padding from the matrix, restoring it to the original size.
func unPadMatrix(matrix [][]int, size int) [][]int {
	unPadded := make([][]int, size)
	for i := 0; i < size; i++ {
		unPadded[i] = matrix[i][:size]
	}
	return unPadded
}

// splitMatrix divides a matrix into four quadrants.
func splitMatrix(matrix [][]int) ([][]int, [][]int, [][]int, [][]int) {
	halfSize := len(matrix) / 2
	a11 := make([][]int, halfSize)
	a12 := make([][]int, halfSize)
	a21 := make([][]int, halfSize)
	a22 := make([][]int, halfSize)
	for i := 0; i < halfSize; i++ {
		a11[i] = matrix[i][:halfSize]
		a12[i] = matrix[i][halfSize:]
		a21[i] = matrix[i+halfSize][:halfSize]
		a22[i] = matrix[i+halfSize][halfSize:]
	}
	return a11, a12, a21, a22
}

// combineQuadrants merges four quadrants into a single matrix.
func combineQuadrants(c11, c12, c21, c22 [][]int) [][]int {
	size := len(c11) * 2
	combined := make([][]int, size)
	for i := 0; i < size/2; i++ {
		combined[i] = append(c11[i], c12[i]...)
		combined[i+size/2] = append(c21[i], c22[i]...)
	}
	return combined
}

// addMatrix adds two matrices element-wise.
func addMatrix(matrixA, matrixB [][]int) [][]int {
	size := len(matrixA)
	result := make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
		for j := 0; j < size; j++ {
			result[i][j] = matrixA[i][j] + matrixB[i][j]
		}
	}
	return result
}

// subtractMatrix subtracts matrixB from matrixA element-wise.
func subtractMatrix(matrixA, matrixB [][]int) [][]int {
	size := len(matrixA)
	result := make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
		for j := 0; j < size; j++ {
			result[i][j] = matrixA[i][j] - matrixB[i][j]
		}
	}
	return result
}

// Main function for testing.
func main() {
	matrixA := [][]int{
		{1, 2},
		{3, 4},
	}
	matrixB := [][]int{
		{5, 6},
		{7, 8},
	}

	multiplier, err := NewStrassenMatrixMultiplier(matrixA, matrixB)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result := multiplier.Multiply()

	fmt.Println("Matrix A:")
	printMatrix(matrixA)
	fmt.Println("Matrix B:")
	printMatrix(matrixB)
	fmt.Println("Result of Strassen's Matrix Multiplication (A * B):")
	printMatrix(result)
}

// printMatrix prints a matrix row by row.
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
