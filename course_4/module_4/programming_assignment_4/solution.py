class Graph:
    """
    Represents a directed graph for the 2-SAT problem.
    Contains adjacency lists for the original and reversed graphs.
    """
    def __init__(self, num_nodes):
        self.num_nodes = num_nodes
        self.adj = {i: [] for i in range(1, num_nodes + 1)}  # Original graph.
        self.adj_inv = {i: [] for i in range(1, num_nodes + 1)}  # Reversed graph.

    def add_edge(self, u, v):
        """
        Adds a directed edge from u to v in the original graph and its reverse in the reversed graph.
        """
        self.adj[u].append(v)
        self.adj_inv[v].append(u)


class SATSolver:
    """
    Encapsulates the logic for solving the 2-SAT problem using Kosaraju's algorithm.
    """
    def __init__(self, num_variables):
        self.num_variables = num_variables
        self.graph = Graph(2 * num_variables)  # 2 * num_variables for positive and negative literals.
        self.visited = set()  # Visited nodes for the original graph.
        self.visited_inv = set()  # Visited nodes for the reversed graph.
        self.stack = []  # Stack for Kosaraju's first pass.
        self.scc = {}  # Strongly Connected Components (SCC) mapping.
        self.counter = 0  # SCC counter.

    def add_clause(self, a, b):
        """
        Adds a clause to the implication graph.
        Each clause (a OR b) is converted to two implications: NOT(a) → b and NOT(b) → a.
        """
        n = self.num_variables
        if a > 0 and b > 0:
            self.graph.add_edge(a + n, b)  # NOT(a) → b
            self.graph.add_edge(b + n, a)  # NOT(b) → a
        elif a > 0 and b < 0:
            self.graph.add_edge(a + n, -b + n)  # NOT(a) → NOT(b)
            self.graph.add_edge(-b, a)  # b → a
        elif a < 0 and b > 0:
            self.graph.add_edge(-a, b)  # a → b
            self.graph.add_edge(b + n, -a + n)  # NOT(b) → NOT(a)
        else:
            self.graph.add_edge(-a, -b + n)  # a → NOT(b)
            self.graph.add_edge(-b, -a + n)  # b → NOT(a)

    def _dfs(self, node):
        """
        Depth-first search on the original graph for Kosaraju's first pass.
        """
        if node in self.visited:
            return
        self.visited.add(node)
        for neighbor in self.graph.adj[node]:
            self._dfs(neighbor)
        self.stack.append(node)  # Push the node to the stack after visiting all its neighbors.

    def _dfs_inv(self, node):
        """
        Depth-first search on the reversed graph for Kosaraju's second pass.
        """
        if node in self.visited_inv:
            return
        self.visited_inv.add(node)
        self.scc[node] = self.counter
        for neighbor in self.graph.adj_inv[node]:
            self._dfs_inv(neighbor)

    def solve(self):
        """
        Determines if the 2-SAT problem is satisfiable using Kosaraju's algorithm.
        """
        # First pass: Perform DFS on the original graph to determine the finishing order.
        for i in range(1, 2 * self.num_variables + 1):
            if i not in self.visited:
                self._dfs(i)

        # Second pass: Perform DFS on the reversed graph to find SCCs.
        while self.stack:
            node = self.stack.pop()
            if node not in self.visited_inv:
                self.counter += 1  # Increment SCC counter.
                self._dfs_inv(node)

        # Check for contradictions: a variable and its negation must not be in the same SCC.
        for i in range(1, self.num_variables + 1):
            if self.scc.get(i) == self.scc.get(i + self.num_variables):
                return False  # Unsatisfiable.
        return True  # Satisfiable.


def parse_file(file_path):
    """
    Reads the input file and constructs a SATSolver object with the clauses from the file.
    """
    with open(file_path, 'r') as file:
        num_variables = int(file.readline().strip())  # Read the number of variables.
        solver = SATSolver(num_variables)

        for line in file:
            a, b = map(int, line.split())
            solver.add_clause(a, b)

    return solver


def main():
    """
    Main function to process all problem instances and determine their satisfiability.
    """
    base_path = ""
    file_prefix = "2sat"
    file_extension = ".txt"
    file_count = 6

    # Generate file paths.
    instances = [f"{base_path}{file_prefix}{i}{file_extension}" for i in range(1, file_count + 1)]

    result = ""

    # Process each instance and determine satisfiability.
    for file_path in instances:
        print(f"Processing file: {file_path}")
        solver = parse_file(file_path)

        if solver.solve():
            result += "1"  # Instance is satisfiable.
        else:
            result += "0"  # Instance is unsatisfiable.

    # Print the final result as a binary string.
    print("Result:", result)


if __name__ == "__main__":
    main()
