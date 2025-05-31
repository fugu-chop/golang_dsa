package datastructures

type binarySearchTree struct {
	currentNode *doubleLinkedNode
}

/*
BinarySearchTree returns a pointer to a binarySearchTree type
with the value of the currentNode set to `val`.
*/
func BinarySearchTree(val int) *binarySearchTree {
	return &binarySearchTree{
		currentNode: &doubleLinkedNode{
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
	if node == nil || node.value == val {
		return node
	}

	if val < node.value {
		return b.Search(val, node.prev)
	}

	if val > node.value {
		return b.Search(val, node.next)
	}

	return nil
}

/*
Insert traverses the binary search tree and inserts a reference to
a doubleLinkedNode to the appropriate node based on provided `val`.
If a node with the value of `val` already exists, Insert is a no-op.
*/
func (b *binarySearchTree) Insert(val int, node *doubleLinkedNode) {
	if node.value < val {
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
	if node.value > val {
		if node.prev == nil {
			node.prev = &doubleLinkedNode{
				value: val,
			}
			return
		}
		b.Insert(val, node.prev)
	}
}
