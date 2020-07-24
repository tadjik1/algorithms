package utils

import (
	"math/rand"
	"time"
)

func IsSorted(users []User) bool {
	for i := 0; i < len(users)-1; i++ {
		if users[i+1].Less(users[i]) {
			return false
		}
	}

	return true
}

func GenerateUsers(n int) Users {
	users := make([]User, n)
	names := []string{"Bob", "Kate", "John", "Michael", "Jenny"}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		users[i] = User{
			Name: names[rand.Intn(len(names))],
			Age:  rand.Intn(100),
		}
	}

	return users
}
