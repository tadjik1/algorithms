package quicksort

import (
	. "algorithms/02_sorting/utils"
	"math/rand"
)

func Sort(s Users) {
	// eliminate dependence on input
	rand.Shuffle(s.Len(), s.Swap)

	sort(s, 0, s.Len()-1)
}

func sort(s Users, lo, hi int) {
	if hi <= lo {
		return
	}
	j := partition(s, lo, hi)
	sort(s, lo, j-1)
	sort(s, j+1, hi)
}

func partition(s Users, lo, hi int) int {
	// left and right scan indices
	i, j := lo, hi

	partitioningItem := s[lo]

	for {
		for !partitioningItem.Less(s[i]) {
			i++
			if i == hi {
				break
			}
		}
		for !s[j].Less(partitioningItem) {
			j--
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		s.Swap(i, j)
	}

	s.Swap(lo, j)

	return j
}
