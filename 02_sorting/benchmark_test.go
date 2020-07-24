package sorting

import (
	"algorithms/02_sorting/insertion"
	"algorithms/02_sorting/mergesort"
	"algorithms/02_sorting/quicksort"
	"algorithms/02_sorting/selection"
	"algorithms/02_sorting/shell"
	. "algorithms/02_sorting/utils"
	"sort"
	"testing"
)

const NumberOfUsers = 100000

func BenchmarkGoSort(b *testing.B) {
	users := GenerateUsers(NumberOfUsers)
	b.ResetTimer()
	sort.Slice(users, func(i, j int) bool {
		return users[i].Less(users[j])
	})
}

func BenchmarkSelectionSort(b *testing.B) {
	users := GenerateUsers(NumberOfUsers)
	b.ResetTimer()
	selection.Sort(users)
}

func BenchmarkInsertionSort(b *testing.B) {
	users := GenerateUsers(NumberOfUsers)
	b.ResetTimer()
	insertion.Sort(users)
}

func BenchmarkShellSort(b *testing.B) {
	users := GenerateUsers(NumberOfUsers)
	b.ResetTimer()
	shell.Sort(users)
}

func BenchmarkMergeSort(b *testing.B) {
	users := GenerateUsers(NumberOfUsers)
	b.ResetTimer()
	mergesort.Sort(users)
}

func BenchmarkMergeSortBU(b *testing.B) {
	users := GenerateUsers(NumberOfUsers)
	b.ResetTimer()
	mergesort.SortBU(users)
}

func BenchmarkQuickSort(b *testing.B) {
	users := GenerateUsers(NumberOfUsers)
	b.ResetTimer()
	quicksort.Sort(users)
}
