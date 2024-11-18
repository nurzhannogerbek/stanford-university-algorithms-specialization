import random
from concurrent.futures import ProcessPoolExecutor, as_completed
import os


class Graph:
    """
    Represents an undirected graph with parallel edges using an adjacency list.
    """

    def __init__(self, adjacency_list):
        """
        Initializes the graph with a shallow copy of the adjacency list.
        :param adjacency_list: A dictionary where keys are vertices and values are lists of adjacent vertices.
        """
        # Create a copy of the adjacency list to avoid modifying the input.
        self.graph = {k: v[:] for k, v in adjacency_list.items()}

    def contract_edge(self, u, v):
        """
        Contracts an edge by merging vertex v into vertex u.
        Updates the adjacency list and removes self-loops.
        :param u: One endpoint of the edge.
        :param v: The other endpoint of the edge.
        """
        # Merge all edges from vertex v into vertex u.
        self.graph[u].extend(self.graph[v])

        # Replace occurrences of v with u in the adjacency list.
        for neighbor in self.graph[v]:
            self.graph[neighbor] = [u if x == v else x for x in self.graph[neighbor]]

        # Remove self-loops from u's adjacency list.
        self.graph[u] = [x for x in self.graph[u] if x != u]

        # Delete vertex v from the graph.
        del self.graph[v]

    def get_random_edge(self):
        """
        Selects a random edge (u, v) from the graph.
        :return: A tuple (u, v) representing an edge.
        """
        # Choose a random vertex u from the graph.
        u = random.choice(list(self.graph.keys()))
        # Choose a random neighbor v of u.
        v = random.choice(self.graph[u])
        return u, v

    def get_min_cut(self):
        """
        Runs the contraction algorithm until only two vertices remain.
        :return: The number of edges between the two remaining vertices.
        """
        while len(self.graph) > 2:
            u, v = self.get_random_edge()
            self.contract_edge(u, v)

        # Get the remaining edges between the two vertices.
        for edges in self.graph.values():
            return len(edges)


def parse_adjacency_list(file_path):
    """
    Parses the adjacency list from the input file.
    :param file_path: Path to the input file containing the adjacency list.
    :return: A dictionary representing the adjacency list.
    """
    adjacency_list = {}
    with open(file_path, 'r') as file:
        for line in file:
            values = list(map(int, line.strip().split()))
            if values:
                adjacency_list[values[0]] = values[1:]
    return adjacency_list


def run_single_trial(adjacency_list, seed):
    """
    Runs a single trial of the randomized contraction algorithm.
    :param adjacency_list: The adjacency list of the graph.
    :param seed: Random seed for reproducibility in this trial.
    :return: The minimum cut for this trial.
    """
    random.seed(seed)  # Ensure reproducibility in parallel processes.
    graph = Graph(adjacency_list)
    return graph.get_min_cut()


def randomized_contraction(file_path, trials=None, workers=None):
    """
    Finds the minimum cut of the graph using the randomized contraction algorithm in parallel.
    :param file_path: Path to the file containing the adjacency list.
    :param trials: Number of trials to run. Defaults to N^2 if not provided.
    :param workers: Number of worker processes for parallel execution.
    :return: The minimum cut found.
    """
    adjacency_list = parse_adjacency_list(file_path)
    n = len(adjacency_list)

    # Set the default number of trials to N^2 if not specified.
    if trials is None:
        trials = n * n

    # Use all available CPU cores if workers are not specified.
    if workers is None:
        workers = os.cpu_count()

    min_cut = float('inf')  # Initialize with a large number.

    # Use ProcessPoolExecutor for true parallelism (multi-core execution).
    with ProcessPoolExecutor(max_workers=workers) as executor:
        # Distribute trials across workers, each with a unique random seed.
        futures = [executor.submit(run_single_trial, adjacency_list, random.randint(0, 2**32 - 1))
                   for _ in range(trials)]

        for i, future in enumerate(as_completed(futures), start=1):
            cut = future.result()
            min_cut = min(min_cut, cut)

            # Print progress every 100 completed trials.
            if i % 100 == 0:
                print(f"Progress: {i}/{trials} trials completed.")

    return min_cut


if __name__ == "__main__":
    random.seed(42)  # Set a seed for reproducibility.
    filepath = "kargerMinCut.txt"  # Update with the correct file path.

    print("Running the optimized randomized contraction algorithm with multiprocessing...")
    result = randomized_contraction(filepath, trials=20000)  # Adjust trials as needed.
    print(f"Minimum cut: {result}")
