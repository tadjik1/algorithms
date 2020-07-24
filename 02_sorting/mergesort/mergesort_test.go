package mergesort

import (
	. "algorithms/02_sorting/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	users := GenerateUsers(1000)
	assert.False(t, IsSorted(users))
	Sort(users)
	assert.True(t, IsSorted(users))
}

func TestSortBU(t *testing.T) {
	users := GenerateUsers(1000)
	assert.False(t, IsSorted(users))
	SortBU(users)
	assert.True(t, IsSorted(users))
}
