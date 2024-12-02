from heapq import heappop, heappush
import math

class Graph:
    """
    A class to represent a directed graph for Johnson's Algorithm.
    """

    def __init__(self, vertices):
        """
        Initialize the graph with a specified number of vertices.

        :param vertices: The number of vertices in the graph.
        """
        self.vertices = vertices
        self.edges = []  # List to store edges as (u, v, weight).

    def add_edge(self, u, v, weight):
        """
        Add a directed edge to the graph.

        :param u: The source vertex.
        :param v: The destination vertex.
        :param weight: The weight of the edge.
        """
        self.edges.append((u, v, weight))

    def _bellman_ford(self, source):
        """
        Perform the Bellman-Ford algorithm to find shortest paths from a source.

        :param source: The source vertex.
        :return: A list of shortest distances from the source to all vertices.
                 Returns None if a negative cycle is detected.
        """
        distances = [math.inf] * self.vertices
        distances[source] = 0

        # Relax edges up to |V|-1 times.
        for _ in range(self.vertices - 1):
            for u, v, weight in self.edges:
                if distances[u] + weight < distances[v]:
                    distances[v] = distances[u] + weight

        # Check for negative weight cycles.
        for u, v, weight in self.edges:
            if distances[u] + weight < distances[v]:
                return None  # Negative cycle detected.

        return distances

    def _dijkstra(self, source, weights):
        """
        Perform Dijkstra's algorithm to find shortest paths from a source.

        :param source: The source vertex.
        :param weights: Adjusted weights of the edges for Dijkstra's algorithm.
        :return: A list of shortest distances from the source to all vertices.
        """
        distances = [math.inf] * self.vertices
        distances[source] = 0
        priority_queue = [(0, source)]

        while priority_queue:
            current_distance, u = heappop(priority_queue)

            if current_distance > distances[u]:
                continue

            for v, weight in weights.get(u, []):
                new_distance = distances[u] + weight
                if new_distance < distances[v]:
                    distances[v] = new_distance
                    heappush(priority_queue, (new_distance, v))

        return distances

    def johnsons_algorithm(self):
        """
        Perform Johnson's Algorithm to find shortest paths between all pairs.

        :return: A matrix of shortest path distances.
                 Returns None if a negative cycle is detected.
        """
        # Step 1: Add a new vertex s connected to all vertices with 0-weight edges.
        self.edges.extend((self.vertices, v, 0) for v in range(self.vertices))
        self.vertices += 1  # Temporarily increase vertex count.

        # Step 2: Perform Bellman-Ford from the new vertex.
        h = self._bellman_ford(self.vertices - 1)
        self.vertices -= 1  # Restore original vertex count.

        if h is None:
            return None  # Negative cycle detected.

        # Step 3: Adjust edge weights using the vertex weights h.
        reweighted_edges = {}
        for u, v, weight in self.edges:
            if u < self.vertices:
                reweighted_edges.setdefault(u, []).append((v, weight + h[u] - h[v]))

        # Step 4: Perform Dijkstra's algorithm for each vertex.
        all_pairs_distances = []
        for u in range(self.vertices):
            distances = self._dijkstra(u, reweighted_edges)
            # Adjust distances back to original weights.
            adjusted_distances = [
                d - h[u] + h[v] if d != math.inf else math.inf for v, d in enumerate(distances)
            ]
            all_pairs_distances.append(adjusted_distances)

        return all_pairs_distances


# Example usage:
if __name__ == "__main__":
    # Create a graph with 5 vertices.
    graph = Graph(5)

    # Add edges with their weights.
    graph.add_edge(0, 1, 3)
    graph.add_edge(0, 2, 8)
    graph.add_edge(1, 3, 1)
    graph.add_edge(2, 3, -4)
    graph.add_edge(3, 4, 2)
    graph.add_edge(4, 0, -1)

    # Run Johnson's Algorithm.
    distances = graph.johnsons_algorithm()

    if distances is None:
        print("The graph contains a negative weight cycle.")
    else:
        print("Shortest distances between all pairs of vertices:")
        for u, row in enumerate(distances):
            print(f"From vertex {u}: {row}")
