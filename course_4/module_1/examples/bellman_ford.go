package main

import (
	"errors"
	"fmt"
	"math"
)

// Graph represents a weighted directed graph.
type Graph struct {
	vertices int
	edges    map[int][]Edge
}

// Edge represents a directed edge with a weight.
type Edge struct {
	to     int
	weight float64
}

// NewGraph creates a new graph with the specified number of vertices.
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edges:    make(map[int][]Edge),
	}
}

// AddEdge adds a directed edge from vertex u to vertex v with the given weight.
func (g *Graph) AddEdge(u, v int, weight float64) error {
	if u < 0 || u >= g.vertices || v < 0 || v >= g.vertices {
		return errors.New("vertex indices must be within valid range")
	}
	g.edges[u] = append(g.edges[u], Edge{to: v, weight: weight})
	return nil
}

// BellmanFord performs the Bellman-Ford algorithm to find shortest paths from the source vertex.
func (g *Graph) BellmanFord(source int) (distances []float64, predecessors []int, err error) {
	// Check if the source vertex is valid.
	if source < 0 || source >= g.vertices {
		return nil, nil, errors.New("source vertex must be within valid range")
	}

	// Initialize distances and predecessors.
	distances = make([]float64, g.vertices)
	predecessors = make([]int, g.vertices)
	for i := 0; i < g.vertices; i++ {
		distances[i] = math.Inf(1) // Initialize all distances to infinity.
		predecessors[i] = -1       // Initialize all predecessors to -1.
	}
	distances[source] = 0 // Distance to the source is 0.

	// Relax edges |V| - 1 times.
	for i := 0; i < g.vertices-1; i++ {
		updated := false // Track if any distance was updated.
		for u, edges := range g.edges {
			for _, edge := range edges {
				if distances[u] != math.Inf(1) && distances[u]+edge.weight < distances[edge.to] {
					distances[edge.to] = distances[u] + edge.weight
					predecessors[edge.to] = u
					updated = true
				}
			}
		}
		if !updated {
			break // Exit early if no updates occurred in this iteration.
		}
	}

	// Check for negative weight cycles.
	for u, edges := range g.edges {
		for _, edge := range edges {
			if distances[u] != math.Inf(1) && distances[u]+edge.weight < distances[edge.to] {
				return nil, nil, errors.New("graph contains a negative weight cycle")
			}
		}
	}

	return distances, predecessors, nil
}

// ReconstructPath reconstructs the shortest path from source to target using predecessor pointers.
func ReconstructPath(source, target int, predecessors []int) ([]int, error) {
	if target < 0 || target >= len(predecessors) {
		return nil, errors.New("target vertex is out of range")
	}

	path := []int{}
	current := target

	// Trace back from target to source using predecessors.
	for current != -1 {
		path = append(path, current)
		if current == source {
			break
		}
		current = predecessors[current]
	}

	if current != source {
		return nil, errors.New("no path exists from source to target")
	}

	// Reverse the path to get the correct order.
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, nil
}

func main() {
	// Create a graph with 5 vertices.
	g := NewGraph(5)

	// Add edges to the graph.
	_ = g.AddEdge(0, 1, 6)
	_ = g.AddEdge(0, 3, 7)
	_ = g.AddEdge(1, 2, 5)
	_ = g.AddEdge(1, 3, 8)
	_ = g.AddEdge(1, 4, -4)
	_ = g.AddEdge(2, 1, -2)
	_ = g.AddEdge(3, 2, -3)
	_ = g.AddEdge(3, 4, 9)
	_ = g.AddEdge(4, 0, 2)
	_ = g.AddEdge(4, 2, 7)

	source := 0

	// Perform the Bellman-Ford algorithm.
	distances, predecessors, err := g.BellmanFord(source)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print distances from the source to each vertex.
	fmt.Println("Shortest distances from source:")
	for i, d := range distances {
		if math.IsInf(d, 1) {
			fmt.Printf("Vertex %d: unreachable\n", i)
		} else {
			fmt.Printf("Vertex %d: %v\n", i, d)
		}
	}

	// Reconstruct and print the path to a specific target.
	target := 4
	path, err := ReconstructPath(source, target, predecessors)
	if err != nil {
		fmt.Println("Error reconstructing path:", err)
	} else {
		fmt.Printf("Shortest path from %d to %d: %v\n", source, target, path)
	}
}
