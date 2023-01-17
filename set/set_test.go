package set_test

import (
	"github.com/likecodingloveproblems/DataStructures-Algorithms/set"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("add two value and check contains", func(t *testing.T) {
		s := set.NewSet[int]()
		s.Add(1)
		s.Add(2)
		got := s.Contains(1)
		want := true
		assert.Equal(t, want, got)
	})

	t.Run("empty set contains check", func(t *testing.T) {
		s := set.NewSet[int]()
		got := s.Contains(1)
		want := false
		assert.Equal(t, want, got)
	})
	t.Run("test remove", func(t *testing.T) {
		s := set.NewSet[int]()
		s.Add(1)
		s.Remove(1)
		got := s.Contains(1)
		want := false
		assert.Equal(t, want, got)
	})
	t.Run("test is empty", func(t *testing.T) {
		s := set.NewSet[int]()
		assert.Equal(t, true, s.IsEmpty())
		s.Add(1)
		assert.Equal(t, false, s.IsEmpty())
		s.Add(2)
		assert.Equal(t, false, s.IsEmpty())
		s.Clear()
		assert.Equal(t, true, s.IsEmpty())
	})
}
