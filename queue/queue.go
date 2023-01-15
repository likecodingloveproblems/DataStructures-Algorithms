package queue

import linkedlist "github.com/likecodingloveproblems/DataStructures-Algorithms/linked-list"

type Queue[T comparable] struct {
	count uint
	items linkedlist.List[T]
}

func (q *Queue[T]) Enqueue(item T) {
	q.items.AddFirst(&linkedlist.Node[T]{Value: item})
}
