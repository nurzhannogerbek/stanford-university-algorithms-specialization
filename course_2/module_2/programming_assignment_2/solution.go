package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const INF = int(^uint(0) >> 1) // Define infinity as the maximum possible integer value.

// Graph represents a weighted graph using an adjacency list.
type Graph struct {
	adjList map[int][]Edge // Adjacency list where the key is the source node and the value is a list of edges.
}

// Edge represents a weighted edge connecting two nodes.
type Edge struct {
	target int // The target node of the edge.
	weight int // The weight of the edge.
}

// NewGraph initializes and returns a new graph.
func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]Edge)}
}

// AddEdge adds a weighted edge from source to target in the graph.
func (g *Graph) AddEdge(source, target, weight int) {
	g.adjList[source] = append(g.adjList[source], Edge{target, weight})
}

// LoadGraph loads a graph from a file and populates the adjacency list.
func (g *Graph) LoadGraph(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue // Skip empty lines.
		}
		parts := strings.Fields(line)
		source, _ := strconv.Atoi(parts[0])
		source-- // Convert to 0-based indexing.

		for _, edge := range parts[1:] {
			targetWeight := strings.Split(edge, ",")
			if len(targetWeight) != 2 {
				return fmt.Errorf("invalid edge format: %s", edge)
			}
			target, _ := strconv.Atoi(targetWeight[0])
			weight, _ := strconv.Atoi(targetWeight[1])
			g.AddEdge(source, target-1, weight) // Convert target to 0-based indexing.
		}
	}
	return nil
}

// Dijkstra runs Dijkstra's algorithm to calculate the shortest distances from the start node.
func (g *Graph) Dijkstra(start int) *DijkstraResult {
	// Initialize distances with infinity and a visited map.
	distances := make(map[int]int)
	visited := make(map[int]bool)
	for node := range g.adjList {
		distances[node] = INF
	}
	distances[start] = 0 // Distance to the start node is 0.

	// Priority queue to process nodes in order of shortest distance.
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, distance: 0})

	for pq.Len() > 0 {
		// Get the node with the smallest distance.
		current := heap.Pop(pq).(*Item)
		if visited[current.node] {
			continue // Skip already visited nodes.
		}
		visited[current.node] = true

		// Update distances for neighboring nodes.
		for _, edge := range g.adjList[current.node] {
			if visited[edge.target] {
				continue
			}
			newDistance := distances[current.node] + edge.weight
			if newDistance < distances[edge.target] {
				distances[edge.target] = newDistance
				heap.Push(pq, &Item{node: edge.target, distance: newDistance})
			}
		}
	}

	// Return the shortest distances as a result.
	return &DijkstraResult{distances: distances}
}

// DijkstraResult holds the shortest distances calculated by Dijkstra's algorithm.
type DijkstraResult struct {
	distances map[int]int // Map of node to shortest distance from the start node.
}

// PrintTargetDistances prints the shortest distances to the specified target nodes.
func (r *DijkstraResult) PrintTargetDistances(targets []int) {
	results := []string{}
	for _, target := range targets {
		if distance, found := r.distances[target]; found {
			results = append(results, strconv.Itoa(distance))
		} else {
			results = append(results, "unreachable") // For nodes that cannot be reached.
		}
	}
	fmt.Println(strings.Join(results, ","))
}

// Item represents an item in the priority queue with a node and its current distance.
type Item struct {
	node     int // The node ID.
	distance int // The current shortest distance to this node.
}

// PriorityQueue implements a priority queue for Dijkstra's algorithm.
type PriorityQueue []*Item

// Len returns the length of the priority queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less compares two items based on their distance.
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }

// Swap swaps two items in the priority queue.
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// Push adds an item to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Item)) }

// Pop removes and returns the item with the smallest distance from the priority queue.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// main is the entry point of the program.
func main() {
	// Create a new graph.
	graph := NewGraph()

	// Load the graph from the file.
	if err := graph.LoadGraph("course_2/module_2/programming_assignment_2/dijkstraData.txt"); err != nil {
		log.Fatalf("Error loading graph: %v", err)
	}

	// Run Dijkstra's algorithm from the start node (0-based index).
	result := graph.Dijkstra(0)

	// Define the target nodes for which distances need to be printed (0-based indexing).
	targets := []int{6, 36, 58, 81, 98, 114, 132, 164, 187, 196}

	// Print the shortest distances to the target nodes.
	result.PrintTargetDistances(targets)
}
