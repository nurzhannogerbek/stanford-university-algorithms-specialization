package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Graph represents an undirected graph with parallel edges.
type Graph struct {
	AdjacencyList map[int][]int
}

// NewGraph creates a new graph based on the given adjacency list.
func NewGraph(adjacencyList map[int][]int) *Graph {
	newList := make(map[int][]int)
	for k, v := range adjacencyList {
		newList[k] = append([]int(nil), v...)
	}
	return &Graph{AdjacencyList: newList}
}

// ContractEdge contracts two nodes (u and v) by merging them into a single node.
func (g *Graph) ContractEdge(u, v int) {
	// Add edges of v to u.
	g.AdjacencyList[u] = append(g.AdjacencyList[u], g.AdjacencyList[v]...)

	// Replace all occurrences of v with u in the graph.
	for _, neighbor := range g.AdjacencyList[v] {
		for i, val := range g.AdjacencyList[neighbor] {
			if val == v {
				g.AdjacencyList[neighbor][i] = u
			}
		}
	}

	// Remove self-loops.
	filtered := g.AdjacencyList[u][:0]
	for _, val := range g.AdjacencyList[u] {
		if val != u {
			filtered = append(filtered, val)
		}
	}
	g.AdjacencyList[u] = filtered

	// Delete vertex v from the graph.
	delete(g.AdjacencyList, v)
}

// GetRandomEdge selects a random edge (u, v) from the graph.
func (g *Graph) GetRandomEdge() (int, int) {
	vertices := make([]int, 0, len(g.AdjacencyList))
	for vertex := range g.AdjacencyList {
		vertices = append(vertices, vertex)
	}

	u := vertices[rand.Intn(len(vertices))]
	v := g.AdjacencyList[u][rand.Intn(len(g.AdjacencyList[u]))]
	return u, v
}

// GetMinCut runs the contraction algorithm to compute the minimum cut of the graph.
func (g *Graph) GetMinCut() int {
	for len(g.AdjacencyList) > 2 {
		u, v := g.GetRandomEdge()
		g.ContractEdge(u, v)
	}

	// After contraction, the remaining edges define the minimum cut.
	for _, edges := range g.AdjacencyList {
		return len(edges)
	}
	return 0
}

// RandomizedContractionAlgorithm represents the randomized contraction algorithm.
type RandomizedContractionAlgorithm struct {
	OriginalAdjacencyList map[int][]int
}

// NewRandomizedContractionAlgorithm creates a new instance of the algorithm.
func NewRandomizedContractionAlgorithm(adjacencyList map[int][]int) *RandomizedContractionAlgorithm {
	return &RandomizedContractionAlgorithm{OriginalAdjacencyList: adjacencyList}
}

// FindMinCut runs multiple trials of the contraction algorithm to find the minimum cut.
func (rca *RandomizedContractionAlgorithm) FindMinCut(trials int, parallel bool) int {
	if trials == 0 {
		n := len(rca.OriginalAdjacencyList)
		trials = n * n
	}

	minCut := int(^uint(0) >> 1) // Initialize with the maximum possible integer value.

	if parallel {
		var wg sync.WaitGroup
		resultChan := make(chan int, trials)

		// Run trials in parallel.
		for i := 0; i < trials; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				graph := NewGraph(rca.OriginalAdjacencyList)
				resultChan <- graph.GetMinCut()
			}()
		}

		// Wait for all trials to complete and collect results.
		wg.Wait()
		close(resultChan)

		for result := range resultChan {
			if result < minCut {
				minCut = result
			}
		}
	} else {
		// Run trials sequentially.
		for i := 0; i < trials; i++ {
			graph := NewGraph(rca.OriginalAdjacencyList)
			if cut := graph.GetMinCut(); cut < minCut {
				minCut = cut
			}
		}
	}

	return minCut
}

func main() {
	// Example adjacency list.
	adjacencyList := map[int][]int{
		1: {2, 3, 4},
		2: {1, 3, 4},
		3: {1, 2, 4},
		4: {1, 2, 3},
	}

	rand.Seed(time.Now().UnixNano())

	// Create an instance of the algorithm and find the minimum cut.
	rca := NewRandomizedContractionAlgorithm(adjacencyList)
	minCut := rca.FindMinCut(0, true) // 0 to calculate trials, true for parallel execution.
	fmt.Printf("Minimum cut: %d\n", minCut)
}
