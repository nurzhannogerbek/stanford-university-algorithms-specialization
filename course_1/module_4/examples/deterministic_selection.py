class DeterministicSelection:
    def __init__(self, input_array):
        """
        Initialize the DeterministicSelection class with an array.
        :param input_array: List of elements on which selection will be performed.
        """
        self.array = input_array

    def _partition(self, left, right, pivot_index):
        """
        Partition the array around the pivot element.
        Elements less than the pivot are moved to its left,
        and elements greater than the pivot are moved to its right.
        :param left: Left boundary of the subarray.
        :param right: Right boundary of the subarray.
        :param pivot_index: Index of the pivot element.
        :return: The final position of the pivot in the partitioned array.
        """
        # Get the value of the pivot element.
        pivot_value = self.array[pivot_index]
        # Move pivot to the end of the current subarray.
        self.array[pivot_index], self.array[right] = self.array[right], self.array[pivot_index]
        store_index = left

        # Iterate through the subarray and rearrange elements based on the pivot value.
        for i in range(left, right):
            if self.array[i] < pivot_value:
                # Swap elements to maintain partition order.
                self.array[i], self.array[store_index] = self.array[store_index], self.array[i]
                store_index += 1

        # Move the pivot to its final sorted position.
        self.array[store_index], self.array[right] = self.array[right], self.array[store_index]
        return store_index

    def _median_of_medians(self, left, right):
        """
        Choose a pivot element deterministically using the "median of medians" strategy.
        :param left: Left boundary of the subarray.
        :param right: Right boundary of the subarray.
        :return: Index of the pivot element.
        """
        n = right - left + 1  # Number of elements in the subarray.
        if n <= 5:
            # If the subarray has 5 or fewer elements, sort it and return the index of the median.
            return left + sorted(range(left, right + 1), key=lambda x: self.array[x])[n // 2]

        # Divide the array into groups of 5 elements and find medians.
        for i in range(left, right + 1, 5):
            # Define the boundary for the current group of 5 elements.
            sub_right = min(i + 4, right)
            # Sort the group and find the median.
            median = left + sorted(range(i, sub_right + 1), key=lambda x: self.array[x])[(sub_right - i + 1) // 2]
            # Move the median to the front of the subarray for later processing.
            self.array[left + (i - left) // 5], self.array[median] = self.array[median], self.array[left + (i - left) // 5]

        # Calculate the middle index for the medians array.
        mid_index = left + (n // 10)
        # Recursively find the median of medians and return its index.
        return self._deterministic_select(left, left + (n // 5) - 1, mid_index + 1)

    def _deterministic_select(self, left, right, k):
        """
        Recursively find the k-th smallest element using the deterministic selection algorithm.
        :param left: Left boundary of the subarray.
        :param right: Right boundary of the subarray.
        :param k: 1-based rank of the desired element.
        :return: Index of the k-th smallest element.
        """
        # Base case: If the subarray contains only one element, return its index.
        if left == right:
            return left

        # Select a pivot deterministically using the "median of medians" strategy.
        pivot_index = self._median_of_medians(left, right)

        # Partition the array around the chosen pivot.
        pivot_index = self._partition(left, right, pivot_index)

        # Determine the rank of the pivot in the current subarray.
        rank = pivot_index - left + 1

        if rank == k:
            # The pivot is the k-th smallest element.
            return pivot_index
        elif k < rank:
            # Recurse on the left subarray to find the k-th smallest element.
            return self._deterministic_select(left, pivot_index - 1, k)
        else:
            # Recurse on the right subarray to find the (k-rank)-th smallest element.
            return self._deterministic_select(pivot_index + 1, right, k - rank)

    def select(self, k):
        """
        Public method to find the k-th smallest element in the array.
        :param k: 1-based index of the desired smallest element.
        :return: The k-th smallest element in the array.
        """
        # Validate that k is within the valid range.
        if k < 1 or k > len(self.array):
            raise ValueError("k is out of bounds of the array.")
        # Perform the deterministic selection algorithm and return the result.
        index = self._deterministic_select(0, len(self.array) - 1, k)
        return self.array[index]


# Example usage.
if __name__ == "__main__":
    array = [10, 4, 5, 8, 6, 11, 26]  # Input array.
    k = 3  # Desired rank (1-based).

    # Create an instance of DeterministicSelection.
    selector = DeterministicSelection(array)
    # Find the k-th smallest element.
    result = selector.select(k)
    # Print the result.
    print(f"The {k}-th smallest element is: {result}")
