package main

import (
	"fmt"
)

// MergeSorter is a structure for sorting arrays using Merge Sort.
type MergeSorter struct {
	array []int
}

// NewMergeSorter initializes a new instance of MergeSorter with the provided array.
func NewMergeSorter(array []int) *MergeSorter {
	if array == nil {
		panic("Input array cannot be nil")
	}
	return &MergeSorter{array: array}
}

// Sort performs the Merge Sort on the array without modifying the original array.
func (ms *MergeSorter) Sort() []int {
	if len(ms.array) <= 1 {
		return ms.array
	}
	return ms.mergeSort(ms.array)
}

// mergeSort recursively splits and sorts the array.
func (ms *MergeSorter) mergeSort(arr []int) []int {
	// Base case: if the array has 1 or 0 elements, it is already sorted.
	if len(arr) <= 1 {
		return arr
	}

	// Split the array into two halves
	mid := len(arr) / 2
	left := ms.mergeSort(arr[:mid])
	right := ms.mergeSort(arr[mid:])

	// Merge the sorted halves
	return ms.merge(left, right)
}

// merge combines two sorted slices into a single sorted slice.
func (ms *MergeSorter) merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Compare elements from both halves and add the smaller one to merged slice.
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	// Append remaining elements from left or right slice.
	if i < len(left) {
		merged = append(merged, left[i:]...)
	} else {
		merged = append(merged, right[j:]...)
	}

	return merged
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	mergeSorter := NewMergeSorter(arr)
	sortedArr := mergeSorter.Sort()
	fmt.Println("Sorted array:", sortedArr)
}
