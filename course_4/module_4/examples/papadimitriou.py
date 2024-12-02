import random
import math
from typing import List, Tuple, Optional


class PapadimitriouTwoSAT:
    def __init__(self, num_variables: int, clauses: List[Tuple[int, int]], seed: Optional[int] = None):
        """
        Initializes the Papadimitriou's 2-SAT solver.
        :param num_variables: Number of variables in the 2-SAT problem.
        :param clauses: List of clauses, where each clause is a tuple of two literals.
                        Positive integers represent variables (e.g., 1 means x1 is True),
                        Negative integers represent negated variables (e.g., -1 means x1 is False).
        :param seed: Optional seed for random number generation for reproducibility.
        """
        self.num_variables = num_variables
        self.clauses = clauses
        if seed is not None:
            random.seed(seed)  # Set the random seed for reproducibility.

    def is_satisfiable(self) -> Tuple[bool, List[Optional[bool]]]:
        """
        Determines if the 2-SAT problem is satisfiable using Papadimitriou's algorithm.
        :return: A tuple (is_satisfiable, assignment).
                 is_satisfiable: True if satisfiable, False otherwise.
                 assignment: A satisfying assignment if one exists, else an empty list.
        """
        max_trials = int(round(math.log2(self.num_variables)))  # Number of independent trials (outer loop).
        max_attempts = 2 * self.num_variables ** 2  # Maximum local moves per trial (inner loop).

        for trial in range(max_trials):
            # Generate a random initial assignment for all variables.
            assignment = [random.choice([True, False]) for _ in range(self.num_variables)]

            for _ in range(max_attempts):
                # Check if the current assignment satisfies all clauses.
                unsatisfied_clauses = [clause for clause in self.clauses if not self._is_clause_satisfied(clause, assignment)]

                if not unsatisfied_clauses:  # If no unsatisfied clauses, return the satisfying assignment.
                    return True, assignment

                # Pick a random unsatisfied clause.
                clause = random.choice(unsatisfied_clauses)

                # Randomly pick one of the two variables in the clause to flip.
                variable_to_flip = abs(random.choice(clause)) - 1  # Convert to 0-based index.
                assignment[variable_to_flip] = not assignment[variable_to_flip]  # Flip the variable.

        # If no satisfying assignment is found, return unsatisfiable.
        return False, []

    @staticmethod
    def _is_clause_satisfied(clause: Tuple[int, int], assignment: List[bool]) -> bool:
        """
        Checks if a clause is satisfied under the current assignment.
        :param clause: A tuple of two literals (e.g., (1, -2)).
        :param assignment: The current variable assignment as a list of booleans.
        :return: True if the clause is satisfied, False otherwise.
        """
        u, v = clause  # Extract the two literals.
        # Determine the truth value of each literal based on the current assignment.
        value_u = assignment[abs(u) - 1] if u > 0 else not assignment[abs(u) - 1]
        value_v = assignment[abs(v) - 1] if v > 0 else not assignment[abs(v) - 1]
        # The clause is satisfied if either literal is True.
        return value_u or value_v


# Example usage:
if __name__ == "__main__":
    # Example 2-SAT problem with 4 variables and 4 clauses.
    num_vars = 4
    example_clauses = [
        (1, 2),     # x1 OR x2
        (-1, 3),    # NOT x1 OR x3
        (3, 4),     # x3 OR x4
        (-2, -4)    # NOT x2 OR NOT x4
    ]

    # Create a solver instance and run the algorithm.
    solver = PapadimitriouTwoSAT(num_vars, example_clauses, seed=42)
    satisfiable, assignment = solver.is_satisfiable()

    if satisfiable:
        print("The 2-SAT problem is satisfiable with the assignment:")
        for i, value in enumerate(assignment, start=1):
            print(f"x{i} = {'True' if value else 'False'}")
    else:
        print("The 2-SAT problem is unsatisfiable.")
