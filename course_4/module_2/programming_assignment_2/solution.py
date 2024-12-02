import math


class TravelingSalesmanProblem:
    """
    A class to represent and solve the Traveling Salesman Problem using recursion and memoization.
    """

    def __init__(self, filename):
        """
        Initialize the TSP solver by reading data from the given file.
        :param filename: Path to the file containing the TSP instance.
        """
        self.filename = filename
        self.num_cities = 0
        self.coordinates = []
        self.distance_matrix = []
        self.VISITED_ALL = 0
        self.memo = {}

        # Read data from the file and prepare the distance matrix.
        self._load_data()
        self._prepare_distance_matrix()
        self.VISITED_ALL = (1 << self.num_cities) - 1  # Bitmask for all cities visited.

    def _load_data(self):
        """
        Load the number of cities and their coordinates from the file.
        """
        with open(self.filename, 'r') as file:
            lines = file.readlines()
            self.num_cities = int(lines[0].strip())  # First line specifies the number of cities.
            for line in lines[1:]:
                x, y = map(float, line.strip().split())
                self.coordinates.append((x, y))  # Store the coordinates as tuples (x, y).

    def _prepare_distance_matrix(self):
        """
        Prepare the distance matrix based on Euclidean distances between cities.
        """
        self.distance_matrix = [[0] * self.num_cities for _ in range(self.num_cities)]
        for i in range(self.num_cities):
            for j in range(self.num_cities):
                if i != j:
                    self.distance_matrix[i][j] = self._euclidean_distance(self.coordinates[i], self.coordinates[j])

    @staticmethod
    def _euclidean_distance(city1, city2):
        """
        Calculate the Euclidean distance between two cities.
        :param city1: Coordinates (x, y) of the first city.
        :param city2: Coordinates (x, y) of the second city.
        :return: Euclidean distance between the two cities.
        """
        x1, y1 = city1
        x2, y2 = city2
        return math.sqrt((x2 - x1) ** 2 + (y2 - y1) ** 2)

    def _solve_recursive(self, mask, pos):
        """
        Recursive function to compute the minimum cost of the TSP using bitmasking and memoization.
        :param mask: Bitmask representing visited cities.
        :param pos: Current city position.
        :return: Minimum cost of visiting all cities and returning to the start.
        """
        # Base case: all cities have been visited.
        if mask == self.VISITED_ALL:
            return self.distance_matrix[pos][0]  # Return the cost to return to the starting city.

        # Check if the result for this state is already memoized.
        key = (mask, pos)
        if key in self.memo:
            return self.memo[key]

        # Try visiting all unvisited cities and find the minimum cost.
        ans = math.inf
        for city in range(self.num_cities):
            if not (mask & (1 << city)):  # If the city has not been visited.
                new_cost = self.distance_matrix[pos][city] + self._solve_recursive(mask | (1 << city), city)
                ans = min(ans, new_cost)

        # Memoize the result for this state.
        self.memo[key] = ans
        return ans

    def solve(self):
        """
        Solve the TSP problem.
        :return: The minimum cost of the traveling salesman tour.
        """
        # Start the recursion from the first city with only the first city visited (mask = 1).
        return self._solve_recursive(1, 0)


# Example usage:
if __name__ == "__main__":
    # Path to the TSP instance file.
    filename = "tsp.txt"

    # Initialize the TSP solver with the input file.
    tsp_solver = TravelingSalesmanProblem(filename)

    # Solve the TSP and print the minimum cost.
    min_cost = tsp_solver.solve()
    print(f"Minimum cost: {math.floor(min_cost)}")
