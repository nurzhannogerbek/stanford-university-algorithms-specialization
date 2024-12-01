from typing import List, Optional, Tuple


class OptimalBinarySearchTree:
    """
    A class to solve the Optimal Binary Search Tree problem using dynamic programming.
    """

    def __init__(self, probabilities: List[float]) -> None:
        """
        Initialize the class with a list of probabilities.

        :param probabilities: List of probabilities (or weights) for each key.
        """
        if not probabilities or any(p < 0 for p in probabilities):
            raise ValueError("Probabilities must be a non-empty list of non-negative numbers.")

        self.probabilities = probabilities
        self.n = len(probabilities)
        self.dp = [[0] * self.n for _ in range(self.n)]  # Table to store the minimum cost.
        self.root = [[0] * self.n for _ in range(self.n)]  # Table to store roots for the intervals.
        self.prefix_sum = self._compute_prefix_sum()  # Precompute prefix sums for fast interval sum calculation.

    def _compute_prefix_sum(self) -> List[float]:
        """
        Compute prefix sums for the probabilities to allow fast interval sum queries.

        :return: List of prefix sums.
        """
        prefix_sum = [0] * (self.n + 1)
        for i in range(self.n):
            prefix_sum[i + 1] = prefix_sum[i] + self.probabilities[i]
        return prefix_sum

    def calculate_sum(self, i: int, j: int) -> float:
        """
        Calculate the sum of probabilities from index i to j, inclusive.

        :param i: Start index.
        :param j: End index.
        :return: Sum of probabilities from i to j.
        """
        return self.prefix_sum[j + 1] - self.prefix_sum[i]

    def solve(self, optimize: bool = False) -> None:
        """
        Solve the Optimal Binary Search Tree problem using dynamic programming.

        :param optimize: Whether to use the O(n^2) optimized approach.
        """
        # Initialize the base case where each interval contains a single key.
        for i in range(self.n):
            self.dp[i][i] = self.probabilities[i]
            self.root[i][i] = i  # Root is the single key itself.

        # Fill the table for intervals of increasing sizes.
        for s in range(1, self.n):  # Size of the interval minus one.
            for i in range(self.n - s):  # Start of the interval.
                j = i + s  # End of the interval.
                self.dp[i][j] = float('inf')  # Initialize to infinity.
                total_prob = self.calculate_sum(i, j)  # Sum of probabilities in the interval.

                # Determine the range of roots to consider.
                root_range = (
                    range(self.root[i][j - 1], self.root[i + 1][j] + 1)
                    if optimize and j - i > 1
                    else range(i, j + 1)
                )

                # Try each key as the root and find the minimum cost.
                for r in root_range:
                    cost_left = self.dp[i][r - 1] if r > i else 0
                    cost_right = self.dp[r + 1][j] if r < j else 0
                    total_cost = cost_left + cost_right + total_prob

                    if total_cost < self.dp[i][j]:
                        self.dp[i][j] = total_cost
                        self.root[i][j] = r  # Store the root for this interval.

    def get_optimal_cost(self) -> float:
        """
        Get the optimal cost for the entire set of keys.

        :return: The minimum cost to construct the optimal binary search tree.
        """
        return self.dp[0][self.n - 1]

    def reconstruct_tree(self, i: int, j: int) -> Optional[Tuple[int, Optional[Tuple], Optional[Tuple]]]:
        """
        Reconstruct the optimal binary search tree from the DP table.

        :param i: Start index of the interval.
        :param j: End index of the interval.
        :return: A tuple representing the tree structure (root, left subtree, right subtree).
        """
        if i > j:
            return None  # No keys in this interval.
        root = self.root[i][j]
        left_subtree = self.reconstruct_tree(i, root - 1)
        right_subtree = self.reconstruct_tree(root + 1, j)
        return root, left_subtree, right_subtree

    def pretty_print_tree(self, tree: Optional[Tuple], depth: int = 0) -> None:
        """
        Pretty print the optimal binary search tree.

        :param tree: The tree structure to print.
        :param depth: The current depth in the tree.
        """
        if tree is None:
            return
        root, left, right = tree
        print("  " * depth + f"Node: {root}")
        self.pretty_print_tree(left, depth + 1)
        self.pretty_print_tree(right, depth + 1)


# Example usage:
if __name__ == "__main__":
    probabilities = [0.1, 0.2, 0.4, 0.3]

    # Solve using cubic approach.
    print("Cubic Approach:")
    obst_cubic = OptimalBinarySearchTree(probabilities)
    obst_cubic.solve(optimize=False)
    print("Optimal cost (cubic):", obst_cubic.get_optimal_cost())
    print("Tree structure (cubic):")
    obst_cubic.pretty_print_tree(obst_cubic.reconstruct_tree(0, obst_cubic.n - 1))

    # Solve using quadratic approach.
    print("\nQuadratic Approach:")
    obst_quadratic = OptimalBinarySearchTree(probabilities)
    obst_quadratic.solve(optimize=True)
    print("Optimal cost (quadratic):", obst_quadratic.get_optimal_cost())
    print("Tree structure (quadratic):")
    obst_quadratic.pretty_print_tree(obst_quadratic.reconstruct_tree(0, obst_quadratic.n - 1))
