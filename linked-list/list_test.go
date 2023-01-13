package linkedlist_test

import (
	"testing"

	linkedlist "github.com/likecodingloveproblems/DataStructures-Algorithms/linked-list"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	t.Run("list is empty initially", func(t *testing.T) {
		l := linkedlist.List[int]{}
		got := l.IsEmpty()
		want := true
		assert.Equal(t, want, got)
	})

	t.Run("added list is not empty", func(t *testing.T) {
		l := linkedlist.List[int]{}
		l.AddFirst(&linkedlist.Node[int]{Value: 1})
		got := l.IsEmpty()
		want := false
		assert.Equal(t, want, got)
	})

	t.Run("pop first items", func(t *testing.T) {
		l := linkedlist.List[int]{}
		l.AddFirst(&linkedlist.Node[int]{Value: 1})
		l.PopFirst()
		got := l.IsEmpty()
		want := true
		assert.Equal(t, want, got)
	})
}

func TestAddFirst(t *testing.T) {
	t.Run("add one item to the first of list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		node1 := linkedlist.Node[int]{Value: 1}
		l.AddFirst(&node1)
		assertCircularDoublyLinkedList(t, []int{1}, l)
		node2 := linkedlist.Node[int]{Value: 2}
		l.AddFirst(&node2)
		assertCircularDoublyLinkedList(t, []int{2, 1}, l)
		assert.Equal(t, uint(2), l.GetCount())
	})
}

func TestAddLast(t *testing.T) {
	t.Run("AddLast to an empty list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		node1 := linkedlist.Node[int]{Value: 1}
		l.AddLast(&node1)
		assertCircularDoublyLinkedList(t, []int{1}, l)
		assert.Equal(t, uint(1), l.GetCount())
	})
	t.Run("AddLast will add a node to the tail of list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		node1 := linkedlist.Node[int]{Value: 1}
		node2 := linkedlist.Node[int]{Value: 2}
		l.AddLast(&node1)
		l.AddLast(&node2)
		assertCircularDoublyLinkedList(t, []int{1, 2}, l)
		assert.Equal(t, uint(2), l.GetCount())
	})
}

func TestPopFirst(t *testing.T) {
	t.Run("empty list will return error", func(t *testing.T) {
		l := linkedlist.List[int]{}
		_, err := l.PopFirst()
		assert.Equal(t, linkedlist.ListIsEmpty, err)
		assert.Equal(t, uint(0), l.GetCount())
	})
	t.Run("pop first item of list with only one item", func(t *testing.T) {
		l := linkedlist.List[int]{}
		node := linkedlist.Node[int]{Value: 1}
		l.AddFirst(&node)
		v, err := l.PopFirst()
		assert.Equal(t, nil, err)
		assert.Equal(t, &node, v)
		assert.Equal(t, uint(0), l.GetCount())
	})

	t.Run("pop first node of a multiple node list", func(t *testing.T) {
		nodes := make([]linkedlist.Node[int], 5)
		l := linkedlist.List[int]{}
		for i := 0; i < 5; i++ {
			nodes[i] = linkedlist.Node[int]{Value: i}
			l.AddLast(&nodes[i])
		}
		node, err := l.PopFirst()
		assert.Equal(t, &nodes[0], node)
		assert.Equal(t, nil, err)
		assertCircularDoublyLinkedList(t, []int{1, 2, 3, 4}, l)
	})
}

func TestPopLast(t *testing.T) {
	t.Run("PopLast from empty list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		_, err := l.PopLast()
		assert.Equal(t, linkedlist.ListIsEmpty, err)
		assert.Equal(t, uint(0), l.GetCount())
	})

	t.Run("PopLast item of one item list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		node := &linkedlist.Node[int]{Value: 1}
		l.AddFirst(node)
		gotNode, err := l.PopLast()
		assert.Equal(t, nil, err)
		assert.Equal(t, node, gotNode)
		assert.Equal(t, uint(0), l.GetCount())
	})

	t.Run("PopLast item of list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		for i := 1; i < 4; i++ {
			l.AddLast(&linkedlist.Node[int]{Value: i})
		}
		n, err := l.PopLast()
		assert.Equal(t, nil, err)
		assert.Equal(t, 3, n.Value)
		assert.Equal(t, uint(2), l.GetCount())
		assertCircularDoublyLinkedList(t, []int{1, 2}, l)
	})
}

func TestContains(t *testing.T) {
	t.Run("node is not in list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		node := &linkedlist.Node[int]{Value: 1}
		got := l.Contains(node)
		want := false
		assert.Equal(t, want, got)
	})

	t.Run("node is the head of list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		node := &linkedlist.Node[int]{Value: 1}
		l.AddFirst(node)
		got := l.Contains(node)
		want := true
		assert.Equal(t, want, got)
	})

	t.Run("node is in the middle of list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		nodes := make([]linkedlist.Node[int], 10)
		for i := 0; i < 10; i++ {
			nodes[i] = linkedlist.Node[int]{Value: i}
			l.AddLast(&nodes[i])
		}
		assert.Equal(t, true, l.Contains(&nodes[5]))
		node := &linkedlist.Node[int]{Value: 5} // the list contains the value but the node pointer isn't equal
		assert.Equal(t, false, l.Contains(node))
	})
}

