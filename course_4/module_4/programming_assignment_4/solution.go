package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Graph represents a directed graph.
// It contains adjacency lists for both the original and reversed graphs.
type Graph struct {
	adj    map[int][]int // Adjacency list for the original graph.
	adjInv map[int][]int // Adjacency list for the reversed graph.
	nodes  int           // Total number of nodes in the graph.
}

// NewGraph initializes a new Graph object with the given number of nodes.
func NewGraph(nodes int) *Graph {
	return &Graph{
		adj:    make(map[int][]int),
		adjInv: make(map[int][]int),
		nodes:  nodes,
	}
}

// AddEdge adds a directed edge to the original graph and its reverse to the reversed graph.
func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)       // Add edge to the original graph.
	g.adjInv[v] = append(g.adjInv[v], u) // Add reversed edge to the reversed graph.
}

// SATSolver encapsulates the logic for solving the 2-SAT problem.
// It uses Kosaraju's algorithm to find strongly connected components (SCCs).
type SATSolver struct {
	graph        *Graph       // The graph representing the problem.
	visited      map[int]bool // Visited array for the original graph.
	visitedInv   map[int]bool // Visited array for the reversed graph.
	scc          map[int]int  // Maps each node to its SCC ID.
	stack        []int        // Stack used for Kosaraju's algorithm.
	counter      int          // Counter for assigning SCC IDs.
	numVariables int          // Number of variables in the 2-SAT problem.
}

// NewSATSolver initializes a new SATSolver object for a given number of variables.
func NewSATSolver(numVariables int) *SATSolver {
	graph := NewGraph(2 * numVariables) // Create a graph with 2 * numVariables nodes (positive and negative literals).
	return &SATSolver{
		graph:        graph,
		visited:      make(map[int]bool),
		visitedInv:   make(map[int]bool),
		scc:          make(map[int]int),
		numVariables: numVariables,
	}
}

// AddClause adds a clause to the implication graph.
// Each clause is represented as (a OR b), which is converted into two implications: NOT(a) → b and NOT(b) → a.
func (solver *SATSolver) AddClause(a, b int) {
	n := solver.numVariables
	if a > 0 && b > 0 {
		solver.graph.AddEdge(a+n, b) // NOT(a) → b.
		solver.graph.AddEdge(b+n, a) // NOT(b) → a.
	} else if a > 0 && b < 0 {
		solver.graph.AddEdge(a+n, -b+n) // NOT(a) → NOT(b).
		solver.graph.AddEdge(-b, a)     // b → a.
	} else if a < 0 && b > 0 {
		solver.graph.AddEdge(-a, b)     // a → b.
		solver.graph.AddEdge(b+n, -a+n) // NOT(b) → NOT(a).
	} else {
		solver.graph.AddEdge(-a, -b+n) // a → NOT(b).
		solver.graph.AddEdge(-b, -a+n) // b → NOT(a).
	}
}

// dfs performs a depth-first search on the original graph.
// It is used in the first pass of Kosaraju's algorithm to determine the order of nodes.
func (solver *SATSolver) dfs(node int) {
	if solver.visited[node] {
		return
	}
	solver.visited[node] = true
	for _, neighbor := range solver.graph.adj[node] {
		solver.dfs(neighbor)
	}
	solver.stack = append(solver.stack, node) // Push the node onto the stack after processing all its neighbors.
}

// dfsInv performs a depth-first search on the reversed graph.
// It is used in the second pass of Kosaraju's algorithm to identify SCCs.
func (solver *SATSolver) dfsInv(node int) {
	if solver.visitedInv[node] {
		return
	}
	solver.visitedInv[node] = true
	solver.scc[node] = solver.counter // Assign the current SCC ID to the node.
	for _, neighbor := range solver.graph.adjInv[node] {
		solver.dfsInv(neighbor)
	}
}

// Solve determines if the 2-SAT problem is satisfiable.
// It uses Kosaraju's algorithm to find SCCs and checks for contradictions.
func (solver *SATSolver) Solve() bool {
	// First pass: Perform DFS on the original graph to determine the finishing order of nodes.
	for i := 1; i <= 2*solver.numVariables; i++ {
		if !solver.visited[i] {
			solver.dfs(i)
		}
	}

	// Second pass: Perform DFS on the reversed graph to identify SCCs.
	for len(solver.stack) > 0 {
		node := solver.stack[len(solver.stack)-1]
		solver.stack = solver.stack[:len(solver.stack)-1]
		if !solver.visitedInv[node] {
			solver.counter++ // Increment the SCC ID counter.
			solver.dfsInv(node)
		}
	}

	// Check for contradictions: a variable and its negation must not be in the same SCC.
	for i := 1; i <= solver.numVariables; i++ {
		if solver.scc[i] == solver.scc[i+solver.numVariables] {
			return false // Unsatisfiable.
		}
	}
	return true // Satisfiable.
}

// ParseFile reads the input file and constructs a SATSolver object.
// The file contains the number of variables on the first line, followed by clauses on subsequent lines.
func ParseFile(filePath string) (*SATSolver, error) {
	file, err := os.Open(filePath) // Open the file.
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the number of variables.
	var numVariables int
	if scanner.Scan() {
		numVariables, err = strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("invalid variable count: %w", err)
		}
	}

	solver := NewSATSolver(numVariables)

	// Read each clause and add it to the solver.
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue // Skip invalid lines.
		}

		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		solver.AddClause(a, b)
	}

	return solver, nil
}

func main() {
	// Base path and configuration for input files.
	basePath := "course_4/module_4/programming_assignment_4/"
	filePrefix := "2sat"
	fileExtension := ".txt"
	fileCount := 6

	// Generate file paths for all instances.
	instances := make([]string, fileCount)
	for i := 1; i <= fileCount; i++ {
		instances[i-1] = fmt.Sprintf("%s%s%d%s", basePath, filePrefix, i, fileExtension)
	}

	var result string

	// Process each file and determine satisfiability.
	for _, filePath := range instances {
		log.Printf("Processing file: %s", filePath)

		solver, err := ParseFile(filePath)
		if err != nil {
			log.Printf("Error reading file %s: %v", filePath, err)
			result += "0" // Mark as unsatisfiable in case of an error.
			continue
		}

		if solver.Solve() {
			result += "1" // Instance is satisfiable.
		} else {
			result += "0" // Instance is unsatisfiable.
		}
	}

	// Print the final result as a binary string.
	fmt.Println("Result:", result)
}
