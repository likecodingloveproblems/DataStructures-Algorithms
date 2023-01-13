package linkedlist

// Doubly circular linked list
// it accept Node struct and is type generic
type List[T any] struct {
	Head  *Node[T]
	Tail  *Node[T]
	count uint
}

// Check list is empty or not, and returns a bool:
//   - false: list is not empty
//   - true: list is empty
func (l *List[T]) IsEmpty() bool {
	return l.count == 0
}

// GetCount is a getter for count
// and will return the count of items in the list
func (l *List[T]) GetCount() uint {
	return l.count
}

// AddFirst will add a node to the first of the list
// the node will become the `Head` of the list
func (l *List[T]) AddFirst(node *Node[T]) {
	defer func() {
		l.count++
	}()
	if l.IsEmpty() {
		l.Head = node
		l.Tail = node
		l.Head.Next = l.Tail
		l.Head.Previous = l.Tail
		l.Tail.Next = l.Head
		l.Tail.Previous = l.Head
	} else if l.Head == l.Tail {
		l.Head = node
		l.Head.Previous = l.Tail
		l.Head.Next = l.Tail
		l.Tail.Previous = l.Head
		l.Tail.Next = l.Head
	} else {
		l.Head.Previous = node
		l.Head = node
		l.Head.Previous = l.Tail
		l.Tail.Next = l.Head
	}
}

func (l *List[T]) AddLast(node *Node[T]) {
	defer func() {
		l.count++
	}()
	if l.IsEmpty() {
		l.Head = node
		l.Tail = node
		l.Head.Next = l.Tail
		l.Head.Previous = l.Tail
		l.Tail.Next = l.Head
		l.Tail.Previous = l.Head
	} else if l.Head == l.Tail {
		l.Tail = node
		l.Tail.Next = l.Head
		l.Tail.Previous = l.Head
		l.Head.Next = l.Tail
		l.Head.Previous = l.Tail
	} else {
		node.Previous = l.Tail
		l.Tail.Next = node
		l.Tail = node
		l.Tail.Next = l.Head
		l.Head.Previous = l.Tail
	}
}

func (l *List[T]) PopFirst() (*Node[T], error) {
	if l.IsEmpty() {
		return nil, ListIsEmpty
	}
	defer func() {
		l.count--
	}()
	node := l.Head
	l.Head = l.Head.Next
	return node, nil
}

func (l *List[T]) PopLast() (*Node[T], error) {
	if l.IsEmpty() {
		return nil, ListIsEmpty
	}
	defer func() {
		l.count--
	}()
	node := l.Tail
	l.Tail = node.Previous
	return node, nil
}

func (l *List[T]) Contains(node *Node[T]) bool {
	if l.Head == node {
		return true
	}
	return false
}

func (l *List[T]) AddAfter(node, newNode *Node[T]) {
	defer func() { l.count++ }()
	newNode.Next = node.Next
	node.Next = newNode
	newNode.Previous = node
	newNode.Next.Previous = newNode
	if l.Tail == node {
		l.Tail = newNode
	}
}

func (l *List[T]) AddBefore(node, newNode *Node[T]) {
	defer func() { l.count++ }()
	newNode.Next = node
	newNode.Previous = node.Previous
	node.Previous = newNode
	newNode.Previous.Next = newNode
	if l.Head == node {
		l.Head = newNode
	}
}
