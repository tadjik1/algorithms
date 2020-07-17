package insertion

import (
	. "algorithms/02_sorting/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	users := []User{
		{"Bob", 31},
		{"Kate", 26},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	Sort(users)

	assert.Equal(t, []int{17, 26, 26, 31, 42}, Map(users, func(user User) int {
		return user.Age
	}))

	assert.Equal(t, "Kate", users[1].Name)
	assert.Equal(t, "Jenny", users[2].Name)
}
