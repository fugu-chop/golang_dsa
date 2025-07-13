package datastructures

/*
The `node` type is intended to be used as the node component of the
linkedList, stack and queue types. Each node has two attributes:

1. A T type; and

2. A pointer to the next `node`.
*/
type node[T any] struct {
	value T
	next  *node[T]
}
