package main

import (
	"fmt"
	"math"
)

// OptimalBinarySearchTree represents the solver for the Optimal Binary Search Tree problem.
type OptimalBinarySearchTree struct {
	probabilities []float64
	n             int
	dp            [][]float64
	root          [][]int
	prefixSum     []float64
}

// NewOptimalBinarySearchTree initializes a new instance of the solver.
func NewOptimalBinarySearchTree(probabilities []float64) *OptimalBinarySearchTree {
	n := len(probabilities)
	dp := make([][]float64, n)
	root := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]float64, n)
		root[i] = make([]int, n)
	}
	return &OptimalBinarySearchTree{
		probabilities: probabilities,
		n:             n,
		dp:            dp,
		root:          root,
		prefixSum:     computePrefixSum(probabilities),
	}
}

// computePrefixSum computes prefix sums of probabilities for fast range sum calculations.
func computePrefixSum(probabilities []float64) []float64 {
	n := len(probabilities)
	prefixSum := make([]float64, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + probabilities[i]
	}
	return prefixSum
}

// calculateSum calculates the sum of probabilities from index i to j (inclusive).
func (obst *OptimalBinarySearchTree) calculateSum(i, j int) float64 {
	return obst.prefixSum[j+1] - obst.prefixSum[i]
}

// Solve solves the Optimal Binary Search Tree problem using a specified approach.
func (obst *OptimalBinarySearchTree) Solve(optimize bool) {
	// Base case: intervals with a single key.
	for i := 0; i < obst.n; i++ {
		obst.dp[i][i] = obst.probabilities[i]
		obst.root[i][i] = i
	}

	// Fill DP table for intervals of increasing size.
	for s := 1; s < obst.n; s++ {
		for i := 0; i < obst.n-s; i++ {
			j := i + s
			obst.dp[i][j] = math.Inf(1)
			totalProb := obst.calculateSum(i, j)

			// Determine root range for optimization or brute force.
			rootStart := i
			rootEnd := j
			if optimize && s > 1 {
				rootStart = obst.root[i][j-1]
				rootEnd = obst.root[i+1][j]
			}

			// Try all roots in the determined range.
			for r := rootStart; r <= rootEnd; r++ {
				costLeft := 0.0
				costRight := 0.0
				if r > i {
					costLeft = obst.dp[i][r-1]
				}
				if r < j {
					costRight = obst.dp[r+1][j]
				}
				totalCost := costLeft + costRight + totalProb

				if totalCost < obst.dp[i][j] {
					obst.dp[i][j] = totalCost
					obst.root[i][j] = r
				}
			}
		}
	}
}

// GetOptimalCost retrieves the optimal cost from the DP table.
func (obst *OptimalBinarySearchTree) GetOptimalCost() float64 {
	return obst.dp[0][obst.n-1]
}

// ReconstructTree reconstructs the tree structure from the DP table.
func (obst *OptimalBinarySearchTree) ReconstructTree(i, j int) *Node {
	if i > j {
		return nil
	}
	root := obst.root[i][j]
	return &Node{
		Value: root,
		Left:  obst.ReconstructTree(i, root-1),
		Right: obst.ReconstructTree(root+1, j),
	}
}

// Node represents a tree node.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// PrettyPrintTree prints the tree structure in a readable format.
func PrettyPrintTree(tree *Node, depth int) {
	if tree == nil {
		return
	}
	fmt.Printf("%sNode: %d\n", getIndent(depth), tree.Value)
	PrettyPrintTree(tree.Left, depth+1)
	PrettyPrintTree(tree.Right, depth+1)
}

// getIndent generates indentation for tree levels.
func getIndent(depth int) string {
	return fmt.Sprintf("%s", string(make([]byte, depth*2)))
}

// Main function demonstrating both approaches.
func main() {
	probabilities := []float64{0.1, 0.2, 0.4, 0.3}

	// Cubic approach.
	fmt.Println("Cubic Approach:")
	obstCubic := NewOptimalBinarySearchTree(probabilities)
	obstCubic.Solve(false)
	fmt.Printf("Optimal cost (cubic): %.1f\n", obstCubic.GetOptimalCost())
	fmt.Println("Tree structure (cubic):")
	PrettyPrintTree(obstCubic.ReconstructTree(0, obstCubic.n-1), 0)

	// Quadratic approach.
	fmt.Println("\nQuadratic Approach:")
	obstQuadratic := NewOptimalBinarySearchTree(probabilities)
	obstQuadratic.Solve(true)
	fmt.Printf("Optimal cost (quadratic): %.1f\n", obstQuadratic.GetOptimalCost())
	fmt.Println("Tree structure (quadratic):")
	PrettyPrintTree(obstQuadratic.ReconstructTree(0, obstQuadratic.n-1), 0)
}
