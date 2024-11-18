import random
from concurrent.futures import ThreadPoolExecutor


class Graph:
    """
    Represents an undirected graph with parallel edges.
    """

    def __init__(self, input_adjacency_list):
        """
        Initializes the graph with an adjacency list.
        :param input_adjacency_list: A dictionary where keys are vertices and values are lists of adjacent vertices.
        """
        self.graph = {k: list(v) for k, v in input_adjacency_list.items()}  # Optimized shallow copy.

    def contract_edge(self, u, v):
        """
        Contracts an edge (u, v), merging vertex v into vertex u.
        :param u: One endpoint of the edge.
        :param v: The other endpoint of the edge.
        """
        # Add all edges of v to u.
        self.graph[u].extend(self.graph[v])

        # Replace all occurrences of v with u in the graph.
        for neighbor in self.graph[v]:
            self.graph[neighbor] = [u if x == v else x for x in self.graph[neighbor]]

        # Remove self-loops.
        self.graph[u] = [x for x in self.graph[u] if x != u]

        # Delete the merged vertex.
        del self.graph[v]

    def get_random_edge(self):
        """
        Selects a random edge (u, v) from the graph.
        :return: A tuple (u, v) representing an edge.
        """
        u = random.choice(list(self.graph.keys()))
        v = random.choice(self.graph[u])
        return u, v

    def get_min_cut(self):
        """
        Runs the randomized contraction algorithm to find a minimum cut.
        :return: The number of crossing edges in the minimum cut.
        """
        while len(self.graph) > 2:
            u, v = self.get_random_edge()
            self.contract_edge(u, v)

        # After contraction, count the edges between the two remaining vertices.
        remaining_vertices = list(self.graph.keys())
        return len(self.graph[remaining_vertices[0]])


class RandomizedContractionAlgorithm:
    """
    Implements the randomized contraction algorithm for finding the minimum cut in a graph.
    """

    def __init__(self, input_adjacency_list):
        """
        Initializes the algorithm with a graph represented as an adjacency list.
        :param input_adjacency_list: A dictionary where keys are vertices and values are lists of adjacent vertices.
        """
        self.original_graph = input_adjacency_list

    def find_min_cut(self, trials=None, parallel=False):
        """
        Finds the minimum cut by running the randomized contraction algorithm multiple times.
        :param trials: Number of trials to run (optional). Defaults to N^2 if not provided.
        :param parallel: Whether to run the trials in parallel.
        :return: The minimum cut found.
        """
        n = len(self.original_graph)
        if trials is None:
            trials = n * n  # Number of trials.

        min_cut = float('inf')

        def run_trial(_):
            """
            Runs a single trial of the randomized contraction algorithm.
            :param _: Ignored argument for compatibility with map function.
            :return: Minimum cut found in this trial.
            """
            graph = Graph(self.original_graph)
            return graph.get_min_cut()

        if parallel:
            # Run trials in parallel.
            with ThreadPoolExecutor() as executor:
                results = executor.map(run_trial, range(trials))
            min_cut = min(results)
        else:
            # Run trials sequentially.
            for _ in range(trials):
                min_cut = min(min_cut, run_trial(None))

        return min_cut


# Example usage.
if __name__ == "__main__":
    # Example adjacency list.
    example_adjacency_list = {
        1: [2, 3, 4],
        2: [1, 3, 4],
        3: [1, 2, 4],
        4: [1, 2, 3]
    }

    rca = RandomizedContractionAlgorithm(example_adjacency_list)
    result = rca.find_min_cut(parallel=True)  # Set parallel=True for parallel execution.
    print(f"Minimum cut: {result}")
