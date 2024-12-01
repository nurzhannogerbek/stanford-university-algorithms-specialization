package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents a pair of points and the distance between them.
type Edge struct {
	Distance float64 // The distance between two points.
	Point1   int     // The index of the first point.
	Point2   int     // The index of the second point.
}

// PriorityQueue implements a min-heap for edges.
type PriorityQueue []Edge

// Len returns the number of elements in the priority queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less compares two edges based on their distances.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

// Swap exchanges two elements in the priority queue.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds a new edge to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Edge))
}

// Pop removes and returns the edge with the smallest distance.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// UnionFind is a structure for managing clusters.
type UnionFind struct {
	Parent []int // Array representing the parent of each element.
	Rank   []int // Array representing the rank of each tree.
}

// NewUnionFind initializes the Union-Find structure.
func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := range parent {
		parent[i] = i // Initially, each element is its own parent.
	}
	return &UnionFind{Parent: parent, Rank: rank}
}

// Find returns the root of the cluster for a given element using path compression.
func (uf *UnionFind) Find(x int) int {
	if uf.Parent[x] != x {
		uf.Parent[x] = uf.Find(uf.Parent[x]) // Path compression to optimize future finds.
	}
	return uf.Parent[x]
}

// Union merges two clusters by rank.
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		// Merge trees based on rank to maintain balance.
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

// ClusteringAlgorithm represents the greedy clustering algorithm.
type ClusteringAlgorithm struct {
	Points          []string      // List of point names.
	DistanceMatrix  [][]float64   // Matrix of distances between points.
	NumClusters     int           // Desired number of clusters.
	UnionFind       *UnionFind    // Union-Find structure to manage clusters.
	PriorityQueue   PriorityQueue // Min-heap for edges.
	CurrentClusters int           // Current number of clusters.
	MinSpacing      float64       // Minimum spacing between clusters.
}

// validateDistanceMatrix validates the input distance matrix for correctness.
func validateDistanceMatrix(matrix [][]float64) error {
	n := len(matrix)
	for i := 0; i < n; i++ {
		if len(matrix[i]) != n {
			return fmt.Errorf("distance matrix is not square")
		}
		if matrix[i][i] != 0 {
			return fmt.Errorf("distance matrix diagonal must be zero")
		}
		for j := i + 1; j < n; j++ {
			if matrix[i][j] != matrix[j][i] {
				return fmt.Errorf("distance matrix is not symmetric")
			}
		}
	}
	return nil
}

// NewClusteringAlgorithm initializes the clustering algorithm.
func NewClusteringAlgorithm(points []string, distanceMatrix [][]float64, numClusters int) *ClusteringAlgorithm {
	if err := validateDistanceMatrix(distanceMatrix); err != nil {
		panic(err)
	}

	pq := PriorityQueue{}
	heap.Init(&pq)
	// Add all edges to the priority queue.
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			heap.Push(&pq, Edge{Distance: distanceMatrix[i][j], Point1: i, Point2: j})
		}
	}
	return &ClusteringAlgorithm{
		Points:          points,
		DistanceMatrix:  distanceMatrix,
		NumClusters:     numClusters,
		UnionFind:       NewUnionFind(len(points)),
		PriorityQueue:   pq,
		CurrentClusters: len(points),
		MinSpacing:      math.MaxFloat64,
	}
}

// Cluster performs the clustering until the desired number of clusters is reached.
func (ca *ClusteringAlgorithm) Cluster() [][]string {
	for ca.CurrentClusters > ca.NumClusters {
		// Get the edge with the smallest distance.
		edge := heap.Pop(&ca.PriorityQueue).(Edge)
		// Merge clusters if they are not already connected.
		if ca.UnionFind.Find(edge.Point1) != ca.UnionFind.Find(edge.Point2) {
			ca.UnionFind.Union(edge.Point1, edge.Point2)
			ca.CurrentClusters--
		}
	}

	// Calculate the minimum spacing between clusters.
	ca.MinSpacing = math.MaxFloat64
	for _, edge := range ca.PriorityQueue {
		if ca.UnionFind.Find(edge.Point1) != ca.UnionFind.Find(edge.Point2) {
			if edge.Distance < ca.MinSpacing {
				ca.MinSpacing = edge.Distance
			}
		}
	}

	// Build the final clusters.
	clusters := make(map[int][]string)
	for i, point := range ca.Points {
		root := ca.UnionFind.Find(i)
		clusters[root] = append(clusters[root], point)
	}

	// Convert the map of clusters into a slice of slices.
	result := make([][]string, 0, len(clusters))
	for _, cluster := range clusters {
		result = append(result, cluster)
	}
	return result
}

// Main function to execute the clustering algorithm.
func main() {
	// Define the points and the distance matrix.
	points := []string{"A", "B", "C", "D", "E"}
	distanceMatrix := [][]float64{
		{0, 2, 6, 10, 9},
		{2, 0, 4, 8, 7},
		{6, 4, 0, 3, 5},
		{10, 8, 3, 0, 6},
		{9, 7, 5, 6, 0},
	}
	numClusters := 2 // Desired number of clusters.

	// Initialize and run the clustering algorithm.
	algorithm := NewClusteringAlgorithm(points, distanceMatrix, numClusters)
	clusters := algorithm.Cluster()

	// Print the clustering result.
	fmt.Println("Clustering result:")
	for i, cluster := range clusters {
		fmt.Printf("Cluster %d: %v\n", i+1, cluster)
	}

	// Print the minimum spacing between clusters.
	fmt.Printf("\nSpacing: %.2f\n", algorithm.MinSpacing)
}
