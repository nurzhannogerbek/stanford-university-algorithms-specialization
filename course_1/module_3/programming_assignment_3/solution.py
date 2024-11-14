class QuickSort:
    # Initializes the QuickSort class with an input array and a comparison counter.
    def __init__(self, input_array):
        self.array = input_array
        self.comparisons = 0

    # Sorts the array using QuickSort with the first element as the pivot.
    def quicksort_first(self, low, high):
        if low < high:
            pivot_index = self.partition_first(low, high)
            self.quicksort_first(low, pivot_index - 1)
            self.quicksort_first(pivot_index + 1, high)

    # Partitions the array using the first element as the pivot.
    def partition_first(self, low, high):
        pivot = self.array[low]
        i = low + 1
        for j in range(low + 1, high + 1):
            if self.array[j] < pivot:
                self.array[i], self.array[j] = self.array[j], self.array[i]
                i += 1
        self.array[low], self.array[i - 1] = self.array[i - 1], self.array[low]
        self.comparisons += (high - low)
        return i - 1

    # Sorts the array using QuickSort with the last element as the pivot.
    def quicksort_last(self, low, high):
        if low < high:
            pivot_index = self.partition_last(low, high)
            self.quicksort_last(low, pivot_index - 1)
            self.quicksort_last(pivot_index + 1, high)

    # Partitions the array using the last element as the pivot.
    def partition_last(self, low, high):
        pivot = self.array[high]
        self.array[low], self.array[high] = self.array[high], self.array[low]
        return self.partition_first(low, high)

    # Sorts the array using QuickSort with the median-of-three as the pivot.
    def quicksort_median(self, low, high):
        if low < high:
            pivot_index = self.partition_median(low, high)
            self.quicksort_median(low, pivot_index - 1)
            self.quicksort_median(pivot_index + 1, high)

    # Partitions the array using the median-of-three as the pivot.
    def partition_median(self, low, high):
        mid = (low + high) // 2
        pivot_candidates = [(self.array[low], low), (self.array[mid], mid), (self.array[high], high)]
        pivot_candidates.sort(key=lambda x: x[0])
        pivot_value, pivot_index = pivot_candidates[1]
        self.array[low], self.array[pivot_index] = self.array[pivot_index], self.array[low]
        return self.partition_first(low, high)

    # Returns the total number of comparisons made during the sort.
    def get_comparisons(self):
        return self.comparisons


# Loads data from the specified file and returns it as a list of integers.
def load_data(filename):
    with open(filename, 'r') as file:
        return [int(line.strip()) for line in file]


# Load the array from the provided txt file.
array = load_data('QuickSort.txt')

# Create an instance of QuickSort using the first element as the pivot.
qs_first = QuickSort(array.copy())
qs_first.quicksort_first(0, len(array) - 1)
print(f"Total comparisons with first element as pivot: {qs_first.get_comparisons()}")

# Create an instance of QuickSort using the last element as the pivot.
qs_last = QuickSort(array.copy())
qs_last.quicksort_last(0, len(array) - 1)
print(f"Total comparisons with last element as pivot: {qs_last.get_comparisons()}")

# Create an instance of QuickSort using the median-of-three as the pivot.
qs_median = QuickSort(array.copy())
qs_median.quicksort_median(0, len(array) - 1)
print(f"Total comparisons with median-of-three as pivot: {qs_median.get_comparisons()}")
