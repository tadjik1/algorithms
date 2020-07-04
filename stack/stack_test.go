package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Size(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)

	assert.Equal(t, 2, s.Size())
}

func TestStack_IsEmpty(t *testing.T) {
	s := New()

	assert.Equal(t, true, s.IsEmpty())
	s.Push(1)
	assert.Equal(t, false, s.IsEmpty())
	s.Pop()
	assert.Equal(t, true, s.IsEmpty())
}

func TestStack_Push(t *testing.T) {
	s := New()

	s.Push(1)
	s.Push(2)

	assert.Equal(t, 2, s.Pop())
}

func TestStack_Iterate(t *testing.T) {
	s := New()
	input := []int{1, 2, 3}
	for _, value := range input {
		s.Push(value)
	}

	var output []int
	for value := range s.Iterate() {
		output = append([]int{(value).(int)}, output...)
	}

	assert.Equal(t, output, input)
}
