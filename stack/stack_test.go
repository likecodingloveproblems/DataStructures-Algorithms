package stack_test

import (
	"github.com/likecodingloveproblems/DataStructures-Algorithms/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("get count of empty stack", func(t *testing.T) {
		s := stack.Stack[int]{}
		count := s.GetCount()
		assert.Equal(t, uint(0), count)
	})
	t.Run("get count of stack with 2 item pushed", func(t *testing.T) {
		s := generateStack(2)
		count := s.GetCount()
		assert.Equal(t, uint(2), count)
	})
	t.Run("check stack with push and to array", func(t *testing.T) {
		s := stack.Stack[int]{}
		s.Push(1)
		s.Push(2)
		items := s.ToArray()
		assert.ElementsMatch(t, []int{1, 2}, items)
	})
	t.Run("pop from empty stack", func(t *testing.T) {
		s := stack.Stack[int]{}
		_, err := s.Pop()
		assert.Equal(t, stack.StackIsEmpty, err)
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		s := stack.Stack[int]{}
		_, err := s.Peek()
		assert.Equal(t, stack.StackIsEmpty, err)
	})
	t.Run("peek only pushed stack", func(t *testing.T) {
		s := stack.Stack[int]{}
		s.Push(1)
		v, err := s.Peek()
		assert.Equal(t, 1, v)
		assert.Equal(t, nil, err)
		s.Push(2)
		v, err = s.Peek()
		assert.Equal(t, 2, v)
		assert.Equal(t, nil, err)
	})
	t.Run("peek popped stack", func(t *testing.T) {
		s := stack.Stack[int]{}
		s.Push(1)
		s.Push(2)
		s.Pop()
		v, err := s.Peek()
		assert.Equal(t, 1, v)
		assert.Equal(t, nil, err)
	})
}

func TestStack_IsEmpty(t *testing.T) {
	t.Run("stack initially is empty", func(t *testing.T) {
		s := stack.Stack[int]{}
		got := s.IsEmpty()
		want := true
		assert.Equal(t, want, got)
	})
	t.Run("pushed stack is not empty", func(t *testing.T) {
		s := stack.Stack[int]{}
		s.Push(1)
		got := s.IsEmpty()
		want := false
		assert.Equal(t, want, got)
	})
	t.Run("pop all item", func(t *testing.T) {
		s := stack.Stack[int]{}
		s.Push(1)
		s.Push(2)
		s.Pop()
		assert.Equal(t, false, s.IsEmpty())
		s.Peek()
		assert.Equal(t, false, s.IsEmpty())
		s.Pop()
		assert.Equal(t, true, s.IsEmpty())
	})
}

func generateStack(numberOfItems int) stack.Stack[int] {
	s := stack.Stack[int]{}
	for i := 0; i < numberOfItems; i++ {
		s.Push(i)
	}
	return s
}
