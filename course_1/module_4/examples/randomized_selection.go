package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// RandomizedSelection struct encapsulates the array and its methods.
type RandomizedSelection struct {
	array []int
}

// NewRandomizedSelection creates a new instance of RandomizedSelection.
func NewRandomizedSelection(array []int) (*RandomizedSelection, error) {
	if len(array) == 0 {
		return nil, errors.New("input array must not be empty")
	}
	return &RandomizedSelection{array: append([]int{}, array...)}, nil
}

// partition partitions the array around the pivot element.
func (rs *RandomizedSelection) partition(left, right, pivotIndex int) int {
	pivotValue := rs.array[pivotIndex]
	// Move pivot to the end.
	rs.array[pivotIndex], rs.array[right] = rs.array[right], rs.array[pivotIndex]
	storeIndex := left

	// Rearrange elements around the pivot.
	for i := left; i < right; i++ {
		if rs.array[i] < pivotValue {
			rs.array[i], rs.array[storeIndex] = rs.array[storeIndex], rs.array[i]
			storeIndex++
		}
	}

	// Move pivot to its final position.
	rs.array[storeIndex], rs.array[right] = rs.array[right], rs.array[storeIndex]
	return storeIndex
}

// randomizedSelect recursively finds the k-th smallest element.
func (rs *RandomizedSelection) randomizedSelect(left, right, k int) int {
	if left == right {
		return rs.array[left] // Base case: only one element.
	}

	// Choose a random pivot index.
	pivotIndex := rand.Intn(right-left+1) + left

	// Partition the array and get the pivot's final index.
	pivotIndex = rs.partition(left, right, pivotIndex)

	// Determine the rank of the pivot.
	rank := pivotIndex - left + 1

	if rank == k {
		return rs.array[pivotIndex]
	} else if k < rank {
		return rs.randomizedSelect(left, pivotIndex-1, k)
	} else {
		return rs.randomizedSelect(pivotIndex+1, right, k-rank)
	}
}

// Select finds the k-th smallest element in the array.
func (rs *RandomizedSelection) Select(k int) (int, error) {
	if k < 1 || k > len(rs.array) {
		return 0, errors.New("k is out of bounds of the array")
	}
	return rs.randomizedSelect(0, len(rs.array)-1, k), nil
}

func main() {
	// Initialize the random seed.
	rand.Seed(time.Now().UnixNano())

	array := []int{10, 4, 5, 8, 6, 11, 26}
	k := 3

	// Create an instance of RandomizedSelection.
	selector, err := NewRandomizedSelection(array)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Find the k-th smallest element.
	result, err := selector.Select(k)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("The %d-th smallest element is: %d\n", k, result)
}
