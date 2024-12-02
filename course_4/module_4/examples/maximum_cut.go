package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Graph represents an undirected graph with vertices and edges.
type Graph struct {
	Vertices []string
	Edges    [][2]string
}

// NewGraph creates a new Graph instance and validates its input.
func NewGraph(vertices []string, edges [][2]string) (*Graph, error) {
	vertexSet := make(map[string]bool)

	for _, vertex := range vertices {
		vertexSet[vertex] = true
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if !vertexSet[u] || !vertexSet[v] {
			return nil, fmt.Errorf("invalid edge (%s, %s): vertices not found", u, v)
		}
		if u == v {
			return nil, fmt.Errorf("self-loop detected at vertex (%s)", u)
		}
	}

	return &Graph{Vertices: vertices, Edges: edges}, nil
}

// String provides a string representation of the Graph.
func (g *Graph) String() string {
	return fmt.Sprintf("Graph(vertices=%d, edges=%d)", len(g.Vertices), len(g.Edges))
}

// MaximumCut represents the Maximum Cut problem solver.
type MaximumCut struct {
	Graph      *Graph
	PartitionA map[string]bool
	PartitionB map[string]bool
}

// NewMaximumCut initializes a new MaximumCut solver for the given graph.
func NewMaximumCut(graph *Graph) *MaximumCut {
	return &MaximumCut{
		Graph:      graph,
		PartitionA: make(map[string]bool),
		PartitionB: make(map[string]bool),
	}
}

// InitializePartition creates an initial random partition of vertices.
func (mc *MaximumCut) InitializePartition() {
	rand.Seed(42) // Fix the random seed for consistency.
	shuffledVertices := append([]string{}, mc.Graph.Vertices...)
	sort.Strings(shuffledVertices) // Sort vertices for consistent processing order.
	rand.Shuffle(len(shuffledVertices), func(i, j int) {
		shuffledVertices[i], shuffledVertices[j] = shuffledVertices[j], shuffledVertices[i]
	})

	mid := len(shuffledVertices) / 2
	for i, vertex := range shuffledVertices {
		if i < mid {
			mc.PartitionA[vertex] = true
		} else {
			mc.PartitionB[vertex] = true
		}
	}
}

// ComputeCutValue calculates the current cut value.
func (mc *MaximumCut) ComputeCutValue() int {
	cutValue := 0
	for _, edge := range mc.Graph.Edges {
		u, v := edge[0], edge[1]
		if (mc.PartitionA[u] && mc.PartitionB[v]) || (mc.PartitionB[u] && mc.PartitionA[v]) {
			cutValue++
		}
	}
	return cutValue
}

// ImprovePartition tries to improve the cut value by moving vertices.
func (mc *MaximumCut) ImprovePartition() bool {
	improved := false
	allVertices := append([]string{}, mc.Graph.Vertices...)
	sort.Strings(allVertices) // Consistent processing order.

	for _, vertex := range allVertices {
		currentGroup := mc.PartitionA
		otherGroup := mc.PartitionB
		if mc.PartitionB[vertex] {
			currentGroup = mc.PartitionB
			otherGroup = mc.PartitionA
		}

		// Count crossing and non-crossing edges for the vertex.
		crossingEdges := 0
		nonCrossingEdges := 0
		for _, edge := range mc.Graph.Edges {
			u, v := edge[0], edge[1]
			if (u == vertex && currentGroup[v]) || (v == vertex && currentGroup[u]) {
				nonCrossingEdges++
			} else if (u == vertex && otherGroup[v]) || (v == vertex && otherGroup[u]) {
				crossingEdges++
			}
		}

		// Move vertex only if it improves the cut value.
		if nonCrossingEdges > crossingEdges {
			delete(currentGroup, vertex)
			otherGroup[vertex] = true
			improved = true
		}
	}
	return improved
}

// Solve solves the Maximum Cut problem using local search.
func (mc *MaximumCut) Solve() int {
	mc.InitializePartition()
	for mc.ImprovePartition() {
	}
	return mc.ComputeCutValue()
}

func main() {
	vertices := []string{"A", "B", "C", "D"}
	edges := [][2]string{
		{"A", "B"}, {"B", "C"}, {"C", "D"}, {"D", "A"}, {"A", "C"},
	}

	graph, err := NewGraph(vertices, edges)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	maxCutSolver := NewMaximumCut(graph)
	maxCutValue := maxCutSolver.Solve()

	fmt.Println(graph) // Print graph representation.
	fmt.Printf("Maximum cut value: %d\n", maxCutValue)
	fmt.Printf("Partition A: %v\n", keys(maxCutSolver.PartitionA))
	fmt.Printf("Partition B: %v\n", keys(maxCutSolver.PartitionB))
}

// Helper function to extract keys from a map as a slice.
func keys(m map[string]bool) []string {
	result := []string{}
	for key := range m {
		result = append(result, key)
	}
	return result
}
