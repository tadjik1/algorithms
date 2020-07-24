package main

import (
	"algorithms/02_sorting/quicksort"
	. "algorithms/02_sorting/utils"
)

func main() {
	users := GenerateUsers(10)
	quicksort.Sort(users)
}
