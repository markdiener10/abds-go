package abds

import (
	"testing"
)

func TestTopLevelSort(t *testing.T) {

	g := New(ARRAY)
	for idx := 0; idx < 10; idx++ {
		g.Add(idx + 1)
	}

	g.Sort(func(pa *AbdsItem, pb *AbdsItem) bool {
		return pa.Vu() < pb.Vu()
	})

	it := NewIter()
	for g.Iter(it) {
		if it.I() != uint(11-it.Vu()) {
			t.Errorf("Sort Not Working:%d:%d", it.I(), it.Vu())
		}
	}
}
