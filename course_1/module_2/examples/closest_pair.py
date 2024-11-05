from math import sqrt
from typing import List, Tuple
import bisect


class ClosestPair:
    def __init__(self, points: List[Tuple[float, float]]):
        """
        Initializes the object with a given list of points.
        """
        self.points = points

    @staticmethod
    def euclidean_distance(p1: Tuple[float, float], p2: Tuple[float, float]) -> float:
        """
        Calculates the Euclidean distance between two points.
        """
        return sqrt((p1[0] - p2[0]) ** 2 + (p1[1] - p2[1]) ** 2)

    def brute_force(self, points: List[Tuple[float, float]]) -> Tuple[
        float, Tuple[Tuple[float, float], Tuple[float, float]]]:
        """
        Solves the closest pair problem using brute-force for a small set of points.
        """
        min_dist = float('inf')
        closest_pair = None
        n = len(points)

        for i in range(n):
            for j in range(i + 1, n):
                dist = self.euclidean_distance(points[i], points[j])
                if dist < min_dist:
                    min_dist = dist
                    closest_pair = (points[i], points[j])

        return min_dist, closest_pair

    def closest_split_pair(self, px: List[Tuple[float, float]], py: List[Tuple[float, float]], delta: float,
                           best_pair: Tuple[Tuple[float, float], Tuple[float, float]]) -> Tuple[
        float, Tuple[Tuple[float, float], Tuple[float, float]]]:
        """
        Handles the case where the closest pair of points is split between two halves.
        """
        mid_x = px[len(px) // 2][0]
        # Use binary search to select points within the strip.
        left = bisect.bisect_left(py, (mid_x - delta, float('-inf')))
        right = bisect.bisect_right(py, (mid_x + delta, float('inf')))
        sy = py[left:right]

        best_dist = delta
        len_sy = len(sy)

        # Check up to the next 7 points by y-coordinate.
        for i in range(len_sy):
            for j in range(1, min(8, len_sy - i)):
                p, q = sy[i], sy[i + j]
                dist = self.euclidean_distance(p, q)
                if dist < best_dist:
                    best_dist = dist
                    best_pair = (p, q)

        return best_dist, best_pair

    def closest_pair_recursive(self, px: List[Tuple[float, float]], py: List[Tuple[float, float]], left: int,
                               right: int) -> Tuple[float, Tuple[Tuple[float, float], Tuple[float, float]]]:
        """
        Recursive divide-and-conquer method to find the closest pair of points.
        """
        if right - left <= 3:
            # Use brute-force if the number of points is small.
            return self.brute_force(px[left:right])

        mid = (left + right) // 2
        midpoint_x = px[mid][0]

        # Divide py into Qy and Ry in one pass.
        Qy = []
        Ry = []
        for point in py:
            if point[0] <= midpoint_x:
                Qy.append(point)
            else:
                Ry.append(point)

        # Recursive calls for the left and right halves.
        dist_left, pair_left = self.closest_pair_recursive(px, Qy, left, mid)
        dist_right, pair_right = self.closest_pair_recursive(px, Ry, mid, right)

        # Choose the minimum distance from the two halves.
        if dist_left < dist_right:
            delta = dist_left
            best_pair = pair_left
        else:
            delta = dist_right
            best_pair = pair_right

        # Check the split pair case.
        dist_split, pair_split = self.closest_split_pair(px, py, delta, best_pair)

        # Return the smallest of the distances found.
        if delta < dist_split:
            return delta, best_pair
        else:
            return dist_split, pair_split

    def find_closest_pair(self) -> Tuple[float, Tuple[Tuple[float, float], Tuple[float, float]]]:
        """
        Finds the closest pair of points in the original set.
        """
        # Check if there are enough points.
        if len(self.points) < 2:
            raise ValueError("At least two points are required to find the closest pair.")

        # Sort points by x and y coordinates.
        px = sorted(self.points, key=lambda p: p[0])
        py = sorted(self.points, key=lambda p: p[1])

        # Start the recursive algorithm.
        return self.closest_pair_recursive(px, py, 0, len(px))


# Example usage and edge case testing.
if __name__ == "__main__":
    sample_points = [(2.1, 3.5), (1.1, 2.9), (3.6, 4.7), (4.2, 2.1), (0.9, 1.5)]
    closest_pair_solver = ClosestPair(sample_points)
    distance, pair = closest_pair_solver.find_closest_pair()
    print(f"The closest pair is {pair} with a distance of {distance:.2f}")

    # Edge case: fewer than two points.
    try:
        single_point = [(2.1, 3.5)]
        closest_pair_solver = ClosestPair(single_point)
        distance, pair = closest_pair_solver.find_closest_pair()
    except ValueError as e:
        print(e)

    # Edge case: large number of points (performance testing).
    import random

    large_points = [(random.uniform(-1000, 1000), random.uniform(-1000, 1000)) for _ in range(10000)]
    closest_pair_solver = ClosestPair(large_points)
    distance, pair = closest_pair_solver.find_closest_pair()
    print(f"Closest pair in large dataset: {pair} with a distance of {distance:.2f}")
