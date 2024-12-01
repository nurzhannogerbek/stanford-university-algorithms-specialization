package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// PathGraphMWIS represents the Maximum Weight Independent Set (MWIS) solver for a path graph.
type PathGraphMWIS struct {
	weights  []int // Weights of the vertices in the path graph.
	n        int   // Number of vertices.
	dp       []int // DP table for storing maximum weights for subproblems.
	solution []int // List of selected vertices in the MWIS (0-based indices).
}

// NewPathGraphMWIS initializes a new MWIS solver with the given vertex weights.
func NewPathGraphMWIS(weights []int) *PathGraphMWIS {
	n := len(weights)
	return &PathGraphMWIS{
		weights:  weights,
		n:        n,
		dp:       make([]int, n+1),
		solution: []int{},
	}
}

// ComputeMWIS computes the maximum weight independent set using dynamic programming.
func (mwis *PathGraphMWIS) ComputeMWIS() {
	if mwis.n == 0 {
		return // No vertices in the graph.
	}
	if mwis.n >= 1 {
		mwis.dp[1] = mwis.weights[0] // Base case for the first vertex.
	}

	// Fill the DP table using the recurrence relation:
	// dp[i] = max(dp[i-1], dp[i-2] + weights[i-1])
	for i := 2; i <= mwis.n; i++ {
		mwis.dp[i] = max(mwis.dp[i-1], mwis.dp[i-2]+mwis.weights[i-1])
	}
}

// ReconstructSolution reconstructs the MWIS from the computed DP table.
func (mwis *PathGraphMWIS) ReconstructSolution() {
	i := mwis.n
	for i >= 1 {
		if i == 1 || mwis.dp[i] != mwis.dp[i-1] {
			// Include the current vertex in the solution.
			mwis.solution = append([]int{i - 1}, mwis.solution...) // Add vertex to the front.
			i -= 2                                                 // Skip the adjacent vertex.
		} else {
			// Exclude the current vertex.
			i -= 1
		}
	}
}

// GetSolutionString generates a binary solution string for the specified vertices.
// Each bit represents whether the corresponding vertex is included in the MWIS.
func (mwis *PathGraphMWIS) GetSolutionString(vertices []int) string {
	// Convert the solution into a set for quick lookup.
	solutionSet := make(map[int]bool)
	for _, v := range mwis.solution {
		solutionSet[v] = true
	}

	// Generate the binary string based on the inclusion of vertices in the solution.
	var result strings.Builder
	for _, v := range vertices {
		if solutionSet[v-1] { // Convert 1-based to 0-based index.
			result.WriteString("1")
		} else {
			result.WriteString("0")
		}
	}
	return result.String()
}

// Solve solves the MWIS problem and generates the binary solution string for the specified vertices.
func (mwis *PathGraphMWIS) Solve(vertices []int) string {
	mwis.ComputeMWIS()         // Compute the DP table.
	mwis.ReconstructSolution() // Reconstruct the MWIS.
	return mwis.GetSolutionString(vertices)
}

// Utility function to compute the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Main function demonstrating the usage of the PathGraphMWIS class.
func main() {
	// Read the input file.
	data, err := ioutil.ReadFile("course_3/module_3/programming_assignment_3/mwis.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the weights from the file.
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	numVertices, _ := strconv.Atoi(lines[0])
	weights := make([]int, numVertices)
	for i := 1; i <= numVertices; i++ {
		weights[i-1], _ = strconv.Atoi(lines[i])
	}

	// Initialize the MWIS solver.
	mwisSolver := NewPathGraphMWIS(weights)

	// Define the vertices to check (1-based indices).
	verticesToCheck := []int{1, 2, 3, 4, 17, 117, 517, 997}

	// Solve the problem and get the solution string.
	solutionString := mwisSolver.Solve(verticesToCheck)

	// Print the result.
	fmt.Println("Solution String:", solutionString)
}
