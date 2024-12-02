class Item:
    """
    Represents an item in the knapsack problem.
    Each item has a value and a weight.
    """
    def __init__(self, value: int, weight: int):
        if value < 0 or weight < 0:
            raise ValueError("Value and weight must be non-negative integers.")
        self.value = value  # Value of the item.
        self.weight = weight  # Weight of the item.

    def value_per_weight(self) -> float:
        """
        Calculates and returns the value-to-weight ratio for the item.
        If the weight is zero, prioritize this item by returning a very high ratio.
        """
        if self.weight == 0:
            return float('inf')  # Zero-weight items are given priority.
        return self.value / self.weight


class Knapsack:
    """
    Represents the knapsack used to store selected items.
    Tracks the capacity of the knapsack and its current usage.
    """
    def __init__(self, capacity: int):
        if capacity <= 0:
            raise ValueError("Knapsack capacity must be a positive integer.")
        self.capacity = capacity  # Total capacity of the knapsack.
        self.items = []  # List of items included in the knapsack.

    def add_item(self, item: Item) -> bool:
        """
        Adds an item to the knapsack if there is enough remaining capacity.
        Returns True if the item was successfully added, False otherwise.
        """
        if self.current_weight() + item.weight <= self.capacity:
            self.items.append(item)  # Add the item to the knapsack.
            return True
        return False

    def current_weight(self) -> int:
        """
        Calculates the current total weight of the items in the knapsack.
        """
        return sum(item.weight for item in self.items)

    def total_value(self) -> int:
        """
        Calculates the total value of the items currently in the knapsack.
        """
        return sum(item.value for item in self.items)


class GreedyKnapsackSolver:
    """
    Solves the knapsack problem using a greedy heuristic.
    Selects items based on their value-to-weight ratio.
    """
    def __init__(self, capacity: int, items: list[Item]):
        if not items:
            raise ValueError("Item list must not be empty.")
        self.knapsack = Knapsack(capacity)  # Initialize the knapsack.
        self.items = items  # List of available items.

    def solve(self, criterion: str = "value_per_weight"):
        """
        Executes the greedy heuristic to solve the knapsack problem.
        Sorts items based on the specified criterion (default: value-to-weight ratio).
        """
        # Validate the sorting criterion.
        if criterion not in {"value_per_weight", "value", "weight"}:
            raise ValueError("Invalid sorting criterion. Must be 'value_per_weight', 'value', or 'weight'.")

        # Sort items by the specified criterion.
        if criterion == "value_per_weight":
            self.items.sort(key=lambda item: item.value_per_weight(), reverse=True)
        elif criterion == "value":
            self.items.sort(key=lambda item: item.value, reverse=True)
        elif criterion == "weight":
            self.items.sort(key=lambda item: item.weight)

        # Add items to the knapsack while respecting capacity constraints.
        for item in self.items:
            self.knapsack.add_item(item)

    def get_result(self) -> dict:
        """
        Returns the result of the heuristic as a dictionary.
        Includes the total value, the total weight, and the list of selected items.
        """
        return {
            "total_value": self.knapsack.total_value(),  # Total value of the selected items.
            "total_weight": self.knapsack.current_weight(),  # Total weight of the knapsack.
            "selected_items": [(item.value, item.weight) for item in self.knapsack.items],  # List of selected items.
        }


# Example usage:
if __name__ == "__main__":
    # Define the items (value, weight).
    items = [
        Item(value=60, weight=10),
        Item(value=100, weight=20),
        Item(value=120, weight=30),
        Item(value=40, weight=0),
    ]

    # Define the capacity of the knapsack.
    capacity = 50

    # Create a greedy knapsack solver.
    solver = GreedyKnapsackSolver(capacity, items)

    # Solve the problem.
    solver.solve()  # Use default sorting by value_per_weight.

    # Get and print the results.
    result = solver.get_result()
    print("Total Value:", result["total_value"])
    print("Total Weight:", result["total_weight"])
    print("Selected Items (Value, Weight):", result["selected_items"])
