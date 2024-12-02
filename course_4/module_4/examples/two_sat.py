import random
from typing import List, Tuple


class TwoSAT:
    def __init__(self, num_variables: int, clauses: List[Tuple[int, int]], seed: int = None, debug: bool = False):
        """
        Initializes a 2-SAT problem instance.
        :param num_variables: Number of boolean variables in the problem.
        :param clauses: List of clauses, where each clause is a tuple of two literals.
                        Positive integers represent variables (e.g., 1 means x1 is True),
                        Negative integers represent negated variables (e.g., -1 means x1 is False).
        :param seed: Seed for random number generation. Used for reproducibility.
        :param debug: If True, prints debug information during execution.
        """
        self.num_variables = num_variables
        self.clauses = clauses
        self.debug = debug
        if seed is not None:
            random.seed(seed)  # Set the random seed if provided for reproducible results.

        # Validate that all literals in clauses reference valid variable indices.
        if any(abs(literal) > num_variables for clause in clauses for literal in clause):
            raise ValueError("Clause contains variable out of bounds.")

    def is_satisfiable(self) -> Tuple[bool, List[bool]]:
        """
        Determines whether the 2-SAT problem is satisfiable using randomized local search.
        :return: A tuple (is_satisfiable, assignment).
                 is_satisfiable: True if satisfiable, False otherwise.
                 assignment: A list representing the variable assignments if satisfiable, else empty.
        """
        # Set the maximum number of flips and trials for the algorithm.
        max_attempts = 2 * self.num_variables ** 2  # Maximum flips per trial.
        max_trials = 10 * self.num_variables  # Maximum number of independent trials.

        # Perform up to max_trials independent runs of the local search algorithm.
        for trial in range(max_trials):
            # Generate a random initial assignment for all variables.
            assignment = [random.choice([True, False]) for _ in range(self.num_variables)]

            # Perform up to max_attempts flips to try to satisfy all clauses.
            for flip in range(max_attempts):
                # Identify all clauses that are not currently satisfied.
                unsatisfied = [clause for clause in self.clauses if not self._is_clause_satisfied(clause, assignment)]

                if not unsatisfied:  # If no unsatisfied clauses remain, the problem is satisfiable.
                    return True, assignment

                # Randomly choose an unsatisfied clause.
                clause = random.choice(unsatisfied)

                # Randomly choose one of the two literals in the clause and flip its value.
                variable_to_flip = abs(random.choice(clause)) - 1  # Convert to 0-based index.
                assignment[variable_to_flip] = not assignment[variable_to_flip]

                # If debug mode is enabled, print the current trial, flip, and assignment.
                if self.debug:
                    print(f"Trial {trial}, Flip {flip}, Assignment: {assignment}")

        # If no satisfying assignment is found after all trials, the problem is unsatisfiable.
        return False, []

    @staticmethod
    def _is_clause_satisfied(clause: Tuple[int, int], assignment: List[bool]) -> bool:
        """
        Checks whether a given clause is satisfied by the current assignment.
        :param clause: A tuple of two literals (e.g., (1, -2)).
        :param assignment: A list of current assignments for all variables.
        :return: True if the clause is satisfied, False otherwise.
        """
        u, v = clause  # Extract the two literals from the clause.
        # Determine the value of the first literal based on its sign.
        value_u = assignment[abs(u) - 1] if u > 0 else not assignment[abs(u) - 1]
        # Determine the value of the second literal based on its sign.
        value_v = assignment[abs(v) - 1] if v > 0 else not assignment[abs(v) - 1]
        # The clause is satisfied if either literal is True.
        return value_u or value_v


# Example usage:
if __name__ == "__main__":
    # Example 2-SAT problem with 4 variables and 4 clauses.
    num_vars = 4  # Number of variables.
    example_clauses = [
        (1, 2),     # x1 OR x2
        (-1, 3),    # NOT x1 OR x3
        (3, 4),     # x3 OR x4
        (-2, -4)    # NOT x2 OR NOT x4
    ]

    # Create a TwoSAT solver instance with debug mode enabled.
    two_sat_solver = TwoSAT(num_vars, example_clauses, debug=True)
    # Run the solver to check satisfiability.
    is_satisfiable, assignment = two_sat_solver.is_satisfiable()

    if is_satisfiable:
        # If satisfiable, print the satisfying assignment for each variable.
        print(f"The 2-SAT problem is satisfiable with the assignment:")
        for i, value in enumerate(assignment, start=1):
            print(f"x{i} = {'True' if value else 'False'}")
    else:
        # If unsatisfiable, print a corresponding message.
        print("The 2-SAT problem is unsatisfiable.")
