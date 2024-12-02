class VertexCover:
    """
    Class to represent the Vertex Cover problem and implement the recursive algorithm.
    """

    def __init__(self, graph):
        """
        Initialize the VertexCover class with an undirected graph.
        :param graph: Dictionary representing an undirected graph. Keys are vertices, and values are lists of adjacent vertices.
        """
        self.graph = graph
        self.memo = {}  # Memoization to avoid redundant computations.

    def vertex_cover(self, graph, k):
        """
        Determine if there exists a vertex cover of size k or smaller.
        :param graph: Dictionary representing the current graph.
        :param k: Integer, the size of the desired vertex cover.
        :return: True if a vertex cover of size k or smaller exists, otherwise False.
        """
        # Check memoized results.
        graph_key = tuple(sorted((v, tuple(sorted(neighbors))) for v, neighbors in graph.items()))
        if (graph_key, k) in self.memo:
            return self.memo[(graph_key, k)]

        # Base cases for recursion.
        if k < 0:
            return False
        if not graph:
            return True
        if k == 0:
            return False

        # Pick an arbitrary edge (u, v) from the graph.
        u = next(iter(graph))
        v = graph[u][0]

        # Create two subgraphs by removing u and v along with their incident edges.
        graph_u = self._create_subgraph_excluding_vertex(graph, u)
        graph_v = self._create_subgraph_excluding_vertex(graph, v)

        # Recursively check both subgraphs for vertex covers of size k-1.
        cover_u = self.vertex_cover(graph_u, k - 1)
        cover_v = self.vertex_cover(graph_v, k - 1)

        # Memoize and return the result.
        result = cover_u or cover_v
        self.memo[(graph_key, k)] = result
        return result

    @staticmethod
    def _create_subgraph_excluding_vertex(graph, vertex):
        """
        Create a new subgraph by removing a given vertex and its incident edges.
        :param graph: The current graph.
        :param vertex: The vertex to remove.
        :return: A new graph with the vertex and its edges removed.
        """
        new_graph = {v: [w for w in neighbors if w != vertex] for v, neighbors in graph.items() if v != vertex}
        return {v: neighbors for v, neighbors in new_graph.items() if neighbors}


# Example usage:
if __name__ == "__main__":
    # Define a simple undirected graph as a dictionary.
    graph = {
        'A': ['B', 'C'],
        'B': ['A', 'D', 'E'],
        'C': ['A', 'F'],
        'D': ['B'],
        'E': ['B'],
        'F': ['C']
    }

    # Create an instance of the VertexCover class.
    vc = VertexCover(graph)

    # Check if there is a vertex cover of size 2 or smaller.
    k = 2
    result = vc.vertex_cover(graph, k)

    # Print the result.
    print(f"Is there a vertex cover of size {k} or smaller? {'Yes' if result else 'No'}.")
