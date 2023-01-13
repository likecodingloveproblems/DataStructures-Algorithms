package linkedlist

// List node that contains:
//
//	Value as the node value
//	Previous node pointer
//	Next node pointer
type Node[T any] struct {
	Value    T
	Previous *Node[T]
	Next     *Node[T]
}
