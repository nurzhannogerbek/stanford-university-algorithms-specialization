package main

import (
	"container/heap"
	"errors"
	"fmt"
	"strings"
)

// Node represents a node in the Huffman Tree.
type Node struct {
	Char  rune  // The character represented by the node.
	Freq  int   // The frequency of the character.
	Left  *Node // Left child node.
	Right *Node // Right child node.
}

// PriorityQueue implements a priority queue of nodes.
type PriorityQueue []*Node

// Methods to implement heap.Interface for PriorityQueue.
func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Freq < pq[j].Freq }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// HuffmanCoding encapsulates the Huffman encoding process.
type HuffmanCoding struct {
	Codes          map[rune]string // Character-to-code mapping.
	ReverseMapping map[string]rune // Code-to-character mapping.
}

// NewHuffmanCoding initializes a new HuffmanCoding instance.
func NewHuffmanCoding() *HuffmanCoding {
	return &HuffmanCoding{
		Codes:          make(map[rune]string),
		ReverseMapping: make(map[string]rune),
	}
}

// BuildFrequencyTable constructs a frequency table for the input text.
func (hc *HuffmanCoding) BuildFrequencyTable(text string) (map[rune]int, error) {
	if len(text) == 0 {
		return nil, errors.New("input text cannot be empty")
	}

	frequency := make(map[rune]int)
	for _, char := range text {
		frequency[char]++
	}
	return frequency, nil
}

// BuildHuffmanTree constructs the Huffman Tree using a priority queue.
func (hc *HuffmanCoding) BuildHuffmanTree(frequencyTable map[rune]int) *Node {
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Add all characters as leaf nodes to the priority queue.
	for char, freq := range frequencyTable {
		heap.Push(pq, &Node{Char: char, Freq: freq})
	}

	// Build the tree by combining nodes with the smallest frequencies.
	for pq.Len() > 1 {
		node1 := heap.Pop(pq).(*Node)
		node2 := heap.Pop(pq).(*Node)

		merged := &Node{
			Freq:  node1.Freq + node2.Freq,
			Left:  node1,
			Right: node2,
		}

		heap.Push(pq, merged)
	}

	// The final node in the priority queue is the root of the Huffman Tree.
	return heap.Pop(pq).(*Node)
}

// GenerateCodes recursively generates Huffman codes for each character.
func (hc *HuffmanCoding) GenerateCodes(node *Node, currentCode string) {
	if node == nil {
		return
	}

	// If the node is a leaf, save its code.
	if node.Char != 0 {
		hc.Codes[node.Char] = currentCode
		hc.ReverseMapping[currentCode] = node.Char
		return
	}

	hc.GenerateCodes(node.Left, currentCode+"0")
	hc.GenerateCodes(node.Right, currentCode+"1")
}

// Encode compresses the input text into a binary string.
func (hc *HuffmanCoding) Encode(text string) (string, error) {
	if len(text) == 0 {
		return "", errors.New("input text cannot be empty")
	}

	var encodedText strings.Builder
	for _, char := range text {
		code, exists := hc.Codes[char]
		if !exists {
			return "", fmt.Errorf("character %q not found in encoding table", char)
		}
		encodedText.WriteString(code)
	}
	return encodedText.String(), nil
}

// Decode decompresses the binary string back into the original text.
func (hc *HuffmanCoding) Decode(encodedText string) (string, error) {
	var decodedText strings.Builder
	currentCode := ""

	for _, bit := range encodedText {
		currentCode += string(bit)
		if char, exists := hc.ReverseMapping[currentCode]; exists {
			decodedText.WriteRune(char)
			currentCode = ""
		}
	}

	if currentCode != "" {
		return "", errors.New("incomplete encoding in the input")
	}

	return decodedText.String(), nil
}

// Compress compresses the input text using Huffman encoding.
func (hc *HuffmanCoding) Compress(text string) (string, error) {
	frequencyTable, err := hc.BuildFrequencyTable(text)
	if err != nil {
		return "", err
	}

	huffmanTree := hc.BuildHuffmanTree(frequencyTable)
	hc.GenerateCodes(huffmanTree, "")

	return hc.Encode(text)
}

// Decompress decompresses the encoded text using Huffman decoding.
func (hc *HuffmanCoding) Decompress(encodedText string) (string, error) {
	return hc.Decode(encodedText)
}

// Main function for testing the HuffmanCoding implementation.
func main() {
	text := "this is an example for huffman encoding"

	huffman := NewHuffmanCoding()

	// Compress the text.
	encodedText, err := huffman.Compress(text)
	if err != nil {
		fmt.Println("Error during compression:", err)
		return
	}
	fmt.Println("Encoded Text:", encodedText)

	// Decompress the text.
	decodedText, err := huffman.Decompress(encodedText)
	if err != nil {
		fmt.Println("Error during decompression:", err)
		return
	}
	fmt.Println("Decoded Text:", decodedText)

	// Verify correctness.
	if text == decodedText {
		fmt.Println("Compression and Decompression are successful!")
	} else {
		fmt.Println("Error: Original and decoded text do not match.")
	}
}
