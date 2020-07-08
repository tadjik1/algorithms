package selection

import (
	"algorithms/02_sorting/sortable"
	"algorithms/02_sorting/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	users := []sortable.User{
		{"Bob", 31},
		{"Kate", 26},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	Sort(sortable.Users(users))

	assert.Equal(t, []int{17, 26, 26, 31, 42}, utils.Map(users, func(user sortable.User) int {
		return user.Age
	}))
}
