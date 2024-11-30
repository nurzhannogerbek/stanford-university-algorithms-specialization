package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Job represents a job with weight and length.
type Job struct {
	Weight int
	Length int
}

// WeightLengthRatio calculates the ratio (weight / length).
func (j Job) WeightLengthRatio() float64 {
	if j.Length == 0 {
		panic("Length of job cannot be zero")
	}
	return float64(j.Weight) / float64(j.Length)
}

// Scheduler handles job scheduling to minimize weighted completion time.
type Scheduler struct {
	Jobs []Job
}

// NewScheduler creates a new Scheduler instance.
func NewScheduler(jobs []Job) *Scheduler {
	return &Scheduler{Jobs: jobs}
}

// ScheduleJobs sorts jobs based on the ratio (weight / length) in descending order.
func (s *Scheduler) ScheduleJobs() {
	sort.Slice(s.Jobs, func(i, j int) bool {
		return s.Jobs[i].WeightLengthRatio() > s.Jobs[j].WeightLengthRatio()
	})
}

// CalculateWeightedCompletionTime calculates the total weighted completion time.
func (s *Scheduler) CalculateWeightedCompletionTime() int {
	totalTime := 0
	completionTime := 0

	for _, job := range s.Jobs {
		completionTime += job.Length
		totalTime += job.Weight * completionTime
	}

	return totalTime
}

// ReadJobsFromFile reads jobs from a file and returns a list of Job objects.
func ReadJobsFromFile(filePath string) ([]Job, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var jobs []Job
	scanner := bufio.NewScanner(file)
	var lineNum int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineNum++

		if lineNum == 1 {
			// First line is the number of jobs, skip it.
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format at line %d: %s", lineNum, line)
		}

		weight, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("invalid weight at line %d: %s", lineNum, parts[0])
		}

		length, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid length at line %d: %s", lineNum, parts[1])
		}

		jobs = append(jobs, Job{Weight: weight, Length: length})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if len(jobs) == 0 {
		return nil, fmt.Errorf("file contains no jobs")
	}

	return jobs, nil
}

func main() {
	// Explicitly set the file path.
	filePath := "course_3/module_1/programming_assignment_1/jobs.txt"

	// Read jobs from the specified file.
	jobs, err := ReadJobsFromFile(filePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	// Initialize the scheduler.
	scheduler := NewScheduler(jobs)

	// Schedule jobs and calculate the weighted completion time.
	scheduler.ScheduleJobs()
	totalWeightedCompletionTime := scheduler.CalculateWeightedCompletionTime()

	// Print the total weighted completion time.
	fmt.Printf("Total Weighted Completion Time: %d\n", totalWeightedCompletionTime)
}
