package stack_test

import (
	"github.com/likecodingloveproblems/DataStructures-Algorithms/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_GetCount(t *testing.T) {
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
}

func TestStack_Pop(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		s := stack.Stack[int]{}
		_, err := s.Pop()
		assert.Equal(t, stack.StackIsEmpty, err)
	})
	t.Run("only one item stack", func(t *testing.T) {
		s := stack.Stack[int]{}
		s.Push(1)
		v, err := s.Pop()
		assert.Equal(t, 1, v)
		assert.Equal(t, nil, err)
	})
	t.Run("pop from stack with multiple items", func(t *testing.T) {
		s := stack.Stack[int]{}
		s.Push(1)
		s.Push(2)
		s.Push(3)
		v, err := s.Pop()
		assert.Equal(t, 3, v)
		assert.Equal(t, nil, err)
	})
	t.Run("pop 2 times from stack", func(t *testing.T) {
		s := stack.Stack[int]{}
		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Pop()
		v, err := s.Pop()
		assert.Equal(t, 2, v)
		assert.Equal(t, nil, err)
	})
	t.Run("pop after a pop and push", func(t *testing.T) {
		s := stack.Stack[int]{}
		s.Push(1)
		s.Push(2)
		s.Push(3)
		s.Pop()
		s.Push(4)
		s.Push(5)
		v, err := s.Pop()
		assert.Equal(t, 5, v)
		assert.Equal(t, nil, err)
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

func TestStack_ToArray(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		s := stack.Stack[int]{}
		v := s.ToArray()
		assert.Equal(t, []int{}, v)
	})
	t.Run("check stack with push and to array", func(t *testing.T) {
		s := stack.Stack[int]{}
		s.Push(1)
		s.Push(2)
		items := s.ToArray()
		assert.ElementsMatch(t, []int{1, 2}, items)
	})
	t.Run("push, pop then push again and pop again", func(t *testing.T) {
		s := generateStack(5)
		s.Pop()
		s.Pop()
		s.Push(5)
		s.Push(6)
		s.Push(7)
		s.Push(8)
		s.Pop()
		s.Pop()
		got := s.ToArray()
		want := []int{0, 1, 2, 5, 6}
		assert.ElementsMatch(t, want, got)
	})
}

func generateStack(numberOfItems int) stack.Stack[int] {
	s := stack.Stack[int]{}
	for i := 0; i < numberOfItems; i++ {
		s.Push(i)
	}
	return s
}
