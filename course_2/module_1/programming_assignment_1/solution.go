package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Graph structure to represent a directed graph.
type Graph struct {
	adj    map[int][]int
	revAdj map[int][]int
}

func newGraph() *Graph {
	return &Graph{
		adj:    make(map[int][]int),
		revAdj: make(map[int][]int),
	}
}

// Adds a directed edge to the graph.
func (g *Graph) addEdge(from, to int) {
	g.adj[from] = append(g.adj[from], to)
	g.revAdj[to] = append(g.revAdj[to], from)
}

// Performs DFS on the graph and fills the stack with the visited nodes.
func (g *Graph) dfs(v int, visited []bool, stack *[]int) {
	visited[v] = true
	for _, neighbor := range g.adj[v] {
		if !visited[neighbor] {
			g.dfs(neighbor, visited, stack)
		}
	}
	*stack = append(*stack, v)
}

// Performs reverse DFS on the reversed graph and returns the size of the SCC.
func (g *Graph) reverseDFS(v int, visited []bool) int {
	visited[v] = true
	size := 1
	for _, neighbor := range g.revAdj[v] {
		if !visited[neighbor] {
			size += g.reverseDFS(neighbor, visited)
		}
	}
	return size
}

// Implements Kosaraju's algorithm to find SCCs.
func (g *Graph) kosaraju(maxVertex int) []int {
	stack := make([]int, 0) // Use make() instead of []int{}.
	visited := make([]bool, maxVertex+1)

	// First pass: Fill the stack with finishing times.
	for node := 1; node <= maxVertex; node++ {
		if !visited[node] {
			g.dfs(node, visited, &stack)
		}
	}

	// Second pass: Find SCCs in the reversed graph.
	visited = make([]bool, maxVertex+1)
	sccSizes := make([]int, 0) // Use make() instead of []int{}.
	for i := len(stack) - 1; i >= 0; i-- {
		node := stack[i]
		if !visited[node] {
			size := g.reverseDFS(node, visited)
			sccSizes = append(sccSizes, size)
		}
	}

	// Sort SCC sizes in descending order.
	sort.Slice(sccSizes, func(i, j int) bool {
		return sccSizes[i] > sccSizes[j]
	})
	return sccSizes
}

func main() {
	// Open the input file.
	file, err := os.Open("course_2/module_1/programming_assignment_1/SCC.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Handle file close error with defer.
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Error closing file: %v\n", closeErr)
		}
	}()

	graph := newGraph()
	scanner := bufio.NewScanner(file)

	// Parse the input file and add edges to the graph.
	maxVertex := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		from, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Error parsing vertex: %s\n", parts[0])
			continue
		}
		to, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Error parsing vertex: %s\n", parts[1])
			continue
		}
		graph.addEdge(from, to)
		if from > maxVertex {
			maxVertex = from
		}
		if to > maxVertex {
			maxVertex = to
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Run Kosaraju's algorithm.
	sccSizes := graph.kosaraju(maxVertex)

	// Output the sizes of the 5 largest SCCs.
	for i := 0; i < 5; i++ {
		if i < len(sccSizes) {
			fmt.Print(sccSizes[i])
		} else {
			fmt.Print(0)
		}
		if i < 4 {
			fmt.Print(",")
		}
	}
	fmt.Println()
}
