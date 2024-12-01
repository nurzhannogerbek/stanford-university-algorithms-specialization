package main

import (
	"fmt"
	"sort"
)

// Edge represents an edge in the graph.
type Edge struct {
	U, V   int     // The vertices connected by the edge.
	Weight float64 // The weight of the edge.
}

// Graph represents a graph with vertices and edges.
type Graph struct {
	Vertices int    // Number of vertices in the graph.
	Edges    []Edge // List of edges in the graph.
}

// AddEdge adds an edge to the graph.
func (g *Graph) AddEdge(u, v int, weight float64) {
	if u < 0 || v < 0 || u >= g.Vertices || v >= g.Vertices {
		panic(fmt.Sprintf("Vertices %d and %d must be between 0 and %d", u, v, g.Vertices-1))
	}
	g.Edges = append(g.Edges, Edge{U: u, V: v, Weight: weight})
}

// UnionFind is a structure for managing disjoint sets.
type UnionFind struct {
	Parent []int // Parent pointers for each element.
	Rank   []int // Rank (depth) of each tree.
}

// NewUnionFind creates a new UnionFind for the given size.
func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{Parent: parent, Rank: rank}
}

// Find finds the root of the set containing x, using path compression.
func (uf *UnionFind) Find(x int) int {
	if uf.Parent[x] != x {
		uf.Parent[x] = uf.Find(uf.Parent[x]) // Path compression.
	}
	return uf.Parent[x]
}

// Union unites the sets containing x and y, using union by rank.
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		if uf.Rank[rootX] > uf.Rank[rootY] {
			uf.Parent[rootY] = rootX
		} else if uf.Rank[rootX] < uf.Rank[rootY] {
			uf.Parent[rootX] = rootY
		} else {
			uf.Parent[rootY] = rootX
			uf.Rank[rootX]++
		}
	}
}

// KruskalMST computes the Minimum Spanning Tree using Kruskal's algorithm.
func (g *Graph) KruskalMST() ([]Edge, float64) {
	// Sort edges by weight.
	sort.Slice(g.Edges, func(i, j int) bool {
		return g.Edges[i].Weight < g.Edges[j].Weight
	})

	uf := NewUnionFind(g.Vertices)

	var mst []Edge
	var totalWeight float64

	// Process edges in order of increasing weight.
	for _, edge := range g.Edges {
		if uf.Find(edge.U) != uf.Find(edge.V) {
			uf.Union(edge.U, edge.V)
			mst = append(mst, edge)
			totalWeight += edge.Weight
		}
	}

	return mst, totalWeight
}

func main() {
	// Create a graph with 5 vertices.
	graph := &Graph{Vertices: 5}

	// Add edges to the graph.
	graph.AddEdge(0, 1, 1)
	graph.AddEdge(0, 2, 3)
	graph.AddEdge(1, 2, 2)
	graph.AddEdge(1, 3, 5)
	graph.AddEdge(2, 3, 4)
	graph.AddEdge(3, 4, 6)

	// Compute the Minimum Spanning Tree.
	mst, totalWeight := graph.KruskalMST()

	// Output the results.
	fmt.Println("Edges in the MST:")
	for _, edge := range mst {
		fmt.Printf("Edge(%d, %d, %.2f)\n", edge.U, edge.V, edge.Weight)
	}
	fmt.Printf("Total weight of the MST: %.2f\n", totalWeight)
}
