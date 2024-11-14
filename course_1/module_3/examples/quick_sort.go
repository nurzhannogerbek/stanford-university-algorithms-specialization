package main

import (
	"fmt"
)

// QuickSort is a struct that holds an array to be sorted and a counter for the number of comparisons.
// The `comparisons` counter tracks the total number of comparisons made during the sorting process.
type QuickSort struct {
	array       []int
	comparisons int
}

// NewQuickSort is a constructor function that initializes a new QuickSort instance.
// It takes an array as input, makes a copy of it to avoid modifying the original array,
// and initializes the comparisons counter to zero.
func NewQuickSort(arr []int) *QuickSort {
	return &QuickSort{array: append([]int(nil), arr...), comparisons: 0}
}

// SortWithFirstPivot sorts the array using the first element as the pivot.
// It calls the `quicksortFirst` method, which performs the recursive quicksort
// operation based on the first element as the pivot.
func (qs *QuickSort) SortWithFirstPivot() {
	qs.quicksortFirst(0, len(qs.array)-1)
}

// SortWithLastPivot sorts the array using the last element as the pivot.
// It calls the `quicksortLast` method, which performs the recursive quicksort
// operation based on the last element as the pivot.
func (qs *QuickSort) SortWithLastPivot() {
	qs.quicksortLast(0, len(qs.array)-1)
}

// SortWithMedianPivot sorts the array using the median-of-three as the pivot.
// It calls the `quicksortMedian` method, which performs the recursive quicksort
// operation based on the median of the first, middle, and last elements.
func (qs *QuickSort) SortWithMedianPivot() {
	qs.quicksortMedian(0, len(qs.array)-1)
}

// quicksortFirst is a recursive method that sorts the array using the first element as the pivot.
// It partitions the array around the first element, then recursively sorts the left and right partitions.
func (qs *QuickSort) quicksortFirst(low, high int) {
	if low < high {
		// Partition the array around the pivot and get the pivot index.
		pivotIndex := qs.partitionFirst(low, high)
		// Recursively sort the elements to the left of the pivot.
		qs.quicksortFirst(low, pivotIndex-1)
		// Recursively sort the elements to the right of the pivot.
		qs.quicksortFirst(pivotIndex+1, high)
	}
}

// partitionFirst partitions the array using the first element as the pivot.
// It places elements smaller than the pivot to its left and larger elements to its right.
// It returns the final index of the pivot after partitioning.
func (qs *QuickSort) partitionFirst(low, high int) int {
	pivot := qs.array[low] // Select the first element as the pivot.
	i := low + 1           // Initialize pointer for the greater element.

	// Iterate over the elements and rearrange them based on the pivot value.
	for j := low + 1; j <= high; j++ {
		if qs.array[j] < pivot {
			// Swap elements to bring smaller elements to the left.
			qs.array[i], qs.array[j] = qs.array[j], qs.array[i]
			i++
		}
	}
	// Place the pivot in its correct sorted position.
	qs.array[low], qs.array[i-1] = qs.array[i-1], qs.array[low]
	// Update the comparison count with the number of elements in the current partition.
	qs.comparisons += high - low
	// Return the index of the pivot after partitioning.
	return i - 1
}

// quicksortLast is a recursive method that sorts the array using the last element as the pivot.
// It swaps the last element with the first element, then partitions and recursively sorts the subarrays.
func (qs *QuickSort) quicksortLast(low, high int) {
	if low < high {
		// Partition the array around the pivot and get the pivot index.
		pivotIndex := qs.partitionLast(low, high)
		// Recursively sort the elements to the left of the pivot.
		qs.quicksortLast(low, pivotIndex-1)
		// Recursively sort the elements to the right of the pivot.
		qs.quicksortLast(pivotIndex+1, high)
	}
}

// partitionLast partitions the array using the last element as the pivot.
// It swaps the last element to the beginning, then calls `partitionFirst` to perform the partitioning.
func (qs *QuickSort) partitionLast(low, high int) int {
	// Swap the last element with the first to use it as the pivot.
	qs.array[low], qs.array[high] = qs.array[high], qs.array[low]
	// Call partitionFirst to partition around the new pivot (initially the last element).
	return qs.partitionFirst(low, high)
}

// quicksortMedian is a recursive method that sorts the array using the median-of-three as the pivot.
// It selects the median of the first, middle, and last elements as the pivot, then partitions and sorts.
func (qs *QuickSort) quicksortMedian(low, high int) {
	if low < high {
		// Partition the array around the median-of-three pivot and get the pivot index.
		pivotIndex := qs.partitionMedian(low, high)
		// Recursively sort the elements to the left of the pivot.
		qs.quicksortMedian(low, pivotIndex-1)
		// Recursively sort the elements to the right of the pivot.
		qs.quicksortMedian(pivotIndex+1, high)
	}
}

// partitionMedian partitions the array using the median of the first, middle, and last elements as the pivot.
// It finds the median among these three elements, swaps it to the beginning, and then partitions the array.
func (qs *QuickSort) partitionMedian(low, high int) int {
	mid := (low + high) / 2 // Calculate the middle index.

	// Create a list of pivot candidates (first, middle, and last) with their values and indices.
	pivotCandidates := []struct {
		value int
		index int
	}{
		{qs.array[low], low},
		{qs.array[mid], mid},
		{qs.array[high], high},
	}

	// Sort the pivot candidates by value to determine the median.
	for i := 0; i < len(pivotCandidates)-1; i++ {
		for j := i + 1; j < len(pivotCandidates); j++ {
			if pivotCandidates[i].value > pivotCandidates[j].value {
				pivotCandidates[i], pivotCandidates[j] = pivotCandidates[j], pivotCandidates[i]
			}
		}
	}

	// Select the median element as the pivot and swap it to the beginning of the array.
	pivotIndex := pivotCandidates[1].index
	qs.array[low], qs.array[pivotIndex] = qs.array[pivotIndex], qs.array[low]
	// Use the partitionFirst method to partition around the median pivot.
	return qs.partitionFirst(low, high)
}

// GetComparisons returns the total number of comparisons made during the sorting process.
func (qs *QuickSort) GetComparisons() int {
	return qs.comparisons
}

// GetSortedArray returns the sorted array after performing quicksort.
// This allows access to the final sorted order after any of the sorting methods is called.
func (qs *QuickSort) GetSortedArray() []int {
	return qs.array
}

// main is the entry point of the program, demonstrating how to use the QuickSort struct
// with different pivot selection strategies for sorting.
func main() {
	// Example array to be sorted.
	array := []int{3, 8, 2, 5, 1, 4, 7, 6}

	// Create a QuickSort instance and sort the array using the first element as the pivot.
	qsFirst := NewQuickSort(array)
	qsFirst.SortWithFirstPivot()
	fmt.Println("Sorted array with first element as pivot:", qsFirst.GetSortedArray())
	fmt.Println("Total comparisons with first element as pivot:", qsFirst.GetComparisons())

	// Create a QuickSort instance and sort the array using the last element as the pivot.
	qsLast := NewQuickSort(array)
	qsLast.SortWithLastPivot()
	fmt.Println("Sorted array with last element as pivot:", qsLast.GetSortedArray())
	fmt.Println("Total comparisons with last element as pivot:", qsLast.GetComparisons())

	// Create a QuickSort instance and sort the array using the median-of-three as the pivot.
	qsMedian := NewQuickSort(array)
	qsMedian.SortWithMedianPivot()
	fmt.Println("Sorted array with median-of-three as pivot:", qsMedian.GetSortedArray())
	fmt.Println("Total comparisons with median-of-three as pivot:", qsMedian.GetComparisons())
}
