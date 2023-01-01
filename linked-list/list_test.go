package linkedlist_test

import (
	linkedlist "dsa/linked-list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	t.Run("list is empty initially", func(t *testing.T) {
		l := linkedlist.NewList[int]()
		got := l.IsEmpty()
		want := true
		assert.Equal(t, want, got)
	})

	t.Run("add one item to the first of list", func(t *testing.T) {
		l := linkedlist.NewList[int]()
		node1 := linkedlist.Node[int]{Value: 1}
		l.AddFirst(node1)
		assert.Equal(t, node1.Value, l.Head.Value)
		assert.Equal(t, node1.Value, l.Tail.Value)
		node2 := linkedlist.Node[int]{Value: 2}
		l.AddFirst(node2)
		assert.Equal(t, node2.Value, l.Head.Value)
		assert.Equal(t, node1.Value, l.Tail.Value)
	})

	t.Run("load an array to list", func(t *testing.T) {
		l := linkedlist.NewList[int]()
		items := []int{1, 2, 3, 4, 5}
		l.LoadFrom(items)
		node := l.Head
		assert.Equal(t, uint(len(items)), l.GetCount())
		for _, item := range items {
			assert.Equal(t, item, node.Value)
			node = *node.Next
		}
	})
}
