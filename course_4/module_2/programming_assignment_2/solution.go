package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// TravelingSalesmanProblem represents a solver for the TSP using recursion and memoization.
type TravelingSalesmanProblem struct {
	numCities      int                // Number of cities.
	coordinates    [][]float64        // Coordinates of the cities.
	distanceMatrix [][]float64        // Distance matrix between cities.
	VISITED_ALL    int                // Bitmask representing all cities visited.
	memo           map[string]float64 // Memoization table for subproblem results.
}

// NewTravelingSalesmanProblem initializes a new TSP solver by reading data from a file.
func NewTravelingSalesmanProblem(filename string) (*TravelingSalesmanProblem, error) {
	tsp := &TravelingSalesmanProblem{
		coordinates:    [][]float64{},
		distanceMatrix: [][]float64{},
		memo:           make(map[string]float64),
	}

	// Load data from the file.
	if err := tsp.loadData(filename); err != nil {
		return nil, err
	}

	// Prepare the distance matrix and the bitmask for all cities visited.
	tsp.prepareDistanceMatrix()
	tsp.VISITED_ALL = (1 << tsp.numCities) - 1

	return tsp, nil
}

// loadData reads the number of cities and their coordinates from the input file.
func (tsp *TravelingSalesmanProblem) loadData(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the number of cities.
	if !scanner.Scan() {
		return fmt.Errorf("file is empty or invalid")
	}
	tsp.numCities, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return fmt.Errorf("invalid number of cities: %v", err)
	}

	// Read the coordinates of the cities.
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return fmt.Errorf("invalid format for coordinates: %v", line)
		}

		x, err1 := strconv.ParseFloat(parts[0], 64)
		y, err2 := strconv.ParseFloat(parts[1], 64)
		if err1 != nil || err2 != nil {
			return fmt.Errorf("invalid coordinates: %v", line)
		}
		tsp.coordinates = append(tsp.coordinates, []float64{x, y})
	}

	return scanner.Err()
}

// prepareDistanceMatrix calculates the distance matrix using Euclidean distances.
func (tsp *TravelingSalesmanProblem) prepareDistanceMatrix() {
	tsp.distanceMatrix = make([][]float64, tsp.numCities)
	for i := 0; i < tsp.numCities; i++ {
		tsp.distanceMatrix[i] = make([]float64, tsp.numCities)
		for j := 0; j < tsp.numCities; j++ {
			if i != j {
				tsp.distanceMatrix[i][j] = tsp.euclideanDistance(tsp.coordinates[i], tsp.coordinates[j])
			}
		}
	}
}

// euclideanDistance calculates the Euclidean distance between two points.
func (tsp *TravelingSalesmanProblem) euclideanDistance(city1, city2 []float64) float64 {
	x1, y1 := city1[0], city1[1]
	x2, y2 := city2[0], city2[1]
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

// solveRecursive is the core recursive function to solve the TSP.
func (tsp *TravelingSalesmanProblem) solveRecursive(mask, pos int) float64 {
	// Base case: all cities have been visited.
	if mask == tsp.VISITED_ALL {
		return tsp.distanceMatrix[pos][0] // Return to the starting city.
	}

	// Check memoization table.
	key := fmt.Sprintf("%d|%d", mask, pos)
	if val, exists := tsp.memo[key]; exists {
		return val
	}

	// Try all unvisited cities and find the minimum cost.
	ans := math.Inf(1)
	for city := 0; city < tsp.numCities; city++ {
		if mask&(1<<city) == 0 { // If the city has not been visited.
			newCost := tsp.distanceMatrix[pos][city] + tsp.solveRecursive(mask|(1<<city), city)
			if newCost < ans {
				ans = newCost
			}
		}
	}

	// Store the result in the memoization table.
	tsp.memo[key] = ans
	return ans
}

// Solve computes the minimum cost of the traveling salesman tour.
func (tsp *TravelingSalesmanProblem) Solve() float64 {
	// Start with the first city visited and position at the first city.
	return tsp.solveRecursive(1, 0)
}

func main() {
	// Path to the TSP instance file.
	filename := "course_4/module_2/programming_assignment_2/tsp.txt"

	// Initialize the TSP solver.
	tsp, err := NewTravelingSalesmanProblem(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Solve the TSP and print the minimum cost.
	minCost := tsp.Solve()
	fmt.Printf("Minimum cost: %.0f\n", math.Floor(minCost))
}
