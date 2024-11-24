package main

import (
	"fmt"
	"sort"
)

// DFSGraph represents a graph using an adjacency list.
type DFSGraph struct {
	adjacencyList map[int][]int
}

// NewDFSGraph creates and initializes a new DFSGraph.
func NewDFSGraph() *DFSGraph {
	return &DFSGraph{
		adjacencyList: make(map[int][]int),
	}
}

// AddEdge adds an edge between two nodes.
// If the graph is undirected, it adds the edge in both directions.
func (g *DFSGraph) AddEdge(u, v int, directed bool) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	if !directed {
		g.adjacencyList[v] = append(g.adjacencyList[v], u)
	}
}

// DFS performs Depth-First Search starting from the given start node.
// It returns the order of visited nodes.
func (g *DFSGraph) DFS(start int) []int {
	visited := make(map[int]bool) // Track visited nodes.
	result := make([]int, 0)      // Store the order of visited nodes.
	stack := []int{start}         // Use a stack for iterative DFS.

	for len(stack) > 0 {
		// Pop the last node from the stack.
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[node] {
			visited[node] = true          // Mark the node as visited.
			result = append(result, node) // Add it to the visited order.

			// Add neighbors to the stack in sorted order.
			neighbors := g.adjacencyList[node]
			sort.Ints(neighbors)
			for i := len(neighbors) - 1; i >= 0; i-- {
				if !visited[neighbors[i]] {
					stack = append(stack, neighbors[i])
				}
			}
		}
	}

	return result
}

// ConnectedComponents finds all connected components in the graph.
// It returns a slice of slices, where each inner slice represents a connected component.
func (g *DFSGraph) ConnectedComponents() [][]int {
	visited := make(map[int]bool)  // Track visited nodes.
	components := make([][]int, 0) // Store all connected components.

	for node := range g.adjacencyList {
		if !visited[node] {
			component := make([]int, 0) // Store the current connected component.
			stack := []int{node}

			for len(stack) > 0 {
				// Pop the last node from the stack.
				current := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if !visited[current] {
					visited[current] = true                // Mark the node as visited.
					component = append(component, current) // Add to the current component.

					// Add neighbors to the stack in sorted order.
					neighbors := g.adjacencyList[current]
					sort.Ints(neighbors)
					for i := len(neighbors) - 1; i >= 0; i-- {
						if !visited[neighbors[i]] {
							stack = append(stack, neighbors[i])
						}
					}
				}
			}

			components = append(components, component)
		}
	}

	return components
}

func main() {
	graph := NewDFSGraph()
	graph.AddEdge(1, 2, false)
	graph.AddEdge(2, 3, false)
	graph.AddEdge(4, 5, false)
	graph.AddEdge(6, 7, false)
	graph.AddEdge(7, 8, false)
	graph.AddEdge(8, 6, false)

	// Perform DFS starting from node 1.
	fmt.Println("Depth-First Search (DFS) Order:")
	dfsOrder := graph.DFS(1)
	fmt.Println(dfsOrder)

	// Find all connected components in the graph.
	fmt.Println("\nConnected Components:")
	components := graph.ConnectedComponents()
	for i, component := range components {
		sort.Ints(component) // Sort each component for consistent output.
		fmt.Printf("Component %d: %v\n", i+1, component)
	}
}
