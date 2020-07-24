package quicksort

import (
	. "algorithms/02_sorting/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	users := GenerateUsers(100)
	assert.False(t, IsSorted(users))
	Sort(users)
	assert.True(t, IsSorted(users))
}
