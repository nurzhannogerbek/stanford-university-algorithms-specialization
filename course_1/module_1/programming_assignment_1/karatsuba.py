import math

class KaratsubaMultiplier:
    @staticmethod
    def karatsuba_multiply(x: int, y: int) -> int:
        """Public method to start the Karatsuba multiplication."""
        return KaratsubaMultiplier._karatsuba(x, y)

    @staticmethod
    def _karatsuba(x: int, y: int) -> int:
        """Implements the recursive Karatsuba algorithm."""
        # Base case: if either number is small enough, perform direct multiplication.
        if x < 100 or y < 100:  # Increased threshold for base case.
            return x * y

        # Calculate the size of the numbers without converting to strings.
        n = max(int(math.log10(x) + 1), int(math.log10(y) + 1))
        m = n // 2

        # Split x and y into high and low parts.
        high_x, low_x = divmod(x, 10**m)
        high_y, low_y = divmod(y, 10**m)

        # Recursively compute three products.
        z0 = KaratsubaMultiplier._karatsuba(low_x, low_y)
        z2 = KaratsubaMultiplier._karatsuba(high_x, high_y)
        sum_x = low_x + high_x
        sum_y = low_y + high_y
        z1 = KaratsubaMultiplier._karatsuba(sum_x, sum_y)

        # Combine the results.
        return (z2 * 10**(2 * m)) + ((z1 - z2 - z0) * 10**m) + z0

# Example usage.
x = 3141592653589793238462643383279502884197169399375105820974944592
y = 2718281828459045235360287471352662497757247093699959574966967627
result = KaratsubaMultiplier.karatsuba_multiply(x, y)
print("Product:", result)
