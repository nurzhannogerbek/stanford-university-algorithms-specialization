package main

import (
	"errors"
	"fmt"
)

// UnionFind represents the Union-Find structure using Lazy Unions, Union by Rank, and Path Compression.
type UnionFind struct {
	parent []int // Parent pointers for each element.
	rank   []int // Rank of each set (helps optimize unions).
}

// NewUnionFind initializes a new Union-Find structure with the given size.
// Each element starts in its own set, where it is its own parent and rank is 0.
func NewUnionFind(size int) *UnionFind {
	if size <= 0 {
		panic("Size must be greater than 0.")
	}
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i // Each element is its own parent.
		rank[i] = 0   // Initial rank for all elements is 0.
	}
	return &UnionFind{parent: parent, rank: rank}
}

// validate ensures that the index x is within the valid range.
func (uf *UnionFind) validate(x int) error {
	if x < 0 || x >= len(uf.parent) {
		return errors.New(fmt.Sprintf("Index %d is out of bounds.", x))
	}
	return nil
}

// Find returns the root of the element x, applying path compression.
// Path compression ensures that all elements along the path point directly to the root,
// flattening the tree and optimizing future operations.
func (uf *UnionFind) Find(x int) (int, error) {
	if err := uf.validate(x); err != nil {
		return -1, err
	}
	if uf.parent[x] != x {
		root, err := uf.Find(uf.parent[x])
		if err != nil {
			return -1, err
		}
		uf.parent[x] = root // Path compression.
	}
	return uf.parent[x], nil
}

// Union merges the sets containing x and y using union by rank.
// The root of the smaller rank tree is made a child of the root of the larger rank tree.
// If ranks are equal, one root is chosen arbitrarily, and its rank is incremented.
func (uf *UnionFind) Union(x, y int) error {
	rootX, err := uf.Find(x)
	if err != nil {
		return err
	}
	rootY, err := uf.Find(y)
	if err != nil {
		return err
	}

	if rootX != rootY {
		// Merge smaller rank tree into the larger rank tree.
		if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else {
			// If ranks are equal, choose one as root and increment its rank.
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
	return nil
}

// Connected checks if two elements x and y are in the same set.
// Returns true if they share the same root, false otherwise.
func (uf *UnionFind) Connected(x, y int) (bool, error) {
	rootX, err := uf.Find(x)
	if err != nil {
		return false, err
	}
	rootY, err := uf.Find(y)
	if err != nil {
		return false, err
	}
	return rootX == rootY, nil
}

// Components returns all connected components as a map.
// The keys of the map are the roots, and the values are slices of elements in each component.
func (uf *UnionFind) Components() map[int][]int {
	components := make(map[int][]int)
	visited := make(map[int]int) // Cache the root of each element.
	for i := 0; i < len(uf.parent); i++ {
		if root, found := visited[i]; found {
			components[root] = append(components[root], i)
		} else {
			root, _ := uf.Find(i) // We already validated indices, no need for extra checks.
			visited[i] = root
			components[root] = append(components[root], i)
		}
	}
	return components
}

// Size returns the number of elements in the Union-Find structure.
func (uf *UnionFind) Size() int {
	return len(uf.parent)
}

// Main function to demonstrate usage of the Union-Find structure.
func main() {
	uf := NewUnionFind(10)

	// Perform union operations.
	_ = uf.Union(1, 2)
	_ = uf.Union(3, 4)
	_ = uf.Union(2, 4)

	// Check connectivity.
	connected, _ := uf.Connected(1, 3)
	fmt.Println("Connected(1, 3):", connected)
	connected, _ = uf.Connected(1, 5)
	fmt.Println("Connected(1, 5):", connected)

	// Print all connected components.
	fmt.Println("Components:")
	components := uf.Components()
	for root, group := range components {
		fmt.Printf("Root %d: %v\n", root, group)
	}
}
