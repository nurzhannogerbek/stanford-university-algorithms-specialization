package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// MinHeap structure.
// A min-heap stores the larger half of the numbers.
type MinHeap struct {
	data []int
}

// Len Implementing heap.Interface methods for MinHeap.
func (h MinHeap) Len() int           { return len(h.data) }                          // Returns the number of elements in the heap.
func (h MinHeap) Less(i, j int) bool { return h.data[i] < h.data[j] }                // Ensures that the heap property is maintained.
func (h MinHeap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] } // Swaps two elements in the heap.

func (h *MinHeap) Push(x interface{}) {
	// Adds a new element to the heap.
	h.data = append(h.data, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	// Removes and returns the smallest element in the heap.
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func (h *MinHeap) Top() int {
	// Returns the smallest element in the heap without removing it.
	if len(h.data) == 0 {
		return 0
	}
	return h.data[0]
}

// MaxHeap structure.
// A max-heap stores the smaller half of the numbers.
type MaxHeap struct {
	data []int
}

// Len Implementing heap.Interface methods for MaxHeap.
func (h MaxHeap) Len() int           { return len(h.data) }                          // Returns the number of elements in the heap.
func (h MaxHeap) Less(i, j int) bool { return h.data[i] > h.data[j] }                // Reverse comparison for max-heap.
func (h MaxHeap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] } // Swaps two elements in the heap.

func (h *MaxHeap) Push(x interface{}) {
	// Adds a new element to the heap.
	h.data = append(h.data, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	// Removes and returns the largest element in the heap.
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func (h *MaxHeap) Top() int {
	// Returns the largest element in the heap without removing it.
	if len(h.data) == 0 {
		return 0
	}
	return h.data[0]
}

// MedianMaintenance structure.
// Maintains the running median using two heaps.
type MedianMaintenance struct {
	minHeap *MinHeap // Min-heap for the larger half of the numbers.
	maxHeap *MaxHeap // Max-heap for the smaller half of the numbers.
}

// NewMedianMaintenance initializes a new instance of MedianMaintenance.
func NewMedianMaintenance() *MedianMaintenance {
	return &MedianMaintenance{
		minHeap: &MinHeap{data: []int{}},
		maxHeap: &MaxHeap{data: []int{}},
	}
}

// AddNumber adds a number to the data structure and maintains the heap balance.
func (m *MedianMaintenance) AddNumber(num int) {
	// Add the number to the appropriate heap.
	if m.maxHeap.Len() == 0 || num <= m.maxHeap.Top() {
		heap.Push(m.maxHeap, num)
	} else {
		heap.Push(m.minHeap, num)
	}

	// Balance the heaps if their sizes differ by more than 1.
	if m.maxHeap.Len() > m.minHeap.Len()+1 {
		heap.Push(m.minHeap, heap.Pop(m.maxHeap))
	} else if m.minHeap.Len() > m.maxHeap.Len() {
		heap.Push(m.maxHeap, heap.Pop(m.minHeap))
	}
}

// GetMedian returns the current median of the numbers.
func (m *MedianMaintenance) GetMedian() int {
	// Median is the top of the larger heap.
	if m.maxHeap.Len() >= m.minHeap.Len() {
		return m.maxHeap.Top()
	}
	return m.minHeap.Top()
}

// readNumbersFromFile reads numbers from a file and returns them as a slice of integers.
func readNumbersFromFile(filePath string) ([]int, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	numbers := make([]int, len(lines))

	for i, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid number on line %d: %v", i+1, err)
		}
		numbers[i] = num
	}

	return numbers, nil
}

// calculateMedianSum calculates the sum of medians modulo 10000.
func calculateMedianSum(numbers []int) int {
	medianMaintenance := NewMedianMaintenance()
	medianSum := 0

	for _, number := range numbers {
		// Add the number to the data structure.
		medianMaintenance.AddNumber(number)
		// Get the current median and update the sum.
		median := medianMaintenance.GetMedian()
		medianSum = (medianSum + median) % 10000
	}

	return medianSum
}

func main() {
	// Replace this path with the actual path to the input file.
	filePath := "course_2/module_3/programming_assignment_3/Median.txt"

	// Read numbers from the file.
	numbers, err := readNumbersFromFile(filePath)
	if err != nil {
		fmt.Printf("Error reading numbers: %v\n", err)
		return
	}

	// Calculate the sum of medians modulo 10000.
	result := calculateMedianSum(numbers)
	fmt.Printf("The sum of the medians modulo 10000 is: %d\n", result)
}
