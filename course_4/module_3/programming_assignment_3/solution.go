package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// City represents a city with x and y coordinates.
type City struct {
	X, Y float64
}

// getDist calculates the Euclidean distance between two cities.
// It computes the square root of the sum of the squared differences
// between the x and y coordinates of the two cities.
func getDist(p1, p2 City) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// loadTSPData reads the TSP data from a file and returns a slice of City objects.
// Each city has x and y coordinates parsed from the file.
// The function ensures that the number of cities declared in the file matches
// the actual number of cities provided in the file.
func loadTSPData(filePath string) ([]City, error) {
	// Open the file for reading.
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the first line, which indicates the number of cities.
	scanner.Scan()
	nCities, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return nil, fmt.Errorf("invalid number of cities: %v", err)
	}

	// Create a slice to hold the cities.
	cities := make([]City, 0, nCities)

	// Read the subsequent lines, each containing city coordinates.
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Split the line into fields.
		parts := strings.Fields(line)
		if len(parts) != 3 { // Ensure each line has three parts: index, x, y.
			return nil, fmt.Errorf("invalid city data: %s", line)
		}

		// Parse the x and y coordinates.
		x, errX := strconv.ParseFloat(parts[1], 64)
		y, errY := strconv.ParseFloat(parts[2], 64)
		if errX != nil || errY != nil {
			return nil, fmt.Errorf("invalid coordinates: %s", line)
		}

		// Append the city to the slice.
		cities = append(cities, City{X: x, Y: y})
	}

	// Check that the number of cities matches the declared count.
	if len(cities) != nCities {
		return nil, fmt.Errorf("mismatch between declared and actual number of cities")
	}

	return cities, nil
}

// nearestNeighborTSP solves the Traveling Salesman Problem using the nearest neighbor heuristic.
// The algorithm starts at the first city and repeatedly visits the closest unvisited city.
// Once all cities are visited, it returns to the starting city to complete the tour.
// The function returns the total distance of the computed tour.
func nearestNeighborTSP(cities []City) float64 {
	nCities := len(cities)           // Total number of cities.
	visited := make([]bool, nCities) // Track whether each city has been visited.
	currCity := 0                    // Start at the first city (index 0).
	visited[currCity] = true         // Mark the starting city as visited.
	nVisited := 1                    // Number of cities visited so far.
	cost := 0.0                      // Total distance of the tour.

	// While there are still unvisited cities, find the nearest neighbor.
	for nVisited < nCities {
		nextCity := -1             // Index of the next city to visit.
		minDist := math.MaxFloat64 // Minimum distance to the next city.

		// Iterate through all cities to find the closest unvisited city.
		for i := 0; i < nCities; i++ {
			if visited[i] { // Skip cities that have already been visited.
				continue
			}
			// Calculate the distance from the current city to city i.
			dist := getDist(cities[currCity], cities[i])
			// Update the next city if this city is closer.
			if dist < minDist {
				minDist = dist
				nextCity = i
			}
		}

		// Add the distance to the total cost.
		cost += minDist
		// Mark the next city as visited.
		visited[nextCity] = true
		// Move to the next city.
		currCity = nextCity
		// Increment the count of visited cities.
		nVisited++
	}

	// Add the distance to return to the starting city to complete the tour.
	cost += getDist(cities[currCity], cities[0])
	return cost
}

func main() {
	// Define the path to the TSP data file.
	filePath := "course_4/module_3/programming_assignment_3/nn.txt"

	// Load the TSP data from the file.
	cities, err := loadTSPData(filePath)
	if err != nil {
		// Print an error message and exit if there is an issue with loading data.
		fmt.Printf("Error loading TSP data: %v\n", err)
		return
	}

	// Solve the TSP using the nearest neighbor heuristic.
	totalCost := nearestNeighborTSP(cities)

	// Output the total distance of the tour, rounded down to the nearest integer.
	fmt.Printf("Total Distance: %d\n", int(math.Floor(totalCost)))
}
