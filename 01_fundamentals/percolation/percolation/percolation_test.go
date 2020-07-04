package percolation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPercolation_IsOpen(t *testing.T) {
	p := New(10)

	for row := 1; row < 9; row++ {
		for col := 0; col < 10; col++ {
			assert.Equal(t, false, p.IsOpen(row, col))
		}
	}

	for col := 0; col < 9; col++ {
		assert.Equal(t, true, p.IsOpen(0, col))
		assert.Equal(t, true, p.IsOpen(9, col))
	}
}

func TestPercolation_Open(t *testing.T) {
	p := New(10)

	for i := 0; i < 10; i++ {
		p.Open(i, i)
	}

	for j := 9; j <= 0; j-- {
		assert.Equal(t, true, p.IsOpen(j, j))
	}
}

func TestPercolation_NumberOfOpenSites(t *testing.T) {
	p := New(10)

	// 2 rows + 2 cells
	assert.Equal(t, 22, p.NumberOfOpenSites())
}

func TestPercolation_IsFull(t *testing.T) {
	p := New(10)

	for i := 0; i < 10; i++ {
		assert.Equal(t, true, p.IsFull(0, i))
	}

	for i := 2; i < 8; i++ {
		p.Open(i, i)
		assert.Equal(t, false, p.IsFull(i, i))
	}

	for i := 1; i < 9; i++ {
		p.Open(i, i)
		p.Open(i, i+1)
		assert.Equal(t, true, p.IsFull(i, i), i)
	}
}

func TestPercolation_Percolates(t *testing.T) {
	p := New(10)

	assert.Equal(t, false, p.Percolates())

	for i := 1; i < 9; i++ {
		p.Open(i, i)
		p.Open(i, i+1)
	}

	assert.Equal(t, true, p.Percolates())
}
