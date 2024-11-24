package main

import (
	"errors"
	"fmt"
)

// TSGraph represents a directed graph using an adjacency list.
type TSGraph struct {
	adjacencyList map[int][]int
}

// NewTSGraph creates and initializes a new TSGraph.
func NewTSGraph() *TSGraph {
	return &TSGraph{
		adjacencyList: make(map[int][]int),
	}
}

// AddEdge adds a directed edge from node u to node v.
func (g *TSGraph) AddEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
}

// TopologicalSort performs a topological sort on the graph.
func (g *TSGraph) TopologicalSort() ([]int, error) {
	visited := make(map[int]bool)  // Track permanently visited nodes.
	tempMark := make(map[int]bool) // Track temporarily marked nodes for cycle detection.
	result := make([]int, 0)       // Store topological ordering.

	var dfs func(node int) error
	dfs = func(node int) error {
		if tempMark[node] {
			return errors.New("graph contains a cycle, topological sort not possible")
		}
		if !visited[node] {
			tempMark[node] = true // Mark the node temporarily.
			for _, neighbor := range g.adjacencyList[node] {
				if err := dfs(neighbor); err != nil {
					return err
				}
			}
			tempMark[node] = false        // Remove the temporary mark.
			visited[node] = true          // Mark the node as permanently visited.
			result = append(result, node) // Add the node to the result.
		}
		return nil
	}

	// Perform DFS for all nodes in the graph.
	for node := range g.adjacencyList {
		if !visited[node] {
			if err := dfs(node); err != nil {
				return nil, err
			}
		}
	}

	// Reverse result for correct topological order.
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result, nil
}

func main() {
	// Example usage of TSGraph.
	graph := NewTSGraph()
	graph.AddEdge(5, 2)
	graph.AddEdge(5, 0)
	graph.AddEdge(4, 0)
	graph.AddEdge(4, 1)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 1)

	order, err := graph.TopologicalSort()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Topological Order:", order)
	}
}
