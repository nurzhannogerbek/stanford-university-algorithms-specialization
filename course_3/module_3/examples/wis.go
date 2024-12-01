package main

import (
	"fmt"
)

// WeightedIndependentSet represents the WIS problem for a path graph.
type WeightedIndependentSet struct {
	weights []int // List of vertex weights in the path graph.
}

// NewWeightedIndependentSet initializes a new WIS solver with the given weights.
func NewWeightedIndependentSet(weights []int) *WeightedIndependentSet {
	return &WeightedIndependentSet{weights: weights}
}

// NaiveRecursive solves the WIS problem using a naive recursive approach.
// This has exponential time complexity and is only suitable for small graphs.
func (wis *WeightedIndependentSet) NaiveRecursive() int {
	var solve func(n int) int
	solve = func(n int) int {
		if n == 0 {
			return 0 // Base case: no vertices.
		}
		if n == 1 {
			return wis.weights[0] // Base case: one vertex.
		}
		// Option 1: Exclude the current vertex.
		exclude := solve(n - 1)
		// Option 2: Include the current vertex.
		include := wis.weights[n-1] + solve(n-2)
		return max(exclude, include)
	}
	return solve(len(wis.weights))
}

// Greedy solves the WIS problem using a greedy approach.
// This approach may not give the optimal solution for all cases.
func (wis *WeightedIndependentSet) Greedy() (int, []int) {
	n := len(wis.weights)
	selected := []int{} // List of selected vertices.
	totalWeight := 0    // Total weight of the selected vertices.

	for i := 0; i < n; i++ {
		// Greedy selection: choose the current vertex if it has a higher weight
		// and does not conflict with the previously selected vertices.
		if i == 0 || (wis.weights[i] > wis.weights[i-1] && (len(selected) == 0 || selected[len(selected)-1] != i-1)) {
			selected = append(selected, i)
			totalWeight += wis.weights[i]
		}
	}

	return totalWeight, selected
}

// DivideAndConquer solves the WIS problem using a divide-and-conquer approach.
// This approach recursively breaks the graph and combines results.
func (wis *WeightedIndependentSet) DivideAndConquer() int {
	var solve func(start, end int) int
	solve = func(start, end int) int {
		if start > end {
			return 0 // No vertices in this range.
		}
		if start == end {
			return wis.weights[start] // Single vertex in this range.
		}
		mid := (start + end) / 2
		// Solve for the left and right sub-problems.
		left := solve(start, mid-1)
		right := solve(mid+1, end)
		return max(left+wis.weights[mid], right)
	}
	return solve(0, len(wis.weights)-1)
}

// DynamicProgramming solves the WIS problem using dynamic programming.
// It returns the maximum weight and the set of selected vertices.
func (wis *WeightedIndependentSet) DynamicProgramming() (int, []int) {
	n := len(wis.weights)
	if n == 0 {
		return 0, []int{}
	}

	// DP table to store the maximum weight for subproblems.
	dp := make([]int, n+1)
	dp[1] = wis.weights[0]

	// Fill the DP table iteratively.
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+wis.weights[i-1])
	}

	// Reconstruct the selected vertices from the DP table.
	selected := []int{}
	for i := n; i >= 1; {
		if dp[i] == dp[i-1] {
			i-- // Exclude the current vertex.
		} else {
			selected = append([]int{i - 1}, selected...) // Include the current vertex.
			i -= 2
		}
	}

	return dp[n], selected
}

// Helper function to compute the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Main function demonstrates the usage of the WIS solver.
func main() {
	weights := []int{1, 4, 5, 4} // Example weights for the path graph.

	wis := NewWeightedIndependentSet(weights)

	fmt.Println("Naive Recursive Approach:")
	fmt.Println("Maximum Weight:", wis.NaiveRecursive())

	fmt.Println("\nGreedy Algorithm:")
	weight, selected := wis.Greedy()
	fmt.Println("Maximum Weight:", weight)
	fmt.Println("Selected Vertices:", selected)

	fmt.Println("\nDivide and Conquer Approach:")
	fmt.Println("Maximum Weight:", wis.DivideAndConquer())

	fmt.Println("\nDynamic Programming Approach:")
	weight, selected = wis.DynamicProgramming()
	fmt.Println("Maximum Weight:", weight)
	fmt.Println("Selected Vertices:", selected)
}
