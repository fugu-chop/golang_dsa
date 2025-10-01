package datastructures

/*
The `doubleLinkedNode` can be used as the node component of
any of the linkedList, stack or queue, but primarily intended
for usage in a binarySearchTree. Each node has the same
attributes as `node` but contains an additional pointer to
the previous `doubleLinkedNode` in the chain.
*/
type doubleLinkedNode struct {
	Value int
	next  *doubleLinkedNode
	prev  *doubleLinkedNode
}

/*
binarySearchTree is an implementation of a Binary Search Tree data structure.
It contains a pointer to a doubleLinkedNode type.
*/
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
			Value: val,
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

	if val < node.Value {
		return b.Search(val, node.prev)
	}

	if val > node.Value {
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
	if val > node.Value {
		if node.next == nil {
			node.next = &doubleLinkedNode{
				Value: val,
			}
			return
		}
		b.Insert(val, node.next)
	}

	// We require the additional conditional check to ensure
	// that node.Value == val is a no-op.
	if val < node.Value {
		if node.prev == nil {
			node.prev = &doubleLinkedNode{
				Value: val,
			}
			return
		}
		b.Insert(val, node.prev)
	}
}

/*
Delete removes a node from the binary search tree. Delete is a no-op if:

  - The value does not exist within the binarySearchTree; or

  - The node being deleted is the root node and it has no child nodes.

Otherwise returns the root node (or the node that replaces the root node).

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
    the deleted node, take the former right child of the successor node and place it where the
    successor node used to be.
*/
func (b *binarySearchTree) Delete(val int, node *doubleLinkedNode) *doubleLinkedNode {
	// node with value of `val` does not exist
	if node == nil {
		return nil
	}

	// Where the value does not match, recursively search for the node for replacement
	if val < node.Value {
		node.prev = b.Delete(val, node.prev)
		return node
	}

	if val > node.Value {
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
with that node's value.

It also eliminates the original successor node by replacing the successor node
with it's right child.

It returns the original node that was passed to it, or nil, if the original node
passed to it had no left children (i.e. it becomes the successor node).
*/
func (b *binarySearchTree) lift(node *doubleLinkedNode, deletionNode *doubleLinkedNode) *doubleLinkedNode {
	// recursively search for the lowest value in the subtree
	if node.prev != nil {
		// eventually replace the "original" successor
		// node with it's original right child
		node.prev = b.lift(node.prev, deletionNode)
		return node
	}

	// When no more left children, this is the successor node
	// We only replace the value, not the reference
	deletionNode.Value = node.Value
	return node.next
}

/*
Traverse collates all the nodes of the tree, starting from the root node
then each subsequent left child recursively until there are no more left child nodes,
then returning to the previous node and collecting each right child recursively.
This is also known as 'pre-order' traversal.
*/

func (b *binarySearchTree) Traverse(node *doubleLinkedNode, results []int) []int {
	if node == nil {
		return results
	}

	results = append(results, node.Value)
	results = b.Traverse(node.prev, results)
	results = b.Traverse(node.next, results)

	return results
}
