package selection

import (
	. "algorithms/02_sorting/utils"
)

func Sort(s Users) {
	for i := 0; i < s.Len(); i++ {
		min := i
		for j := i + 1; j < s.Len(); j++ {
			if s[j].Less(s[min]) {
				min = j
			}
		}
		s.Swap(i, min)
	}
}
