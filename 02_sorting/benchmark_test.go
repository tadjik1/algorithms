package sorting

import (
	"algorithms/02_sorting/insertion"
	"algorithms/02_sorting/selection"
	"algorithms/02_sorting/shell"
	"algorithms/02_sorting/sortable"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkSelectionSort(b *testing.B) {
	users := generate(10000)
	b.ResetTimer()
	selection.Sort(users)
}

func BenchmarkInsertionSort(b *testing.B) {
	users := generate(10000)
	b.ResetTimer()
	insertion.Sort(users)
}

func BenchmarkShellSort(b *testing.B) {
	users := generate(10000)
	b.ResetTimer()
	shell.Sort(users)
}

func generate(n int) sortable.Sortable {
	users := make([]sortable.User, n)
	names := []string{"Bob", "Kate", "John", "Michael", "Jenny"}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		users[i] = sortable.User{
			Name: names[rand.Intn(len(names))],
			Age:  rand.Intn(100),
		}
	}

	return sortable.Users(users)
}
