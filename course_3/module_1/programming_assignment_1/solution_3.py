import heapq


class Edge:
    """
    Represents an edge in a graph with a source, target, and cost.
    """
    def __init__(self, source, target, cost):
        self.source = source
        self.target = target
        self.cost = cost

    def __lt__(self, other):
        """
        Compares two edges by their cost for use in a priority queue.
        """
        return self.cost < other.cost

    def __repr__(self):
        return f"Edge({self.source} -- {self.cost} --> {self.target})"


class Graph:
    """
    Represents an undirected weighted graph.
    """
    def __init__(self, num_nodes):
        """
        Initializes the graph with a specified number of nodes.
        """
        self.num_nodes = num_nodes
        self.adjacency_list = {i: [] for i in range(1, num_nodes + 1)}

    def add_edge(self, source, target, cost):
        """
        Adds an undirected edge to the graph.
        """
        self.adjacency_list[source].append(Edge(source, target, cost))
        self.adjacency_list[target].append(Edge(target, source, cost))


class PrimMST:
    """
    Implements Prim's algorithm to find the Minimum Spanning Tree (MST).
    """
    def __init__(self, graph):
        """
        Initializes the PrimMST with a graph.
        """
        self.graph = graph
        self.total_cost = 0

    def find_mst(self, start_node=1):
        """
        Finds the Minimum Spanning Tree (MST) and its total cost.
        :param start_node: The node to start the MST algorithm from.
        :return: Total cost of the MST as an integer.
        """
        if start_node not in self.graph.adjacency_list:
            raise ValueError(f"Start node {start_node} is not in the graph.")

        visited = set()
        min_heap = []
        self.total_cost = 0

        # Start with the given node.
        visited.add(start_node)
        for edge in self.graph.adjacency_list[start_node]:
            heapq.heappush(min_heap, edge)

        # Process the heap.
        while min_heap and len(visited) < self.graph.num_nodes:
            edge = heapq.heappop(min_heap)
            if edge.target in visited:
                continue

            # Include this edge in the MST.
            self.total_cost += edge.cost
            visited.add(edge.target)

            # Add all edges from the newly visited node.
            for next_edge in self.graph.adjacency_list[edge.target]:
                if next_edge.target not in visited:
                    heapq.heappush(min_heap, next_edge)

        if len(visited) < self.graph.num_nodes:
            raise ValueError("Graph is not connected. MST cannot be formed.")

        return self.total_cost


if __name__ == "__main__":
    try:
        # Read graph data from a file.
        with open("edges.txt", "r") as file:
            lines = file.readlines()
            num_nodes, num_edges = map(int, lines[0].strip().split())
            graph = Graph(num_nodes)

            for line in lines[1:]:
                source, target, cost = map(int, line.strip().split())
                graph.add_edge(source, target, cost)

        # Run Prim's algorithm.
        prim = PrimMST(graph)
        total_cost = prim.find_mst()

        # Print only the total cost of the MST.
        print(f"Total cost of the MST: {total_cost}")

    except FileNotFoundError:
        print("Error: File 'edges.txt' not found.")
    except ValueError as e:
        print(f"Error: {e}")