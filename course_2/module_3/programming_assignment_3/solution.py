import heapq

class MaxHeap:
    """Wrapper for a max-heap using heapq."""
    def __init__(self):
        self.heap = []

    def push(self, value):
        heapq.heappush(self.heap, -value)

    def pop(self):
        return -heapq.heappop(self.heap)

    def top(self):
        return -self.heap[0] if self.heap else None

    def __len__(self):
        return len(self.heap)

class MedianMaintenance:
    """Maintains a running median using two heaps."""
    def __init__(self):
        self.min_heap = []  # Min-heap for the larger half.
        self.max_heap = MaxHeap()  # Max-heap for the smaller half.

    def insert(self, num):
        """Insert a number and rebalance heaps."""
        if not self.max_heap.heap or num <= self.max_heap.top():
            self.max_heap.push(num)
        else:
            heapq.heappush(self.min_heap, num)

        # Balance the sizes of the heaps.
        if len(self.max_heap) > len(self.min_heap) + 1:
            heapq.heappush(self.min_heap, self.max_heap.pop())
        elif len(self.min_heap) > len(self.max_heap):
            self.max_heap.push(heapq.heappop(self.min_heap))

    def find_median(self):
        """Find the current median."""
        return self.max_heap.top()

def read_numbers_from_file(file_path):
    """Read numbers from a file, yielding one at a time."""
    try:
        with open(file_path, 'r') as file:
            for line in file:
                yield int(line.strip())
    except FileNotFoundError:
        raise FileNotFoundError(f"The file '{file_path}' does not exist.")
    except ValueError:
        raise ValueError("The file contains non-integer values.")

def calculate_median_sum(input_numbers):
    """Calculate the sum of medians modulo 10000."""
    median_maintenance = MedianMaintenance()
    median_sum = 0

    for number in input_numbers:
        median_maintenance.insert(number)
        median_sum += median_maintenance.find_median()
        median_sum %= 10000  # Keep only the last 4 digits.

    return median_sum

# File path for the input file.
file_path = 'Median.txt'
numbers = read_numbers_from_file(file_path)

# Calculate and print the result.
result = calculate_median_sum(numbers)
print(f"The sum of the medians modulo 10000 is: {result}")
