class Edge:
    def __init__(self, u, v, weight):
        self.u = u  # The starting vertex of the edge.
        self.v = v  # The ending vertex of the edge.
        self.weight = weight  # The weight of the edge.

    def __repr__(self):
        return f"Edge({self.u}, {self.v}, {self.weight})"


class Graph:
    def __init__(self, vertices):
        self.vertices = vertices  # The number of vertices in the graph.
        self.edges = []  # The list of edges in the graph.

    def add_edge(self, u, v, weight):
        # Add an edge between vertex u and vertex v with a given weight.
        if u < 0 or v < 0 or u >= self.vertices or v >= self.vertices:
            raise ValueError(f"Vertices {u} and {v} must be between 0 and {self.vertices - 1}.")
        self.edges.append(Edge(u, v, weight))

    def kruskal_mst(self):
        # Sort edges by weight.
        sorted_edges = sorted(self.edges, key=lambda edge: edge.weight)

        # Create a Union-Find instance to manage connected components.
        uf = UnionFind(self.vertices)

        mst = []  # The edges in the minimum spanning tree.
        total_weight = 0  # The total weight of the MST.

        for edge in sorted_edges:
            # Check if the edge creates a cycle.
            if uf.find(edge.u) != uf.find(edge.v):
                uf.union(edge.u, edge.v)
                mst.append(edge)
                total_weight += edge.weight

        return mst, total_weight


class UnionFind:
    def __init__(self, size):
        self.parent = list(range(size))  # Parent pointers for each element.
        self.rank = [0] * size  # The rank (tree depth) for each element.

    def find(self, x):
        # Find the root of the set containing x with path compression.
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])  # Path compression.
        return self.parent[x]

    def union(self, x, y):
        # Union two sets by rank.
        root_x = self.find(x)
        root_y = self.find(y)

        if root_x != root_y:
            if self.rank[root_x] > self.rank[root_y]:
                self.parent[root_y] = root_x
            elif self.rank[root_x] < self.rank[root_y]:
                self.parent[root_x] = root_y
            else:
                self.parent[root_y] = root_x
                self.rank[root_x] += 1


# Example usage.
if __name__ == "__main__":
    # Create a graph with 5 vertices.
    graph = Graph(5)

    # Add edges (u, v, weight).
    graph.add_edge(0, 1, 1)
    graph.add_edge(0, 2, 3)
    graph.add_edge(1, 2, 2)
    graph.add_edge(1, 3, 5)
    graph.add_edge(2, 3, 4)
    graph.add_edge(3, 4, 6)

    # Find the minimum spanning tree.
    mst, total_weight = graph.kruskal_mst()

    print("Edges in the MST:")
    for edge in mst:
        print(edge)
    print(f"Total weight of the MST: {total_weight}")
