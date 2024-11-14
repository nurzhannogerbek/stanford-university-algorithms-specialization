package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// QuickSort is a struct that holds an array to be sorted and a counter for comparisons.
type QuickSort struct {
	array       []int
	comparisons int
}

// NewQuickSort is a constructor that initializes QuickSort with a copy of the input array.
// It prevents modification of the original array by creating a new copy.
func NewQuickSort(arr []int) *QuickSort {
	return &QuickSort{array: append([]int(nil), arr...), comparisons: 0}
}

// QuickSortFirst sorts the array using the first element as the pivot.
// It recursively divides the array around the chosen pivot.
func (qs *QuickSort) QuickSortFirst(low, high int) {
	if low < high {
		pivotIndex := qs.PartitionFirst(low, high)
		qs.QuickSortFirst(low, pivotIndex-1)  // Sort elements before the pivot.
		qs.QuickSortFirst(pivotIndex+1, high) // Sort elements after the pivot.
	}
}

// PartitionFirst partitions the array using the first element as the pivot.
// It moves elements smaller than the pivot to the left and larger to the right.
func (qs *QuickSort) PartitionFirst(low, high int) int {
	pivot := qs.array[low] // Choose the first element as the pivot.
	i := low + 1           // Start pointer for the greater element.

	// Loop through the subarray and rearrange elements around the pivot.
	for j := low + 1; j <= high; j++ {
		if qs.array[j] < pivot {
			qs.array[i], qs.array[j] = qs.array[j], qs.array[i]
			i++
		}
	}
	qs.array[low], qs.array[i-1] = qs.array[i-1], qs.array[low] // Place the pivot in the correct position.
	qs.comparisons += high - low                                // Count comparisons for the current partition.
	return i - 1                                                // Return the pivot's final position.
}

// QuickSortLast sorts the array using the last element as the pivot.
// It recursively divides the array around the chosen pivot.
func (qs *QuickSort) QuickSortLast(low, high int) {
	if low < high {
		pivotIndex := qs.PartitionLast(low, high)
		qs.QuickSortLast(low, pivotIndex-1)  // Sort elements before the pivot.
		qs.QuickSortLast(pivotIndex+1, high) // Sort elements after the pivot.
	}
}

// PartitionLast partitions the array using the last element as the pivot.
// It swaps the last element to the beginning, then uses the first element as the pivot.
func (qs *QuickSort) PartitionLast(low, high int) int {
	qs.array[low], qs.array[high] = qs.array[high], qs.array[low] // Move the last element to the beginning.
	return qs.PartitionFirst(low, high)                           // Use PartitionFirst logic.
}

// QuickSortMedian sorts the array using the median of the first, middle, and last elements as the pivot.
// It recursively divides the array around the chosen pivot.
func (qs *QuickSort) QuickSortMedian(low, high int) {
	if low < high {
		pivotIndex := qs.PartitionMedian(low, high)
		qs.QuickSortMedian(low, pivotIndex-1)  // Sort elements before the pivot.
		qs.QuickSortMedian(pivotIndex+1, high) // Sort elements after the pivot.
	}
}

// PartitionMedian partitions the array using the median-of-three as the pivot.
// It finds the median of the first, middle, and last elements, then uses it as the pivot.
func (qs *QuickSort) PartitionMedian(low, high int) int {
	mid := (low + high) / 2 // Calculate the middle index.

	// Create a slice of pivot candidates and their positions.
	pivotCandidates := []struct {
		value int
		index int
	}{
		{qs.array[low], low},
		{qs.array[mid], mid},
		{qs.array[high], high},
	}

	// Sort the pivot candidates to find the median.
	for i := 0; i < len(pivotCandidates)-1; i++ {
		for j := i + 1; j < len(pivotCandidates); j++ {
			if pivotCandidates[i].value > pivotCandidates[j].value {
				pivotCandidates[i], pivotCandidates[j] = pivotCandidates[j], pivotCandidates[i]
			}
		}
	}

	// Choose the median element as the pivot.
	pivotIndex := pivotCandidates[1].index
	qs.array[low], qs.array[pivotIndex] = qs.array[pivotIndex], qs.array[low] // Place pivot at the start.
	return qs.PartitionFirst(low, high)                                       // Partition using the selected pivot.
}

// GetComparisons returns the total number of comparisons made during sorting.
func (qs *QuickSort) GetComparisons() int {
	return qs.comparisons
}

// LoadData loads integers from a file and returns them as a slice of integers.
// Each line in the file is expected to contain a single integer.
func LoadData(filename string) ([]int, error) {
	file, err := os.Open(filename) // Open the file.
	if err != nil {
		return nil, err // Return an error if the file cannot be opened.
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Warning: error closing file: %v\n", err)
		}
	}()

	var arr []int
	scanner := bufio.NewScanner(file) // Create a scanner to read the file line by line.
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text()) // Convert each line to an integer.
		arr = append(arr, num)                 // Add the integer to the array.
	}
	if err := scanner.Err(); err != nil {
		return nil, err // Return an error if reading fails.
	}
	return arr, nil // Return the populated array.
}

// main is the entry point of the program.
func main() {
	// Load the array from the specified file.
	array, err := LoadData("course_1/module_3/programming_assignment_3/QuickSort.txt")
	if err != nil {
		fmt.Println("Error loading data:", err) // Print an error if file loading fails.
		return
	}

	// Sort using the first element as the pivot and print the number of comparisons.
	qsFirst := NewQuickSort(array)
	qsFirst.QuickSortFirst(0, len(array)-1)
	fmt.Printf("Total comparisons with first element as pivot: %d\n", qsFirst.GetComparisons())

	// Sort using the last element as the pivot and print the number of comparisons.
	qsLast := NewQuickSort(array)
	qsLast.QuickSortLast(0, len(array)-1)
	fmt.Printf("Total comparisons with last element as pivot: %d\n", qsLast.GetComparisons())

	// Sort using the median-of-three as the pivot and print the number of comparisons.
	qsMedian := NewQuickSort(array)
	qsMedian.QuickSortMedian(0, len(array)-1)
	fmt.Printf("Total comparisons with median-of-three as pivot: %d\n", qsMedian.GetComparisons())
}
