class Graph:
    def __init__(self):
        """Initialize the graph as an adjacency list."""
        self.adjacency_list = {}

    def add_edge(self, u, v):
        """
        Add a directed edge from node u to node v.

        Args:
            u: The start node of the directed edge.
            v: The end node of the directed edge.
        """
        if u not in self.adjacency_list:
            self.adjacency_list[u] = []
        if v not in self.adjacency_list:
            self.adjacency_list[v] = []  # Ensure v exists in the graph.
        self.adjacency_list[u].append(v)

    def topological_sort(self):
        """
        Perform topological sorting of the graph using DFS.

        Returns:
            List of nodes in topological order. If the graph is cyclic, raises a ValueError.
        """
        if not self.adjacency_list:
            return []  # Return an empty list for an empty graph.

        visited = set()  # To track visited nodes.
        temp_mark = set()  # To detect cycles.
        result = []  # To store the topological order.

        def dfs(node):
            """
            Recursive helper function to perform DFS.

            Args:
                node: The current node being visited.

            Raises:
                ValueError: If a cycle is detected in the graph.
            """
            if node in temp_mark:
                raise ValueError(f"Cycle detected at node: {node}")
            if node not in visited:
                temp_mark.add(node)  # Temporarily mark the node.
                for neighbor in self.adjacency_list.get(node, []):  # Explore neighbors.
                    dfs(neighbor)
                temp_mark.remove(node)  # Remove the temporary mark.
                visited.add(node)  # Mark the node as permanently visited.
                result.append(node)  # Add to the topological order.

        # Perform DFS for all nodes in the graph.
        for node in self.adjacency_list:
            if node not in visited:
                dfs(node)

        # Reverse the result to get the correct topological order.
        return result[::-1]


# Example usage.
if __name__ == "__main__":
    g = Graph()
    g.add_edge(5, 2)
    g.add_edge(5, 0)
    g.add_edge(4, 0)
    g.add_edge(4, 1)
    g.add_edge(2, 3)
    g.add_edge(3, 1)

    try:
        print("Topological Order:", g.topological_sort())
    except ValueError as e:
        print(e)
