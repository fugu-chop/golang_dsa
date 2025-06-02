package datastructures

type binarySearchTree struct {
	CurrentNode *doubleLinkedNode
}

/*
BinarySearchTree returns a pointer to a binarySearchTree type
with the value of the currentNode set to `val`.
*/
func BinarySearchTree(val int) *binarySearchTree {
	return &binarySearchTree{
		CurrentNode: &doubleLinkedNode{
			value: val,
		},
	}
}

/*
Search returns a pointer to a doubleLinkedNode if a node in the
tree contains the value of `val`. If a node does not contain `val`,
`nil` is returned.
*/
func (b *binarySearchTree) Search(val int, node *doubleLinkedNode) *doubleLinkedNode {
	if node == nil {
		return nil
	}

	if val < node.value {
		return b.Search(val, node.prev)
	}

	if val > node.value {
		return b.Search(val, node.next)
	}

	return node
}

/*
Insert traverses the binary search tree and inserts a reference to
a doubleLinkedNode to the appropriate node based on provided `val`.
If a node with the value of `val` already exists, Insert is a no-op.
*/
func (b *binarySearchTree) Insert(val int, node *doubleLinkedNode) {
	if val > node.value {
		if node.next == nil {
			node.next = &doubleLinkedNode{
				value: val,
			}
			return
		}
		b.Insert(val, node.next)
	}

	// We require the additional conditional check to ensure
	// that node.value == val is a no-op.
	if val < node.value {
		if node.prev == nil {
			node.prev = &doubleLinkedNode{
				value: val,
			}
			return
		}
		b.Insert(val, node.prev)
	}
}

/*
Delete removes a node from the binary search tree. If the value does not
exist within the binarySearchTree, Delete is a no-op. Otherwise returns the
root node (or the node that replaces the root node).

Delete follows these rules:

  - If the node being deleted has no children, simply delete it.

  - If the node being deleted has one child, delete the node and plug the child
    into the spot where the deleted node was.

  - When deleting a node with two children, replace the deleted node with the successor node.
    The successor node is the child node whose value is the least of all values that are greater
    than the deleted node.

  - To find the successor node: visit the right child of the deleted value, and
    then keep on visiting the left child of each subsequent child until there
    are no more left children. The bottom value is the successor node.

  - If the successor node has a right child, after plugging the successor node into the spot of
    the deleted node, take the former right child of the successor node and turn it into the
    left child of the former parent of the successor node.
*/
func (b *binarySearchTree) Delete(val int, node *doubleLinkedNode) *doubleLinkedNode {
	// node with value of `val` does not exist
	if node == nil {
		return nil
	}

	// Where the value does not match, recursively search for the node for replacement
	if val < node.value {
		node.prev = b.Delete(val, node.prev)
		return node
	}

	if val > node.value {
		node.next = b.Delete(val, node.next)
		return node
	}

	// Node value matches
	// Replace the node with it's child
	if node.prev == nil {
		return node.next
	}

	if node.next == nil {
		return node.prev
	}

	// Node has two children
	node.next = b.lift(node.next, node)

	return node
}

/*
lift finds the node with the lowest value that is still larger than the value
of `deletionNode` (the "successor" node) and replaces the value of `deletionNode`
with that node's value. It also eliminates the original successor node by
replacing the successor node with it's right child for the succesor node's parent.
*/
func (b *binarySearchTree) lift(node *doubleLinkedNode, deletionNode *doubleLinkedNode) *doubleLinkedNode {
	// recursively search for the lowest value in the subtree
	if node.prev != nil {
		node.prev = b.lift(node.prev, deletionNode)
		return node
	}

	// When no more left children, this is the successor node
	deletionNode.value = node.value
	return node.next
}
