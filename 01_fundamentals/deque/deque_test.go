package deque

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeque_PushLeft(t *testing.T) {
	d := New()

	d.PushLeft(1)
	d.PushLeft(2)

	assert.Equal(t, 2, d.PopLeft())
	assert.Equal(t, 1, d.PopLeft())
}

func TestDeque_PushRight(t *testing.T) {
	d := New()

	d.PushRight(1)
	d.PushRight(2)

	assert.Equal(t, 2, d.PopRight())
	assert.Equal(t, 1, d.PopRight())
}

func TestDeque_PopLeft(t *testing.T) {
	d := New()

	d.PushRight(1)
	d.PushRight(2)

	assert.Equal(t, 1, d.PopLeft())
	assert.Equal(t, 2, d.PopLeft())
}

func TestDeque_PopRight(t *testing.T) {
	d := New()

	d.PushLeft(1)
	d.PushLeft(2)

	assert.Equal(t, 1, d.PopRight())
	assert.Equal(t, 2, d.PopRight())
}
