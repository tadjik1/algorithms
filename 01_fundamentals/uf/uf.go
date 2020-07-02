package uf

type UF struct {
	id []int
	sz []int
	count int
}

func New(n int) *UF {
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sz := make([]int, n)
	for i := range id {
		sz[i] = 1
	}
	return &UF{id, sz, n}
}

func (uf *UF) Count() int {
	return uf.count
}

func (uf *UF) Connected(p int, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UF) Find(p int) int {
	nodes := make([]int, 0)

	for p != uf.id[p] {
		nodes = append(nodes, p)
		p = uf.id[p]
	}

	for _, i := range nodes {
		uf.id[i] = p
	}

	return p
}

func (uf *UF) Union(p int, q int) {
	i := uf.Find(p)
	j := uf.Find(q)
	if i == j {
		return
	}

	if uf.sz[i] < uf.sz[j] {
		uf.id[i] = j
		uf.sz[j] = uf.sz[j] + uf.sz[i]
	} else {
		uf.id[j] = i
		uf.sz[i] = uf.sz[i] + uf.sz[j]
	}
	uf.count = uf.count - 1
}
