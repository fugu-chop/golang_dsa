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
