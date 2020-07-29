package utils

type User struct {
	Name string
	Age  int
}

func (user User) CompareTo(other User) int {
	if user.Age > other.Age {
		return 1
	} else if user.Age < other.Age {
		return -1
	} else {
		return 0
	}
}

func (user User) Less(other User) bool {
	return user.CompareTo(other) < 0
}

type Users []User

func (users Users) Len() int {
	return len(users)
}

func (users Users) Swap(i int, j int) {
	users[i], users[j] = users[j], users[i]
}
