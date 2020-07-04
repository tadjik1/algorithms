package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Size(t *testing.T) {
	q := New()

	q.Enqueue(1)
	q.Enqueue(2)

	assert.Equal(t, 2, q.Size())
}

func TestQueue_IsEmpty(t *testing.T) {
	q := New()

	assert.Equal(t, true, q.IsEmpty())
	q.Enqueue(1)
	assert.Equal(t, false, q.IsEmpty())
	q.Dequeue()
	assert.Equal(t, true, q.IsEmpty())
}

func TestQueue_Enqueue(t *testing.T) {
	q := New()

	q.Enqueue(1)
	q.Enqueue(2)

	assert.Equal(t, 1, q.Dequeue())
}

func TestQueue_Iterate(t *testing.T) {
	q := New()
	input := []int{1, 2, 3}
	for _, value := range input {
		q.Enqueue(value)
	}

	var output []int
	for value := range q.Iterate() {
		output = append(output, (value).(int))
	}

	assert.Equal(t, output, input)
}
