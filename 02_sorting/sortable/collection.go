package sortable

type User struct {
	Name string
	Age  int
}

type Users []User

func (users Users) Len() int {
	return len(users)
}

func (users Users) Swap(i int, j int) {
	users[i], users[j] = users[j], users[i]
}

func (users Users) Less(i int, j int) bool {
	return users[i].Age < users[j].Age
}
