package main

import "algorithms/01_fundamentals/uf"

type Percolation struct {
	uf *uf.UF
	sites *[]bool
	n int
	numberOfOpenSites int
}

func New(n int) *Percolation {
	// create 2 additional sites which connected to top and bottom
	u := uf.New((n * n) + 2)
	sites := make([]bool, n * n)

	for i := range sites {
		sites[i] = false
	}

	for i := 0; i < n; i++ {
		// connect to top
		u.Union(i, n * n)
		// connect to bottom
		u.Union(n * (n - 1) + i, n * n + 1)
	}

	return &Percolation{uf: u, sites: &sites, n: n}
}

func (p *Percolation) Open(row int, col int) {
	if p.IsOpen(row, col) {
		return
	}

	(*p.sites)[row * (p.n - 1) + col] = true

	if row != 0 && p.IsOpen(row - 1, col) {
		p.uf.Union((row - 1) * (p.n - 1) + col, row * (p.n - 1) + col)
	}

	if col != p.n && p.IsOpen(row, col + 1) {
		p.uf.Union(row * (p.n - 1) + col + 1, row * (p.n - 1) + col)
	}

	if row != p.n && p.IsOpen(row + 1, col) {
		p.uf.Union(row * p.n + col, row * (p.n - 1) + col)
	}

	if col != 0 && p.IsOpen(row, col - 1) {
		p.uf.Union(row * (p.n - 1) + col - 1, row * (p.n - 1) + col)
	}
}

func (p *Percolation) IsOpen(row int, col int) bool {
	return (*p.sites)[row * (p.n - 1) + col]
}

func (p *Percolation) IsFull(row int, col int) bool {
	return p.uf.Connected(row * (row - 1) + col, p.n * p.n)
}

func (p *Percolation) NumberOfOpenSites() int {
	return p.numberOfOpenSites
}

func (p *Percolation) Percolates() bool {
	return p.uf.Connected(p.n * p.n, p.n * p.n + 1)
}
