package insertion

import (
	. "algorithms/02_sorting/utils"
)

func Sort(s Users) {
	for i := 1; i < s.Len(); i++ {
		for j := i; j > 0 && s[j].Less(s[j - 1]); j-- {
			s.Swap(j, j-1)
		}
	}
}
