package main

import (
	"bufio"
	"fmt"
	"log"
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

// WeightLengthDifference calculates the difference (weight - length).
func (j Job) WeightLengthDifference() int {
	return j.Weight - j.Length
}

// Scheduler handles scheduling jobs to minimize weighted completion time.
type Scheduler struct {
	Jobs []Job
}

// NewScheduler creates a new Scheduler instance.
func NewScheduler(jobs []Job) *Scheduler {
	return &Scheduler{Jobs: jobs}
}

// ScheduleJobs sorts jobs based on the difference (weight - length).
// If two jobs have the same difference, it prioritizes by weight.
func (s *Scheduler) ScheduleJobs() {
	sort.Slice(s.Jobs, func(i, j int) bool {
		diff1 := s.Jobs[i].WeightLengthDifference()
		diff2 := s.Jobs[j].WeightLengthDifference()
		if diff1 == diff2 {
			return s.Jobs[i].Weight > s.Jobs[j].Weight
		}
		return diff1 > diff2
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

		// Skip the first line, which contains the number of jobs.
		if lineNum == 1 {
			continue
		}

		// Parse job details from the line.
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

		// Add the job to the list.
		jobs = append(jobs, Job{Weight: weight, Length: length})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return jobs, nil
}

func main() {
	// Define the input file path.
	filePath := "course_3/module_1/programming_assignment_1/jobs.txt"

	// Read jobs from the file.
	jobs, err := ReadJobsFromFile(filePath)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	// Ensure we have jobs to process.
	if len(jobs) == 0 {
		log.Fatalf("Error: No jobs found in the file.")
	}

	// Initialize the scheduler.
	scheduler := NewScheduler(jobs)

	// Schedule jobs and calculate the weighted completion time.
	scheduler.ScheduleJobs()
	totalWeightedCompletionTime := scheduler.CalculateWeightedCompletionTime()

	// Print the result.
	fmt.Printf("Total Weighted Completion Time: %d\n", totalWeightedCompletionTime)
}
