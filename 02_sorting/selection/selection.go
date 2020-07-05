package selection

import "algorithms/02_sorting/utils"

func Sort(s []int) {
	for i := 0; i < len(s); i++ {
		min := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		utils.Exchange(s, i, min)
	}
}
