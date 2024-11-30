package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Edge represents an edge in a graph with source, target, and cost.
type Edge struct {
	Source int
	Target int
	Cost   int
}

// PriorityQueue implements a priority queue for edges.
type PriorityQueue []*Edge

// Len returns the length of the priority queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less compares two edges by their cost.
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Cost < pq[j].Cost }

// Swap swaps two edges in the priority queue.
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// Push adds an edge to the priority queue.
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Edge)) }

// Pop removes and returns the smallest edge in the priority queue.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Graph represents an undirected weighted graph.
type Graph struct {
	AdjacencyList map[int][]*Edge
}

// NewGraph creates a new graph with the specified number of nodes.
func NewGraph(numNodes int) *Graph {
	return &Graph{
		AdjacencyList: make(map[int][]*Edge, numNodes),
	}
}

// AddEdge adds an undirected edge to the graph.
func (g *Graph) AddEdge(source, target, cost int) {
	g.AdjacencyList[source] = append(g.AdjacencyList[source], &Edge{Source: source, Target: target, Cost: cost})
	g.AdjacencyList[target] = append(g.AdjacencyList[target], &Edge{Source: target, Target: source, Cost: cost})
}

// PrimMST implements Prim's algorithm to find the MST.
func PrimMST(graph *Graph, numNodes int) int {
	visited := make(map[int]bool)
	pq := &PriorityQueue{}
	heap.Init(pq)
	totalCost := 0

	// Start with an arbitrary node, e.g., node 1.
	startNode := 1
	visited[startNode] = true
	for _, edge := range graph.AdjacencyList[startNode] {
		heap.Push(pq, edge)
	}

	// Process the priority queue.
	for pq.Len() > 0 && len(visited) < numNodes {
		edge := heap.Pop(pq).(*Edge)
		if visited[edge.Target] {
			continue
		}

		// Add edge to the MST.
		totalCost += edge.Cost
		visited[edge.Target] = true

		// Add all edges from the newly visited node.
		for _, nextEdge := range graph.AdjacencyList[edge.Target] {
			if !visited[nextEdge.Target] {
				heap.Push(pq, nextEdge)
			}
		}
	}

	// Check if the graph is connected.
	if len(visited) != numNodes {
		fmt.Println("Error: The graph is not connected.")
		os.Exit(1)
	}

	return totalCost
}

func main() {
	// Укажите путь к файлу явно.
	filePath := "course_3/module_1/programming_assignment_1/edges.txt"

	// Read graph data from the specified file.
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer file.Close()

	var numNodes, numEdges int
	fmt.Fscanf(file, "%d %d\n", &numNodes, &numEdges)

	graph := NewGraph(numNodes)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) != 3 {
			fmt.Printf("Invalid line format: %v\n", line)
			continue
		}
		source, err1 := strconv.Atoi(line[0])
		target, err2 := strconv.Atoi(line[1])
		cost, err3 := strconv.Atoi(line[2])
		if err1 != nil || err2 != nil || err3 != nil {
			fmt.Printf("Invalid edge data: %v\n", line)
			continue
		}
		graph.AddEdge(source, target, cost)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Calculate the total cost of the MST.
	totalCost := PrimMST(graph, numNodes)

	// Print only the total cost of the MST with descriptive text.
	fmt.Printf("Total cost of the MST: %d\n", totalCost)
}
