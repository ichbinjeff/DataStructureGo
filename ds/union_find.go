package ds

type UnionFind struct {
	store []int
	size  []int
	count int
}

func NewUnionFind(size int) UnionFind {
	store := make([]int, size, size)
	for i := 0; i < size; i++ {
		store[i] = i
	}
	return UnionFind{
		store: store,
		size: make([]int, size, size),
		count: size,
	}
}

// 1  9
// 9  9
func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.Root(p) == uf.Root(q)
}

// 1, 2, 3, 4, 5
func (uf *UnionFind) Root(p int) int {
	store := uf.store
	for store[p] != p {
		p = store[p]
	}
	return p
}

func (uf *UnionFind) Union(p int, q int) {
	rootP := uf.Root(p)
	rootQ := uf.Root(q)

	if rootP == rootQ {
		return
	}

	if uf.size[rootP] < uf.size[rootQ] {
		uf.store[rootP] = rootQ
		uf.size[rootP] += uf.size[rootQ]
	} else {
		uf.store[rootQ] = rootP
		uf.size[rootQ] += uf.size[rootP]
	}
	uf.count--
}

func (uf *UnionFind) Store() []int {
	return uf.store
}

func (uf *UnionFind) Count() int {
	return uf.count
}

