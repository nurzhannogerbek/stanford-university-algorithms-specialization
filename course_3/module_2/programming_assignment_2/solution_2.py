class UnionFind:
    """Union-Find structure with path compression."""
    def __init__(self, nodes):
        # Initialize Union-Find structure.
        # Each node is its own parent initially.
        self.parent = {node: node for node in nodes}
        # Initialize rank for each node to 0 for union by rank optimization.
        self.rank = {node: 0 for node in nodes}

    def find(self, node):
        """
        Find the root of the node using path compression.
        Path compression ensures that the tree structure is flattened,
        making future operations faster.
        """
        if self.parent[node] != node:
            self.parent[node] = self.find(self.parent[node])  # Path compression.
        return self.parent[node]

    def union(self, node1, node2):
        """
        Union two nodes by rank. Attach the smaller tree under the larger tree.
        If ranks are equal, increase the rank of the resulting root.
        """
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
    """Clustering algorithm using Hamming distance."""
    def __init__(self, vertices, bit_length):
        # Initialize clustering with a list of vertices and the bit length of each vertex.
        self.vertices = vertices
        self.bit_length = bit_length
        # Use Union-Find to manage clusters.
        self.uf = UnionFind(vertices)

    @staticmethod
    def invert(bit):
        """
        Invert a single bit.
        Returns '1' if the bit is '0' and '0' if the bit is '1'.
        """
        return '1' if bit == '0' else '0'

    def generate_neighbors(self, vertex):
        """
        Generate all neighbors of a vertex within Hamming distance 1 or 2.
        For each bit in the vertex:
        - Flip one bit to generate a neighbor.
        - Flip two bits to generate another neighbor.
        """
        neighbors = []
        for i in range(self.bit_length):
            # Flip one bit.
            flipped = vertex[:i] + self.invert(vertex[i]) + vertex[i+1:]
            neighbors.append(flipped)
            for j in range(i + 1, self.bit_length):
                # Flip two bits.
                flipped_two = flipped[:j] + self.invert(flipped[j]) + flipped[j+1:]
                neighbors.append(flipped_two)
        return neighbors

    def cluster(self):
        """
        Perform clustering by iterating over all vertices.
        For each vertex, find its neighbors with Hamming distance <= 2.
        Merge clusters if neighbors belong to different clusters.
        """
        for vertex in self.vertices:
            vertex_root = self.uf.find(vertex)  # Find the root of the current vertex.
            for neighbor in self.generate_neighbors(vertex):
                if neighbor in self.uf.parent:  # Check if the neighbor exists in the graph.
                    neighbor_root = self.uf.find(neighbor)
                    if vertex_root != neighbor_root:
                        self.uf.union(vertex_root, neighbor_root)

    def count_clusters(self):
        """
        Count the number of unique clusters.
        A cluster is identified by its root in the Union-Find structure.
        """
        roots = {self.uf.find(vertex) for vertex in self.vertices}
        return len(roots)


def main(file_path):
    """Main function to load data, perform clustering, and output the result."""
    # Load data from the file.
    with open(file_path, "r") as file:
        lines = file.read().splitlines()

    # The first line contains metadata: number of nodes and bit length.
    num_nodes, bit_length = map(int, lines[0].split())
    # Convert each subsequent line into a cleaned bit string.
    vertices = ["".join(line.split()) for line in lines[1:]]

    # Initialize clustering and perform the clustering process.
    clustering = Clustering(vertices, bit_length)
    clustering.cluster()

    # Count the number of clusters and print the result.
    num_clusters = clustering.count_clusters()
    print(f"The largest value of k with spacing at least 3 is: {num_clusters}")


if __name__ == "__main__":
    file_path = "clustering_big.txt"
    main(file_path)
