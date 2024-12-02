package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Item represents an item with a value and weight.
type Item struct {
	Value  int
	Weight int
}

// Knapsack represents a knapsack with a given capacity and a list of items.
type Knapsack struct {
	Capacity int
	Items    []Item
}

// NewKnapsack creates a new knapsack with the specified capacity.
func NewKnapsack(capacity int) *Knapsack {
	return &Knapsack{
		Capacity: capacity,
		Items:    []Item{},
	}
}

// AddItem adds an item to the knapsack.
func (k *Knapsack) AddItem(item Item) {
	k.Items = append(k.Items, item)
}

// Solve solves the knapsack problem using dynamic programming.
func (k *Knapsack) Solve() int {
	numItems := len(k.Items)
	// Create a DP table with dimensions (numItems+1) x (Capacity+1).
	dp := make([][]int, numItems+1)
	for i := range dp {
		dp[i] = make([]int, k.Capacity+1)
	}

	// Fill the DP table.
	for i := 1; i <= numItems; i++ {
		for w := 0; w <= k.Capacity; w++ {
			if k.Items[i-1].Weight > w {
				dp[i][w] = dp[i-1][w] // Item cannot be included.
			} else {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-k.Items[i-1].Weight]+k.Items[i-1].Value)
			}
		}
	}

	return dp[numItems][k.Capacity]
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ReadInputFile reads the input file and constructs a Knapsack object.
func ReadInputFile(filePath string) (*Knapsack, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the first line for knapsack capacity and number of items.
	scanner.Scan()
	firstLine := strings.Fields(scanner.Text())
	capacity, _ := strconv.Atoi(firstLine[0])

	// Create the knapsack object.
	knapsack := NewKnapsack(capacity)

	// Read each subsequent line for item values and weights.
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		value, _ := strconv.Atoi(line[0])
		weight, _ := strconv.Atoi(line[1])
		knapsack.AddItem(Item{Value: value, Weight: weight})
	}

	return knapsack, nil
}

func main() {
	// File path to the input data.
	filePath := "course_3/module_4/programming_assignment_4/knapsack1.txt"

	// Read the knapsack data from the file.
	knapsack, err := ReadInputFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Solve the knapsack problem and print the optimal value.
	optimalValue := knapsack.Solve()
	fmt.Printf("The optimal value is: %d\n", optimalValue)
}