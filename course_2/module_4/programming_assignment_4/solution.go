package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type TwoSumSolver struct {
	filePath   string
	lowerBound int
	upperBound int
	numbers    map[int]struct{}
}

// NewTwoSumSolver creates a new instance of the TwoSumSolver.
func NewTwoSumSolver(filePath string, lowerBound, upperBound int) *TwoSumSolver {
	return &TwoSumSolver{
		filePath:   filePath,
		lowerBound: lowerBound,
		upperBound: upperBound,
		numbers:    make(map[int]struct{}),
	}
}

// LoadData reads integers from the file into a map for fast lookup.
func (solver *TwoSumSolver) LoadData() error {
	file, err := os.Open(solver.filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Printf("Warning: failed to close file: %v\n", cerr)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return fmt.Errorf("failed to parse number: %w", err)
		}
		solver.numbers[num] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error while reading file: %w", err)
	}

	return nil
}

// countTargetsInChunk calculates valid target values for a chunk of numbers.
func (solver *TwoSumSolver) countTargetsInChunk(numbersChunk []int, results chan<- map[int]struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	targets := make(map[int]struct{})

	for _, num := range numbersChunk {
		for t := solver.lowerBound; t <= solver.upperBound; t++ {
			complement := t - num
			if complement != num {
				if _, exists := solver.numbers[complement]; exists {
					targets[t] = struct{}{}
				}
			}
		}
	}
	results <- targets
}

// CountTargetValues calculates the number of distinct target values in parallel.
func (solver *TwoSumSolver) CountTargetValues() int {
	// Convert map keys to a slice for chunking.
	numbers := make([]int, 0, len(solver.numbers))
	for num := range solver.numbers {
		numbers = append(numbers, num)
	}

	// Divide the numbers into chunks for parallel processing.
	numWorkers := 8
	chunkSize := (len(numbers) + numWorkers - 1) / numWorkers
	results := make(chan map[int]struct{}, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < len(numbers); i += chunkSize {
		end := i + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		wg.Add(1)
		go solver.countTargetsInChunk(numbers[i:end], results, &wg)
	}

	// Close the results channel after all goroutines finish.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Combine results from all goroutines.
	finalTargets := make(map[int]struct{})
	for partialTargets := range results {
		for t := range partialTargets {
			finalTargets[t] = struct{}{}
		}
	}

	return len(finalTargets)
}

func main() {
	// Path to the input file.
	filePath := "course_2/module_4/programming_assignment_4/2sum.txt" // Replace with the actual file path.
	lowerBound := -10000                                              // Lower bound of the range.
	upperBound := 10000                                               // Upper bound of the range.

	// Initialize the TwoSumSolver.
	solver := NewTwoSumSolver(filePath, lowerBound, upperBound)

	// Load data from the file.
	err := solver.LoadData()
	if err != nil {
		fmt.Printf("Error loading data: %v\n", err)
		return
	}

	// Compute the number of distinct target values.
	result := solver.CountTargetValues()
	fmt.Printf("The number of target values in the range [%d, %d] is: %d\n", lowerBound, upperBound, result)
}
