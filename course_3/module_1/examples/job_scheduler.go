package main

import (
	"fmt"
	"sort"
)

// Job represents a task with a weight and a length.
type Job struct {
	ID                  int
	Weight, Length      float64
	WeightToLengthRatio float64 // Pre-computed for optimization
}

// NewJob creates a new Job with validation.
func NewJob(id int, weight, length float64) (Job, error) {
	if weight < 0 {
		return Job{}, fmt.Errorf("weight must be non-negative")
	}
	if length <= 0 {
		return Job{}, fmt.Errorf("length must be greater than zero")
	}
	return Job{
		ID:                  id,
		Weight:              weight,
		Length:              length,
		WeightToLengthRatio: weight / length,
	}, nil
}

// Scheduler manages a list of jobs and schedules them using a greedy algorithm.
type Scheduler struct {
	jobs []Job
}

// NewScheduler creates a new Scheduler with a list of jobs.
func NewScheduler(jobs []Job) (*Scheduler, error) {
	if len(jobs) == 0 {
		return nil, fmt.Errorf("job list cannot be empty")
	}
	return &Scheduler{jobs: jobs}, nil
}

// ScheduleJobs sorts the jobs in descending order of their weight-to-length ratio.
func (s *Scheduler) ScheduleJobs() []Job {
	// Create a copy of jobs to avoid modifying the original slice.
	scheduledJobs := make([]Job, len(s.jobs))
	copy(scheduledJobs, s.jobs)

	// Sort jobs by pre-computed weight-to-length ratio in descending order.
	sort.SliceStable(scheduledJobs, func(i, j int) bool {
		return scheduledJobs[i].WeightToLengthRatio > scheduledJobs[j].WeightToLengthRatio
	})
	return scheduledJobs
}

// CalculateWeightedCompletionTime calculates the total weighted sum of completion times.
func (s *Scheduler) CalculateWeightedCompletionTime(jobs []Job) float64 {
	totalCompletionTime := 0.0
	currentTime := 0.0

	for _, job := range jobs {
		currentTime += job.Length
		totalCompletionTime += job.Weight * currentTime
	}
	return totalCompletionTime
}

func main() {
	// Define a list of jobs.
	jobs := []Job{
		{ID: 1, Weight: 3, Length: 2, WeightToLengthRatio: 3.0 / 2.0},
		{ID: 2, Weight: 1, Length: 3, WeightToLengthRatio: 1.0 / 3.0},
		{ID: 3, Weight: 2, Length: 1, WeightToLengthRatio: 2.0 / 1.0},
	}

	// Create a Scheduler instance.
	scheduler, err := NewScheduler(jobs)
	if err != nil {
		fmt.Println("Error creating scheduler:", err)
		return
	}

	// Schedule jobs using the greedy algorithm.
	scheduledJobs := scheduler.ScheduleJobs()

	// Print the scheduled jobs.
	fmt.Println("Scheduled jobs (in execution order):")
	for _, job := range scheduledJobs {
		fmt.Printf("Job(ID=%d, Weight=%.2f, Length=%.2f)\n", job.ID, job.Weight, job.Length)
	}

	// Calculate the total weighted sum of completion times.
	totalWeightedCompletionTime := scheduler.CalculateWeightedCompletionTime(scheduledJobs)
	fmt.Printf("\nWeighted sum of completion times: %.2f\n", totalWeightedCompletionTime)
}
