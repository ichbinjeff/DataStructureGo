package ds

import "testing"

type pair struct {
	x int
	y int
}

//TODO: TEST BALANCED
func TestUnionFindConnected(t *testing.T) {
	tests := []struct {
		pairs []pair
	}{
		{
			[]pair{
				{1,2},
				{3,4},
			},
		},
		{
			[]pair{
				{1,2},
				{2,4},
				{4,5},
			},
		},
		{
			[]pair{
				{1,2},
			},
		},

	}

	for _, test := range tests {
		var uf = NewUnionFind(10)
		for _, pair := range test.pairs {
			uf.Union(pair.x, pair.y)
			if !uf.Connected(pair.x, pair.y) {
				t.Fatal("x, y should connect but does not")
			}
		}
	}
}

func TestUnionFindCount(t *testing.T) {
	tests := []struct {
		pairs []pair
		count []int
	}{
		{
			[]pair{
				{1,2},
				{3,4},
			},
			[]int{9,8},
		},
		{
			[]pair{
				{1,2},
				{2,4},
				{4,5},
			},
			[]int{9,8,7},
		},
		{
			[]pair{
				{1,2},
			},
			[]int{9},
		},

	}

	for _, test := range tests {
		var uf = NewUnionFind(10)
		if uf.Count() != 10 {
			t.Fatal("fail to get the correct size")
		}

		for i, pair := range test.pairs {
			uf.Union(pair.x, pair.y)
			if uf.Count() != test.count[i] {
				t.Fatal("fail to get the correct size")
			}
		}
	}
}


func TestUnionFindRoot(t *testing.T) {
	tests := []struct {
		pairs []pair
	}{
		{
			[]pair{
				{1,2},
				{3,4},
			},
		},
		{
			[]pair{
				{1,2},
				{2,4},
				{4,5},
			},
		},
		{
			[]pair{
				{1,2},
			},
		},

	}

	for _, test := range tests {
		var uf = NewUnionFind(10)
		for _, pair := range test.pairs {
			uf.Union(pair.x, pair.y)
			if uf.Root(pair.x) != uf.Root(pair.y) {
				t.Fatal("connected items should point to the same root")
			}
		}
	}
}