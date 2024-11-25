class Node:
    """Represents a node in the Red-Black Tree."""
    def __init__(self, key, color="red", parent=None, left=None, right=None):
        self.key = key  # The key value of the node.
        self.color = color  # The color of the node: "red" or "black".
        self.parent = parent  # Pointer to the parent node.
        self.left = left  # Pointer to the left child.
        self.right = right  # Pointer to the right child.


class RedBlackTree:
    """Implements a Red-Black Tree with standard operations."""
    def __init__(self):
        self.TNULL = Node(key=None, color="black")  # Sentinel node for leaves.
        self.root = self.TNULL  # Initialize the tree as empty.

    def search(self, key, node=None):
        """Searches for a key in the Red-Black Tree."""
        if node is None:  # If no starting node is specified, start at the root.
            node = self.root
        while node != self.TNULL and node.key != key:
            if key < node.key:  # Traverse to the left subtree.
                node = node.left
            else:  # Traverse to the right subtree.
                node = node.right
        return node  # Return the node if found, or TNULL if not found.

    def left_rotate(self, x):
        """Performs a left rotation to balance the tree."""
        y = x.right  # Set y as the right child of x.
        x.right = y.left  # Reassign the left subtree of y to the right of x.
        if y.left != self.TNULL:
            y.left.parent = x  # Update parent pointer for y's left subtree.
        y.parent = x.parent  # Update y's parent to x's parent.
        if x.parent is None:  # If x was the root, update the root.
            self.root = y
        elif x == x.parent.left:  # If x was a left child, update the parent's left pointer.
            x.parent.left = y
        else:  # Otherwise, update the parent's right pointer.
            x.parent.right = y
        y.left = x  # Make x the left child of y.
        x.parent = y  # Update x's parent pointer to y.

    def right_rotate(self, x):
        """Performs a right rotation to balance the tree."""
        y = x.left  # Set y as the left child of x.
        x.left = y.right  # Reassign the right subtree of y to the left of x.
        if y.right != self.TNULL:
            y.right.parent = x  # Update parent pointer for y's right subtree.
        y.parent = x.parent  # Update y's parent to x's parent.
        if x.parent is None:  # If x was the root, update the root.
            self.root = y
        elif x == x.parent.right:  # If x was a right child, update the parent's right pointer.
            x.parent.right = y
        else:  # Otherwise, update the parent's left pointer.
            x.parent.left = y
        y.right = x  # Make x the right child of y.
        x.parent = y  # Update x's parent pointer to y.

    def insert(self, key):
        """Inserts a new key into the Red-Black Tree."""
        new_node = Node(key, color="red", left=self.TNULL, right=self.TNULL)
        y = None  # Parent pointer.
        x = self.root  # Start at the root.

        while x != self.TNULL:  # Traverse the tree to find the insert location.
            y = x
            if new_node.key < x.key:
                x = x.left
            else:
                x = x.right

        new_node.parent = y  # Assign the parent.
        if y is None:  # If the tree was empty, make the new node the root.
            self.root = new_node
        elif new_node.key < y.key:  # Insert as the left child.
            y.left = new_node
        else:  # Insert as the right child.
            y.right = new_node

        if new_node.parent is None:  # If the new node is the root, color it black.
            new_node.color = "black"
            return

        if new_node.parent.parent is None:  # If the grandparent is None, nothing to fix.
            return

        self._fix_insert(new_node)  # Fix the tree to maintain Red-Black properties.

    def _fix_insert(self, node):
        """Fixes the tree after insertion to maintain Red-Black properties."""
        while node.parent and node.parent.color == "red":
            if node.parent == node.parent.parent.right:  # If parent is the right child.
                uncle = node.parent.parent.left  # Get the uncle node.
                if uncle.color == "red":  # Case 1: Uncle is red.
                    uncle.color = "black"
                    node.parent.color = "black"
                    node.parent.parent.color = "red"
                    node = node.parent.parent  # Move up the tree.
                else:
                    if node == node.parent.left:  # Case 2: Node is left child.
                        node = node.parent
                        self.right_rotate(node)
                    # Case 3: Node is right child.
                    node.parent.color = "black"
                    node.parent.parent.color = "red"
                    self.left_rotate(node.parent.parent)
            else:  # If parent is the left child.
                uncle = node.parent.parent.right  # Get the uncle node.
                if uncle.color == "red":  # Case 1: Uncle is red.
                    uncle.color = "black"
                    node.parent.color = "black"
                    node.parent.parent.color = "red"
                    node = node.parent.parent  # Move up the tree.
                else:
                    if node == node.parent.right:  # Case 2: Node is right child.
                        node = node.parent
                        self.left_rotate(node)
                    # Case 3: Node is left child.
                    node.parent.color = "black"
                    node.parent.parent.color = "red"
                    self.right_rotate(node.parent.parent)
            if node == self.root:  # If we reached the root, stop.
                break
        self.root.color = "black"  # Ensure the root is always black.

    def inorder_traversal(self, node=None, result=None):
        """Returns the in-order traversal of the tree."""
        if node is None:
            node = self.root  # Start at the root if no node is specified.
        if result is None:
            result = []  # Initialize the result list.
        if node != self.TNULL:  # Traverse the tree recursively.
            self.inorder_traversal(node.left, result)
            result.append(node.key)  # Append the current key.
            self.inorder_traversal(node.right, result)
        return result  # Return the sorted list of keys.


# Example usage
if __name__ == "__main__":
    rbt = RedBlackTree()

    # Insert elements.
    for key in [20, 15, 25, 10, 5, 30]:
        rbt.insert(key)

    # In-order traversal.
    print("In-order Traversal:", rbt.inorder_traversal())

    # Search for a key.
    search_key = 15
    found_node = rbt.search(search_key)
    if found_node != rbt.TNULL:
        print(f"Key {search_key} found with color {found_node.color}.")
    else:
        print(f"Key {search_key} not found.")
