import random

class RandomizedSelection:
    """
    Class for finding the k-th smallest element using a randomized selection algorithm.

    Attributes:
        array (list): The input array of distinct elements.
    """

    def __init__(self, array):
        """
        Initialize the RandomizedSelection class with an array of distinct elements.
        """
        if not array:
            raise ValueError("Input array must not be empty.")
        if len(set(array)) != len(array):
            raise ValueError("Array must contain distinct elements.")
        self.array = array[:]

    def _partition(self, left, right, pivot_index):
        """
        Partition the array around the pivot element.
        """
        pivot_value = self.array[pivot_index]
        self.array[pivot_index], self.array[right] = self.array[right], self.array[pivot_index]
        store_index = left
        for i in range(left, right):
            if self.array[i] < pivot_value:
                self.array[i], self.array[store_index] = self.array[store_index], self.array[i]
                store_index += 1
        self.array[right], self.array[store_index] = self.array[store_index], self.array[right]
        return store_index

    def _randomized_select(self, left, right, k):
        """
        Recursively find the k-th smallest element in the array.
        """
        if left == right:
            return self.array[left]

        pivot_index = random.randint(left, right)
        pivot_index = self._partition(left, right, pivot_index)
        rank = pivot_index - left + 1

        if rank == k:
            return self.array[pivot_index]
        elif k < rank:
            return self._randomized_select(left, pivot_index - 1, k)
        else:
            return self._randomized_select(pivot_index + 1, right, k - rank)

    def select(self, k):
        """
        Find the k-th smallest element in the array.

        Args:
            k (int): 1-based index of the desired smallest element.

        Returns:
            int: The k-th smallest element in the array.
        """
        if k < 1 or k > len(self.array):
            raise ValueError("k is out of bounds of the array.")
        return self._randomized_select(0, len(self.array) - 1, k)


# Example usage.
if __name__ == "__main__":
    array = [10, 4, 5, 8, 6, 11, 26]
    k = 3
    selector = RandomizedSelection(array)
    result = selector.select(k)
    print(f"The {k}-th smallest element is: {result}")
