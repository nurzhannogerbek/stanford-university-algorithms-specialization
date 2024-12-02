from collections import deque


class Graph:
    """
    Class to represent a weighted directed graph using adjacency list.
    """

    INF = float("inf")  # Constant to represent infinity.

    def __init__(self, vertices):
        """
        Initialize the graph with the given number of vertices.

        :param vertices: Number of vertices in the graph.
        """
        self.V = vertices  # Number of vertices in the graph.
        self.adjacency_list = {i: [] for i in range(vertices)}  # Adjacency list.

    def add_edge(self, u, v, weight):
        """
        Add a directed edge from vertex u to vertex v with the given weight.

        :param u: Source vertex.
        :param v: Destination vertex.
        :param weight: Weight of the edge.
        """
        if not (0 <= u < self.V) or not (0 <= v < self.V):
            raise ValueError("Vertex indices must be within valid range.")
        self.adjacency_list[u].append((v, weight))  # Add edge to adjacency list.

    def bellman_ford(self, source):
        """
        Perform the Bellman-Ford algorithm to find shortest paths from the source vertex.

        :param source: The source vertex.
        :return: Tuple containing a list of distances and a list of predecessor pointers.
        """
        # Initialize distances to all vertices as infinity and source as 0.
        distances = [self.INF] * self.V
        distances[source] = 0

        # Initialize predecessor pointers for path reconstruction.
        predecessors = [None] * self.V

        # Relax edges |V| - 1 times.
        for _ in range(self.V - 1):
            for u in self.adjacency_list:
                for v, weight in self.adjacency_list[u]:
                    if distances[u] != self.INF and distances[u] + weight < distances[v]:
                        distances[v] = distances[u] + weight
                        predecessors[v] = u

        # Check for negative weight cycles.
        for u in self.adjacency_list:
            for v, weight in self.adjacency_list[u]:
                if distances[u] != self.INF and distances[u] + weight < distances[v]:
                    return None, None  # Indicates a negative weight cycle.

        return distances, predecessors

    def reconstruct_path(self, source, target, predecessors):
        """
        Reconstruct the shortest path from source to target using predecessor pointers.

        :param source: The source vertex.
        :param target: The target vertex.
        :param predecessors: The list of predecessor pointers.
        :return: List representing the shortest path from source to target.
        """
        if predecessors[target] is None:
            raise ValueError("Target vertex is not reachable from the source.")

        path = deque()  # Use deque for efficient prepend operations.
        current = target

        # Trace back from target to source using predecessors.
        while current is not None:
            path.appendleft(current)
            current = predecessors[current]

        if path[0] != source:
            raise ValueError("No path exists from source to target.")

        return list(path)  # Convert deque to list for final output.


# Example usage:
if __name__ == "__main__":
    # Create a graph with 5 vertices.
    g = Graph(5)

    # Add edges to the graph.
    g.add_edge(0, 1, 6)
    g.add_edge(0, 3, 7)
    g.add_edge(1, 2, 5)
    g.add_edge(1, 3, 8)
    g.add_edge(1, 4, -4)
    g.add_edge(2, 1, -2)
    g.add_edge(3, 2, -3)
    g.add_edge(3, 4, 9)
    g.add_edge(4, 0, 2)
    g.add_edge(4, 2, 7)

    source = 0

    try:
        # Perform the Bellman-Ford algorithm.
        distances, predecessors = g.bellman_ford(source)

        if distances is None:
            print("The graph contains a negative weight cycle.")
        else:
            # Print distances from the source to each vertex.
            print("Shortest distances from source:", distances)

            # Reconstruct and print the path to a specific target.
            target = 4
            path = g.reconstruct_path(source, target, predecessors)
            print(f"Shortest path from {source} to {target}:", path)
    except ValueError as e:
        print(e)
