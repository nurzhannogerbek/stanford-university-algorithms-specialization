package main

import (
	"fmt"
	"math"
)

// TravelingSalesmanProblem represents a solver for the TSP using bitmasking and recursion with memoization.
type TravelingSalesmanProblem struct {
	costMatrix [][]int        // Cost matrix representing the graph.
	numCities  int            // Number of cities (vertices).
	visitedAll int            // Bitmask representing all cities visited.
	memo       map[string]int // Memoization table to store results of subproblems.
}

// NewTravelingSalesmanProblem initializes a new TSP solver with the given cost matrix.
func NewTravelingSalesmanProblem(costMatrix [][]int) *TravelingSalesmanProblem {
	numCities := len(costMatrix)
	visitedAll := (1 << numCities) - 1 // Bitmask where all cities are visited.
	return &TravelingSalesmanProblem{
		costMatrix: costMatrix,
		numCities:  numCities,
		visitedAll: visitedAll,
		memo:       make(map[string]int),
	}
}

// tspRecursive is the core recursive function to solve TSP.
// mask: Bitmask representing visited cities.
// pos: Current position (city).
func (tsp *TravelingSalesmanProblem) tspRecursive(mask, pos int) int {
	// Base case: all cities have been visited.
	if mask == tsp.visitedAll {
		return tsp.costMatrix[pos][0] // Return cost to go back to the starting city.
	}

	// Generate a unique key for the current state (mask and position).
	key := fmt.Sprintf("%d|%d", mask, pos)

	// Check if the result for this state is already memoized.
	if val, exists := tsp.memo[key]; exists {
		return val
	}

	// Initialize the minimum cost for this state.
	ans := math.MaxInt

	// Try visiting all unvisited cities.
	for city := 0; city < tsp.numCities; city++ {
		if mask&(1<<city) == 0 { // Check if city is unvisited.
			// Calculate the new cost.
			newCost := tsp.costMatrix[pos][city] + tsp.tspRecursive(mask|(1<<city), city)
			// Update the minimum cost.
			if newCost < ans {
				ans = newCost
			}
		}
	}

	// Memoize the result for the current state.
	tsp.memo[key] = ans
	return ans
}

// Solve computes the minimum cost of the traveling salesman tour.
func (tsp *TravelingSalesmanProblem) Solve() int {
	// Start from the first city with only the first city visited (mask = 1).
	return tsp.tspRecursive(1, 0)
}

// Main function to demonstrate the usage of the TSP solver.
func main() {
	// Define the cost matrix where costMatrix[i][j] is the cost of traveling from city i to city j.
	costMatrix := [][]int{
		{0, 20, 42, 25},
		{20, 0, 30, 34},
		{42, 30, 0, 10},
		{25, 34, 10, 0},
	}

	// Initialize the TSP solver.
	tsp := NewTravelingSalesmanProblem(costMatrix)

	// Solve the TSP and print the minimum cost.
	minCost := tsp.Solve()
	fmt.Printf("Minimum cost: %d\n", minCost)
}
