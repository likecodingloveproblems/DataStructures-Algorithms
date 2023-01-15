package queue

import linkedlist "github.com/likecodingloveproblems/DataStructures-Algorithms/linked-list"

type Queue[T comparable] struct {
	count uint
	items linkedlist.List[T]
}

func (q *Queue[T]) Enqueue(item T) {
	q.items.AddLast(&linkedlist.Node[T]{Value: item})
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var noop T
		return noop, QueueIsEmpty
	}
	node, err := q.items.PopFirst()
	if err != nil {
		var noop T
		return noop, err
	}
	return node.Value, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var noop T
		return noop, QueueIsEmpty
	}
	return q.items.Head.Value, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.items.IsEmpty()
}

func (q *Queue[T]) Contains(item T) bool {
	return q.items.ContainsValue(item)
}
