import random
from typing import Any, List, Optional


class UniversalHashFunction:
    def __init__(self, p: int, m: int, k: int, a: Optional[List[int]] = None, b: Optional[int] = None):
        """
        Initialize a universal hash function.
        :param p: A prime number larger than the universe size U.
        :param m: Number of buckets in the hash table.
        :param k: Dimensionality of keys (e.g., 4 for IP addresses).
        :param a: Optionally provided coefficients a.
        :param b: Optionally provided value b.
        """
        if not self._is_prime(p):
            raise ValueError(f"The number {p} must be prime.")

        self.p = p
        self.m = m
        self.k = k
        self.a = a if a else [random.randint(1, p - 1) for _ in range(k)]
        self.b = b if b else random.randint(0, p - 1)

    def hash(self, key: List[int]) -> int:
        """
        Compute the hash for a given key.
        :param key: A key represented as a list of integers.
        :return: The bucket index.
        """
        if len(key) != self.k:
            raise ValueError(f"Key size must be {self.k}, got {len(key)}.")
        return (sum(a * x for a, x in zip(self.a, key)) + self.b) % self.p % self.m

    @staticmethod
    def _is_prime(num: int) -> bool:
        """Check if a number is prime."""
        if num < 2:
            return False
        for i in range(2, int(num**0.5) + 1):
            if num % i == 0:
                return False
        return True


class HashTableWithChaining:
    def __init__(self, initial_buckets: int, p: int, k: int, load_factor_threshold: float = 0.7):
        """
        Initialize a hash table with chaining.
        :param initial_buckets: Initial number of buckets.
        :param p: A prime number for the hash function.
        :param k: Dimensionality of keys.
        :param load_factor_threshold: Load factor threshold for resizing.
        """
        self.p = p
        self.k = k
        self.load_factor_threshold = load_factor_threshold
        self.num_buckets = initial_buckets
        self.size = 0
        self.table = [[] for _ in range(initial_buckets)]
        self.hash_function = UniversalHashFunction(p, initial_buckets, k)

    def insert(self, key: List[int], value: Any) -> None:
        """
        Insert a key-value pair into the hash table.
        """
        if len(key) != self.k:
            raise ValueError(f"Key size must be {self.k}.")
        if self.size / self.num_buckets > self.load_factor_threshold:
            self._resize()

        bucket_index = self.hash_function.hash(key)
        for pair in self.table[bucket_index]:
            if pair[0] == key:
                pair[1] = value  # Update value
                return
        self.table[bucket_index].append([key, value])
        self.size += 1

    def search(self, key: List[int]) -> Optional[Any]:
        """
        Search for a value by key.
        :param key: The key to search for.
        :return: The associated value or None if the key is not found.
        """
        bucket_index = self.hash_function.hash(key)
        for pair in self.table[bucket_index]:
            if pair[0] == key:
                return pair[1]
        return None

    def delete(self, key: List[int]) -> bool:
        """
        Delete an element by key.
        :return: True if the element was deleted, False otherwise.
        """
        bucket_index = self.hash_function.hash(key)
        for i, pair in enumerate(self.table[bucket_index]):
            if pair[0] == key:
                del self.table[bucket_index][i]
                self.size -= 1
                return True
        return False

    def _resize(self) -> None:
        """
        Resize the hash table when the load factor exceeds the threshold.
        """
        old_table = self.table
        self.num_buckets *= 2
        self.table = [[] for _ in range(self.num_buckets)]
        self.hash_function = UniversalHashFunction(self.p, self.num_buckets, self.k)
        self.size = 0

        for bucket in old_table:
            for key, value in bucket:
                self.insert(key, value)

    def display(self, limit: int = 10) -> None:
        """
        Display the contents of the hash table.
        :param limit: Maximum number of buckets to display.
        """
        print(f"Hash Table (size: {self.size}, buckets: {self.num_buckets}):")
        displayed = 0
        for i, bucket in enumerate(self.table):
            if bucket:
                print(f"Bucket {i}: {bucket}")
                displayed += 1
            if displayed >= limit:
                print(f"... and more buckets.")
                break


# Example usage:
if __name__ == "__main__":
    hash_table = HashTableWithChaining(initial_buckets=10, p=31, k=4)

    # Add key-value pairs.
    hash_table.insert([192, 168, 0, 1], "Device A")
    hash_table.insert([192, 168, 0, 2], "Device B")
    hash_table.insert([10, 0, 0, 1], "Device C")

    # Search for values.
    print("Search [192, 168, 0, 1]:", hash_table.search([192, 168, 0, 1]))
    print("Search [192, 168, 0, 3]:", hash_table.search([192, 168, 0, 3]))

    # Delete a value.
    print("Delete [192, 168, 0, 1]:", hash_table.delete([192, 168, 0, 1]))

    # Display the hash table.
    hash_table.display()
