package main

import (
	"errors"
	"fmt"
	"sort"
)

// DeterministicSelection struct encapsulates the array and provides methods for the selection algorithm.
type DeterministicSelection struct {
	array []int
}

// NewDeterministicSelection creates a new instance of DeterministicSelection.
func NewDeterministicSelection(inputArray []int) *DeterministicSelection {
	return &DeterministicSelection{array: inputArray}
}

// partition rearranges the array such that elements less than the pivot are on the left
// and elements greater than the pivot are on the right.
func (ds *DeterministicSelection) partition(left, right, pivotIndex int) int {
	pivotValue := ds.array[pivotIndex]
	// Move pivot to the end.
	ds.array[pivotIndex], ds.array[right] = ds.array[right], ds.array[pivotIndex]
	storeIndex := left

	// Rearrange elements based on the pivot value.
	for i := left; i < right; i++ {
		if ds.array[i] < pivotValue {
			ds.array[i], ds.array[storeIndex] = ds.array[storeIndex], ds.array[i]
			storeIndex++
		}
	}

	// Move pivot to its final position.
	ds.array[storeIndex], ds.array[right] = ds.array[right], ds.array[storeIndex]
	return storeIndex
}

// medianOfMedians selects a pivot deterministically using the "median of medians" strategy.
func (ds *DeterministicSelection) medianOfMedians(left, right int) int {
	n := right - left + 1

	// If the subarray has 5 or fewer elements, return the median directly.
	if n <= 5 {
		return ds.medianOfSmallGroup(left, right)
	}

	// Divide array into groups of 5 and find their medians.
	medians := make([]int, 0, (n+4)/5) // +4 to ensure the division rounds up.
	for i := left; i <= right; i += 5 {
		subRight := i + 4
		if subRight > right {
			subRight = right
		}
		median := ds.medianOfSmallGroup(i, subRight)
		medians = append(medians, ds.array[median])
	}

	// Recursively find the median of medians.
	newDS := NewDeterministicSelection(medians)
	return left + newDS.deterministicSelect(0, len(medians)-1, len(medians)/2)
}

// medianOfSmallGroup finds the median of a small group of up to 5 elements.
func (ds *DeterministicSelection) medianOfSmallGroup(left, right int) int {
	// Sort the group and find the median index.
	subArray := ds.array[left : right+1]
	sort.Ints(subArray)
	return left + (right-left)/2
}

// deterministicSelect recursively finds the k-th smallest element.
func (ds *DeterministicSelection) deterministicSelect(left, right, k int) int {
	// Base case: if the array has only one element.
	if left == right {
		return left
	}

	// Choose a pivot deterministically.
	pivotIndex := ds.medianOfMedians(left, right)

	// Partition the array around the pivot.
	pivotIndex = ds.partition(left, right, pivotIndex)

	// Determine the rank of the pivot.
	rank := pivotIndex - left + 1

	if rank == k {
		// The pivot is the k-th smallest element.
		return pivotIndex
	} else if k < rank {
		// Recurse on the left subarray.
		return ds.deterministicSelect(left, pivotIndex-1, k)
	} else {
		// Recurse on the right subarray.
		return ds.deterministicSelect(pivotIndex+1, right, k-rank)
	}
}

// Select is the public method to find the k-th smallest element in the array.
func (ds *DeterministicSelection) Select(k int) (int, error) {
	if k < 1 || k > len(ds.array) {
		return 0, errors.New("k is out of bounds of the array")
	}
	index := ds.deterministicSelect(0, len(ds.array)-1, k)
	return ds.array[index], nil
}

func main() {
	// Input array.
	array := []int{10, 4, 5, 8, 6, 11, 26}
	k := 3 // Desired rank (1-based).

	// Create a new instance of DeterministicSelection.
	selector := NewDeterministicSelection(array)

	// Find the k-th smallest element.
	result, err := selector.Select(k)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result.
	fmt.Printf("The %d-th smallest element is: %d\n", k, result)
}
