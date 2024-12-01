class WeightedIndependentSet:
    """
    Class to solve the Weighted Independent Set problem for path graphs.
    """

    def __init__(self, weights):
        """
        Initialize the problem with the weights of the vertices.
        :param weights: List of weights for the vertices in the path graph.
        """
        self.weights = weights
        self.n = len(weights)

    def naive_recursive(self):
        """
        Solve the WIS problem using a naive recursive approach.
        :return: Maximum weight of an independent set.
        """

        def _solve(n):
            if n == 0:
                return 0
            if n == 1:
                return self.weights[0]
            # Compare excluding or including the current vertex.
            exclude = _solve(n - 1)
            include = self.weights[n - 1] + _solve(n - 2)
            return max(exclude, include)

        return _solve(self.n)

    def greedy(self):
        """
        Solve the WIS problem using a greedy approach.
        :return: Maximum weight and the selected vertices.
        """
        selected = []
        for i in range(self.n):
            if i == 0 or (self.weights[i] > self.weights[i - 1] and (i - 1 not in selected)):
                selected.append(i)
        total_weight = sum(self.weights[i] for i in selected)
        return total_weight, selected

    def divide_and_conquer(self):
        """
        Solve the WIS problem using a divide-and-conquer approach.
        :return: Maximum weight of an independent set.
        """

        def _solve(start, end):
            if start > end:
                return 0
            if start == end:
                return self.weights[start]
            mid = (start + end) // 2
            # Solve for left and right parts.
            left = _solve(start, mid - 1)
            right = _solve(mid + 1, end)
            return max(left + self.weights[mid], right)

        return _solve(0, self.n - 1)

    def dynamic_programming(self):
        """
        Solve the WIS problem using dynamic programming.
        :return: Maximum weight and the selected vertices.
        """
        if self.n == 0:
            return 0, []

        dp = [0] * (self.n + 1)  # Table to store the maximum weight at each step.
        dp[1] = self.weights[0]

        # Fill the DP table.
        for i in range(2, self.n + 1):
            dp[i] = max(dp[i - 1], dp[i - 2] + self.weights[i - 1])

        # Reconstruct the solution.
        selected = []
        i = self.n
        while i >= 1:
            if dp[i] == dp[i - 1]:
                i -= 1  # Exclude the current vertex.
            else:
                selected.append(i - 1)  # Include the current vertex.
                i -= 2

        selected.reverse()
        return dp[self.n], selected


# Example usage:
if __name__ == "__main__":
    weights = [1, 4, 5, 4]  # Example graph weights.

    wis_solver = WeightedIndependentSet(weights)

    print("Naive Recursive Approach:")
    print("Maximum Weight:", wis_solver.naive_recursive())

    print("\nGreedy Algorithm:")
    weight, selected = wis_solver.greedy()
    print("Maximum Weight:", weight)
    print("Selected Vertices:", selected)

    print("\nDivide and Conquer Approach:")
    print("Maximum Weight:", wis_solver.divide_and_conquer())

    print("\nDynamic Programming Approach:")
    weight, selected = wis_solver.dynamic_programming()
    print("Maximum Weight:", weight)
    print("Selected Vertices:", selected)
