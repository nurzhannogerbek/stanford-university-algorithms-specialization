package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Graph represents an undirected graph with parallel edges.
// The adjacency list maps each vertex to its list of connected vertices.
type Graph struct {
	AdjacencyList map[int][]int
}

// NewGraph creates a new graph from the given adjacency list.
// This function makes a deep copy of the adjacency list to ensure
// the original graph remains unmodified during operations.
func NewGraph(adjacencyList map[int][]int) *Graph {
	newList := make(map[int][]int)
	for key, value := range adjacencyList {
		newList[key] = append([]int(nil), value...)
	}
	return &Graph{AdjacencyList: newList}
}

// ContractEdge merges two vertices (u and v) into one.
// It combines all edges of v into u and removes self-loops caused by the contraction.
func (g *Graph) ContractEdge(u, v int) {
	// Add all edges of v to u.
	g.AdjacencyList[u] = append(g.AdjacencyList[u], g.AdjacencyList[v]...)

	// Replace all occurrences of v with u in the adjacency list.
	for _, neighbor := range g.AdjacencyList[v] {
		for i, node := range g.AdjacencyList[neighbor] {
			if node == v {
				g.AdjacencyList[neighbor][i] = u
			}
		}
	}

	// Remove self-loops (edges from u to itself).
	filtered := g.AdjacencyList[u][:0]
	for _, node := range g.AdjacencyList[u] {
		if node != u {
			filtered = append(filtered, node)
		}
	}
	g.AdjacencyList[u] = filtered

	// Delete the merged vertex v from the graph.
	delete(g.AdjacencyList, v)
}

// GetRandomEdge selects a random edge (u, v) from the graph.
// It chooses a random vertex u and then randomly selects one of its neighbors v.
func (g *Graph) GetRandomEdge() (int, int) {
	// Collect all vertices into a slice for random selection.
	vertices := make([]int, 0, len(g.AdjacencyList))
	for vertex := range g.AdjacencyList {
		vertices = append(vertices, vertex)
	}

	// Select a random vertex u and a random neighbor v.
	u := vertices[rand.Intn(len(vertices))]
	v := g.AdjacencyList[u][rand.Intn(len(g.AdjacencyList[u]))]
	return u, v
}

// GetMinCut runs the randomized contraction algorithm on the graph.
// It repeatedly contracts edges until only two vertices remain.
// The number of edges between the two remaining vertices represents the minimum cut.
func (g *Graph) GetMinCut() int {
	for len(g.AdjacencyList) > 2 {
		u, v := g.GetRandomEdge()
		g.ContractEdge(u, v)
	}

	// Return the number of edges between the two remaining vertices.
	for _, edges := range g.AdjacencyList {
		return len(edges)
	}
	return 0
}

// RandomizedContractionAlgorithm represents the overall algorithm
// and stores the original graph's adjacency list.
type RandomizedContractionAlgorithm struct {
	OriginalAdjacencyList map[int][]int
}

// NewRandomizedContractionAlgorithm creates a new instance of the algorithm,
// initializing it with the original adjacency list.
func NewRandomizedContractionAlgorithm(adjacencyList map[int][]int) *RandomizedContractionAlgorithm {
	return &RandomizedContractionAlgorithm{OriginalAdjacencyList: adjacencyList}
}

// FindMinCut performs multiple trials of the contraction algorithm to find the minimum cut.
// It can run in parallel if the `parallel` flag is set to true.
func (rca *RandomizedContractionAlgorithm) FindMinCut(trials int, parallel bool) int {
	// If the number of trials is not specified, default to n^2, where n is the number of vertices.
	if trials == 0 {
		n := len(rca.OriginalAdjacencyList)
		trials = n * n
	}

	minCut := int(^uint(0) >> 1) // Initialize to the maximum possible integer value.

	if parallel {
		// Run trials in parallel using goroutines.
		var wg sync.WaitGroup
		resultChan := make(chan int, trials)

		for i := 0; i < trials; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				graph := NewGraph(rca.OriginalAdjacencyList)
				resultChan <- graph.GetMinCut()
			}()
		}

		wg.Wait()
		close(resultChan)

		// Collect results from the channel and update the minimum cut.
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

// ParseAdjacencyList parses the adjacency list from the input file.
// Each line in the file represents a vertex and its neighbors.
func ParseAdjacencyList(filepath string) (map[int][]int, error) {
	adjacencyList := make(map[int][]int)
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue // Skip invalid lines.
		}

		vertex, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		neighbors := make([]int, len(parts)-1)
		for i, neighbor := range parts[1:] {
			neighbors[i], err = strconv.Atoi(neighbor)
			if err != nil {
				return nil, err
			}
		}
		adjacencyList[vertex] = neighbors
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return adjacencyList, nil
}

func main() {
	// Seed the random number generator for reproducibility.
	rand.Seed(time.Now().UnixNano())

	// Specify the path to the adjacency list file.
	filepath := "course_1/module_4/programming_assignment_4/kargerMinCut.txt"

	// Parse the adjacency list from the file.
	adjacencyList, err := ParseAdjacencyList(filepath)
	if err != nil {
		fmt.Printf("Error reading adjacency list: %v\n", err)
		return
	}

	// Create an instance of the RandomizedContractionAlgorithm.
	rca := NewRandomizedContractionAlgorithm(adjacencyList)
	fmt.Println("Running randomized contraction algorithm...")

	// Run the algorithm to find the minimum cut.
	minCut := rca.FindMinCut(0, true) // 0 for auto-calculating trials, true for parallel execution.
	fmt.Printf("Minimum cut: %d\n", minCut)
}
