package quicksort3way


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
	lt, gt := lo, hi
	partitioningItem := s[lo]

	i := lo+1

	for i <= gt {
		cmp := s[i].CompareTo(partitioningItem)
		if cmp < 0 {
			s.Swap(lt, i)
			lt, i = lt+1, i+1
		} else if cmp > 0 {
			s.Swap(i, gt)
			gt = gt-1
		} else {
			i++
		}
	}

	sort(s, lo, lt-1)
	sort(s, gt+1, hi)
}
