package main

import (
	"fmt"
	"math/rand"
	//"os"
	//"strconv"
)

func main() {
	//args := os.Args[1:]

	//n, _ := strconv.Atoi(args[0])
	//T, _ := strconv.Atoi(args[1])
	n := 10
	p := New(n)

	//results := make([]int, 0)
	//
	//for i := 0; i < T; i++ {
	//
	//}

	for p.Percolates() {
		site := rand.Intn(n * n)
		row := site / n
		col := site % n
		if !p.IsOpen(row, col) {
			p.Open(row, col)
		}
	}
	fmt.Printf("percolates! #of open sites: %d, total sites: %d\n", p.NumberOfOpenSites(), n)
}