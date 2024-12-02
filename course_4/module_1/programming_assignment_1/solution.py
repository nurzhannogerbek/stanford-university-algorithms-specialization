from math import inf


class Graph:
    """
    Represents a graph using an adjacency matrix.
    Implements the Floyd-Warshall algorithm for shortest path calculations.
    """

    def __init__(self, vertices):
        """
        Initializes the graph with a given number of vertices.
        :param vertices: Number of vertices in the graph.
        """
        self.vertices = vertices
        self.edges = [[inf] * vertices for _ in range(vertices)]
        for i in range(vertices):
            self.edges[i][i] = 0  # Distance to itself is always 0.

    def add_edge(self, u, v, weight):
        """
        Adds an edge to the graph.
        :param u: Source vertex.
        :param v: Destination vertex.
        :param weight: Edge weight.
        """
        self.edges[u][v] = weight

    def floyd_warshall(self):
        """
        Executes the Floyd-Warshall algorithm to find shortest paths between all pairs of vertices.
        :return: A tuple containing the distance matrix and a boolean indicating if a negative cycle exists.
        """
        dist = [row[:] for row in self.edges]  # Create a copy of the adjacency matrix.

        for k in range(self.vertices):
            for i in range(self.vertices):
                for j in range(self.vertices):
                    if dist[i][k] != inf and dist[k][j] != inf:
                        dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])

        # Check for negative weight cycles.
        has_negative_cycle = any(dist[i][i] < 0 for i in range(self.vertices))
        return dist, has_negative_cycle

    def get_shortest_shortest_path(self):
        """
        Computes the shortest path among all vertex pairs.
        :return: The shortest path length or None if a negative weight cycle exists.
        """
        dist, has_negative_cycle = self.floyd_warshall()
        if has_negative_cycle:
            return None

        shortest = inf
        for i in range(self.vertices):
            for j in range(self.vertices):
                if i != j and dist[i][j] < shortest:
                    shortest = dist[i][j]

        return shortest if shortest != inf else None


def read_graph_from_file(filename):
    """
    Reads a graph from a file.
    :param filename: Path to the file containing graph data.
    :return: A Graph object.
    """
    with open(filename, 'r') as file:
        lines = file.readlines()
        vertices, _ = map(int, lines[0].split())
        graph = Graph(vertices)
        for line in lines[1:]:
            u, v, weight = map(int, line.split())
            graph.add_edge(u - 1, v - 1, weight)  # Convert to 0-based indexing.
    return graph


def main():
    """
    Main function to calculate shortest paths for multiple graphs.
    """
    filenames = ["g1.txt", "g2.txt", "g3.txt"]

    for i, filename in enumerate(filenames, 1):
        graph = read_graph_from_file(filename)
        result = graph.get_shortest_shortest_path()
        if result is None:
            print(f"Graph g{i}: Negative cycle detected.")
        else:
            print(f"Graph g{i}: Shortest shortest path = {result:.2f}")


if __name__ == "__main__":
    main()
