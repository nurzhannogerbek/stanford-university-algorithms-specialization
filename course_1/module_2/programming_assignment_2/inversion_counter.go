package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type InversionCounter struct{}

// countInversions initiates the inversion counting process.
func (ic *InversionCounter) countInversions(array []int) int {
	tempArray := make([]int, len(array)) // Temporary array for merging
	return ic.countInversionsRecursive(array, tempArray, 0, len(array)-1)
}

// countInversionsRecursive recursively counts inversions using divide and conquer.
func (ic *InversionCounter) countInversionsRecursive(array, tempArray []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := (left + right) / 2

	invCount := ic.countInversionsRecursive(array, tempArray, left, mid)
	invCount += ic.countInversionsRecursive(array, tempArray, mid+1, right)
	invCount += ic.mergeAndCount(array, tempArray, left, mid, right)

	return invCount
}

// mergeAndCount merges two halves and counts the inversions.
func (ic *InversionCounter) mergeAndCount(array, tempArray []int, left, mid, right int) int {
	i, j, k := left, mid+1, left
	invCount := 0

	// Merge the two halves while counting inversions.
	for i <= mid && j <= right {
		if array[i] <= array[j] {
			tempArray[k] = array[i]
			i++
		} else {
			tempArray[k] = array[j]
			invCount += (mid - i + 1) // All remaining elements in the left half are greater than array[j].
			j++
		}
		k++
	}

	// Copy any remaining elements from the left half.
	for i <= mid {
		tempArray[k] = array[i]
		i++
		k++
	}

	// Copy any remaining elements from the right half.
	for j <= right {
		tempArray[k] = array[j]
		j++
		k++
	}

	// Copy the sorted and merged subarray back into the original array.
	copy(array[left:right+1], tempArray[left:right+1])

	return invCount
}

func main() {
	// Get the path to the current directory where the script is located.
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Specify the path to the file 'IntegerArray.txt' relative to the current directory.
	filePath := filepath.Join(currentDir, "course_1", "module_2", "programming_assignment_2", "IntegerArray.txt")

	// Read the file and parse integers into an array.
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var inputArray []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Skipping invalid integer: %s\n", scanner.Text())
			continue
		}
		inputArray = append(inputArray, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Create an instance of InversionCounter and count inversions.
	counter := InversionCounter{}
	inversionCount := counter.countInversions(inputArray)

	fmt.Println("Number of inversions:", inversionCount)
}
