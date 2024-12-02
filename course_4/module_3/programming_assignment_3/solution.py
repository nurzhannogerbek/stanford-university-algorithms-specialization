import math


class City:
    """
    Represents a city with x and y coordinates.
    """
    def __init__(self, x: float, y: float):
        self.x = x  # X-coordinate of the city.
        self.y = y  # Y-coordinate of the city.

    def distance_to(self, other: "City") -> float:
        """
        Calculates the Euclidean distance to another city.
        """
        dx = self.x - other.x
        dy = self.y - other.y
        return math.sqrt(dx * dx + dy * dy)


def load_tsp_data(file_path: str) -> list[City]:
    """
    Loads the TSP data from the file and returns a list of City objects.
    """
    with open(file_path, "r") as file:
        lines = file.readlines()

    # Read the number of cities from the first line.
    n_cities = int(lines[0].strip())

    # Parse city coordinates.
    cities = []
    for line in lines[1:]:
        parts = line.strip().split()
        if len(parts) == 3:
            _, x, y = map(float, parts)  # Skip the city index.
            cities.append(City(x, y))

    # Validate the number of cities matches the declared count.
    if len(cities) != n_cities:
        raise ValueError("Mismatch between declared and actual number of cities.")

    return cities


def nearest_neighbor_tsp(cities: list[City]) -> float:
    """
    Solves the Traveling Salesman Problem using the nearest neighbor heuristic.
    Returns the total distance of the computed tour.
    """
    n_cities = len(cities)
    visited = [False] * n_cities  # Track whether each city has been visited.
    curr_city_index = 0  # Start at the first city.
    visited[curr_city_index] = True
    n_visited = 1
    total_distance = 0.0

    while n_visited < n_cities:
        next_city_index = -1
        min_dist = float("inf")

        # Find the nearest unvisited city.
        for i in range(n_cities):
            if not visited[i]:
                dist = cities[curr_city_index].distance_to(cities[i])
                if dist < min_dist:
                    min_dist = dist
                    next_city_index = i

        # Move to the next city.
        total_distance += min_dist
        visited[next_city_index] = True
        curr_city_index = next_city_index
        n_visited += 1

    # Add the distance to return to the starting city.
    total_distance += cities[curr_city_index].distance_to(cities[0])
    return total_distance


def main():
    file_path = "nn.txt"

    try:
        # Load the TSP data.
        cities = load_tsp_data(file_path)

        # Solve the TSP using the nearest neighbor heuristic.
        total_distance = nearest_neighbor_tsp(cities)

        # Output the total distance, rounded down to the nearest integer.
        print(f"Total Distance: {int(math.floor(total_distance))}")
    except Exception as e:
        print(f"Error: {e}")


if __name__ == "__main__":
    main()
