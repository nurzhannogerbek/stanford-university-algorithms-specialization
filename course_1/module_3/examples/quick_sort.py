class QuickSort:
    def __init__(self, array):
        # Initializes the QuickSort class with the array to be sorted and a counter for comparisons.
        self.array = array
        self.comparisons = 0

    def sort_with_first_pivot(self):
        # Sorts the array using the first element as the pivot.
        self._quicksort_first(0, len(self.array) - 1)

    def sort_with_last_pivot(self):
        # Sorts the array using the last element as the pivot.
        self._quicksort_last(0, len(self.array) - 1)

    def sort_with_median_pivot(self):
        # Sorts the array using the median of the first, middle, and last elements as the pivot.
        self._quicksort_median(0, len(self.array) - 1)

    def _quicksort_first(self, low, high):
        # Recursively sorts the array using the first element as the pivot.
        if low < high:
            pivot_index = self._partition_first(low, high)
            self._quicksort_first(low, pivot_index - 1)
            self._quicksort_first(pivot_index + 1, high)

    def _partition_first(self, low, high):
        # Partitions the array around the first element as the pivot.
        pivot = self.array[low]
        i = low + 1
        for j in range(low + 1, high + 1):
            if self.array[j] < pivot:
                self.array[i], self.array[j] = self.array[j], self.array[i]
                i += 1
        self.array[low], self.array[i - 1] = self.array[i - 1], self.array[low]
        self.comparisons += high - low
        return i - 1

    def _quicksort_last(self, low, high):
        # Recursively sorts the array using the last element as the pivot.
        if low < high:
            pivot_index = self._partition_last(low, high)
            self._quicksort_last(low, pivot_index - 1)
            self._quicksort_last(pivot_index + 1, high)

    def _partition_last(self, low, high):
        # Partitions the array around the last element as the pivot.
        self.array[low], self.array[high] = self.array[high], self.array[low]  # Swap last element to the front
        return self._partition_first(low, high)

    def _quicksort_median(self, low, high):
        # Recursively sorts the array using the median-of-three as the pivot.
        if low < high:
            pivot_index = self._partition_median(low, high)
            self._quicksort_median(low, pivot_index - 1)
            self._quicksort_median(pivot_index + 1, high)

    def _partition_median(self, low, high):
        # Partitions the array around the median-of-three as the pivot.
        mid = (low + high) // 2
        pivot_candidates = [(self.array[low], low), (self.array[mid], mid), (self.array[high], high)]
        pivot_candidates.sort(key=lambda x: x[0])
        pivot_value, pivot_index = pivot_candidates[1]
        self.array[low], self.array[pivot_index] = self.array[pivot_index], self.array[low]
        return self._partition_first(low, high)

    def get_comparisons(self):
        # Returns the total number of comparisons made during the sorting.
        return self.comparisons

    def get_sorted_array(self):
        # Returns the sorted array.
        return self.array


# Example usage
if __name__ == "__main__":
    array = [3, 8, 2, 5, 1, 4, 7, 6]

    # Sorting with the first element as the pivot
    qs_first = QuickSort(array.copy())
    qs_first.sort_with_first_pivot()
    print("Sorted array with first element as pivot:", qs_first.get_sorted_array())
    print("Total comparisons with first element as pivot:", qs_first.get_comparisons())

    # Sorting with the last element as the pivot
    qs_last = QuickSort(array.copy())
    qs_last.sort_with_last_pivot()
    print("Sorted array with last element as pivot:", qs_last.get_sorted_array())
    print("Total comparisons with last element as pivot:", qs_last.get_comparisons())

    # Sorting with the median-of-three as the pivot
    qs_median = QuickSort(array.copy())
    qs_median.sort_with_median_pivot()
    print("Sorted array with median-of-three as pivot:", qs_median.get_sorted_array())
    print("Total comparisons with median-of-three as pivot:", qs_median.get_comparisons())
