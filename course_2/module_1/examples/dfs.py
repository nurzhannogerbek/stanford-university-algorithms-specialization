class Graph:
    def __init__(self):
        # Initialize the graph as an adjacency list.
        self.adjacency_list = {}

    def add_edge(self, u, v, directed=False):
        """
        Add an edge between nodes u and v.
        For undirected graphs, the edge is added in both directions.

        Args:
            u: The starting node of the edge.
            v: The ending node of the edge.
            directed: Boolean indicating whether the edge is directed.
        """
        self.adjacency_list.setdefault(u, []).append(v)
        if not directed:
            self.adjacency_list.setdefault(v, []).append(u)

    def dfs(self, start):
        """
        Perform Depth-First Search (DFS) starting from the node `start`.
        Uses an iterative stack-based approach.

        Args:
            start: The starting node for DFS.

        Returns:
            visited_order: A list of nodes in the order they are visited.
        """
        if not self.adjacency_list:
            raise ValueError("Graph is empty.")
        if start not in self.adjacency_list:
            raise ValueError(f"Start node {start} does not exist in the graph.")

        visited = set()  # Set to track visited nodes.
        visited_order = []  # List to store the order of visited nodes.
        stack = [start]  # Use a stack for iterative DFS.

        while stack:
            node = stack.pop()  # Take the last node added to the stack.
            if node not in visited:
                visited.add(node)  # Mark the node as visited.
                visited_order.append(node)  # Add the node to the visit order.
                # Add unvisited neighbors to the stack.
                stack.extend(neighbor for neighbor in self.adjacency_list.get(node, []) if neighbor not in visited)

        return visited_order

    def _dfs_collect(self, node, visited, component):
        """
        Helper function to collect nodes in a connected component using DFS.

        Args:
            node: The current node being explored.
            visited: Set of already visited nodes.
            component: List to store the current connected component.
        """
        visited.add(node)  # Mark the node as visited.
        component.append(node)  # Add the node to the current component.

        for neighbor in self.adjacency_list.get(node, []):  # Explore neighbors.
            if neighbor not in visited:  # If the neighbor is not visited.
                self._dfs_collect(neighbor, visited, component)  # Recursively explore the neighbor.

    def connected_components(self):
        """
        Identify all connected components in the graph.
        Returns a list of connected components, where each component is a list of nodes.

        Returns:
            components: A list of lists, where each inner list represents a connected component.
        """
        if not self.adjacency_list:
            raise ValueError("Graph is empty.")

        visited = set()  # Set to track visited nodes.
        components = []  # List to store all connected components.

        for node in self.adjacency_list:  # Iterate through all nodes in the graph.
            if node not in visited:  # Start a new DFS if the node is unexplored.
                component = []  # List to store the current connected component.
                self._dfs_collect(node, visited, component)
                components.append(component)  # Add the collected component to the list.

        return components


if __name__ == "__main__":
    # Example usage of the Graph class.
    g = Graph()
    g.add_edge(1, 2)
    g.add_edge(2, 3)
    g.add_edge(4, 5)
    g.add_edge(6, 7)
    g.add_edge(7, 8)
    g.add_edge(8, 6)

    print("Depth-First Search (DFS) Order:")
    print(g.dfs(1))  # Perform DFS starting from node 1.

    print("\nConnected Components:")
    print(g.connected_components())  # Find all connected components in the graph.
