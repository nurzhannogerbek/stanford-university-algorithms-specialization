import random
from typing import List, Tuple, Set


class Graph:
    def __init__(self, vertices: List[str], edges: List[Tuple[str, str]]):
        """
        Creates a graph with the given vertices and edges.
        :param vertices: List of graph vertices.
        :param edges: List of edges as tuples (u, v).
        """
        self.vertices: List[str] = vertices
        self.edges: List[Tuple[str, str]] = []
        self._validate_and_add_edges(edges)

    def _validate_and_add_edges(self, edges: List[Tuple[str, str]]):
        """
        Validates edges to ensure they reference existing vertices and have no self-loops or duplicates.
        :param edges: List of edges to validate and add.
        """
        vertex_set = set(self.vertices)  # Set of all valid vertices for quick lookup.
        for u, v in edges:
            if u not in vertex_set or v not in vertex_set:
                raise ValueError(f"Edge ({u}, {v}) contains vertices not in the vertex set.")  # Check for invalid vertices.
            if u == v:
                raise ValueError(f"Self-loop detected on vertex ({u}).")  # Ensure there are no self-loops.
            self.edges.append((u, v))  # Add valid edge to the list.

    def __repr__(self):
        """
        Returns a string representation of the graph, displaying the number of vertices and edges.
        """
        return f"Graph(vertices={len(self.vertices)}, edges={len(self.edges)})"


class MaximumCut:
    def __init__(self, graph: Graph):
        """
        Initializes the Maximum Cut problem solver.
        :param graph: An instance of the Graph class.
        """
        self.graph = graph
        self.partition_a: Set[str] = set()  # Set to store vertices in partition A.
        self.partition_b: Set[str] = set()  # Set to store vertices in partition B.

    def initialize_partition(self):
        """
        Creates an initial random partition of vertices.
        Ensures roughly equal-sized groups for better performance.
        """
        random.seed(42)  # Fix the random seed for consistency.
        shuffled_vertices = sorted(self.graph.vertices)  # Sort vertices for consistent processing order.
        random.shuffle(shuffled_vertices)  # Shuffle vertices randomly.
        mid = len(shuffled_vertices) // 2
        self.partition_a = set(shuffled_vertices[:mid])  # Assign the first half to partition A.
        self.partition_b = set(shuffled_vertices[mid:])  # Assign the second half to partition B.

    def compute_cut_value(self) -> int:
        """
        Calculates the current cut value, which is the number of crossing edges.
        A crossing edge is an edge where its endpoints belong to different partitions.
        :return: The cut value.
        """
        return sum(
            1 for u, v in self.graph.edges
            if (u in self.partition_a and v in self.partition_b) or
               (u in self.partition_b and v in self.partition_a)  # Check if the edge crosses the partition.
        )

    def improve_partition(self) -> bool:
        """
        Attempts to improve the current partition by moving vertices between groups.
        A vertex is moved if it increases the number of crossing edges.
        :return: True if any improvement was made, False otherwise.
        """
        improved = False  # Flag to indicate if any improvement was made.
        for vertex in sorted(self.partition_a | self.partition_b):  # Consistent processing order.
            # Determine the current and the other group for the vertex.
            current_group = self.partition_a if vertex in self.partition_a else self.partition_b
            other_group = self.partition_b if vertex in self.partition_a else self.partition_a

            # Calculate the number of crossing and non-crossing edges for this vertex.
            crossing_edges = sum(
                1 for u, v in self.graph.edges
                if (u == vertex and v in other_group) or (v == vertex and u in other_group)
            )
            non_crossing_edges = sum(
                1 for u, v in self.graph.edges
                if (u == vertex and v in current_group) or (v == vertex and u in current_group)
            )

            # Move the vertex to the other group if it improves the cut value.
            if non_crossing_edges > crossing_edges:
                current_group.remove(vertex)
                other_group.add(vertex)
                improved = True
        return improved

    def solve(self) -> int:
        """
        Solves the Maximum Cut problem using a local search algorithm.
        :return: The final cut value.
        """
        self.initialize_partition()  # Start with a random initial partition.
        while self.improve_partition():  # Continue improving the partition until no further improvement is possible.
            pass
        return self.compute_cut_value()  # Return the cut value of the final partition.

    def get_partitions(self) -> Tuple[Set[str], Set[str]]:
        """
        Returns the current partition of vertices.
        :return: A tuple of sets representing the two partitions (A and B).
        """
        return self.partition_a, self.partition_b


if __name__ == "__main__":
    # Example graph with vertices and edges.
    vertices = ['A', 'B', 'C', 'D']
    edges = [('A', 'B'), ('B', 'C'), ('C', 'D'), ('D', 'A'), ('A', 'C')]

    # Create the graph and solve the Maximum Cut problem.
    graph = Graph(vertices, edges)
    max_cut_solver = MaximumCut(graph)
    max_cut_value = max_cut_solver.solve()

    # Get and display the final partitions and cut value.
    partition_a, partition_b = max_cut_solver.get_partitions()

    # Print the results.
    print(f"Graph: {graph}")
    print(f"Maximum cut value: {max_cut_value}.")
    print(f"Group A: {partition_a}.")
    print(f"Group B: {partition_b}.")
