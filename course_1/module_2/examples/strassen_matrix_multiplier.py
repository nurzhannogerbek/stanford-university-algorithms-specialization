class StrassenMatrixMultiplier:
    def __init__(self, matrix_a, matrix_b):
        # Check that matrices A and B are square and of the same dimensions.
        if not (len(matrix_a) == len(matrix_a[0]) == len(matrix_b) == len(matrix_b[0])):
            raise ValueError("Matrices must be square and of the same dimensions")

        # If the size is not a power of 2, pad matrices to the nearest power of 2.
        self.size = len(matrix_a)
        if not (self.size & (self.size - 1) == 0 and self.size != 0):  # Check if size is a power of 2.
            self.A = self._pad_to_power_of_two(matrix_a)
            self.B = self._pad_to_power_of_two(matrix_b)
        else:
            self.A = matrix_a
            self.B = matrix_b

    def multiply(self):
        # Perform Strassen's multiplication and remove padding if it was applied.
        final_result = self._strassen(self.A, self.B)
        # Remove padding if any was added.
        return [final_row[:self.size] for final_row in final_result[:self.size]]

    def _strassen(self, matrix_a, matrix_b):
        n = len(matrix_a)
        # Base case: If the matrix is 1x1, return scalar product directly.
        if n == 1:
            return [[matrix_a[0][0] * matrix_b[0][0]]]

        # Split matrices into quadrants.
        a11, a12, a21, a22 = self._split_matrix(matrix_a)
        b11, b12, b21, b22 = self._split_matrix(matrix_b)

        # Calculate the 7 products using Strassen's algorithm.
        p1 = self._strassen(a11, self._elementwise_subtract(b12, b22))
        p2 = self._strassen(self._elementwise_add(a11, a12), b22)
        p3 = self._strassen(self._elementwise_add(a21, a22), b11)
        p4 = self._strassen(a22, self._elementwise_subtract(b21, b11))
        p5 = self._strassen(self._elementwise_add(a11, a22), self._elementwise_add(b11, b22))
        p6 = self._strassen(self._elementwise_subtract(a12, a22), self._elementwise_add(b21, b22))
        p7 = self._strassen(self._elementwise_subtract(a11, a21), self._elementwise_add(b11, b12))

        # Combine the products to get the final quadrants of the resulting matrix.
        c11 = self._elementwise_add(self._elementwise_subtract(self._elementwise_add(p5, p4), p2), p6)
        c12 = self._elementwise_add(p1, p2)
        c21 = self._elementwise_add(p3, p4)
        c22 = self._elementwise_subtract(self._elementwise_subtract(self._elementwise_add(p5, p1), p3), p7)

        # Combine quadrants into a single matrix.
        return self._combine_quadrants(c11, c12, c21, c22)

    @staticmethod
    def _split_matrix(matrix):
        # Split a matrix into four quadrants.
        mid = len(matrix) // 2
        a11 = [matrix_row[:mid] for matrix_row in matrix[:mid]]
        a12 = [matrix_row[mid:] for matrix_row in matrix[:mid]]
        a21 = [matrix_row[:mid] for matrix_row in matrix[mid:]]
        a22 = [matrix_row[mid:] for matrix_row in matrix[mid:]]
        return a11, a12, a21, a22

    @staticmethod
    def _combine_quadrants(c11, c12, c21, c22):
        # Combine four quadrants into a single matrix.
        top = [row_c11 + row_c12 for row_c11, row_c12 in zip(c11, c12)]
        bottom = [row_c21 + row_c22 for row_c21, row_c22 in zip(c21, c22)]
        return top + bottom

    @staticmethod
    def _pad_to_power_of_two(matrix):
        """Pad the matrix with zeros to the nearest power of two size."""
        n = len(matrix)
        new_size = 1 << (n - 1).bit_length()
        padded_matrix = [[0] * new_size for _ in range(new_size)]
        for i in range(n):
            for j in range(n):
                padded_matrix[i][j] = matrix[i][j]
        return padded_matrix

    @staticmethod
    def _elementwise_add(matrix_a, matrix_b):
        # Add two matrices element-wise.
        n = len(matrix_a)
        return [[matrix_a[i][j] + matrix_b[i][j] for j in range(n)] for i in range(n)]

    @staticmethod
    def _elementwise_subtract(matrix_a, matrix_b):
        # Subtract matrix_b from matrix_a element-wise.
        n = len(matrix_a)
        return [[matrix_a[i][j] - matrix_b[i][j] for j in range(n)] for i in range(n)]


# Example usage.
if __name__ == "__main__":
    # Define two square matrices of arbitrary size.
    A = [[1, 2], [3, 4]]
    B = [[5, 6], [7, 8]]

    multiplier = StrassenMatrixMultiplier(A, B)
    result = multiplier.multiply()

    print("Matrix A:")
    for row in A:
        print(row)
    print("Matrix B:")
    for row in B:
        print(row)
    print("Result of Strassen's Matrix Multiplication (A * B):")
    for row in result:
        print(row)
