package main

import (
	"errors"
	"fmt"
)

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Key    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
	Size   int // Size of the subtree, including this node.
}

// BinarySearchTree represents the binary search tree structure.
type BinarySearchTree struct {
	Root *TreeNode
}

// NewTreeNode creates a new tree node.
func NewTreeNode(key int) *TreeNode {
	return &TreeNode{
		Key:  key,
		Size: 1, // A single node has size 1.
	}
}

// Insert adds a new key to the BST. Ignores duplicates.
func (bst *BinarySearchTree) Insert(key int) {
	if bst.Search(key) != nil {
		return // Ignore duplicates.
	}
	newNode := NewTreeNode(key)
	if bst.Root == nil {
		bst.Root = newNode
		return
	}
	current := bst.Root
	for current != nil {
		current.Size++ // Update size of the subtree.
		if key < current.Key {
			if current.Left == nil {
				current.Left = newNode
				newNode.Parent = current
				bst.updateSize(current)
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				newNode.Parent = current
				bst.updateSize(current)
				return
			}
			current = current.Right
		}
	}
}

// Search finds a node with the given key in the BST.
func (bst *BinarySearchTree) Search(key int) *TreeNode {
	current := bst.Root
	for current != nil {
		if key == current.Key {
			return current
		} else if key < current.Key {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return nil
}

// Minimum finds the node with the smallest key in the BST or subtree.
func (bst *BinarySearchTree) Minimum(node *TreeNode) *TreeNode {
	if node == nil {
		node = bst.Root
	}
	for node != nil && node.Left != nil {
		node = node.Left
	}
	return node
}

// Maximum finds the node with the largest key in the BST or subtree.
func (bst *BinarySearchTree) Maximum(node *TreeNode) *TreeNode {
	if node == nil {
		node = bst.Root
	}
	for node != nil && node.Right != nil {
		node = node.Right
	}
	return node
}

// Successor finds the next largest key after the given key.
func (bst *BinarySearchTree) Successor(key int) (*int, error) {
	node := bst.Search(key)
	if node == nil {
		return nil, errors.New("key not found")
	}
	if node.Right != nil {
		min := bst.Minimum(node.Right)
		return &min.Key, nil
	}
	current := node.Parent
	for current != nil && node == current.Right {
		node = current
		current = current.Parent
	}
	if current == nil {
		return nil, errors.New("no successor found")
	}
	return &current.Key, nil
}

// Predecessor finds the next smaller key before the given key.
func (bst *BinarySearchTree) Predecessor(key int) (*int, error) {
	node := bst.Search(key)
	if node == nil {
		return nil, errors.New("key not found")
	}
	if node.Left != nil {
		max := bst.Maximum(node.Left)
		return &max.Key, nil
	}
	current := node.Parent
	for current != nil && node == current.Left {
		node = current
		current = current.Parent
	}
	if current == nil {
		return nil, errors.New("no predecessor found")
	}
	return &current.Key, nil
}

// Delete removes a node with the given key from the BST.
func (bst *BinarySearchTree) Delete(key int) error {
	node := bst.Search(key)
	if node == nil {
		return errors.New("key not found")
	}
	if node.Left == nil && node.Right == nil { // Case 1: No children.
		bst.transplant(node, nil)
	} else if node.Right == nil { // Case 2: One child (left).
		bst.transplant(node, node.Left)
	} else if node.Left == nil { // Case 2: One child (right).
		bst.transplant(node, node.Right)
	} else { // Case 3: Two children.
		successor := bst.Minimum(node.Right)
		if successor.Parent != node {
			bst.transplant(successor, successor.Right)
			successor.Right = node.Right
			successor.Right.Parent = successor
		}
		bst.transplant(node, successor)
		successor.Left = node.Left
		successor.Left.Parent = successor
		successor.Size = node.Size
	}
	bst.updateSize(node.Parent) // Update sizes up the tree.
	return nil
}

// transplant replaces one subtree with another.
func (bst *BinarySearchTree) transplant(u, v *TreeNode) {
	if u.Parent == nil {
		bst.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

// updateSize updates the sizes of all ancestor nodes.
func (bst *BinarySearchTree) updateSize(node *TreeNode) {
	for node != nil {
		node.Size = 1
		if node.Left != nil {
			node.Size += node.Left.Size
		}
		if node.Right != nil {
			node.Size += node.Right.Size
		}
		node = node.Parent
	}
}

// InOrderTraversal returns the keys in sorted order.
func (bst *BinarySearchTree) InOrderTraversal() []int {
	var result []int
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		result = append(result, node.Key)
		inOrder(node.Right)
	}
	inOrder(bst.Root)
	return result
}

// Main function to demonstrate the BinarySearchTree operations.
func main() {
	bst := &BinarySearchTree{}

	// Insert keys.
	bst.Insert(15)
	bst.Insert(10)
	bst.Insert(20)
	bst.Insert(8)
	bst.Insert(12)

	// In-order traversal.
	fmt.Println("In-order Traversal:", bst.InOrderTraversal())

	// Find minimum and maximum.
	fmt.Println("Minimum:", bst.Minimum(nil).Key)
	fmt.Println("Maximum:", bst.Maximum(nil).Key)

	// Find successor and predecessor.
	successor, _ := bst.Successor(10)
	fmt.Println("Successor of 10:", *successor)
	predecessor, _ := bst.Predecessor(15)
	fmt.Println("Predecessor of 15:", *predecessor)

	// Delete a node and display updated tree.
	_ = bst.Delete(10)
	fmt.Println("In-order Traversal after deleting 10:", bst.InOrderTraversal())
}
