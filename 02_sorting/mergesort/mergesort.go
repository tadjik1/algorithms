package mergesort

import . "algorithms/02_sorting/utils"

// top-down ("default") mergesort
func Sort(s Users) {
	// allocate space just once for auxiliary array for merges
	temp := make(Users, s.Len())
	sort(s, temp, 0, s.Len() - 1)
}

// bottom-up mergesort
func SortBU(s Users) {
	temp := make(Users, s.Len())
	for l := 1; l < s.Len(); l *= 2 {
		for lo := 0; lo < s.Len() - l; lo += l + l {
			merge(s, temp, lo, lo + l - 1, min(lo + l + l - 1, s.Len() - 1))
		}
	}
}

func sort(s, temp Users, lo, hi int) {
	if hi <= lo {
		return
	}

	mid := lo + (hi - lo) / 2

	// sort left half
	sort(s, temp, lo, mid)
	// sort right half
	sort(s, temp, mid + 1, hi)
	// merge results
	merge(s, temp, lo, mid, hi)
}

func merge(s, temp Users, lo, mid, hi int) {
	i, j := lo, mid + 1
	var k int

	// copy s[lo..hi] to temp[lo..hi]
	for k = lo; k <= hi; k++ {
		temp[k] = s[k]
	}

	// merge back to s[lo..hi]
	for k = lo; k <= hi; k++ {
		if i > mid {
			s[k] = temp[j]
			j++
		} else if j > hi {
			s[k] = temp[i]
			i++
		} else if temp[j].Less(temp[i]) {
			s[k] = temp[j]
			j++
		} else {
			s[k] = temp[i]
			i++
		}
	}
}

func min(x, y int) int {
	if y < x {
		return y
	}
	return x
}