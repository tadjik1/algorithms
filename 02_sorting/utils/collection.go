package utils

type User struct {
	Name string
	Age  int
}

func (user User) Less(other User) bool {
	return user.Age < other.Age
}

type Users []User

func (users Users) Len() int {
	return len(users)
}

func (users Users) Swap(i int, j int) {
	users[i], users[j] = users[j], users[i]
}
