package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Edge represents an edge in the graph with two nodes and a cost.
type Edge struct {
	Node1 int // The first node of the edge.
	Node2 int // The second node of the edge.
	Cost  int // The cost or weight of the edge.
}

// UnionFind represents the union-find (disjoint-set) structure.
type UnionFind struct {
	Parent []int // Array to track the parent of each node.
	Rank   []int // Array to track the rank (depth) of each set.
}

// NewUnionFind initializes a new UnionFind structure with the specified size.
func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i // Initially, each node is its own parent.
	}
	return &UnionFind{Parent: parent, Rank: rank}
}

// Find finds the root of the node using path compression for efficiency.
func (uf *UnionFind) Find(node int) int {
	if uf.Parent[node] != node {
		uf.Parent[node] = uf.Find(uf.Parent[node]) // Recursively find the root and compress the path.
	}
	return uf.Parent[node]
}

// Union unions two subsets based on their rank to maintain a balanced tree structure.
func (uf *UnionFind) Union(node1, node2 int) {
	root1 := uf.Find(node1) // Find the root of the first node.
	root2 := uf.Find(node2) // Find the root of the second node.

	if root1 != root2 {
		if uf.Rank[root1] > uf.Rank[root2] {
			uf.Parent[root2] = root1 // Attach the smaller tree under the larger tree.
		} else if uf.Rank[root1] < uf.Rank[root2] {
			uf.Parent[root1] = root2
		} else {
			uf.Parent[root2] = root1 // If ranks are equal, attach one tree and increase the rank.
			uf.Rank[root1]++
		}
	}
}

// Clustering represents the clustering algorithm.
type Clustering struct {
	Nodes int    // Number of nodes in the graph.
	Edges []Edge // List of edges in the graph.
}

// NewClustering initializes a new Clustering instance with nodes and edges.
func NewClustering(nodes int, edges []Edge) *Clustering {
	return &Clustering{Nodes: nodes, Edges: edges}
}

// ComputeMaxSpacing computes the maximum spacing of k-clustering.
func (c *Clustering) ComputeMaxSpacing(k int) (int, error) {
	// Validate the value of k.
	if k <= 0 || k > c.Nodes {
		return 0, fmt.Errorf("invalid value of k: must be between 1 and %d", c.Nodes)
	}

	// Sort edges by cost in ascending order.
	sort.Slice(c.Edges, func(i, j int) bool {
		return c.Edges[i].Cost < c.Edges[j].Cost
	})

	uf := NewUnionFind(c.Nodes) // Initialize the Union-Find structure.
	numClusters := c.Nodes      // Start with each node as its own cluster.

	// Process edges to form clusters.
	for _, edge := range c.Edges {
		// Check if the nodes belong to different clusters.
		if uf.Find(edge.Node1) != uf.Find(edge.Node2) {
			if numClusters == k {
				// If the desired number of clusters is reached, return the next edge cost.
				return edge.Cost, nil
			}
			uf.Union(edge.Node1, edge.Node2) // Merge the clusters.
			numClusters--                    // Decrease the number of clusters.
		}
	}

	return 0, fmt.Errorf("unable to determine max spacing")
}

// ReadInput reads the input data from a file and returns the nodes and edges.
func ReadInput(filePath string) (int, []Edge, error) {
	file, err := os.Open(filePath) // Open the input file.
	if err != nil {
		return 0, nil, fmt.Errorf("unable to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return 0, nil, fmt.Errorf("file is empty or missing number of nodes")
	}

	numNodes, err := strconv.Atoi(scanner.Text()) // Read the number of nodes.
	if err != nil {
		return 0, nil, fmt.Errorf("invalid number of nodes: %w", err)
	}

	var edges []Edge
	lineNumber := 1
	// Read each line as an edge.
	for scanner.Scan() {
		lineNumber++
		parts := strings.Fields(scanner.Text())
		if len(parts) != 3 {
			return 0, nil, fmt.Errorf("invalid line %d: expected 3 fields, got %d", lineNumber, len(parts))
		}

		node1, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, nil, fmt.Errorf("invalid node1 at line %d: %w", lineNumber, err)
		}

		node2, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, nil, fmt.Errorf("invalid node2 at line %d: %w", lineNumber, err)
		}

		cost, err := strconv.Atoi(parts[2])
		if err != nil {
			return 0, nil, fmt.Errorf("invalid cost at line %d: %w", lineNumber, err)
		}

		// Add the edge to the list.
		edges = append(edges, Edge{Node1: node1 - 1, Node2: node2 - 1, Cost: cost})
	}

	if err := scanner.Err(); err != nil {
		return 0, nil, fmt.Errorf("error reading file: %w", err)
	}

	return numNodes, edges, nil
}

func main() {
	// Define the input file path.
	filePath := "course_3/module_2/programming_assignment_2/clustering1.txt"

	// Read input data from the file.
	nodes, edges, err := ReadInput(filePath)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	// Initialize the clustering algorithm.
	clustering := NewClustering(nodes, edges)

	// Compute the maximum spacing for 4 clusters.
	maxSpacing, err := clustering.ComputeMaxSpacing(4)
	if err != nil {
		fmt.Printf("Error computing max spacing: %v\n", err)
		return
	}

	// Print the result.
	fmt.Printf("The maximum spacing of a 4-clustering is: %d\n", maxSpacing)
}
