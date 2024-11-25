class TreeNode:
    def __init__(self, key):
        self.key = key
        self.left = None
        self.right = None
        self.parent = None
        self.size = 1  # Size of the subtree, including this node.


class BinarySearchTree:
    def __init__(self):
        self.root = None

    def insert(self, key):
        """Insert a key into the tree, ignoring duplicates."""
        if self.search(key):  # Skip if the key already exists.
            return
        new_node = TreeNode(key)
        if not self.root:
            self.root = new_node  # Set root if tree is empty.
            return
        current = self.root
        while current:
            current.size += 1  # Update the size of each ancestor node.
            if key < current.key:
                if not current.left:
                    current.left = new_node
                    new_node.parent = current
                    return
                current = current.left
            else:
                if not current.right:
                    current.right = new_node
                    new_node.parent = current
                    return
                current = current.right

    def search(self, key):
        """Search for a key in the tree and return its node, or None if not found."""
        current = self.root
        while current and current.key != key:
            current = current.left if key < current.key else current.right
        return current

    def minimum(self, node=None):
        """Return the node with the minimum key in the tree or subtree."""
        if node is None:
            node = self.root
        while node and node.left:
            node = node.left  # Traverse left until null.
        return node

    def maximum(self, node=None):
        """Return the node with the maximum key in the tree or subtree."""
        if node is None:
            node = self.root
        while node and node.right:
            node = node.right  # Traverse right until null.
        return node

    def successor(self, key):
        """Return the key of the next largest element after the given key."""
        node = self.search(key)
        if not node:
            return None
        if node.right:
            return self.minimum(node.right).key  # The successor is the minimum in the right subtree.
        current = node.parent
        while current and node == current.right:
            node = current
            current = current.parent
        return current.key if current else None  # Return successor or None if not found.

    def predecessor(self, key):
        """Return the key of the previous smaller element before the given key."""
        node = self.search(key)
        if not node:
            return None
        if node.left:
            return self.maximum(node.left).key  # The predecessor is the maximum in the left subtree.
        current = node.parent
        while current and node == current.left:
            node = current
            current = current.parent
        return current.key if current else None  # Return predecessor or None if not found.

    def delete(self, key):
        """Delete the node with the given key from the tree."""
        node = self.search(key)
        if not node:
            raise ValueError(f"Key {key} not found in the tree.")

        # Update the sizes of ancestor nodes.
        current = node
        while current:
            current.size -= 1
            current = current.parent

        if not node.left and not node.right:  # Case 1: No children.
            self._transplant(node, None)
        elif not node.right:  # Case 2: One child (left).
            self._transplant(node, node.left)
        elif not node.left:  # Case 2: One child (right).
            self._transplant(node, node.right)
        else:  # Case 3: Two children.
            successor = self.minimum(node.right)
            if successor.parent != node:
                self._transplant(successor, successor.right)
                successor.right = node.right
                successor.right.parent = successor
            self._transplant(node, successor)
            successor.left = node.left
            successor.left.parent = successor
            successor.size = node.size  # Update the size of the new root.

    def _transplant(self, u, v):
        """Replace the subtree rooted at u with the subtree rooted at v."""
        if not u.parent:
            self.root = v  # If u is root, replace root.
        elif u == u.parent.left:
            u.parent.left = v
        else:
            u.parent.right = v
        if v:
            v.parent = u.parent  # Update parent pointer of v.

    def select(self, i):
        """Return the key of the i-th smallest element in the tree (1-based index)."""
        def _select(node, i):
            if not node:
                return None
            left_size = node.left.size if node.left else 0
            if i == left_size + 1:
                return node
            elif i <= left_size:
                return _select(node.left, i)
            else:
                return _select(node.right, i - left_size - 1)
        node = _select(self.root, i)
        return node.key if node else None

    def rank(self, key):
        """Return the number of keys less than or equal to the given key."""
        node = self.search(key)
        if not node:
            return None
        rank = 1
        if node.left:
            rank += node.left.size
        while node != self.root:
            if node == node.parent.right:
                rank += 1
                if node.parent.left:
                    rank += node.parent.left.size
            node = node.parent
        return rank

    def inorder_traversal(self):
        """Return a list of keys in sorted order."""
        result = []

        def _inorder(node):
            if not node:
                return
            _inorder(node.left)
            result.append(node.key)
            _inorder(node.right)

        _inorder(self.root)
        return result


if __name__ == "__main__":
    # Example usage of BinarySearchTree.
    bst = BinarySearchTree()
    bst.insert(15)
    bst.insert(10)
    bst.insert(20)
    bst.insert(8)
    bst.insert(12)

    print("In-order Traversal:", bst.inorder_traversal())
    print("Minimum:", bst.minimum().key)
    print("Maximum:", bst.maximum().key)
    print("Rank of 12:", bst.rank(12))
    print("Select 3rd smallest:", bst.select(3))
    print("Successor of 10:", bst.successor(10))
    print("Predecessor of 15:", bst.predecessor(15))

    bst.delete(10)
    print("In-order Traversal after deleting 10:", bst.inorder_traversal())
