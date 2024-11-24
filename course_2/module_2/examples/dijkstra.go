package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strings"
)

// PriorityQueue implements a priority queue for Dijkstra's algorithm.
type PriorityQueue []*Item

// Item represents a node and its priority in the priority queue.
type Item struct {
	Node     string // Node label.
	Priority int    // Priority of the node, used for shortest path computation.
	Index    int    // Index of the item in the priority queue.
}

// Len returns the length of the priority queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less compares two items in the priority queue based on their priority.
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }

// Swap swaps two items in the priority queue.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x any) {
	item := x.(*Item)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}

// Pop removes and returns the item with the highest priority (lowest value).
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // Mark as removed.
	*pq = old[0 : n-1]
	return item
}

// Graph represents a weighted graph.
type Graph struct {
	Edges map[string]map[string]int // Adjacency list to store edges and their weights.
}

// NewGraph creates and initializes a new Graph.
func NewGraph() *Graph {
	return &Graph{
		Edges: make(map[string]map[string]int),
	}
}

// AddNode adds a node to the graph without edges.
func (g *Graph) AddNode(node string) {
	if _, exists := g.Edges[node]; !exists {
		g.Edges[node] = make(map[string]int) // Initialize an empty adjacency list.
	}
}

// AddEdge adds a directed edge to the graph.
func (g *Graph) AddEdge(from, to string, weight int) {
	g.AddNode(from)
	g.AddNode(to)
	g.Edges[from][to] = weight // Add the edge with its weight.
}

// Dijkstra computes shortest paths from the start node to all other nodes.
func (g *Graph) Dijkstra(start string) (map[string]int, map[string]string) {
	// Initialize distances map with maximum integer values (infinity).
	distances := make(map[string]int)
	previous := make(map[string]string) // To reconstruct paths.

	for node := range g.Edges {
		distances[node] = math.MaxInt
	}
	distances[start] = 0 // Distance to the start node is 0.

	// Initialize the priority queue and add the start node.
	pq := &PriorityQueue{}
	heap.Push(pq, &Item{Node: start, Priority: 0})

	// Process nodes in the priority queue.
	for pq.Len() > 0 {
		// Extract the node with the smallest distance.
		current := heap.Pop(pq).(*Item)

		// Update distances to neighboring nodes.
		for neighbor, weight := range g.Edges[current.Node] {
			newDistance := distances[current.Node] + weight
			if newDistance < distances[neighbor] { // If a shorter path is found, update it.
				distances[neighbor] = newDistance
				previous[neighbor] = current.Node
				heap.Push(pq, &Item{Node: neighbor, Priority: newDistance}) // Add or update the neighbor in the queue.
			}
		}
	}

	return distances, previous // Return the computed shortest distances and previous nodes for paths.
}

// ReconstructPath reconstructs the shortest path from start to a given target node.
func ReconstructPath(previous map[string]string, start, target string) []string {
	path := []string{}
	for at := target; at != ""; at = previous[at] {
		path = append([]string{at}, path...) // Prepend the node to the path.
		if at == start {
			break
		}
	}
	if len(path) == 0 || path[0] != start {
		return []string{} // Return an empty slice if no path exists.
	}
	return path
}

// joinPath formats a slice of nodes into a string path like "A -> B -> C".
func joinPath(path []string) string {
	return strings.Join(path, " -> ")
}

func main() {
	// Create a new graph.
	graph := NewGraph()

	// Add all nodes explicitly to ensure inclusion of isolated nodes.
	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddNode("E")

	// Add edges to the graph.
	graph.AddEdge("A", "B", 1)
	graph.AddEdge("A", "C", 4)
	graph.AddEdge("B", "C", 2)
	graph.AddEdge("B", "D", 6)
	graph.AddEdge("C", "D", 3)
	graph.AddEdge("C", "E", 5)
	graph.AddEdge("D", "E", 1)

	// Compute shortest distances and paths from node "A".
	distances, previous := graph.Dijkstra("A")

	// Collect all nodes and sort them.
	nodes := make([]string, 0, len(graph.Edges))
	for node := range graph.Edges {
		nodes = append(nodes, node)
	}
	sort.Strings(nodes) // Sort nodes alphabetically.

	// Print the shortest distances from node "A".
	fmt.Println("Shortest distances from A:")
	for _, node := range nodes {
		distance := distances[node]
		if distance == math.MaxInt {
			fmt.Printf("To %s: unreachable\n", node) // Print unreachable nodes.
		} else {
			fmt.Printf("To %s: %d\n", node, distance) // Print the shortest distance.
		}
	}

	// Print the shortest paths from node "A".
	fmt.Println("\nShortest paths from A:")
	for _, node := range nodes {
		path := ReconstructPath(previous, "A", node)
		if len(path) > 0 {
			fmt.Printf("To %s: %s\n", node, joinPath(path))
		} else {
			fmt.Printf("To %s: No path found\n", node)
		}
	}
}
