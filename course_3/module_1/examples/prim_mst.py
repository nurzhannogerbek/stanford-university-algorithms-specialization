import heapq

class Edge:
    """
    Represents an edge in a graph with a source, target, and weight.
    """
    def __init__(self, source, target, weight):
        self.source = source
        self.target = target
        self.weight = weight

    def __repr__(self):
        return f"Edge({self.source} -- {self.weight} --> {self.target})"

    # Comparison methods for heapq to compare edges by weight.
    def __lt__(self, other):
        return self.weight < other.weight

class Graph:
    """
    Represents an undirected weighted graph.
    """
    def __init__(self, num_vertices):
        """
        Initializes the graph with a given number of vertices.
        :param num_vertices: Total number of vertices in the graph.
        """
        self.num_vertices = num_vertices
        self.adjacency_list = {i: [] for i in range(num_vertices)}

    def add_edge(self, source, target, weight):
        """
        Adds an undirected edge to the graph.
        :param source: Starting vertex of the edge.
        :param target: Ending vertex of the edge.
        :param weight: Weight of the edge.
        """
        if source < 0 or source >= self.num_vertices or target < 0 or target >= self.num_vertices:
            raise ValueError("Source or target vertex out of range.")
        edge = Edge(source, target, weight)
        self.adjacency_list[source].append(edge)
        self.adjacency_list[target].append(Edge(target, source, weight))  # Add reverse edge.

    def get_edges(self, vertex):
        """
        Returns all edges connected to a specific vertex.
        :param vertex: The vertex whose edges need to be retrieved.
        :return: A list of edges connected to the vertex.
        """
        return self.adjacency_list[vertex]

class PrimMST:
    """
    Implements Prim's algorithm to find the Minimum Spanning Tree (MST).
    """
    def __init__(self, graph):
        """
        Initializes the PrimMST object with a graph.
        :param graph: An instance of the Graph class.
        """
        self.graph = graph
        self.mst_edges = []  # List to store edges included in the MST.
        self.total_cost = 0  # Total weight of the MST.
        self.visited = []  # Tracks whether a vertex is in the MST.
        self.min_heap = []  # Min-heap to store edges based on their weights.

    def push_edges(self, vertex):
        """
        Pushes all edges of a vertex into the min-heap if they lead to unvisited vertices.
        :param vertex: The vertex whose edges are pushed into the heap.
        """
        self.visited[vertex] = True
        for edge in self.graph.get_edges(vertex):
            if not self.visited[edge.target]:
                heapq.heappush(self.min_heap, edge)

    def find_mst(self, start_vertex=0):
        """
        Finds the Minimum Spanning Tree (MST) starting from a specified vertex.
        :param start_vertex: The vertex to start building the MST from (default is 0).
        :return: A tuple (mst_edges, total_cost), where mst_edges is a list of edges in the MST
                 and total_cost is the sum of their weights.
        """
        self.visited = [False] * self.graph.num_vertices  # Tracks whether a vertex is in the MST.
        self.min_heap = []  # Min-heap to store edges based on their weights.
        self.mst_edges = []
        self.total_cost = 0

        # Push all edges from the start vertex.
        self.push_edges(start_vertex)

        # Main loop: Process edges from the heap.
        while self.min_heap:
            edge = heapq.heappop(self.min_heap)
            if self.visited[edge.target]:
                continue  # Skip edges that lead to already visited vertices.

            # Add the edge to the MST.
            self.mst_edges.append(edge)
            self.total_cost += edge.weight

            # Push edges from the newly added vertex.
            self.push_edges(edge.target)

        return self.mst_edges, self.total_cost

# Example usage:
if __name__ == "__main__":
    # Create a graph with 6 vertices.
    graph = Graph(6)
    graph.add_edge(0, 1, 4)
    graph.add_edge(0, 2, 4)
    graph.add_edge(1, 2, 2)
    graph.add_edge(1, 3, 5)
    graph.add_edge(2, 3, 8)
    graph.add_edge(2, 4, 10)
    graph.add_edge(3, 4, 2)
    graph.add_edge(3, 5, 6)
    graph.add_edge(4, 5, 3)

    # Run Prim's algorithm.
    prim_mst = PrimMST(graph)
    mst_edges, total_cost = prim_mst.find_mst(start_vertex=0)

    # Print the result.
    print("Edges in the MST:")
    for edge in mst_edges:
        print(edge)

    print(f"Total cost of the MST: {total_cost}")
