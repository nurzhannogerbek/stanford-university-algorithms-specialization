import hashlib
import math


class BloomFilter:
    def __init__(self, num_items: int, false_positive_rate: float):
        """
        Initialize the Bloom Filter.
        :param num_items: Estimated number of items to store.
        :param false_positive_rate: Desired false positive rate.
        """
        if num_items <= 0:
            raise ValueError("Number of items must be greater than 0.")
        if not (0 < false_positive_rate < 1):
            raise ValueError("False positive rate must be between 0 and 1 (exclusive).")

        self.num_items = num_items
        self.false_positive_rate = false_positive_rate

        # Calculate the optimal size of the bit array and the number of hash functions.
        self.size = self._optimal_size(num_items, false_positive_rate)
        self.num_hashes = self._optimal_num_hashes(self.size, num_items)

        # Create a bit array represented as a list of integers.
        self.bit_array = [0] * ((self.size + 31) // 32)  # Use 32 bits per integer.

    @staticmethod
    def _optimal_size(n: int, p: float) -> int:
        """
        Calculate the optimal size of the bit array.
        :param n: Number of items expected to be stored.
        :param p: Desired false positive rate.
        :return: Size of the bit array.
        """
        return int(-n * math.log(p) / (math.log(2) ** 2))

    @staticmethod
    def _optimal_num_hashes(m: int, n: int) -> int:
        """
        Calculate the optimal number of hash functions.
        :param m: Size of the bit array.
        :param n: Number of items expected to be stored.
        :return: Number of hash functions.
        """
        return max(1, int((m / n) * math.log(2)))

    def _hash(self, element: str, seed: int) -> int:
        """
        Generate a hash value for the given element using a specific seed.
        :param element: The element to hash.
        :param seed: The seed for the hash function.
        :return: Hash value modulo the size of the bit array.
        """
        data = f"{seed}-{element}".encode("utf-8")
        hash_value = int(hashlib.sha256(data).hexdigest(), 16)
        return hash_value % self.size

    def _set_bit(self, index: int) -> None:
        """
        Set a specific bit in the bit array.
        :param index: Index of the bit to set.
        """
        self.bit_array[index // 32] |= (1 << (index % 32))

    def _get_bit(self, index: int) -> bool:
        """
        Get the value of a specific bit in the bit array.
        :param index: Index of the bit to check.
        :return: True if the bit is set, False otherwise.
        """
        return (self.bit_array[index // 32] & (1 << (index % 32))) != 0

    def add(self, element: str) -> None:
        """
        Add an element to the Bloom Filter.
        :param element: The element to add. Must be convertible to a string.
        """
        if not isinstance(element, str):
            element = str(element)

        for i in range(self.num_hashes):
            index = self._hash(element, i)
            self._set_bit(index)

    def contains(self, element: str) -> bool:
        """
        Check if an element is in the Bloom Filter.
        :param element: The element to check. Must be convertible to a string.
        :return: True if the element might be in the Bloom Filter, False otherwise.
        """
        if not isinstance(element, str):
            element = str(element)

        for i in range(self.num_hashes):
            index = self._hash(element, i)
            if not self._get_bit(index):
                return False
        return True

    def __str__(self) -> str:
        """
        Return a string representation of the Bloom Filter's parameters.
        """
        return f"BloomFilter(size={self.size}, num_hashes={self.num_hashes})"


# Example usage:
if __name__ == "__main__":
    # Create a Bloom Filter for 100 items with a false positive rate of 1%.
    bloom_filter = BloomFilter(num_items=100, false_positive_rate=0.01)

    print(bloom_filter)

    # Add items to the Bloom Filter.
    items_to_add = ["apple", "banana", "orange", "grape"]
    for item in items_to_add:
        bloom_filter.add(item)
        print(f"Added: {item}")

    # Check for items in the Bloom Filter.
    check_items = ["apple", "cherry", "orange", "watermelon"]
    for item in check_items:
        result = bloom_filter.contains(item)
        print(f"Item '{item}' is in the Bloom Filter: {result}")
