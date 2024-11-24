from collections import deque

class Graph:
    def __init__(self):
        # Initialize the graph as an adjacency list.
        self.adjacency_list = {}

    def add_edge(self, u, v):
        """
        Add an undirected edge between nodes u and v.
        If the nodes do not exist in the graph, initialize them.
        """
        self.adjacency_list.setdefault(u, []).append(v)
        self.adjacency_list.setdefault(v, []).append(u)

    def bfs(self, start):
        """
        Perform Breadth-First Search (BFS) starting from the node 'start'.
        Returns a dictionary with the shortest distance from 'start' to all reachable nodes.

        Args:
            start: The starting node for BFS.

        Returns:
            distances: A dictionary where the keys are nodes and values are their distances from 'start'.
        """
        if start not in self.adjacency_list:
            raise ValueError("Start node not found in the graph.")  # Handle invalid starting node.

        explored = {start}  # Set to track explored nodes.
        queue = deque([start])  # Queue for BFS.
        distances = {start: 0}  # Dictionary to store distances from 'start'.

        while queue:
            current = queue.popleft()  # Dequeue the next node.
            for neighbor in self.adjacency_list.get(current, []):  # Explore neighbors.
                if neighbor not in explored:  # If the neighbor hasn't been visited.
                    explored.add(neighbor)  # Mark it as visited.
                    queue.append(neighbor)  # Enqueue the neighbor.
                    distances[neighbor] = distances[current] + 1  # Update its distance.

        return distances

    def connected_components(self):
        """
        Identify all connected components in the graph.
        Returns a list of connected components, where each component is a list of nodes.

        Returns:
            components: A list of lists, where each inner list represents a connected component.
        """
        explored = set()  # Set to track explored nodes.
        components = []  # List to store all connected components.

        for node in self.adjacency_list:  # Iterate through all nodes in the graph.
            if node not in explored:  # Start a new BFS if the node is unexplored.
                component = []  # List to store the current connected component.
                queue = deque([node])  # Queue for BFS.

                while queue:
                    current = queue.popleft()  # Dequeue the next node.
                    if current not in explored:  # If the node hasn't been visited.
                        explored.add(current)  # Mark it as visited.
                        component.append(current)  # Add it to the current component.
                        # Enqueue all unexplored neighbors.
                        queue.extend(
                            neighbor for neighbor in self.adjacency_list.get(current, [])
                            if neighbor not in explored
                        )

                components.append(component)  # Add the current component to the list.

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

    print("Breadth-First Search (BFS) Distances:")
    print(g.bfs(1))  # Perform BFS starting from node 1.

    print("\nConnected Components:")
    print(g.connected_components())  # Find all connected components in the graph.
