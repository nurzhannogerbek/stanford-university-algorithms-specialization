package main

import (
	"fmt"
	"math"
)

// FloydWarshall represents a graph and contains methods to compute shortest paths.
type FloydWarshall struct {
	vertices  int
	distances [][]float64
	next      [][]int
}

// NewFloydWarshall creates a new instance of the FloydWarshall algorithm.
func NewFloydWarshall(vertices int) *FloydWarshall {
	distances := make([][]float64, vertices)
	next := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		distances[i] = make([]float64, vertices)
		next[i] = make([]int, vertices)
		for j := 0; j < vertices; j++ {
			if i == j {
				distances[i][j] = 0
			} else {
				distances[i][j] = math.Inf(1)
			}
			next[i][j] = -1
		}
	}
	return &FloydWarshall{vertices, distances, next}
}

// AddEdge adds an edge to the graph with a specified weight.
func (fw *FloydWarshall) AddEdge(u, v int, weight float64) error {
	if err := fw.validateVertex(u); err != nil {
		return err
	}
	if err := fw.validateVertex(v); err != nil {
		return err
	}
	fw.distances[u][v] = weight
	fw.next[u][v] = v
	return nil
}

// validateVertex checks if a vertex index is valid.
func (fw *FloydWarshall) validateVertex(v int) error {
	if v < 0 || v >= fw.vertices {
		return fmt.Errorf("vertex index %d is out of range 0 to %d", v, fw.vertices-1)
	}
	return nil
}

// RunAlgorithm runs the Floyd-Warshall algorithm to compute shortest paths.
func (fw *FloydWarshall) RunAlgorithm() {
	for k := 0; k < fw.vertices; k++ {
		for i := 0; i < fw.vertices; i++ {
			for j := 0; j < fw.vertices; j++ {
				if fw.distances[i][k]+fw.distances[k][j] < fw.distances[i][j] {
					fw.distances[i][j] = fw.distances[i][k] + fw.distances[k][j]
					fw.next[i][j] = fw.next[i][k]
				}
			}
		}
	}
}

// HasNegativeCycle checks if the graph contains a negative weight cycle.
func (fw *FloydWarshall) HasNegativeCycle() bool {
	for i := 0; i < fw.vertices; i++ {
		if fw.distances[i][i] < 0 {
			return true
		}
	}
	return false
}

// GetShortestDistance returns the shortest distance between two vertices.
func (fw *FloydWarshall) GetShortestDistance(u, v int) (float64, error) {
	if err := fw.validateVertex(u); err != nil {
		return math.Inf(1), err
	}
	if err := fw.validateVertex(v); err != nil {
		return math.Inf(1), err
	}
	return fw.distances[u][v], nil
}

// ReconstructPath reconstructs the shortest path between two vertices.
func (fw *FloydWarshall) ReconstructPath(u, v int) ([]int, error) {
	if err := fw.validateVertex(u); err != nil {
		return nil, err
	}
	if err := fw.validateVertex(v); err != nil {
		return nil, err
	}

	if fw.distances[u][v] == math.Inf(1) {
		return []int{}, nil
	}

	path := []int{u}
	for u != v {
		u = fw.next[u][v]
		if u == -1 {
			return nil, fmt.Errorf("no path exists between the given vertices")
		}
		path = append(path, u)
	}

	return path, nil
}

func main() {
	graph := NewFloydWarshall(4)

	if err := graph.AddEdge(0, 1, 5); err != nil {
		fmt.Println("Error adding edge:", err)
	}
	if err := graph.AddEdge(0, 3, 10); err != nil {
		fmt.Println("Error adding edge:", err)
	}
	if err := graph.AddEdge(1, 2, 3); err != nil {
		fmt.Println("Error adding edge:", err)
	}
	if err := graph.AddEdge(2, 3, 1); err != nil {
		fmt.Println("Error adding edge:", err)
	}

	graph.RunAlgorithm()

	if graph.HasNegativeCycle() {
		fmt.Println("The graph contains a negative weight cycle.")
	} else {
		dist, err := graph.GetShortestDistance(0, 3)
		if err != nil {
			fmt.Println("Error getting shortest distance:", err)
		} else {
			fmt.Printf("Shortest distance from 0 to 3: %v\n", dist)
		}

		path, err := graph.ReconstructPath(0, 3)
		if err != nil {
			fmt.Println("Error reconstructing path:", err)
		} else {
			fmt.Printf("Shortest path from 0 to 3: %v\n", path)
		}
	}
}
