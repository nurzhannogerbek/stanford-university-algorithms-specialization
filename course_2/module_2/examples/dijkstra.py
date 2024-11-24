import heapq
import math


class Graph:
    """
    Represents a weighted graph using an adjacency list.
    """

    def __init__(self):
        """Initialize the graph with an empty adjacency list."""
        self.edges = {}

    def add_node(self, node):
        """
        Add a node to the graph.

        Args:
            node: The node to add.
        """
        if node not in self.edges:
            self.edges[node] = {}

    def add_edge(self, from_node, to_node, weight):
        """
        Add a directed edge with a weight to the graph.

        Args:
            from_node: The starting node.
            to_node: The ending node.
            weight: The weight of the edge.
        """
        self.add_node(from_node)
        self.add_node(to_node)
        self.edges[from_node][to_node] = weight


class Dijkstra:
    """
    Implements Dijkstra's algorithm for finding shortest paths in a graph.
    """

    @staticmethod
    def compute_shortest_paths(graph, start_node):
        """
        Compute the shortest paths from the start node to all other nodes.

        Args:
            graph: The graph to process.
            start_node: The starting node.

        Returns:
            distances: A dictionary where the keys are nodes and the values are the shortest distances.
            previous: A dictionary for reconstructing the shortest paths.
        """
        if start_node not in graph.edges:
            raise ValueError(f"Start node '{start_node}' is not in the graph.")

        # Initialize distances to all nodes as infinity.
        distances = {node: math.inf for node in graph.edges}
        distances[start_node] = 0

        # Dictionary to reconstruct paths.
        previous = {node: None for node in graph.edges}

        # Priority queue to store (distance, node) pairs.
        pq = []
        heapq.heappush(pq, (0, start_node))

        while pq:
            # Extract the node with the smallest distance.
            current_distance, current_node = heapq.heappop(pq)

            # If the distance is outdated, skip processing.
            if current_distance > distances[current_node]:
                continue

            # Update distances to neighbors.
            for neighbor, weight in graph.edges.get(current_node, {}).items():
                new_distance = current_distance + weight
                if new_distance < distances[neighbor]:
                    distances[neighbor] = new_distance
                    previous[neighbor] = current_node
                    heapq.heappush(pq, (new_distance, neighbor))

        return distances, previous

    @staticmethod
    def reconstruct_path(previous, start_node, end_node):
        """
        Reconstruct the shortest path from the start node to the end node.

        Args:
            previous: The dictionary of predecessors.
            start_node: The starting node.
            end_node: The ending node.

        Returns:
            A list representing the shortest path.
        """
        path = []
        current = end_node
        while current is not None:
            path.insert(0, current)
            current = previous[current]

        if path[0] != start_node:  # If the path does not start with the start_node.
            return []  # No valid path exists.

        return path


if __name__ == "__main__":
    # Create a graph and add edges.
    graph = Graph()
    graph.add_edge("A", "B", 1)
    graph.add_edge("A", "C", 4)
    graph.add_edge("B", "C", 2)
    graph.add_edge("B", "D", 6)
    graph.add_edge("C", "D", 3)
    graph.add_edge("C", "E", 5)
    graph.add_edge("D", "E", 1)

    # Compute shortest distances and paths from node "A".
    dijkstra = Dijkstra()
    distances, previous = dijkstra.compute_shortest_paths(graph, "A")

    # Print the shortest distances from node "A".
    print("Shortest distances from A:")
    for node, distance in distances.items():
        if distance == math.inf:
            print(f"To {node}: unreachable")
        else:
            print(f"To {node}: {distance}")

    # Print the shortest paths from "A" to all other nodes.
    print("\nShortest paths from A:")
    for node in graph.edges:
        path = dijkstra.reconstruct_path(previous, "A", node)
        if path:
            print(f"To {node}: {' -> '.join(path)}")
        else:
            print(f"To {node}: No path found")
