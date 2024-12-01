import heapq
from collections import defaultdict


class Node:
    """Represents a node in the Huffman Tree."""

    def __init__(self, char, freq):
        """
        Initialize a node with a character and its frequency.

        Args:
            char (str): The character represented by the node.
            freq (int): The frequency of the character.
        """
        self.char = char
        self.freq = freq
        self.left = None
        self.right = None

    def __lt__(self, other):
        """
        Less-than operator for priority queue comparisons.

        Args:
            other (Node): Another node to compare with.

        Returns:
            bool: True if this node's frequency is less than the other.
        """
        return self.freq < other.freq

    def __str__(self):
        """Return a string representation of the node for debugging."""
        return f"Node(char={self.char}, freq={self.freq})"


class HuffmanCoding:
    """
    Implements Huffman encoding and decoding for lossless compression.

    Attributes:
        codes (dict): Maps characters to their binary Huffman codes.
        reverse_mapping (dict): Maps binary Huffman codes to their characters.
    """

    def __init__(self):
        """Initialize the HuffmanCoding class with empty mappings."""
        self.codes = {}
        self.reverse_mapping = {}

    def build_frequency_table(self, text):
        """
        Builds a frequency table from the input text.

        Args:
            text (str): The text to analyze.

        Returns:
            dict: A dictionary where keys are characters and values are frequencies.
        """
        if not text:
            raise ValueError("Input text cannot be empty.")

        frequency = defaultdict(int)
        for char in text:
            frequency[char] += 1
        return frequency

    def build_huffman_tree(self, frequency_table):
        """
        Constructs the Huffman Tree from a given frequency table.

        Args:
            frequency_table (dict): A dictionary where keys are characters and values are their frequencies.

        Returns:
            Node: The root of the Huffman tree.
        """
        heap = [Node(char, freq) for char, freq in frequency_table.items()]
        heapq.heapify(heap)

        while len(heap) > 1:
            node1 = heapq.heappop(heap)
            node2 = heapq.heappop(heap)

            merged = Node(None, node1.freq + node2.freq)
            merged.left = node1
            merged.right = node2

            heapq.heappush(heap, merged)

        return heap[0]

    def generate_codes(self, root, current_code=""):
        """
        Recursively generates Huffman codes for each character.

        Args:
            root (Node): The root of the Huffman Tree.
            current_code (str): The binary code being generated (used internally).
        """
        if root is None:
            return

        if root.char is not None:  # Leaf node
            self.codes[root.char] = current_code
            self.reverse_mapping[current_code] = root.char
            return

        self.generate_codes(root.left, current_code + "0")
        self.generate_codes(root.right, current_code + "1")

    def encode(self, text):
        """
        Encodes the input text into a binary string using Huffman codes.

        Args:
            text (str): The text to encode.

        Returns:
            str: The encoded binary string.
        """
        if not text:
            return ""

        encoded_text = "".join(self.codes[char] for char in text)
        return encoded_text

    def decode(self, encoded_text):
        """
        Decodes a binary string back into the original text.

        Args:
            encoded_text (str): The encoded binary string.

        Returns:
            str: The original text.
        """
        current_code = ""
        decoded_text = []

        for bit in encoded_text:
            current_code += bit
            if current_code in self.reverse_mapping:
                decoded_text.append(self.reverse_mapping[current_code])
                current_code = ""

        return "".join(decoded_text)

    def compress(self, text):
        """
        Compresses the input text using Huffman encoding.

        Args:
            text (str): The text to compress.

        Returns:
            str: The encoded binary string.
        """
        frequency_table = self.build_frequency_table(text)
        huffman_tree = self.build_huffman_tree(frequency_table)
        self.generate_codes(huffman_tree)
        return self.encode(text)

    def decompress(self, encoded_text):
        """
        Decompresses an encoded binary string back into the original text.

        Args:
            encoded_text (str): The encoded binary string.

        Returns:
            str: The original text.
        """
        return self.decode(encoded_text)


# Example usage:
if __name__ == "__main__":
    text = "this is an example for huffman encoding"
    huffman = HuffmanCoding()

    # Compress the text.
    try:
        encoded_text = huffman.compress(text)
        print("Encoded Text:", encoded_text)

        # Decompress the text.
        decoded_text = huffman.decompress(encoded_text)
        print("Decoded Text:", decoded_text)

        # Check if the original text matches the decoded text.
        assert text == decoded_text, "Error: Original and decoded text do not match."
        print("Compression and Decompression are successful!")
    except ValueError as e:
        print(f"Error: {e}")
