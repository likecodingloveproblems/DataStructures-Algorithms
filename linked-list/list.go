package linkedlist

type list[T any] struct {
	Head  Node[T]
	Tail  Node[T]
	count uint
}

func NewList[T any]() list[T] {
	return list[T]{count: 0}
}

func (l *list[T]) IsEmpty() bool {
	return l.count == 0
}

func (l *list[T]) GetCount() uint {
	return l.count
}

func (l *list[T]) AddFirst(node Node[T]) {
	defer func() {
		l.count++
	}()
	if l.IsEmpty() {
		l.Head = node
		l.Tail = node
		l.Head.Next = &l.Tail
		l.Head.Previous = &l.Tail
		l.Tail.Next = &l.Head
		l.Tail.Previous = &l.Head
	} else {
		prevHead := l.Head
		node.Next = &prevHead
		l.Head = node
		l.Head.Previous = &l.Tail
	}
}

func (l *list[T]) LoadFrom(items []T) {
	if !l.IsEmpty() {
		panic("only empty list can be loaded")
	}
	defer func() {
		l.count += uint(len(items))
	}()
}
