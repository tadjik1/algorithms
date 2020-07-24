# Sorting

```bash
goos: darwin
goarch: amd64
pkg: algorithms/02_sorting
BenchmarkGoSort
BenchmarkGoSort-8          	1000000000	         0.0177 ns/op
BenchmarkSelectionSort
BenchmarkSelectionSort-8   	       1	         9551462977 ns/op
BenchmarkInsertionSort
BenchmarkInsertionSort-8   	       1	         6700196672 ns/op
BenchmarkShellSort
BenchmarkShellSort-8       	1000000000	         0.0100 ns/op
BenchmarkMergeSort
BenchmarkMergeSort-8       	1000000000	         0.0176 ns/op
BenchmarkMergeSortBU
BenchmarkMergeSortBU-8     	1000000000	         0.0175 ns/op
BenchmarkQuickSort
BenchmarkQuickSort-8       	1000000000	         0.0571 ns/op
PASS
```