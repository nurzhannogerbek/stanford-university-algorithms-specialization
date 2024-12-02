package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// INF is a placeholder for infinite distance. It is defined as a variable instead of a constant
// because Go does not allow math.Inf(1) to be used as a constant.
var INF = math.Inf(1)

// Graph represents a graph structure with vertices and edges.
type Graph struct {
	vertices int
	edges    [][]float64
}

// NewGraph initializes a new Graph object with a given number of vertices.
func NewGraph(vertices int) *Graph {
	// Create a 2D slice for edges with initial distances as INF.
	edges := make([][]float64, vertices)
	for i := range edges {
		edges[i] = make([]float64, vertices)
		for j := range edges[i] {
			if i == j {
				edges[i][j] = 0 // Distance to itself is always 0.
			} else {
				edges[i][j] = INF // Initialize other distances as INF.
			}
		}
	}
	return &Graph{
		vertices: vertices,
		edges:    edges,
	}
}

// AddEdge adds an edge with a given weight to the graph.
func (g *Graph) AddEdge(u, v int, weight float64) {
	g.edges[u][v] = weight
}

// FloydWarshall computes the shortest paths between all pairs of vertices using
// the Floyd-Warshall algorithm and checks for negative weight cycles.
func (g *Graph) FloydWarshall() ([][]float64, bool) {
	// Initialize the distance matrix by copying the current graph edges.
	dist := make([][]float64, g.vertices)
	for i := range dist {
		dist[i] = make([]float64, g.vertices)
		copy(dist[i], g.edges[i])
	}

	// Apply the Floyd-Warshall algorithm.
	for k := 0; k < g.vertices; k++ {
		for i := 0; i < g.vertices; i++ {
			for j := 0; j < g.vertices; j++ {
				if dist[i][k] != INF && dist[k][j] != INF {
					dist[i][j] = math.Min(dist[i][j], dist[i][k]+dist[k][j])
				}
			}
		}
	}

	// Check for negative weight cycles.
	for i := 0; i < g.vertices; i++ {
		if dist[i][i] < 0 {
			return dist, true // Negative cycle detected.
		}
	}
	return dist, false
}

// GetShortestShortestPath computes the shortest path among all vertex pairs
// or indicates that a negative cycle exists in the graph.
func (g *Graph) GetShortestShortestPath() (float64, bool) {
	dist, hasNegativeCycle := g.FloydWarshall()
	if hasNegativeCycle {
		return 0, true // Indicate the presence of a negative cycle.
	}

	// Find the shortest path in the distance matrix.
	shortest := INF
	for i := 0; i < g.vertices; i++ {
		for j := 0; j < g.vertices; j++ {
			if i != j && dist[i][j] < shortest {
				shortest = dist[i][j]
			}
		}
	}

	// If no valid path is found, return 0.
	if shortest == INF {
		return 0, false
	}
	return shortest, false
}

// ReadGraphFromFile reads a graph from a file and returns a Graph object.
func ReadGraphFromFile(filename string) (*Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Read the first line for the number of vertices and edges.
	scanner.Scan()
	header := strings.Fields(scanner.Text())
	vertices, err := strconv.Atoi(header[0])
	if err != nil {
		return nil, fmt.Errorf("invalid number of vertices: %v", err)
	}

	// Initialize the graph with the given number of vertices.
	graph := NewGraph(vertices)
	// Read edges line by line.
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) != 3 {
			return nil, fmt.Errorf("invalid edge format: %s", line)
		}
		u, err1 := strconv.Atoi(line[0])
		v, err2 := strconv.Atoi(line[1])
		weight, err3 := strconv.ParseFloat(line[2], 64)
		if err1 != nil || err2 != nil || err3 != nil {
			return nil, fmt.Errorf("invalid edge data: %v, %v, %v", err1, err2, err3)
		}
		graph.AddEdge(u-1, v-1, weight) // Convert to 0-based indexing.
	}
	return graph, nil
}

func main() {
	// Define the base directory to avoid repeating common paths.
	basePath := "course_4/module_1/programming_assignment_1"
	filenames := []string{
		"g1.txt",
		"g2.txt",
		"g3.txt",
	}

	// Process each file in the given directory.
	for i, file := range filenames {
		filePath := filepath.Join(basePath, file) // Combine base path and file name.
		graph, err := ReadGraphFromFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			continue
		}
		// Get the shortest path or detect a negative cycle.
		shortest, hasNegativeCycle := graph.GetShortestShortestPath()
		if hasNegativeCycle {
			fmt.Printf("Graph g%d: Negative cycle detected.\n", i+1)
		} else {
			fmt.Printf("Graph g%d: Shortest shortest path = %.2f\n", i+1, shortest)
		}
	}
}
