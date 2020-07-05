package insertion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	Sort(values)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, values)
}
