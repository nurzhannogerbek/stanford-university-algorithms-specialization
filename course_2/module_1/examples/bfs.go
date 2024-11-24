package main

import (
	"container/list"
	"fmt"
)

// BFSGraph represents an undirected graph using an adjacency list.
type BFSGraph struct {
	adjacencyList map[int][]int // The adjacency list to store graph edges.
}

// NewBFSGraph creates and initializes a new BFSGraph.
func NewBFSGraph() *BFSGraph {
	return &BFSGraph{
		adjacencyList: make(map[int][]int),
	}
}

// AddEdge adds an undirected edge between nodes u and v.
//
// Args:
//
//	u: The start node of the edge.
//	v: The end node of the edge.
func (g *BFSGraph) AddEdge(u, v int) {
	g.adjacencyList[u] = append(g.adjacencyList[u], v)
	g.adjacencyList[v] = append(g.adjacencyList[v], u)
}

// BFS performs Breadth-First Search starting from a given node.
// It calculates the shortest distance from the start node to all reachable nodes.
//
// Args:
//
//	start: The starting node for the BFS traversal.
//
// Returns:
//
//	A map where keys are nodes and values are their distances from the start node.
//	If the start node does not exist in the graph, returns an error.
func (g *BFSGraph) BFS(start int) (map[int]int, error) {
	if _, exists := g.adjacencyList[start]; !exists {
		return nil, fmt.Errorf("start node %d not found in the graph", start)
	}

	distances := make(map[int]int) // Store distances from the start node.
	visited := make(map[int]bool)  // Track visited nodes.
	queue := list.New()            // Queue for BFS.

	// Initialize BFS.
	queue.PushBack(start)
	visited[start] = true
	distances[start] = 0

	// Perform BFS traversal.
	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)

		current := element.Value.(int)

		// Explore all neighbors of the current node.
		for _, neighbor := range g.adjacencyList[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
				distances[neighbor] = distances[current] + 1
			}
		}
	}

	return distances, nil
}

// ConnectedComponents identifies all connected components in the graph.
//
// Returns:
//
//	A slice of slices, where each inner slice represents a connected component.
func (g *BFSGraph) ConnectedComponents() [][]int {
	visited := make(map[int]bool)  // Track visited nodes.
	components := make([][]int, 0) // Use `make` to initialize the slice.

	// Explore all nodes in the graph.
	for node := range g.adjacencyList {
		if !visited[node] {
			component := make([]int, 0) // Initialize the component slice.
			queue := list.New()         // Queue for BFS.

			// Initialize BFS for this component.
			queue.PushBack(node)
			visited[node] = true

			// Perform BFS for the current component.
			for queue.Len() > 0 {
				element := queue.Front()
				queue.Remove(element)

				current := element.Value.(int)
				component = append(component, current)

				for _, neighbor := range g.adjacencyList[current] {
					if !visited[neighbor] {
						visited[neighbor] = true
						queue.PushBack(neighbor)
					}
				}
			}

			// Append the component to the result.
			components = append(components, component)
		}
	}

	return components
}

func main() {
	// Create a new graph.
	graph := NewBFSGraph()

	// Add edges to the graph.
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(4, 5)
	graph.AddEdge(6, 7)
	graph.AddEdge(7, 8)
	graph.AddEdge(8, 6)

	// Perform BFS starting from node 1 and print distances.
	fmt.Println("Breadth-First Search (BFS) Distances:")
	if distances, err := graph.BFS(1); err != nil {
		fmt.Println(err)
	} else {
		// Format distances as {node: distance}.
		fmt.Print("{")
		first := true
		for node, distance := range distances {
			if !first {
				fmt.Print(", ")
			}
			fmt.Printf("%d: %d", node, distance)
			first = false
		}
		fmt.Println("}")
	}

	// Find all connected components in the graph and print them.
	fmt.Println("\nConnected Components:")
	components := graph.ConnectedComponents()
	fmt.Print("[")
	for i, component := range components {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(component)
	}
	fmt.Println("]")
}
