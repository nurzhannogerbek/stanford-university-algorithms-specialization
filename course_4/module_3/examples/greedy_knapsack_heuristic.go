package main

import (
	"fmt"
	"sort"
)

// Item represents an item in the knapsack problem.
type Item struct {
	Value  int     // Value of the item.
	Weight int     // Weight of the item.
	Ratio  float64 // Value-to-weight ratio of the item.
}

// Knapsack represents the knapsack used to store selected items.
type Knapsack struct {
	Capacity      int    // Total capacity of the knapsack.
	CurrentWeight int    // Current weight of the knapsack.
	Items         []Item // List of items included in the knapsack.
	TotalValue    int    // Total value of the items in the knapsack.
}

// NewKnapsack initializes a new Knapsack.
func NewKnapsack(capacity int) *Knapsack {
	if capacity <= 0 {
		panic("Knapsack capacity must be positive.")
	}
	return &Knapsack{
		Capacity: capacity,
	}
}

// AddItem tries to add an item to the knapsack if it fits.
func (k *Knapsack) AddItem(item Item) bool {
	if k.CurrentWeight+item.Weight <= k.Capacity {
		k.Items = append(k.Items, item)
		k.CurrentWeight += item.Weight
		k.TotalValue += item.Value
		return true
	}
	return false
}

// GreedyKnapsackSolver solves the knapsack problem using a greedy heuristic.
type GreedyKnapsackSolver struct {
	Knapsack *Knapsack // The knapsack to store selected items.
	Items    []Item    // List of available items.
}

// NewGreedyKnapsackSolver initializes a new solver with the given capacity and items.
func NewGreedyKnapsackSolver(capacity int, items []Item) *GreedyKnapsackSolver {
	if len(items) == 0 {
		panic("Item list must not be empty.")
	}
	for i := range items {
		if items[i].Weight < 0 || items[i].Value < 0 {
			panic("Item values and weights must be non-negative.")
		}
		// Calculate value-to-weight ratio for sorting.
		if items[i].Weight == 0 {
			items[i].Ratio = 1e9 // Prioritize zero-weight items with a high ratio.
		} else {
			items[i].Ratio = float64(items[i].Value) / float64(items[i].Weight)
		}
	}
	return &GreedyKnapsackSolver{
		Knapsack: NewKnapsack(capacity),
		Items:    items,
	}
}

// Solve executes the greedy heuristic to fill the knapsack.
func (solver *GreedyKnapsackSolver) Solve() {
	// Sort items by value-to-weight ratio in descending order.
	sort.Slice(solver.Items, func(i, j int) bool {
		return solver.Items[i].Ratio > solver.Items[j].Ratio
	})

	// Add items to the knapsack while respecting capacity constraints.
	for _, item := range solver.Items {
		solver.Knapsack.AddItem(item)
	}
}

// Result provides the results of the heuristic solution.
func (solver *GreedyKnapsackSolver) Result() {
	fmt.Printf("Total Value: %d\n", solver.Knapsack.TotalValue)
	fmt.Printf("Total Weight: %d\n", solver.Knapsack.CurrentWeight)
	fmt.Println("Selected Items (Value, Weight):")
	for _, item := range solver.Knapsack.Items {
		fmt.Printf("  Value: %d, Weight: %d\n", item.Value, item.Weight)
	}
}

func main() {
	// Define the items (value, weight).
	items := []Item{
		{Value: 60, Weight: 10},
		{Value: 100, Weight: 20},
		{Value: 120, Weight: 30},
		{Value: 40, Weight: 0}, // Zero-weight item.
	}

	// Define the capacity of the knapsack.
	capacity := 50

	// Create a greedy knapsack solver.
	solver := NewGreedyKnapsackSolver(capacity, items)

	// Solve the problem.
	solver.Solve()

	// Display the results.
	solver.Result()
}
