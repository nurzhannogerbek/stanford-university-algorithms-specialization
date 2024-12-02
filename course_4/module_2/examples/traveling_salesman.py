import math

class TravelingSalesmanProblem:
    """
    Solve the Traveling Salesman Problem using bitmasking and recursion with memoization.
    """

    def __init__(self, cost_matrix):
        """
        Initialize the TSP solver.
        :param cost_matrix: A 2D list representing the cost matrix. cost_matrix[i][j] is the cost of traveling from vertex i to vertex j.
        """
        self.cost_matrix = cost_matrix
        self.num_vertices = len(cost_matrix)
        self.VISITED_ALL = (1 << self.num_vertices) - 1
        self.memo = {}

    def tsp(self, mask, pos):
        """
        Recursive function to solve TSP using bitmasking and memoization.
        :param mask: Bitmask representing visited cities.
        :param pos: Current position (city).
        :return: Minimum cost of visiting all cities and returning to the starting point.
        """
        # Base case: all cities have been visited
        if mask == self.VISITED_ALL:
            return self.cost_matrix[pos][0]  # Return to the starting city.

        # Check if result is already memoized
        if (mask, pos) in self.memo:
            return self.memo[(mask, pos)]

        # Try visiting all unvisited cities and take the minimum cost path
        ans = math.inf
        for city in range(self.num_vertices):
            if not (mask & (1 << city)):  # If city is not visited
                new_cost = self.cost_matrix[pos][city] + self.tsp(mask | (1 << city), city)
                ans = min(ans, new_cost)

        # Memoize and return the result
        self.memo[(mask, pos)] = ans
        return ans


# Example usage:
if __name__ == "__main__":
    cost_matrix = [
        [0, 20, 42, 25],
        [20, 0, 30, 34],
        [42, 30, 0, 10],
        [25, 34, 10, 0]
    ]
    tsp_solver = TravelingSalesmanProblem(cost_matrix)
    min_cost = tsp_solver.tsp(1, 0)  # Start from city 0 with only city 0 visited
    print(f"Minimum cost: {min_cost}")
