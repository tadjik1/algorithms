package main

import (
	"algorithms/01_fundamentals/percolation/percolation"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type PercolationStats struct {
	runs []float64
}

func (p *PercolationStats) Mean() float64 {
	var sum float64
	for i := 0; i < len(p.runs); i++ {
		sum += p.runs[i]
	}
	return sum / float64(len(p.runs))
}

func (p *PercolationStats) Stddev() float64 {
	mean := p.Mean()
	var sum float64
	for i := 0; i < len(p.runs); i++ {
		sum += math.Pow(p.runs[i]-mean, 2)
	}
	return math.Sqrt(sum / float64(len(p.runs)))
}

func main() {

	args := os.Args[1:]

	n, _ := strconv.Atoi(args[0])
	T, _ := strconv.Atoi(args[1])

	stats := PercolationStats{}
	for i := 0; i < T; i++ {
		p := percolation.New(n)
		rand.Seed(time.Now().UnixNano())
		for !p.Percolates() {
			site := rand.Intn(n * n)
			row := site / n
			col := site % n
			if !p.IsOpen(row, col) {
				p.Open(row, col)
			}
		}
		stats.runs = append(stats.runs, float64(p.NumberOfOpenSites())/float64(n*n))
	}

	fmt.Printf("mean: %.4f, stddev: %.4f\n", stats.Mean(), stats.Stddev())
}
