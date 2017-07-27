package ds

type UnionFind struct {
	store []int
	size  []int
	count int
}

// NewUnionFind constructs a new unionfind data structure
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

// Connected checks if two components are connected
func (uf *UnionFind) Connected(p int, q int) bool {
	return uf.Root(p) == uf.Root(q)
}

// Root checks the root of the component
func (uf *UnionFind) Root(p int) int {
	store := uf.store
	for store[p] != p {
		p = store[p]
	}
	return p
}

// Union connects two components together
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

// Store is the getter for store
func (uf *UnionFind) Store() []int {
	return uf.store
}

// Count is the getter for count
func (uf *UnionFind) Count() int {
	return uf.count
}

