class UnionFind:
    def __init__(self, size: int):
        """
        Initialize the Union-Find structure with Lazy Unions, Union by Rank, and Path Compression.
        :param size: Number of elements.
        """
        if size <= 0:
            raise ValueError("Size must be greater than 0.")
        self.parent = list(range(size))  # Parent pointers.
        self.rank = [0] * size           # Rank of each set.

    def find(self, x: int) -> int:
        """
        Find the root of the element x with path compression.
        :param x: Element to find.
        :return: Root of the element.
        """
        if x < 0 or x >= len(self.parent):
            raise ValueError(f"Index {x} is out of bounds.")
        if self.parent[x] != x:
            # Path compression: update parent to the root.
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x: int, y: int):
        """
        Union by rank of the sets containing x and y.
        :param x: First element.
        :param y: Second element.
        """
        root_x = self.find(x)
        root_y = self.find(y)

        if root_x != root_y:
            # Union by rank: attach smaller tree under larger tree.
            if self.rank[root_x] < self.rank[root_y]:
                self.parent[root_x] = root_y
            elif self.rank[root_x] > self.rank[root_y]:
                self.parent[root_y] = root_x
            else:
                # If ranks are equal, choose one as root and increment its rank.
                self.parent[root_y] = root_x
                self.rank[root_x] += 1

    def connected(self, x: int, y: int) -> bool:
        """
        Check if two elements are in the same set.
        :param x: First element.
        :param y: Second element.
        :return: True if they are in the same set, False otherwise.
        """
        return self.find(x) == self.find(y)

    def components(self) -> dict:
        """
        Get all connected components as a dictionary.
        :return: Dictionary of root to list of elements.
        """
        groups = {}
        for i in range(len(self.parent)):
            root = self.find(i)
            if root not in groups:
                groups[root] = []
            groups[root].append(i)
        return groups


# Example usage:
if __name__ == "__main__":
    uf = UnionFind(10)

    # Perform union operations.
    uf.union(1, 2)
    uf.union(3, 4)
    uf.union(2, 4)

    # Check connectivity.
    print(uf.connected(1, 3))
    print(uf.connected(1, 5))

    # Print all components in the desired format.
    print("Components:")
    components = uf.components()
    for root, group in components.items():
        print(f"Root {root}: {group}")
