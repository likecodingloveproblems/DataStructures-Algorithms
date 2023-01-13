package linkedlist

// Doubly circular linked list
// it accept Node struct and is type generic
type List[T comparable] struct {
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

// AddFirst will add the node to the head of the list
// so the node will become the `Head` of the list
func (l *List[T]) AddFirst(node *Node[T]) {
	defer func() { l.count++ }()
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

// AddLast will add the node to the tail of the list
// so the node will become the `Tail` of the list
func (l *List[T]) AddLast(node *Node[T]) {
	defer func() { l.count++ }()
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

// Pop the first node of the list that is `Head`
func (l *List[T]) PopFirst() (*Node[T], error) {
	if l.IsEmpty() {
		return nil, ListIsEmpty
	}
	defer func() {
		l.count--
	}()
	node := l.Head
	l.Head = l.Head.Next
	l.Head.Previous = l.Tail
	l.Tail.Next = l.Head
	return node, nil
}

// Pop the last node of the list that is the `Tail`
func (l *List[T]) PopLast() (*Node[T], error) {
	if l.IsEmpty() {
		return nil, ListIsEmpty
	}
	defer func() {
		l.count--
	}()
	node := l.Tail
	l.Tail = node.Previous
	l.Tail.Next = l.Head
	l.Head.Previous = l.Tail
	return node, nil
}

// Add new node after the node
// Caution: it won't check the node is in the list or not
// so if the user send a node that is not in the list,
// the list count will become wrong
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

// Add new node before the node
// Caution: it won't check the node is in the list or not
// so if the user send a node that is not in the list,
// the list count will become wrong
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

// Check node is in the list
// it will traverse all the list's node and return bool
// Caution: it will check the node pointer
func (l *List[T]) Contains(node *Node[T]) bool {
	if l.IsEmpty() {
		return false
	}
	n := l.Head
	for {
		if n == node {
			return true
		}
		n = n.Next
		if n == l.Head {
			break
		}
	}
	return false
}

// Check the list contains the value
// It will traverse list nodes and check nodes value with the given value
func (l *List[T]) ContainsValue(value T) bool {
	if l.IsEmpty() {
		return false
	}
	for node := l.Head; ; {
		if node.Value == value {
			return true
		}
		node = node.Next
		if node == l.Head {
			break
		}
	}
	return false
}

// Remove the head of the list
func (l *List[T]) RemoveFirst() error {
	_, err := l.PopFirst()
	return err
}

// Remove the tail node of the list
func (l *List[T]) RemoveLast() error {
	_, err := l.PopLast()
	return err
}
