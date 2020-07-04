package percolation

import (
	"algorithms/01_fundamentals/uf"
)

type Percolation struct {
	uf                *uf.UF
	sites             *[]bool
	n                 int
	numberOfOpenSites int
}

func New(n int) *Percolation {
	// create 2 additional sites which connected to top and bottom
	u := uf.New((n * n) + 2)
	sites := make([]bool, n*n)

	for i := range sites {
		sites[i] = false
	}

	for i := 0; i < n; i++ {
		// connect 1st row to top and open
		u.Union(i, n*n)
		sites[i] = true
		// connect last row to bottom and open
		u.Union((n-1)*n+i, n*n+1)
		sites[(n-1)*n+i] = true
	}

	return &Percolation{uf: u, sites: &sites, n: n, numberOfOpenSites: n*2 + 2}
}

func (p *Percolation) Open(row int, col int) {
	if p.IsOpen(row, col) {
		return
	}

	(*p.sites)[row*p.n+col] = true
	p.numberOfOpenSites = p.numberOfOpenSites + 1

	if row != 0 && p.IsOpen(row-1, col) {
		p.uf.Union((row-1)*p.n+col, row*p.n+col)
	}

	if col != p.n-1 && p.IsOpen(row, col+1) {
		p.uf.Union(row*p.n+col+1, row*p.n+col)
	}

	if row != p.n-1 && p.IsOpen(row+1, col) {
		p.uf.Union((row+1)*p.n+col, row*p.n+col)
	}

	if col != 0 && p.IsOpen(row, col-1) {
		p.uf.Union(row*p.n+col-1, row*p.n+col)
	}
}

func (p *Percolation) IsOpen(row int, col int) bool {
	return (*p.sites)[row*p.n+col]
}

func (p *Percolation) IsFull(row int, col int) bool {
	return p.uf.Connected(row*p.n+col, p.n*p.n)
}

func (p *Percolation) NumberOfOpenSites() int {
	return p.numberOfOpenSites
}

func (p *Percolation) Percolates() bool {
	return p.IsFull(p.n, 1)
}
