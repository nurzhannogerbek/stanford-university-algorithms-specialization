package main

import "fmt"

// Item represents an item with a value and a weight.
type Item struct {
	Value  int
	Weight int
}

// Knapsack represents a knapsack with a given capacity and a list of items.
type Knapsack struct {
	Capacity int
	Items    []Item
}

// NewKnapsack creates a new knapsack with a given capacity.
func NewKnapsack(capacity int) *Knapsack {
	return &Knapsack{
		Capacity: capacity,
		Items:    []Item{},
	}
}

// AddItem adds a new item to the knapsack.
func (k *Knapsack) AddItem(value, weight int) {
	k.Items = append(k.Items, Item{Value: value, Weight: weight})
}

// Solve solves the Knapsack Problem using dynamic programming.
// It returns the maximum value and the list of items included in the optimal solution.
func (k *Knapsack) Solve() (int, []Item) {
	n := len(k.Items)
	W := k.Capacity

	// Create a 2D slice for dynamic programming.
	// dp[i][w] stores the maximum value for the first i items with a capacity of w.
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	// Fill the dp table.
	for i := 1; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if k.Items[i-1].Weight <= w {
				// If the current item's weight fits, take the maximum value
				// of including or excluding the item.
				dp[i][w] = max(
					dp[i-1][w], // Exclude the current item
					dp[i-1][w-k.Items[i-1].Weight]+k.Items[i-1].Value, // Include the current item
				)
			} else {
				// Otherwise, exclude the item.
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// Backtrack to find the items included in the optimal solution.
	w := W
	includedItems := []Item{}
	for i := n; i > 0; i-- {
		if dp[i][w] != dp[i-1][w] { // If the value changed, the item was included.
			includedItems = append(includedItems, k.Items[i-1])
			w -= k.Items[i-1].Weight // Reduce the remaining capacity.
		}
	}

	// Return the maximum value and the included items.
	return dp[n][W], includedItems
}

// Utility function to get the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Main function to demonstrate the usage of the Knapsack implementation.
func main() {
	// Create a knapsack with a capacity of 6.
	knapsack := NewKnapsack(6)

	// Add items (value, weight).
	knapsack.AddItem(3, 4)
	knapsack.AddItem(2, 3)
	knapsack.AddItem(4, 2)
	knapsack.AddItem(4, 3)

	// Solve the knapsack problem.
	maxValue, includedItems := knapsack.Solve()

	// Output the results.
	fmt.Printf("Maximum value: %d\n", maxValue)
	fmt.Println("Items included in the optimal solution:")
	for _, item := range includedItems {
		fmt.Printf("Value: %d, Weight: %d\n", item.Value, item.Weight)
	}
}
