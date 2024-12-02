package main

import (
	"fmt"
	"math/rand"
)

// TwoSAT represents a 2-SAT problem instance.
type TwoSAT struct {
	numVariables int      // Number of variables in the problem.
	clauses      [][2]int // List of clauses, where each clause is a pair of literals.
	debug        bool     // If true, enables debug output.
}

// NewTwoSAT initializes a new 2-SAT problem instance.
func NewTwoSAT(numVariables int, clauses [][2]int, debug bool) (*TwoSAT, error) {
	// Validate that all literals in clauses are within the valid range of variables.
	for _, clause := range clauses {
		for _, literal := range clause {
			if abs(literal) > numVariables {
				return nil, fmt.Errorf("invalid literal %d: exceeds number of variables", literal)
			}
		}
	}
	return &TwoSAT{
		numVariables: numVariables,
		clauses:      clauses,
		debug:        debug,
	}, nil
}

// IsSatisfiable determines whether the 2-SAT problem instance is satisfiable.
// It uses randomized local search to find a satisfying assignment if it exists.
func (ts *TwoSAT) IsSatisfiable() (bool, []bool) {
	maxAttempts := 2 * ts.numVariables * ts.numVariables // Maximum flips per trial.
	maxTrials := 10 * ts.numVariables                    // Maximum number of independent trials.

	// Perform up to maxTrials independent runs of the local search algorithm.
	for trial := 0; trial < maxTrials; trial++ {
		// Generate a random initial assignment for all variables.
		assignment := make([]bool, ts.numVariables)
		for i := range assignment {
			assignment[i] = rand.Intn(2) == 1 // Randomly set each variable to true or false.
		}

		// Perform up to maxAttempts flips to try to satisfy all clauses.
		for flip := 0; flip < maxAttempts; flip++ {
			// Identify unsatisfied clauses.
			unsatisfied := ts.getUnsatisfiedClauses(assignment)
			if len(unsatisfied) == 0 {
				// If no unsatisfied clauses remain, the problem is satisfiable.
				return true, assignment
			}

			// Randomly choose an unsatisfied clause.
			clause := unsatisfied[rand.Intn(len(unsatisfied))]

			// Randomly choose one of the two literals in the clause and flip its value.
			variableToFlip := abs(clause[rand.Intn(2)]) - 1
			assignment[variableToFlip] = !assignment[variableToFlip]

			// Print debug information if enabled.
			if ts.debug {
				fmt.Printf("Trial %d, Flip %d, Assignment: %v\n", trial, flip, assignment)
			}
		}
	}

	// If no satisfying assignment is found after all trials, the problem is unsatisfiable.
	return false, nil
}

// getUnsatisfiedClauses returns a list of unsatisfied clauses for the current assignment.
func (ts *TwoSAT) getUnsatisfiedClauses(assignment []bool) [][2]int {
	unsatisfied := [][2]int{}
	for _, clause := range ts.clauses {
		if !ts.isClauseSatisfied(clause, assignment) {
			unsatisfied = append(unsatisfied, clause)
		}
	}
	return unsatisfied
}

// isClauseSatisfied checks whether a given clause is satisfied by the current assignment.
func (ts *TwoSAT) isClauseSatisfied(clause [2]int, assignment []bool) bool {
	u, v := clause[0], clause[1]
	valueU := assignment[abs(u)-1]
	if u < 0 {
		valueU = !valueU
	}
	valueV := assignment[abs(v)-1]
	if v < 0 {
		valueV = !valueV
	}
	return valueU || valueV
}

// Helper function to compute the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Example 2-SAT problem with 4 variables and 4 clauses.
	numVars := 4
	clauses := [][2]int{
		{1, 2},   // x1 OR x2
		{-1, 3},  // NOT x1 OR x3
		{3, 4},   // x3 OR x4
		{-2, -4}, // NOT x2 OR NOT x4
	}

	// Initialize the TwoSAT solver.
	twoSAT, err := NewTwoSAT(numVars, clauses, true)
	if err != nil {
		fmt.Printf("Error initializing 2-SAT: %v\n", err)
		return
	}

	// Check satisfiability.
	isSatisfiable, assignment := twoSAT.IsSatisfiable()
	if isSatisfiable {
		// Print the satisfying assignment if one is found.
		fmt.Println("The 2-SAT problem is satisfiable with the following assignment:")
		for i, value := range assignment {
			fmt.Printf("x%d = %v\n", i+1, value)
		}
	} else {
		// Print a message indicating that the problem is unsatisfiable.
		fmt.Println("The 2-SAT problem is unsatisfiable.")
	}
}
