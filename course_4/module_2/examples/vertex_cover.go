package main

import (
	"fmt"
)

// Graph represents an undirected graph using an adjacency list.
type Graph map[string][]string

// VertexCover represents the vertex cover algorithm with memoization.
type VertexCover struct {
	memo map[string]bool // Memoization to store intermediate results.
}

// NewVertexCover initializes a VertexCover instance.
func NewVertexCover() *VertexCover {
	return &VertexCover{
		memo: make(map[string]bool),
	}
}

// VertexCoverSolver determines if there exists a vertex cover of size k or smaller.
func (vc *VertexCover) VertexCoverSolver(graph Graph, k int) bool {
	// Serialize the graph and k into a key for memoization.
	graphKey := vc.serializeGraph(graph) + fmt.Sprintf("_%d", k)

	// Check if the result is already memoized.
	if result, exists := vc.memo[graphKey]; exists {
		return result
	}

	// Base cases for recursion.
	if k < 0 {
		return false // Negative k is invalid.
	}
	if len(graph) == 0 {
		return true // Empty graph is trivially covered.
	}
	if k == 0 {
		return false // Non-empty graph cannot have a vertex cover of size 0.
	}

	// Pick a vertex with the highest degree (for better performance).
	var u string
	maxDegree := 0
	for vertex, neighbors := range graph {
		if len(neighbors) > maxDegree {
			u = vertex
			maxDegree = len(neighbors)
		}
	}

	// Select the first neighbor of u.
	v := graph[u][0]

	// Create subgraphs by removing u and v along with their incident edges.
	graphU := vc.createSubgraphExcludingVertices(graph, []string{u})
	graphV := vc.createSubgraphExcludingVertices(graph, []string{v})

	// Recursively check for vertex covers of size k-1.
	coverU := vc.VertexCoverSolver(graphU, k-1)
	coverV := vc.VertexCoverSolver(graphV, k-1)

	// Memoize and return the result.
	result := coverU || coverV
	vc.memo[graphKey] = result
	return result
}

// createSubgraphExcludingVertices creates a new subgraph by removing given vertices and their incident edges.
func (vc *VertexCover) createSubgraphExcludingVertices(graph Graph, vertices []string) Graph {
	// Create a set of vertices to remove for fast lookup.
	remove := make(map[string]bool)
	for _, v := range vertices {
		remove[v] = true
	}

	// Create a new graph excluding specified vertices and their edges.
	newGraph := make(Graph)
	for v, neighbors := range graph {
		if remove[v] {
			continue
		}
		filteredNeighbors := []string{}
		for _, neighbor := range neighbors {
			if !remove[neighbor] {
				filteredNeighbors = append(filteredNeighbors, neighbor)
			}
		}
		if len(filteredNeighbors) > 0 {
			newGraph[v] = filteredNeighbors
		}
	}
	return newGraph
}

// serializeGraph creates a string representation of the graph for memoization.
func (vc *VertexCover) serializeGraph(graph Graph) string {
	serialized := ""
	for v, neighbors := range graph {
		serialized += v + ":"
		for _, neighbor := range neighbors {
			serialized += neighbor + ","
		}
		serialized += ";"
	}
	return serialized
}

// Main function to test the implementation.
func main() {
	// Define a simple undirected graph as an adjacency list.
	graph := Graph{
		"A": {"B", "C"},
		"B": {"A", "D", "E"},
		"C": {"A", "F"},
		"D": {"B"},
		"E": {"B"},
		"F": {"C"},
	}

	// Create an instance of the VertexCover struct.
	vc := NewVertexCover()

	// Check if there is a vertex cover of size 2 or smaller.
	k := 2
	result := vc.VertexCoverSolver(graph, k)

	// Print the result.
	if result {
		fmt.Printf("Is there a vertex cover of size %d or smaller? Yes.\n", k)
	} else {
		fmt.Printf("Is there a vertex cover of size %d or smaller? No.\n", k)
	}
}
