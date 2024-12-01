class Item:
    """
    Represents an item with a value and a weight.
    """
    def __init__(self, value, weight):
        self.value = value
        self.weight = weight


class Knapsack:
    """
    Represents a knapsack and provides a method to solve the Knapsack Problem
    using dynamic programming.
    """
    def __init__(self, capacity):
        self.capacity = capacity  # Maximum weight the knapsack can hold
        self.items = []          # List of items that can be placed in the knapsack

    def add_item(self, value, weight):
        """
        Adds an item to the knapsack's list of items.
        """
        self.items.append(Item(value, weight))

    def solve(self):
        """
        Solves the Knapsack Problem using dynamic programming.
        Returns the maximum value and the items included in the optimal solution.
        """
        n = len(self.items)  # Number of items.
        W = self.capacity    # Knapsack capacity.

        # Create a 2D array for dynamic programming.
        # dp[i][w] represents the maximum value obtainable with the first i items
        # and a knapsack capacity of w.
        dp = [[0] * (W + 1) for _ in range(n + 1)]

        # Fill the dp table.
        for i in range(1, n + 1):
            for w in range(W + 1):
                if self.items[i - 1].weight <= w:
                    # If the current item's weight is less than or equal to the current capacity,
                    # consider the maximum value by including or excluding the item.
                    dp[i][w] = max(
                        dp[i - 1][w],  # Exclude the current item
                        dp[i - 1][w - self.items[i - 1].weight] + self.items[i - 1].value  # Include the current item.
                    )
                else:
                    # If the current item's weight exceeds the current capacity, exclude it.
                    dp[i][w] = dp[i - 1][w]

        # Backtracking to find the items included in the optimal solution.
        w = W
        included_items = []
        for i in range(n, 0, -1):
            if dp[i][w] != dp[i - 1][w]:  # If the value differs, the item was included.
                included_items.append(self.items[i - 1])
                w -= self.items[i - 1].weight  # Reduce the remaining capacity.

        # Return the maximum value and the list of included items.
        return dp[n][W], included_items


# Example usage:
if __name__ == "__main__":
    # Create a knapsack with a capacity of 6.
    knapsack = Knapsack(capacity=6)

    # Add items (value, weight).
    knapsack.add_item(value=3, weight=4)
    knapsack.add_item(value=2, weight=3)
    knapsack.add_item(value=4, weight=2)
    knapsack.add_item(value=4, weight=3)

    # Solve the problem.
    max_value, included_items = knapsack.solve()

    # Output the results.
    print(f"Maximum value: {max_value}")
    print("Items included in the optimal solution:")
    for item in included_items:
        print(f"Value: {item.value}, Weight: {item.weight}")
