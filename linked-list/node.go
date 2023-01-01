package linkedlist

type INode[T any] interface {
	Equals(INode[T])
}

type Node[T any] struct {
	Value    T
	Previous *Node[T]
	Next     *Node[T]
}
