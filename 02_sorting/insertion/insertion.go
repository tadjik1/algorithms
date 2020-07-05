package insertion

import "algorithms/02_sorting/utils"

func Sort(s []int) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			utils.Exchange(s, j, j-1)
		}
	}
}
