import os
from typing import List


class InversionCounter:
    def __init__(self, array: List[int]):
        # Initialize the array to count inversions in.
        self.array = array

    def count_inversions(self) -> int:
        # Initialize a temporary array for storing intermediate results.
        temp_array = [0] * len(self.array)
        # Start the recursive inversion counting process.
        return self._count_inversions(temp_array, 0, len(self.array) - 1)

    def _count_inversions(self, temp_array: List[int], left: int, right: int) -> int:
        # Base case: if the subarray has one or no elements, there are no inversions.
        if left >= right:
            return 0

        # Find the midpoint of the current subarray.
        mid = (left + right) // 2

        # Count inversions in the left and right halves, then merge them.
        inv_count = self._count_inversions(temp_array, left, mid)
        inv_count += self._count_inversions(temp_array, mid + 1, right)
        inv_count += self._merge_and_count(temp_array, left, mid, right)

        return inv_count

    def _merge_and_count(self, temp_array: List[int], left: int, mid: int, right: int) -> int:
        # Initialize pointers for the two halves and a pointer for the temp_array.
        i, j, k = left, mid + 1, left
        inv_count = 0

        # Merge the two halves while counting inversions.
        while i <= mid and j <= right:
            if self.array[i] <= self.array[j]:
                temp_array[k] = self.array[i]
                i += 1
            else:
                temp_array[k] = self.array[j]
                # All remaining elements in the left half are greater than array[j].
                inv_count += (mid - i + 1)
                j += 1
            k += 1

        # Copy any remaining elements from the left half.
        while i <= mid:
            temp_array[k] = self.array[i]
            i += 1
            k += 1

        # Copy any remaining elements from the right half.
        while j <= right:
            temp_array[k] = self.array[j]
            j += 1
            k += 1

        # Copy the sorted and merged subarray back into the original array.
        self.array[left:right + 1] = temp_array[left:right + 1]

        return inv_count


# Get the path to the current directory where the script is located.
current_dir = os.path.dirname(os.path.abspath(__file__))

# Specify the path to the file 'IntegerArray.txt' relative to the current directory.
file_path = os.path.join(current_dir, 'IntegerArray.txt')

# Read the file.
with open(file_path, 'r') as file:
    input_array = [int(line.strip()) for line in file]

# Create an instance of InversionCounter and count inversions.
counter = InversionCounter(input_array)
inversion_count = counter.count_inversions()

print("Number of inversions:", inversion_count)
