package insertion

import "algorithms/02_sorting/sortable"

func Sort(s sortable.Sortable) {
	for i := 1; i < s.Len(); i++ {
		for j := i; j > 0 && s.Less(j, j - 1); j-- {
			s.Swap(j, j - 1)
		}
	}
}
