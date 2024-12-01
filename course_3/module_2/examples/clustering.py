import heapq


class UnionFind:
    """Data structure to manage clusters using union-find operations."""
    def __init__(self, n):
        """
        Initialize the Union-Find structure.

        :param n: The number of elements to manage.
        """
        # Each element starts as its own parent, indicating separate clusters.
        self.parent = list(range(n))
        # Rank is used to keep the tree flat during union operations.
        self.rank = [0] * n

    def find(self, x):
        """
        Find the root of the cluster for the given element.

        :param x: The element to find the root for.
        :return: The root of the cluster.
        """
        # Path compression: make each node point directly to the root.
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x, y):
        """
        Merge two clusters by connecting their roots.

        :param x: The first element.
        :param y: The second element.
        """
        # Find the roots of both clusters.
        root_x = self.find(x)
        root_y = self.find(y)

        # If roots are different, merge them.
        if root_x != root_y:
            # Attach the smaller tree under the larger tree to keep it balanced.
            if self.rank[root_x] > self.rank[root_y]:
                self.parent[root_y] = root_x
            elif self.rank[root_x] < self.rank[root_y]:
                self.parent[root_x] = root_y
            else:
                # If ranks are equal, choose one root and increase its rank.
                self.parent[root_y] = root_x
                self.rank[root_x] += 1


class ClusteringAlgorithm:
    """Greedy clustering algorithm using union-find."""
    def __init__(self, points, distance_matrix, num_clusters):
        """
        Initialize the clustering algorithm.

        :param points: List of points to cluster.
        :param distance_matrix: Matrix of distances between points.
        :param num_clusters: Desired number of clusters.
        """
        self.points = points
        self.distance_matrix = distance_matrix
        self.num_clusters = num_clusters
        # Union-Find structure to manage clusters.
        self.uf = UnionFind(len(points))
        # Priority queue (heap) to store edges sorted by distance.
        self.heap = self._build_heap()

    def _build_heap(self):
        """
        Build a heap to efficiently access the smallest distances.

        :return: A heap containing all pairwise distances.
        """
        heap = []
        num_points = len(self.points)
        # Iterate over all pairs of points (i, j) and store their distances in the heap.
        for i in range(num_points):
            for j in range(i + 1, num_points):
                distance = self.distance_matrix[i][j]
                heapq.heappush(heap, (distance, i, j))  # Push (distance, point1, point2).
        return heap

    def cluster(self):
        """
        Perform clustering until the desired number of clusters is reached.

        :return: A list of clusters, where each cluster is a list of points.
        """
        current_clusters = len(self.points)  # Start with each point as its own cluster.

        # Merge clusters until the desired number of clusters is reached.
        while current_clusters > self.num_clusters:
            # Extract the smallest distance from the heap.
            distance, i, j = heapq.heappop(self.heap)

            # If the points are in different clusters, merge them.
            if self.uf.find(i) != self.uf.find(j):
                self.uf.union(i, j)
                current_clusters -= 1

        # Build the final clusters from the Union-Find structure.
        clusters = {}
        for idx, point in enumerate(self.points):
            root = self.uf.find(idx)  # Find the root of the cluster.
            if root not in clusters:
                clusters[root] = []
            clusters[root].append(point)

        # Return the clusters as a list of lists.
        return list(clusters.values())

    def get_spacing(self):
        """
        Calculate the smallest distance between points in different clusters.

        :return: The spacing value (minimum inter-cluster distance).
        """
        min_distance = float('inf')
        num_points = len(self.points)

        # Iterate over all pairs of points to find the minimum distance between clusters.
        for i in range(num_points):
            for j in range(num_points):
                # Only consider points in different clusters.
                if self.uf.find(i) != self.uf.find(j):
                    min_distance = min(min_distance, self.distance_matrix[i][j])

        return min_distance


# Example usage:
if __name__ == "__main__":
    # Define the points and their pairwise distances.
    points = ["A", "B", "C", "D", "E"]
    distance_matrix = [
        [0, 2, 6, 10, 9],
        [2, 0, 4, 8, 7],
        [6, 4, 0, 3, 5],
        [10, 8, 3, 0, 6],
        [9, 7, 5, 6, 0]
    ]
    num_clusters = 2  # Desired number of clusters.

    # Perform clustering.
    clustering = ClusteringAlgorithm(points, distance_matrix, num_clusters)
    clusters = clustering.cluster()

    # Print the resulting clusters.
    print("Clustering result:")
    for i, cluster in enumerate(clusters):
        print(f"Cluster {i + 1}: {cluster}")

    # Print the spacing between clusters.
    print("\nSpacing:", clustering.get_spacing())
