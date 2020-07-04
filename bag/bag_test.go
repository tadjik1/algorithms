package bag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBag_Size(t *testing.T) {
	b := New()

	b.Add(1)
	b.Add(2)

	assert.Equal(t, 2, b.Size())
}

func TestBag_IsEmpty(t *testing.T) {
	b := New()

	assert.Equal(t, true, b.IsEmpty())
	b.Add(1)
	assert.Equal(t, false, b.IsEmpty())
}

func TestBag_Iterate(t *testing.T) {
	b := New()
	input := []int{1, 2, 3}
	for _, value := range input {
		b.Add(value)
	}

	var output []int
	for value := range b.Iterate() {
		output = append([]int{(value).(int)}, output...)
	}

	assert.Equal(t, output, input)
}
