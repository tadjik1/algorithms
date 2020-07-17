package shell

import (
	. "algorithms/02_sorting/utils"
)

func Sort(s Users) {
	// 3x+1 increment sequence:  1, 4, 13, 40, 121, 364, 1093, ...
	h := 1
	for h < s.Len()/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < s.Len(); i++ {
			for j := i; j >= h && s[j].Less(s[j-h]); j -= h {
				s.Swap(j, j-h)
			}
		}
		h /= 3
	}
}
