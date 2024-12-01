class PathGraphMWIS:
    """
    Class to solve the Maximum Weight Independent Set (MWIS) problem in a path graph.
    """

    def __init__(self, weights):
        """
        Initialize the graph with the vertex weights.
        :param weights: List of weights for the vertices in the path graph.
        """
        self.weights = weights
        self.n = len(weights)
        self.dp = [0] * (self.n + 1)  # DP table for storing max weights.
        self.solution = []  # List to store the indices of selected vertices.

    def compute_mwis(self):
        """
        Compute the maximum weight independent set using dynamic programming.
        """
        if self.n == 0:
            return
        if self.n >= 1:
            self.dp[1] = self.weights[0]

        # Fill the DP table.
        for i in range(2, self.n + 1):
            self.dp[i] = max(self.dp[i - 1], self.dp[i - 2] + self.weights[i - 1])

    def reconstruct_solution(self):
        """
        Reconstruct the MWIS from the computed DP table.
        """
        i = self.n
        while i >= 1:
            if i == 1 or self.dp[i] != self.dp[i - 1]:  # Current vertex is included.
                self.solution.append(i - 1)  # Convert to 0-based index.
                i -= 2
            else:  # Current vertex is excluded.
                i -= 1

        self.solution.reverse()  # Reverse to get the solution in ascending order.

    def get_solution_string(self, vertices):
        """
        Generate the 8-bit solution string for the specified vertices.
        :param vertices: List of vertices (1-based indices) to check.
        :return: An 8-bit string where 1 indicates inclusion in MWIS.
        """
        solution_set = set(self.solution)
        return "".join("1" if v - 1 in solution_set else "0" for v in vertices)

    def solve(self, vertices):
        """
        Solve the MWIS problem and generate the solution string for the specified vertices.
        :param vertices: List of vertices to check.
        :return: 8-bit solution string.
        """
        self.compute_mwis()
        self.reconstruct_solution()
        return self.get_solution_string(vertices)


# Example usage
if __name__ == "__main__":
    # Read weights from the file.
    with open("mwis.txt", "r") as file:
        lines = file.readlines()
        num_vertices = int(lines[0].strip())
        weights = [int(line.strip()) for line in lines[1:]]

    # Initialize the MWIS solver.
    mwis_solver = PathGraphMWIS(weights)

    # Define the vertices to check (1-based indices).
    vertices_to_check = [1, 2, 3, 4, 17, 117, 517, 997]

    # Solve the problem and get the solution string.
    solution_string = mwis_solver.solve(vertices_to_check)

    # Print the result.
    print("Solution String:", solution_string)
