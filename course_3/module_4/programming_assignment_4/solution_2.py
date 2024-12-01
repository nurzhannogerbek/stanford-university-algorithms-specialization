class Item:
    """Represents an item with a value and weight."""
    def __init__(self, value: int, weight: int):
        self.value = value
        self.weight = weight


class Knapsack:
    """Represents a knapsack with a given capacity and a list of items."""
    def __init__(self, capacity: int, items: list[Item]):
        self.capacity = capacity
        self.items = items

    def solve(self) -> int:
        """Solves the knapsack problem using optimizations."""
        dp = [0] * (self.capacity + 1)

        # Sort items by value-to-weight ratio.
        self.items.sort(key=lambda x: x.value / x.weight, reverse=True)

        # Update dp array for each item.
        for item in self.items:
            for w in range(self.capacity, item.weight - 1, -1):
                dp[w] = max(dp[w], dp[w - item.weight] + item.value)

        return dp[self.capacity]


def filter_items(items: list[Item], capacity: int) -> list[Item]:
    """Filters out irrelevant items."""
    return [item for item in items if item.weight <= capacity]


def read_input_file(file_path: str) -> Knapsack:
    """Reads the input file and constructs a Knapsack object."""
    with open(file_path, 'r') as file:
        lines = file.readlines()

    first_line = lines[0].split()
    capacity = int(first_line[0])

    items = []
    for line in lines[1:]:
        value, weight = map(int, line.split())
        items.append(Item(value, weight))

    # Filter items and return Knapsack instance.
    items = filter_items(items, capacity)
    return Knapsack(capacity, items)


if __name__ == "__main__":
    file_path = "knapsack_big.txt"
    knapsack = read_input_file(file_path)
    optimal_value = knapsack.solve()
    print(f"The optimal value is: {optimal_value}.")
