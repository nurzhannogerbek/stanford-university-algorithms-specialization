package main

import (
	"container/heap"
	"fmt"
)

// Edge represents an edge in a graph with a source, target, and weight.
type Edge struct {
	Source int
	Target int
	Weight int
}

// Graph represents an undirected weighted graph.
type Graph struct {
	NumVertices int
	AdjList     map[int][]Edge
}

// NewGraph creates a new graph with a specified number of vertices.
func NewGraph(numVertices int) *Graph {
	return &Graph{
		NumVertices: numVertices,
		AdjList:     make(map[int][]Edge),
	}
}

// AddEdge adds an undirected edge to the graph.
func (g *Graph) AddEdge(source, target, weight int) {
	g.AdjList[source] = append(g.AdjList[source], Edge{Source: source, Target: target, Weight: weight})
	g.AdjList[target] = append(g.AdjList[target], Edge{Source: target, Target: source, Weight: weight})
}

// PriorityQueue implements a priority queue for edges.
type PriorityQueue []Edge

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Weight < pq[j].Weight }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

// Push adds an element to the queue.
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Edge))
}

// Pop removes and returns the smallest element from the queue.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// PrimMST represents the Prim's algorithm for finding MST.
type PrimMST struct {
	Graph     *Graph
	TotalCost int
	MSTEdges  []Edge
}

// NewPrimMST initializes a new instance of PrimMST.
func NewPrimMST(graph *Graph) *PrimMST {
	return &PrimMST{
		Graph:     graph,
		TotalCost: 0,
		MSTEdges:  []Edge{},
	}
}

// FindMST computes the minimum spanning tree using Prim's algorithm.
func (pmst *PrimMST) FindMST(startVertex int) {
	visited := make([]bool, pmst.Graph.NumVertices)
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Add initial edges from the start vertex.
	pmst.addEdges(startVertex, visited, pq)

	for pq.Len() > 0 {
		// Get the edge with the smallest weight.
		edge := heap.Pop(pq).(Edge)

		// Skip if the target vertex is already visited.
		if visited[edge.Target] {
			continue
		}

		// Include the edge in the MST.
		pmst.MSTEdges = append(pmst.MSTEdges, edge)
		pmst.TotalCost += edge.Weight

		// Add edges from the newly added vertex.
		pmst.addEdges(edge.Target, visited, pq)
	}
}

// addEdges pushes all valid edges from a vertex into the priority queue.
func (pmst *PrimMST) addEdges(vertex int, visited []bool, pq *PriorityQueue) {
	visited[vertex] = true
	for _, edge := range pmst.Graph.AdjList[vertex] {
		if !visited[edge.Target] {
			heap.Push(pq, edge)
		}
	}
}

func main() {
	// Create a graph with 6 vertices.
	graph := NewGraph(6)
	graph.AddEdge(0, 1, 4)
	graph.AddEdge(0, 2, 4)
	graph.AddEdge(1, 2, 2)
	graph.AddEdge(1, 3, 5)
	graph.AddEdge(2, 3, 8)
	graph.AddEdge(2, 4, 10)
	graph.AddEdge(3, 4, 2)
	graph.AddEdge(3, 5, 6)
	graph.AddEdge(4, 5, 3)

	// Run Prim's algorithm.
	primMST := NewPrimMST(graph)
	primMST.FindMST(0)

	// Print the result.
	fmt.Println("Edges in the MST:")
	for _, edge := range primMST.MSTEdges {
		fmt.Printf("%d -- %d (%d)\n", edge.Source, edge.Target, edge.Weight)
	}
	fmt.Printf("Total cost of the MST: %d\n", primMST.TotalCost)
}
