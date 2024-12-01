package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// UnionFind represents the Union-Find (disjoint-set) structure.
type UnionFind struct {
	parent map[string]string
	rank   map[string]int
}

// NewUnionFind initializes a new Union-Find structure.
func NewUnionFind(nodes []string) *UnionFind {
	parent := make(map[string]string)
	rank := make(map[string]int)

	for _, node := range nodes {
		parent[node] = node
		rank[node] = 0
	}

	return &UnionFind{parent: parent, rank: rank}
}

// Find finds the root of a node using path compression.
func (uf *UnionFind) Find(node string) string {
	if uf.parent[node] != node {
		uf.parent[node] = uf.Find(uf.parent[node]) // Path compression.
	}
	return uf.parent[node]
}

// Union unites two nodes by rank.
func (uf *UnionFind) Union(node1, node2 string) {
	root1 := uf.Find(node1)
	root2 := uf.Find(node2)

	if root1 != root2 {
		if uf.rank[root1] > uf.rank[root2] {
			uf.parent[root2] = root1
		} else if uf.rank[root1] < uf.rank[root2] {
			uf.parent[root1] = root2
		} else {
			uf.parent[root2] = root1
			uf.rank[root1]++
		}
	}
}

// Clustering represents the clustering algorithm.
type Clustering struct {
	vertices  []string
	bitLength int
	unionFind *UnionFind
}

// NewClustering initializes a new Clustering instance.
func NewClustering(vertices []string, bitLength int) *Clustering {
	return &Clustering{
		vertices:  vertices,
		bitLength: bitLength,
		unionFind: NewUnionFind(vertices),
	}
}

// invert flips a single bit ('0' -> '1' or '1' -> '0').
func invert(bit byte) byte {
	if bit == '0' {
		return '1'
	}
	return '0'
}

// GenerateNeighbors generates all neighbors of a vertex with Hamming distance 1 or 2.
func (c *Clustering) GenerateNeighbors(vertex string) []string {
	neighbors := []string{}

	for i := 0; i < c.bitLength; i++ {
		// Flip one bit.
		flipped := vertex[:i] + string(invert(vertex[i])) + vertex[i+1:]
		neighbors = append(neighbors, flipped)

		for j := i + 1; j < c.bitLength; j++ {
			// Flip two bits.
			flippedTwo := flipped[:j] + string(invert(flipped[j])) + flipped[j+1:]
			neighbors = append(neighbors, flippedTwo)
		}
	}

	return neighbors
}

// Cluster performs the clustering operation.
func (c *Clustering) Cluster() {
	for _, vertex := range c.vertices {
		for _, neighbor := range c.GenerateNeighbors(vertex) {
			if _, exists := c.unionFind.parent[neighbor]; exists {
				c.unionFind.Union(vertex, neighbor)
			}
		}
	}
}

// CountClusters counts the number of unique clusters.
func (c *Clustering) CountClusters() int {
	clusterRoots := make(map[string]struct{})

	for _, vertex := range c.vertices {
		root := c.unionFind.Find(vertex)
		clusterRoots[root] = struct{}{}
	}

	return len(clusterRoots)
}

// Main function to load data, perform clustering, and output the result.
func main() {
	filePath := "course_3/module_2/programming_assignment_2/clustering_big.txt"

	// Open the file.
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the first line (metadata).
	scanner.Scan()
	meta := strings.Split(scanner.Text(), " ")
	bitLength, err := strconv.Atoi(meta[1])
	if err != nil {
		fmt.Println("Error parsing bit length:", err)
		return
	}

	// Read vertices.
	var vertices []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		vertices = append(vertices, strings.ReplaceAll(line, " ", ""))
	}

	// Initialize clustering and perform it.
	clustering := NewClustering(vertices, bitLength)
	clustering.Cluster()

	// Count the number of clusters and print the result.
	numClusters := clustering.CountClusters()
	fmt.Printf("The largest value of k with spacing at least 3 is: %d\n", numClusters)
}