func TestAddAfter(t *testing.T) {
	t.Run("add after a node in normal situation", func(t *testing.T) {
		l := linkedlist.List[int]{}
		node2 := &linkedlist.Node[int]{Value: 2}
		node3 := &linkedlist.Node[int]{Value: 3}
		l.AddLast(&linkedlist.Node[int]{Value: 1})
		l.AddLast(node2)
		l.AddLast(&linkedlist.Node[int]{Value: 4})
		l.AddAfter(node2, node3)
		assertCircularDoublyLinkedList(t, []int{1, 2, 3, 4}, l)
	})
	t.Run("add after tail", func(t *testing.T) {
		l := linkedlist.List[int]{}
		l.AddFirst(&linkedlist.Node[int]{Value: 1})
		node := &linkedlist.Node[int]{Value: 2}
		l.AddAfter(l.Head, node)
		assertCircularDoublyLinkedList(t, []int{1, 2}, l)
	})
}

func TestAddBefore(t *testing.T) {
	t.Run("add before a node", func(t *testing.T) {
		l := linkedlist.List[int]{}
		node2 := &linkedlist.Node[int]{Value: 2}
		node3 := &linkedlist.Node[int]{Value: 3}
		l.AddLast(&linkedlist.Node[int]{Value: 1})
		l.AddLast(node3)
		l.AddLast(&linkedlist.Node[int]{Value: 4})
		l.AddBefore(node3, node2)
		assertCircularDoublyLinkedList(t, []int{1, 2, 3, 4}, l)
	})

	t.Run("add before head", func(t *testing.T) {
		l := linkedlist.List[int]{}
		l.AddFirst(&linkedlist.Node[int]{Value: 1})
		l.AddBefore(l.Head, &linkedlist.Node[int]{Value: 0})
		assertCircularDoublyLinkedList[int](t, []int{0, 1}, l)
	})
}

func TestContainsValue(t *testing.T) {
	t.Run("list is empty", func(t *testing.T) {
		l := linkedlist.List[int]{}
		want := false
		got := l.ContainsValue(1)
		assert.Equal(t, want, got)
	})

	t.Run("list contain value", func(t *testing.T) {
		l := linkedlist.List[int]{}
		l.AddLast(&linkedlist.Node[int]{Value: 1})
		got := l.ContainsValue(1)
		want := true
		assert.Equal(t, want, got)
	})
}

func TestRemoveFirst(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		got := l.RemoveFirst()
		want := linkedlist.ListIsEmpty
		assert.Equal(t, want, got)
	})

	t.Run("remove the first node of a one node list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		l.AddFirst(&linkedlist.Node[int]{Value: 1})
		err := l.RemoveFirst()
		assert.Equal(t, nil, err)
		assert.Equal(t, true, l.IsEmpty())
	})

	t.Run("remove the first node of multiple node list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		l.AddLast(&linkedlist.Node[int]{Value: 1})
		l.AddLast(&linkedlist.Node[int]{Value: 2})
		l.AddLast(&linkedlist.Node[int]{Value: 3})
		l.AddLast(&linkedlist.Node[int]{Value: 4})
		l.RemoveFirst()
		assertCircularDoublyLinkedList(t, []int{2, 3, 4}, l)
	})
}

func TestRemoveLast(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		err := l.RemoveLast()
		assert.Equal(t, linkedlist.ListIsEmpty, err)
	})

	t.Run("remove last of a one node list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		l.AddFirst(&linkedlist.Node[int]{Value: 1})
		err := l.RemoveLast()
		assert.Equal(t, nil, err)
		assert.Equal(t, uint(0), l.GetCount())
	})

	t.Run("remvoe the last node of a multiple node list", func(t *testing.T) {
		l := linkedlist.List[int]{}
		for i := 0; i < 5; i++ {
			l.AddLast(&linkedlist.Node[int]{Value: i})
		}
		err := l.RemoveLast()
		assert.Equal(t, nil, err)
		assertCircularDoublyLinkedList(t, []int{0, 1, 2, 3}, l)
	})
}

func assertCircularDoublyLinkedList[T comparable](t testing.TB, values []T, list linkedlist.List[T]) {
	t.Helper()
	assert.Equal(t, uint(len(values)), list.GetCount())
	assert.Equal(t, values[0], list.Head.Value)
	assert.Equal(t, values[len(values)-1], list.Tail.Value)
	head := list.Head
	node := head
	for index := range values {
		// Check previous node
		if index == 0 {
			assert.Equal(t, values[len(values)-1], node.Previous.Value)
		} else {
			assert.Equal(t, values[index-1], node.Previous.Value)
		}
		// Check node itself
		assert.Equal(t, values[index], node.Value)
		// Check next node
		if index == len(values)-1 {
			assert.Equal(t, values[0], node.Next.Value)
		} else {
			assert.Equal(t, values[index+1], node.Next.Value)
		}
		node = node.Next
	}
}
