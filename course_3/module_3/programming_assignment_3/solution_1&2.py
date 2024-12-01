import heapq

class HuffmanNode:
    """
    Class representing a node in the Huffman tree.
    """
    def __init__(self, weight, symbol=None):
        self.weight = weight  # Weight of the node.
        self.symbol = symbol  # Symbol, if the node is a leaf.
        self.left = None  # Left child node.
        self.right = None  # Right child node.

    def __lt__(self, other):
        """
        Define comparison for heapq (based on weight).
        """
        return self.weight < other.weight


class HuffmanCoding:
    """
    Class to generate Huffman codes using a greedy algorithm.
    """
    def __init__(self, weights):
        """
        Initialize with a list of weights for symbols.
        :param weights: List of weights for the symbols.
        """
        self.weights = weights
        self.root = None  # Root of the Huffman tree.
        self.codes = {}  # Dictionary to store the codes for each symbol.

    def build_tree(self):
        """
        Build the Huffman tree using a priority queue (min-heap).
        """
        # Create a min-heap of leaf nodes.
        heap = [HuffmanNode(weight, symbol=i) for i, weight in enumerate(self.weights)]
        heapq.heapify(heap)

        # Combine nodes until there is only one tree.
        while len(heap) > 1:
            # Remove the two nodes with the smallest weights.
            left = heapq.heappop(heap)
            right = heapq.heappop(heap)

            # Create a new internal node with their combined weight.
            merged = HuffmanNode(left.weight + right.weight)
            merged.left = left
            merged.right = right

            # Add the merged node back to the heap.
            heapq.heappush(heap, merged)

        # The last node in the heap is the root of the Huffman tree.
        self.root = heap[0]

    def generate_codes(self):
        """
        Generate Huffman codes by traversing the tree.
        """
        def traverse(node, code):
            if node is None:
                return

            # If it's a leaf node, assign the code.
            if node.symbol is not None:
                self.codes[node.symbol] = code
                return

            # Traverse left and right children.
            traverse(node.left, code + "0")
            traverse(node.right, code + "1")

        traverse(self.root, "")

    def get_max_min_code_length(self):
        """
        Get the maximum and minimum lengths of codewords in the Huffman code.
        :return: Tuple (max_length, min_length).
        """
        if not self.codes:
            raise ValueError("Huffman codes have not been generated.")

        lengths = [len(code) for code in self.codes.values()]
        return max(lengths), min(lengths)

    def solve(self):
        """
        Build the Huffman tree, generate codes, and return the max and min code lengths.
        :return: Tuple (max_length, min_length).
        """
        self.build_tree()
        self.generate_codes()
        return self.get_max_min_code_length()


# Example usage:
if __name__ == "__main__":
    # Read weights from the file (assumes the file is formatted as described in the task).
    with open("huffman.txt", "r") as file:
        lines = file.readlines()
        num_symbols = int(lines[0].strip())
        weights = [int(line.strip()) for line in lines[1:]]

    # Solve the Huffman coding problem.
    huffman = HuffmanCoding(weights)
    max_length, min_length = huffman.solve()

    print("Maximum length of a codeword:", max_length)
    print("Minimum length of a codeword:", min_length)
