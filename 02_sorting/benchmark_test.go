package sorting

import (
	"algorithms/02_sorting/insertion"
	"algorithms/02_sorting/mergesort"
	"algorithms/02_sorting/selection"
	"algorithms/02_sorting/shell"
	. "algorithms/02_sorting/utils"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSelectionSort(b *testing.B) {
	users := generate()
	b.ResetTimer()
	selection.Sort(users)
}

func BenchmarkInsertionSort(b *testing.B) {
	users := generate()
	b.ResetTimer()
	insertion.Sort(users)
}

func BenchmarkShellSort(b *testing.B) {
	users := generate()
	b.ResetTimer()
	shell.Sort(users)
}

func BenchmarkMergeSort(b *testing.B) {
	users := generate()
	b.ResetTimer()
	mergesort.Sort(users)
}

func BenchmarkMergeSortBU(b *testing.B) {
	users := generate()
	b.ResetTimer()
	mergesort.SortBU(users)
}

func generate() Users {
	n := 100000
	users := make([]User, n)
	names := []string{"Bob", "Kate", "John", "Michael", "Jenny"}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		users[i] = User{
			Name: names[rand.Intn(len(names))],
			Age:  rand.Intn(100),
		}
	}

	return users
}
