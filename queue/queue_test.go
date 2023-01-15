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

func TestQueue_Contains(t *testing.T) {
	type sampleData struct {
		values []int
		item   int
		result bool
	}
	samples := []sampleData{
		{[]int{}, 0, false},
		{[]int{0}, 0, true},
		{[]int{1, 3, 4, 5}, 2, false},
		{[]int{-10, -1, 0, 10}, 10, true},
		{[]int{-10, -1, 0, 10}, -10, true},
		{[]int{-10, -1, 0, 10}, 1, false},
	}
	for _, sampleData := range samples {
		q := generateQueue(sampleData.values)
		result := q.Contains(sampleData.item)
		assert.Equal(t, sampleData.result, result)
	}
}

func TestQueue_Peek(t *testing.T) {
	type sampleData struct {
		values          []int
		numberOfDequeue int
		expected        int
		err             error
	}
	samples := []sampleData{
		{[]int{}, 0, 0, queue.QueueIsEmpty},
		{[]int{0}, 1, 0, queue.QueueIsEmpty},
		{[]int{0, 3}, 2, 0, queue.QueueIsEmpty},
		{[]int{1, 3}, 0, 1, nil},
		{[]int{-1, 0, 1, 3}, 0, -1, nil},
		{[]int{-1, 0, 1, 3}, 2, 1, nil},
	}
	for _, sample := range samples {
		q := generateQueue(sample.values)
		for i := 0; i < sample.numberOfDequeue; i++ {
			q.Dequeue()
		}
		v, err := q.Peek()
		assert.Equal(t, sample.expected, v)
		assert.Equal(t, sample.err, err)
	}
}

func generateQueue(values []int) queue.Queue[int] {
	q := queue.Queue[int]{}
	for _, value := range values {
		q.Enqueue(value)
	}
	return q
}
