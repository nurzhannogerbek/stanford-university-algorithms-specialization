import sys
from collections import defaultdict, deque
from typing import List
import argparse
import heapq

sys.setrecursionlimit(10**6)

class Graph:
    """Class to represent a directed graph."""

    def __init__(self, vertices: int):
        """Initialize the graph with the given number of vertices."""
        self.vertices = vertices
        self.graph = defaultdict(list)
        self.rev_graph = defaultdict(list)

    def add_edge(self, from_vertex: int, to_vertex: int):
        """Add a directed edge to the graph."""
        self.graph[from_vertex].append(to_vertex)
        self.rev_graph[to_vertex].append(from_vertex)

    def _dfs(self, vertex: int, visited: set, stack: deque):
        """Perform depth-first search iteratively and store the finishing times."""
        local_stack = deque([vertex])
        while local_stack:
            node = local_stack[-1]
            if node not in visited:
                visited.add(node)
                for neighbor in self.graph[node]:
                    if neighbor not in visited:
                        local_stack.append(neighbor)
            else:
                local_stack.pop()
                stack.append(node)

    def _reverse_dfs(self, vertex: int, visited: set) -> int:
        """Perform depth-first search iteratively on the reversed graph."""
        local_stack = deque([vertex])
        size = 0
        while local_stack:
            node = local_stack.pop()
            if node not in visited:
                visited.add(node)
                size += 1
                for neighbor in self.rev_graph[node]:
                    if neighbor not in visited:
                        local_stack.append(neighbor)
        return size

    def kosaraju(self) -> List[int]:
        """Perform Kosaraju's algorithm to find strongly connected components."""
        stack = deque()
        visited = set()

        # Step 1: Perform a DFS to determine the finishing order.
        for vertex in range(1, self.vertices + 1):
            if vertex not in visited:
                self._dfs(vertex, visited, stack)

        visited = set()
        scc_sizes = []

        # Step 2: Perform DFS on the reversed graph to find SCCs.
        while stack:
            vertex = stack.pop()
            if vertex not in visited:
                scc_size = self._reverse_dfs(vertex, visited)
                scc_sizes.append(scc_size)

        # Return the 5 largest SCCs.
        return heapq.nlargest(5, scc_sizes)


def parse_args():
    """Parse command-line arguments."""
    parser = argparse.ArgumentParser(description="Find SCCs in a graph using Kosaraju's algorithm.")
    parser.add_argument(
        "input_file",
        nargs="?",
        default="SCC.txt",  # Default value if no argument is provided.
        help="Path to the input file containing the graph.",
    )
    return parser.parse_args()


def main():
    args = parse_args()
    input_file = args.input_file

    try:
        with open(input_file, "r") as file:
            edges = file.readlines()
    except FileNotFoundError:
        print(f"Error: File '{input_file}' not found.")
        sys.exit(1)

    if not edges:
        print("Error: The input file is empty.")
        sys.exit(1)

    edges_parsed = []
    max_vertex = 0

    for line in edges:
        try:
            from_vertex, to_vertex = map(int, line.strip().split())
            max_vertex = max(max_vertex, from_vertex, to_vertex)
            edges_parsed.append((from_vertex, to_vertex))
        except ValueError:
            print(f"Warning: Skipping malformed line: {line.strip()}")

    # Initialize the graph with the maximum vertex index.
    graph = Graph(max_vertex)

    # Add edges to the graph.
    for from_vertex, to_vertex in edges_parsed:
        graph.add_edge(from_vertex, to_vertex)

    # Run Kosaraju's algorithm.
    scc_sizes = graph.kosaraju()

    # Output the sizes of the 5 largest SCCs.
    result = [str(scc_sizes[i]) if i < len(scc_sizes) else "0" for i in range(5)]
    print(",".join(result))


if __name__ == "__main__":
    main()
