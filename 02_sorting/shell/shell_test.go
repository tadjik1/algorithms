package shell

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type values []int

func (v values) Len() int {
	return len(v)
}

func (v values) Swap(i int, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v values) Less(i int, j int) bool {
	return v[i] < v[j]
}

func TestSort(t *testing.T) {
	v := []int{5, 4, 3, 2, 1}
	Sort(values(v))

	assert.Equal(t, []int{1, 2, 3, 4, 5}, v)
}
