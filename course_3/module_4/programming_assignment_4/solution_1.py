class Item:
    """Represents an item with a value and weight."""

    def __init__(self, value: int, weight: int):
        self.value = value
        self.weight = weight


class Knapsack:
    """Represents a knapsack with a given capacity and a list of items."""

    def __init__(self, capacity: int):
        self.capacity = capacity
        self.items = []

    def add_item(self, item: Item):
        """Adds an item to the knapsack."""
        self.items.append(item)

    def solve(self) -> int:
        """Solves the knapsack problem using dynamic programming."""
        num_items = len(self.items)
        # Create a 2D DP table initialized with zeros.
        dp = [[0] * (self.capacity + 1) for _ in range(num_items + 1)]

        # Fill the DP table.
        for i in range(1, num_items + 1):
            for w in range(self.capacity + 1):
                if self.items[i - 1].weight > w:
                    dp[i][w] = dp[i - 1][w]  # Item cannot be included.
                else:
                    # Take the maximum of including or excluding the item.
                    dp[i][w] = max(dp[i - 1][w],
                                   dp[i - 1][w - self.items[i - 1].weight] + self.items[i - 1].value)

        return dp[num_items][self.capacity]


def read_input_file(file_path: str) -> Knapsack:
    """Reads the input file and constructs a Knapsack object."""
    with open(file_path, 'r') as file:
        lines = file.readlines()

    # Read the knapsack capacity and number of items from the first line.
    first_line = lines[0].split()
    capacity = int(first_line[0])

    # Create the Knapsack object.
    knapsack = Knapsack(capacity)

    # Add items to the knapsack.
    for line in lines[1:]:
        value, weight = map(int, line.split())
        knapsack.add_item(Item(value, weight))

    return knapsack


if __name__ == "__main__":
    # File path to the input data.
    file_path = "knapsack1.txt"

    # Read the knapsack data from the file.
    knapsack = read_input_file(file_path)

    # Solve the knapsack problem and print the optimal value.
    optimal_value = knapsack.solve()
    print(f"The optimal value is: {optimal_value}.")
