package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// PapadimitriouTwoSAT represents the Papadimitriou's 2-SAT solver.
type PapadimitriouTwoSAT struct {
	NumVariables int      // Number of boolean variables in the problem.
	Clauses      [][2]int // List of clauses, each clause is a pair of literals.
}

// NewPapadimitriouTwoSAT initializes a new instance of Papadimitriou's 2-SAT solver.
// numVariables: Number of variables.
// clauses: List of clauses, where positive integers represent variables (e.g., 1 means x1 is true),
//
//	and negative integers represent negated variables (e.g., -1 means x1 is false).
func NewPapadimitriouTwoSAT(numVariables int, clauses [][2]int) *PapadimitriouTwoSAT {
	return &PapadimitriouTwoSAT{
		NumVariables: numVariables,
		Clauses:      clauses,
	}
}

// IsSatisfiable determines if the 2-SAT problem is satisfiable using Papadimitriou's algorithm.
// Returns a boolean indicating satisfiability and a satisfying assignment if one exists.
func (solver *PapadimitriouTwoSAT) IsSatisfiable() (bool, []bool) {
	maxTrials := int(math.Ceil(math.Log2(float64(solver.NumVariables)))) // Number of independent trials (outer loop).
	maxAttempts := 2 * solver.NumVariables * solver.NumVariables         // Maximum local moves per trial (inner loop).

	rand.Seed(time.Now().UnixNano()) // Seed for random number generation.

	for trial := 0; trial < maxTrials; trial++ {
		// Generate a random initial assignment for all variables.
		assignment := make([]bool, solver.NumVariables)
		for i := range assignment {
			assignment[i] = rand.Intn(2) == 1
		}

		for attempt := 0; attempt < maxAttempts; attempt++ {
			// Find all unsatisfied clauses.
			unsatisfiedClauses := solver.getUnsatisfiedClauses(assignment)

			// If no unsatisfied clauses remain, the problem is satisfiable.
			if len(unsatisfiedClauses) == 0 {
				return true, assignment
			}

			// Pick a random unsatisfied clause.
			clause := unsatisfiedClauses[rand.Intn(len(unsatisfiedClauses))]

			// Randomly pick one of the two variables in the clause to flip.
			variableToFlip := abs(clause[rand.Intn(2)]) - 1          // Convert to 0-based index.
			assignment[variableToFlip] = !assignment[variableToFlip] // Flip the variable.
		}
	}

	// If no satisfying assignment is found after all trials, return unsatisfiable.
	return false, nil
}

// getUnsatisfiedClauses identifies clauses that are not satisfied by the current assignment.
func (solver *PapadimitriouTwoSAT) getUnsatisfiedClauses(assignment []bool) [][2]int {
	var unsatisfied [][2]int
	for _, clause := range solver.Clauses {
		if !isClauseSatisfied(clause, assignment) {
			unsatisfied = append(unsatisfied, clause)
		}
	}
	return unsatisfied
}

// isClauseSatisfied checks if a given clause is satisfied under the current assignment.
// clause: A pair of literals (e.g., (1, -2)).
// assignment: Current assignment of variables.
// Returns true if the clause is satisfied, false otherwise.
func isClauseSatisfied(clause [2]int, assignment []bool) bool {
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

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Example usage.
func main() {
	// Example 2-SAT problem with 4 variables and 4 clauses.
	numVars := 4
	clauses := [][2]int{
		{1, 2},   // x1 OR x2
		{-1, 3},  // NOT x1 OR x3
		{3, 4},   // x3 OR x4
		{-2, -4}, // NOT x2 OR NOT x4
	}

	// Create a solver instance.
	solver := NewPapadimitriouTwoSAT(numVars, clauses)

	// Run the algorithm to check satisfiability.
	satisfiable, assignment := solver.IsSatisfiable()

	// Print the results.
	if satisfiable {
		fmt.Println("The 2-SAT problem is satisfiable with the assignment:")
		for i, value := range assignment {
			fmt.Printf("x%d = %t\n", i+1, value)
		}
	} else {
		fmt.Println("The 2-SAT problem is unsatisfiable.")
	}
}
