package queue_test

import (
	"github.com/likecodingloveproblems/DataStructures-Algorithms/queue"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	t.Run("enqueue to empty queue", func(t *testing.T) {
		q := queue.Queue[int]{}
		q.Enqueue(1)
	})
}
