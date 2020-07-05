package utils

func Exchange(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}
