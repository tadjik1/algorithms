package steque

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSteque_Size(t *testing.T) {
	s := New()

	assert.Equal(t, 0, s.Size())
}

func TestSteque_Push(t *testing.T) {
	s := New()

	s.Push(1)

	assert.Equal(t, 1, s.Pop())
}

func TestSteque_Enqueue(t *testing.T) {
	s := New()

	s.Enqueue(1)
	s.Enqueue(2)

	assert.Equal(t, 1, s.Pop())
}

func TestSteque_Iterate(t *testing.T) {
	s := New()

	input := []int{0, 1, 2}

	s.Enqueue(1)
	s.Enqueue(2)
	s.Push(0)

	var output []int
	for value := range s.Iterate() {
		output = append(output, (value).(int))
	}

	assert.Equal(t, output, input)
}
