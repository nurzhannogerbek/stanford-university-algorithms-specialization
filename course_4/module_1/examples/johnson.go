package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents a directed edge in the graph.
type Edge struct {
	From, To int
	Weight   float64
}

// Graph represents a directed graph with a list of edges and vertices.
type Graph struct {
	Vertices int
	Edges    []Edge
}

// NewGraph creates a new graph with the specified number of vertices.
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		Edges:    []Edge{},
	}
}

// AddEdge adds a directed edge to the graph.
func (g *Graph) AddEdge(from, to int, weight float64) {
	g.Edges = append(g.Edges, Edge{From: from, To: to, Weight: weight})
}

// BellmanFord computes shortest paths from a source vertex using the Bellman-Ford algorithm.
// Returns a slice of distances or nil if a negative weight cycle is detected.
func (g *Graph) BellmanFord(source int) []float64 {
	distances := make([]float64, g.Vertices)
	for i := range distances {
		distances[i] = math.Inf(1)
	}
	distances[source] = 0

	// Relax edges |V|-1 times.
	for i := 0; i < g.Vertices-1; i++ {
		for _, edge := range g.Edges {
			if distances[edge.From]+edge.Weight < distances[edge.To] {
				distances[edge.To] = distances[edge.From] + edge.Weight
			}
		}
	}

	// Check for negative weight cycles.
	for _, edge := range g.Edges {
		if distances[edge.From]+edge.Weight < distances[edge.To] {
			return nil
		}
	}

	return distances
}

// Dijkstra computes shortest paths from a source vertex using Dijkstra's algorithm.
// Returns a slice of distances.
func (g *Graph) Dijkstra(source int, adjustedWeights map[int]map[int]float64) []float64 {
	distances := make([]float64, g.Vertices)
	for i := range distances {
		distances[i] = math.Inf(1)
	}
	distances[source] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{vertex: source, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		u := current.vertex

		// Skip if this distance is outdated.
		if current.priority > distances[u] {
			continue
		}

		// Relax neighbors.
		for v, weight := range adjustedWeights[u] {
			if distances[u]+weight < distances[v] {
				distances[v] = distances[u] + weight
				heap.Push(pq, &Item{vertex: v, priority: distances[v]})
			}
		}
	}

	return distances
}

// Johnson computes shortest paths between all pairs of vertices using Johnson's Algorithm.
// Returns a 2D slice of distances or nil if a negative weight cycle is detected.
func (g *Graph) Johnson() [][]float64 {
	// Step 1: Add a new vertex connected to all other vertices with zero-weight edges.
	extendedGraph := &Graph{
		Vertices: g.Vertices + 1,
		Edges:    append(g.Edges, createZeroWeightEdges(g.Vertices)...),
	}

	// Step 2: Run Bellman-Ford from the new vertex.
	h := extendedGraph.BellmanFord(g.Vertices)
	if h == nil {
		return nil // Negative weight cycle detected.
	}

	// Step 3: Reweight the edges to eliminate negative weights.
	adjustedWeights := make(map[int]map[int]float64)
	for _, edge := range g.Edges {
		if _, exists := adjustedWeights[edge.From]; !exists {
			adjustedWeights[edge.From] = make(map[int]float64)
		}
		adjustedWeights[edge.From][edge.To] = edge.Weight + h[edge.From] - h[edge.To]
	}

	// Step 4: Run Dijkstra for each vertex.
	allPairsDistances := make([][]float64, g.Vertices)
	for u := 0; u < g.Vertices; u++ {
		shortestFromU := g.Dijkstra(u, adjustedWeights)
		allPairsDistances[u] = make([]float64, g.Vertices)
		for v := 0; v < g.Vertices; v++ {
			// Adjust back to original weights.
			if shortestFromU[v] != math.Inf(1) {
				allPairsDistances[u][v] = shortestFromU[v] - h[u] + h[v]
			} else {
				allPairsDistances[u][v] = math.Inf(1)
			}
		}
	}

	return allPairsDistances
}

// createZeroWeightEdges creates zero-weight edges from a new vertex to all existing vertices.
func createZeroWeightEdges(vertices int) []Edge {
	edges := make([]Edge, vertices)
	for i := 0; i < vertices; i++ {
		edges[i] = Edge{From: vertices, To: i, Weight: 0}
	}
	return edges
}

// Priority Queue Implementation for Dijkstra's Algorithm.
type Item struct {
	vertex   int
	priority float64
	index    int
}

// PriorityQueue is a priority queue for Dijkstra's algorithm.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // For safety.
	*pq = old[0 : n-1]
	return item
}

// Main function demonstrating Johnson's Algorithm.
func main() {
	graph := NewGraph(5)

	// Add edges with their weights.
	graph.AddEdge(0, 1, 3)
	graph.AddEdge(0, 2, 8)
	graph.AddEdge(1, 3, 1)
	graph.AddEdge(2, 3, -4)
	graph.AddEdge(3, 4, 2)
	graph.AddEdge(4, 0, -1)

	// Run Johnson's Algorithm.
	distances := graph.Johnson()

	if distances == nil {
		fmt.Println("The graph contains a negative weight cycle.")
	} else {
		fmt.Println("Shortest distances between all pairs of vertices:")
		for i, row := range distances {
			fmt.Printf("From vertex %d: %v\n", i, row)
		}
	}
}
