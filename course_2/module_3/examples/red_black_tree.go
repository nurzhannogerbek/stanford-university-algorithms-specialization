package main

import "fmt"

// Color constants for the Red-Black Tree nodes.
const (
	Red   = true
	Black = false
)

// Node represents a node in the Red-Black Tree.
type Node struct {
	Key    int   // The key of the node.
	Color  bool  // The color of the node: true for red, false for black.
	Parent *Node // Pointer to the parent node.
	Left   *Node // Pointer to the left child.
	Right  *Node // Pointer to the right child.
}

// RedBlackTree represents the Red-Black Tree.
type RedBlackTree struct {
	Root  *Node // Root node of the tree.
	TNULL *Node // Sentinel node (black).
}

// NewRedBlackTree initializes an empty Red-Black Tree.
func NewRedBlackTree() *RedBlackTree {
	tNull := &Node{Color: Black} // Initialize sentinel node.
	return &RedBlackTree{
		Root:  tNull,
		TNULL: tNull,
	}
}

// Search searches for a key in the Red-Black Tree.
func (rbt *RedBlackTree) Search(key int) *Node {
	current := rbt.Root
	for current != rbt.TNULL && key != current.Key {
		if key < current.Key {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return current
}

// LeftRotate performs a left rotation around the given node.
func (rbt *RedBlackTree) LeftRotate(x *Node) {
	y := x.Right
	x.Right = y.Left
	if y.Left != rbt.TNULL {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		rbt.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

// RightRotate performs a right rotation around the given node.
func (rbt *RedBlackTree) RightRotate(x *Node) {
	y := x.Left
	x.Left = y.Right
	if y.Right != rbt.TNULL {
		y.Right.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		rbt.Root = y
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}
	y.Right = x
	x.Parent = y
}

// Insert inserts a new key into the Red-Black Tree.
func (rbt *RedBlackTree) Insert(key int) {
	newNode := &Node{
		Key:    key,
		Color:  Red,
		Left:   rbt.TNULL,
		Right:  rbt.TNULL,
		Parent: nil,
	}

	y := (*Node)(nil)
	x := rbt.Root

	// Traverse the tree to find the correct position.
	for x != rbt.TNULL {
		y = x
		if newNode.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	newNode.Parent = y
	if y == nil {
		rbt.Root = newNode
	} else if newNode.Key < y.Key {
		y.Left = newNode
	} else {
		y.Right = newNode
	}

	// If the new node is the root, color it black.
	if newNode.Parent == nil {
		newNode.Color = Black
		return
	}

	// If the grandparent is nil, no need to fix.
	if newNode.Parent.Parent == nil {
		return
	}

	rbt.fixInsert(newNode)
}

// fixInsert fixes the Red-Black Tree after an insertion.
func (rbt *RedBlackTree) fixInsert(node *Node) {
	for node.Parent != nil && node.Parent.Color == Red {
		if node.Parent == node.Parent.Parent.Right {
			uncle := node.Parent.Parent.Left
			if uncle != nil && uncle.Color == Red { // Case 1: Uncle is red.
				uncle.Color = Black
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left { // Case 2: Node is a left child.
					node = node.Parent
					rbt.RightRotate(node)
				}
				// Case 3: Node is a right child.
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				rbt.LeftRotate(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Right
			if uncle != nil && uncle.Color == Red { // Case 1: Uncle is red.
				uncle.Color = Black
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right { // Case 2: Node is a right child.
					node = node.Parent
					rbt.LeftRotate(node)
				}
				// Case 3: Node is a left child.
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				rbt.RightRotate(node.Parent.Parent)
			}
		}
		if node == rbt.Root {
			break
		}
	}
	rbt.Root.Color = Black // Ensure the root is always black.
}

// GetInOrderTraversal returns the in-order traversal of the tree as a slice of keys.
func (rbt *RedBlackTree) GetInOrderTraversal() []int {
	var result []int
	var inorder func(node *Node)
	inorder = func(node *Node) {
		if node == rbt.TNULL {
			return
		}
		inorder(node.Left)
		result = append(result, node.Key)
		inorder(node.Right)
	}
	inorder(rbt.Root)
	return result
}

// Main function to demonstrate the Red-Black Tree.
func main() {
	rbt := NewRedBlackTree()

	// Insert elements.
	keys := []int{20, 15, 25, 10, 5, 30}
	for _, key := range keys {
		rbt.Insert(key)
	}

	// Print in-order traversal.
	fmt.Println("In-order Traversal:", rbt.GetInOrderTraversal())

	// Search for a key.
	searchKey := 15
	node := rbt.Search(searchKey)
	if node != rbt.TNULL {
		fmt.Printf("Key %d found with color %s.\n", searchKey, colorToString(node.Color))
	} else {
		fmt.Printf("Key %d not found.\n", searchKey)
	}
}

func colorToString(color bool) string {
	if color {
		return "Red"
	}
	return "Black"
}
