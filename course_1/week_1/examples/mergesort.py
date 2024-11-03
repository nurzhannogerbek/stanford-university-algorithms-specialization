class MergeSort:
    def __init__(self, array):
        if not isinstance(array, list):
            raise TypeError("Input should be a list.")
        self.array = array

    def sort(self):
        """Public method to start the merge sort."""
        if len(self.array) <= 1:
            return self.array
        return self._merge_sort(self.array)

    @staticmethod
    def _merge_sort(sub_array):
        """Recursively splits and sorts the array."""
        # Base case: if the list has 1 or 0 elements, it's already sorted.
        if len(sub_array) <= 1:
            return sub_array

        # Split the list into two halves.
        mid = len(sub_array) // 2
        left_half = MergeSort._merge_sort(sub_array[:mid])
        right_half = MergeSort._merge_sort(sub_array[mid:])

        # Merge the sorted halves.
        return MergeSort._merge(left_half, right_half)

    @staticmethod
    def _merge(left, right):
        """Merges two sorted halves into a single sorted list."""
        merged_array = []
        i = j = 0

        # Merge the two halves by comparing elements.
        while i < len(left) and j < len(right):
            if left[i] < right[j]:
                merged_array.append(left[i])
                i += 1
            else:
                merged_array.append(right[j])
                j += 1

        # Append any remaining elements from both halves.
        merged_array.extend(left[i:])
        merged_array.extend(right[j:])

        return merged_array

# Example usage.
arr = [38, 27, 43, 3, 9, 82, 10]
merge_sort_instance = MergeSort(arr)
sorted_arr = merge_sort_instance.sort()
print("Sorted array:", sorted_arr)
