class UnionFind:
    def __init__(self, size):
        # Initialize the Union-Find structure with the given size.
        # Each node is initially its own parent, and the rank is set to 0.
        if size <= 0:
            raise ValueError("Size of UnionFind must be positive.")
        self.parent = list(range(size))
        self.rank = [0] * size

    def find(self, node):
        # Find the root of the node using path compression.
        # Path compression flattens the structure for faster future queries.
        if self.parent[node] != node:
            self.parent[node] = self.find(self.parent[node])  # Path compression.
        return self.parent[node]

    def union(self, node1, node2):
        # Perform the union operation by rank.
        # Connects two subsets into a single subset.
        root1 = self.find(node1)
        root2 = self.find(node2)

        if root1 != root2:
            if self.rank[root1] > self.rank[root2]:
                self.parent[root2] = root1
            elif self.rank[root1] < self.rank[root2]:
                self.parent[root1] = root2
            else:
                self.parent[root2] = root1
                self.rank[root1] += 1


class Clustering:
    def __init__(self, nodes, edges):
        # Initialize the clustering algorithm with the number of nodes and edges.
        if nodes <= 0:
            raise ValueError("Number of nodes must be positive.")
        if not edges:
            raise ValueError("Edge list cannot be empty.")
        self.nodes = nodes
        self.edges = edges

    def compute_max_spacing(self, k):
        # Compute the maximum spacing of k-clustering.
        if k > self.nodes or k <= 0:
            raise ValueError(f"Invalid number of clusters: {k}. Must be between 1 and {self.nodes}.")

        # Sort the edges by their cost in ascending order.
        self.edges.sort(key=lambda edge: edge[2])  # Edge is now a tuple (node1, node2, cost).
        uf = UnionFind(self.nodes)
        num_clusters = self.nodes

        # Process edges until the number of clusters equals k.
        for node1, node2, cost in self.edges:
            if uf.find(node1) != uf.find(node2):  # Check if nodes are in different clusters.
                if num_clusters == k:
                    # Return the cost of the next edge that connects two different clusters.
                    return cost
                uf.union(node1, node2)  # Merge the clusters.
                num_clusters -= 1

        raise RuntimeError("Unable to determine max spacing. Check input data.")


def read_input(file_path):
    """Reads the input file and parses the nodes and edges."""
    # Open the input file and read its contents line by line.
    with open(file_path, 'r') as file:
        lines = file.readlines()

    # The first line contains the number of nodes.
    num_nodes = int(lines[0].strip())
    edges = []

    # Parse the rest of the lines as edges.
    for line in lines[1:]:
        node1, node2, cost = map(int, line.strip().split())
        edges.append((node1 - 1, node2 - 1, cost))  # Convert to 0-based index.

    return num_nodes, edges


def main():
    # File path to the input data.
    file_path = "clustering1.txt"

    try:
        # Read the input data from the specified file.
        nodes, edges = read_input(file_path)
        clustering = Clustering(nodes, edges)

        # Compute the maximum spacing for 4 clusters.
        max_spacing = clustering.compute_max_spacing(k=4)

        # Print the result to the console.
        print(f"The maximum spacing of a 4-clustering is: {max_spacing}")
    except FileNotFoundError:
        # Handle the case where the input file is not found.
        print(f"Error: File '{file_path}' not found.")
    except Exception as e:
        # Handle any other unexpected errors.
        print(f"Error: {e}")


if __name__ == "__main__":
    # Entry point of the script.
    main()
