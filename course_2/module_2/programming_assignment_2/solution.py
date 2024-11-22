import heapq
import sys
from collections import defaultdict

class DijkstraAlgorithm:
    def __init__(self):
        # Initialize the graph as an adjacency list.
        self.graph = defaultdict(list)

    def load_graph(self, filename):
        # Load the graph from a file and populate the adjacency list.
        try:
            with open(filename, "r") as file:
                for line in file:
                    parts = line.split()
                    source = int(parts[0]) - 1
                    for edge in parts[1:]:
                        dest, length = map(int, edge.split(","))
                        self.graph[source].append((dest - 1, length))
        except FileNotFoundError:
            print(f"Error: File '{filename}' not found.")
            sys.exit(1)
        except ValueError as e:
            print(f"Error: Invalid file format. Details: {e}")
            sys.exit(1)

    def dijkstra(self, start_vertex):
        # Implement Dijkstra's algorithm using a priority queue.
        INF = sys.maxsize
        shortest_distances = [INF] * len(self.graph)
        shortest_distances[start_vertex] = 0

        priority_queue = [(0, start_vertex)]  # (distance, vertex)
        while priority_queue:
            current_distance, current_vertex = heapq.heappop(priority_queue)

            # Skip processing if we've already found a shorter path.
            if current_distance > shortest_distances[current_vertex]:
                continue

            for neighbor, weight in self.graph[current_vertex]:
                distance = current_distance + weight
                if distance < shortest_distances[neighbor]:
                    shortest_distances[neighbor] = distance
                    heapq.heappush(priority_queue, (distance, neighbor))

        return shortest_distances

    @staticmethod
    def print_target_distances(distances, target_vertices):
        # Print the shortest distances to the specified target vertices.
        results = [distances[target - 1] for target in target_vertices]
        print(",".join(map(str, results)))


if __name__ == "__main__":
    # Initialize the DijkstraAlgorithm object.
    dijkstra_solver = DijkstraAlgorithm()

    # Load the graph data from the input file.
    input_file = "dijkstraData.txt"
    dijkstra_solver.load_graph(input_file)

    # Run Dijkstra's algorithm from the start vertex (0-based index).
    shortest_distances = dijkstra_solver.dijkstra(start_vertex=0)

    # List of target vertices to print distances for (1-based indices).
    target_vertices = [7, 37, 59, 82, 99, 115, 133, 165, 188, 197]

    # Print the shortest distances to the target vertices.
    dijkstra_solver.print_target_distances(shortest_distances, target_vertices)

