package queue_test

import (
	"github.com/likecodingloveproblems/DataStructures-Algorithms/queue"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	t.Run("enqueue to empty queue", func(t *testing.T) {
		q := queue.Queue[int]{}
		q.Enqueue(1)
	})
}

func TestQueue_Dequeue(t *testing.T) {
	type queueSampleData struct {
		values []int
		value  int
		err    error
	}
	samples := []queueSampleData{
		{[]int{}, 0, queue.QueueIsEmpty},
		{[]int{0}, 0, nil},
		{[]int{1, 2, 4, 5}, 1, nil},
		{[]int{10, 2, -1}, 10, nil},
	}
	for _, sampleData := range samples {
		q := generateQueue(sampleData.values)
		v, err := q.Dequeue()
		assert.Equal(t, sampleData.value, v)
		assert.Equal(t, sampleData.err, err)
	}
}

func generateQueue(values []int) queue.Queue[int] {
	q := queue.Queue[int]{}
	for _, value := range values {
		q.Enqueue(value)
	}
	return q
}
