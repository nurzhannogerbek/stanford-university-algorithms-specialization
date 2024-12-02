class FloydWarshall:
    """
    A class to represent and solve the All-Pairs Shortest Paths problem using the Floyd-Warshall algorithm.
    """

    def __init__(self, vertices):
        """
        Initialize the FloydWarshall object.

        :param vertices: The number of vertices in the graph.
        """
        self.vertices = vertices
        # Initialize the distance matrix with infinity (no direct path).
        self.distances = [[float('inf')] * vertices for _ in range(vertices)]
        # Initialize the next matrix to reconstruct paths.
        self.next = [[None] * vertices for _ in range(vertices)]
        # Distance to self is zero.
        for i in range(vertices):
            self.distances[i][i] = 0

    def add_edge(self, u, v, weight):
        """
        Add an edge to the graph with a specified weight.

        :param u: The source vertex (0-indexed).
        :param v: The destination vertex (0-indexed).
        :param weight: The weight of the edge.
        :raises ValueError: If vertex indices are out of range.
        """
        if not (0 <= u < self.vertices and 0 <= v < self.vertices):
            raise ValueError("Vertex indices must be within valid range.")
        self.distances[u][v] = weight
        self.next[u][v] = v  # Set the next vertex in the shortest path.

    def run_algorithm(self):
        """
        Run the Floyd-Warshall algorithm to compute shortest paths between all pairs of vertices.
        """
        # Loop over each intermediate vertex.
        for k in range(self.vertices):
            # Loop over each source vertex.
            for i in range(self.vertices):
                # Loop over each destination vertex.
                for j in range(self.vertices):
                    # Check if going through vertex k offers a shorter path.
                    if self.distances[i][k] + self.distances[k][j] < self.distances[i][j]:
                        # Update the shortest path distance.
                        self.distances[i][j] = self.distances[i][k] + self.distances[k][j]
                        # Update the next vertex in the path.
                        self.next[i][j] = self.next[i][k]

    def has_negative_cycle(self):
        """
        Check if the graph contains a negative cycle.

        :return: True if a negative cycle exists, False otherwise.
        """
        for i in range(self.vertices):
            if self.distances[i][i] < 0:
                return True
        return False

    def get_shortest_distance(self, u, v):
        """
        Get the shortest distance between two vertices.

        :param u: The source vertex (0-indexed).
        :param v: The destination vertex (0-indexed).
        :return: The shortest distance, or float('inf') if no path exists.
        :raises ValueError: If vertex indices are out of range.
        """
        if not (0 <= u < self.vertices and 0 <= v < self.vertices):
            raise ValueError("Vertex indices must be within valid range.")
        return self.distances[u][v]

    def reconstruct_path(self, u, v):
        """
        Reconstruct the shortest path between two vertices.

        :param u: The source vertex (0-indexed).
        :param v: The destination vertex (0-indexed).
        :return: A list representing the path from u to v, or an empty list if no path exists.
        :raises ValueError: If vertex indices are out of range.
        """
        if not (0 <= u < self.vertices and 0 <= v < self.vertices):
            raise ValueError("Vertex indices must be within valid range.")

        # If there's no path, return an empty list.
        if self.distances[u][v] == float('inf'):
            return []

        # Initialize the path with the source vertex.
        path = [u]
        while u != v:
            u = self.next[u][v]  # Follow the next pointers to reconstruct the path.
            path.append(u)

        return path


# Example usage:
if __name__ == "__main__":
    # Create a graph with 4 vertices.
    graph = FloydWarshall(4)

    # Add edges with their weights.
    graph.add_edge(0, 1, 5)
    graph.add_edge(0, 3, 10)
    graph.add_edge(1, 2, 3)
    graph.add_edge(2, 3, 1)

    # Run the Floyd-Warshall algorithm.
    graph.run_algorithm()

    # Check for negative cycles.
    if graph.has_negative_cycle():
        print("The graph contains a negative weight cycle.")
    else:
        # Get the shortest distance between vertices 0 and 3.
        print("Shortest distance from 0 to 3:", graph.get_shortest_distance(0, 3))

        # Reconstruct the shortest path from vertex 0 to 3.
        print("Shortest path from 0 to 3:", graph.reconstruct_path(0, 3))
