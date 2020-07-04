package uf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUF_Count(t *testing.T) {
	u := New(10)

	assert.Equal(t, 10, u.Count())
}

func TestUF_Find(t *testing.T) {
	u := New(10)
	nodes := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for node := range nodes {
		assert.Equal(t, node, u.Find(node))
	}
}

func TestUF_Union(t *testing.T) {
	u := New(10)

	assert.Equal(t, false, u.Connected(0, 1))

	u.Union(0, 2)
	assert.Equal(t, false, u.Connected(0, 1))
	u.Union(2, 1)
	assert.Equal(t, true, u.Connected(0, 1))
}

func TestUF_Find2(t *testing.T) {
	u := New(10)

	u.Union(0, 2)
	u.Union(2, 3)
	u.Union(3, 1)
	u.Union(1, 0)

	for node := range []int{0, 1, 2, 3} {
		assert.Equal(t, 0, u.Find(node))
	}
}

func TestUF_Connected(t *testing.T) {
	u := New(10)

	u.Union(0, 2)
	u.Union(2, 3)
	u.Union(3, 1)
	u.Union(1, 0)

	assert.Equal(t, true, u.Connected(0, 3))
}

func TestUF_Count2(t *testing.T) {
	u := New(10)

	u.Union(0, 1)
	u.Union(1, 2)
	u.Union(2, 3)
	u.Union(3, 4)
	u.Union(4, 5)
	u.Union(5, 6)
	u.Union(6, 7)
	u.Union(7, 8)
	u.Union(8, 9)

	assert.Equal(t, 1, u.Count())
}
