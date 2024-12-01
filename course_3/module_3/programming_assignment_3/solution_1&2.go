package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// HuffmanNode represents a node in the Huffman tree.
type HuffmanNode struct {
	Weight int          // Weight of the node.
	Symbol int          // Symbol index (if it's a leaf node).
	Left   *HuffmanNode // Pointer to the left child.
	Right  *HuffmanNode // Pointer to the right child.
}

// PriorityQueue implements a min-heap for Huffman nodes.
type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*HuffmanNode))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// HuffmanTree represents the Huffman tree.
type HuffmanTree struct {
	Root *HuffmanNode // Root of the Huffman tree.
}

// Build constructs the Huffman tree using the given weights.
func (ht *HuffmanTree) Build(weights []int) {
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Create a leaf node for each symbol and add it to the priority queue.
	for i, weight := range weights {
		node := &HuffmanNode{Weight: weight, Symbol: i}
		heap.Push(pq, node)
	}

	// Combine nodes until there is only one tree.
	for pq.Len() > 1 {
		// Remove the two nodes with the smallest weights.
		left := heap.Pop(pq).(*HuffmanNode)
		right := heap.Pop(pq).(*HuffmanNode)

		// Create a new internal node with their combined weight.
		merged := &HuffmanNode{
			Weight: left.Weight + right.Weight,
			Left:   left,
			Right:  right,
		}

		// Add the merged node back to the priority queue.
		heap.Push(pq, merged)
	}

	// The remaining node is the root of the Huffman tree.
	ht.Root = heap.Pop(pq).(*HuffmanNode)
}

// GetCodeLengths computes the lengths of Huffman codes for all symbols.
func (ht *HuffmanTree) GetCodeLengths() (int, int) {
	var traverse func(node *HuffmanNode, depth int)
	maxLength := 0
	minLength := 1 << 31 // A large value for initialization.

	traverse = func(node *HuffmanNode, depth int) {
		if node == nil {
			return
		}

		// If it's a leaf node, record the depth as the code length.
		if node.Left == nil && node.Right == nil {
			if depth > maxLength {
				maxLength = depth
			}
			if depth < minLength {
				minLength = depth
			}
			return
		}

		// Recursively traverse the left and right children.
		traverse(node.Left, depth+1)
		traverse(node.Right, depth+1)
	}

	// Start traversal from the root at depth 0.
	traverse(ht.Root, 0)
	return maxLength, minLength
}

func main() {
	// Open and read the input file.
	file, err := os.ReadFile("course_3/module_3/programming_assignment_3/huffman.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the weights from the file.
	lines := strings.Split(string(file), "\n")
	numSymbols, _ := strconv.Atoi(lines[0])
	weights := make([]int, numSymbols)
	for i := 1; i <= numSymbols; i++ {
		weights[i-1], _ = strconv.Atoi(lines[i])
	}

	// Build the Huffman tree and compute code lengths.
	huffmanTree := &HuffmanTree{}
	huffmanTree.Build(weights)
	maxLength, minLength := huffmanTree.GetCodeLengths()

	// Print the results.
	fmt.Println("Maximum length of a codeword:", maxLength)
	fmt.Println("Minimum length of a codeword:", minLength)
}
