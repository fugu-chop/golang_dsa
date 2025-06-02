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
exist within the binarySearchTree, Delete is a no-op.
*/
func (b *binarySearchTree) Delete(val int, node *doubleLinkedNode) *doubleLinkedNode {
	// node with value of `val` does not exist
	if node == nil {
		return nil
	}

	if val < node.value {
		node.prev = b.Delete(val, node.prev)
	}

	if val > node.value {
		node.next = b.Delete(val, node.next)
	}

	if node.prev == nil {
		return node.next
	}

	if node.next == nil {
		return node.prev
	}

	node.next = b.lift(node.next, node)

	return node
}

func (b *binarySearchTree) lift(node *doubleLinkedNode, deletionNode *doubleLinkedNode) *doubleLinkedNode {
	if node.prev != nil {
		node.prev = b.lift(node.prev, deletionNode)
		return node
	}

	deletionNode.value = node.value
	return node.next
}
