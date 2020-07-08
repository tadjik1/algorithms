package utils

import "algorithms/02_sorting/sortable"

func Map(vs []sortable.User, f func(user sortable.User) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
