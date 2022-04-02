package abds

import (
	"testing"
)

func TestTopLevelSort(t *testing.T) {

	g := New()
	for idx := 0; idx < 10; idx++ {
		g.S(idx + 1)
	}

	g.Sort(func(pa *AbdsIter, pb *AbdsIter) bool {
		return pa.Vu() < pb.Vu()
	})

	it := g.NewIter()
	for g.Iter(it) {
		if it.I() != uint(11-it.Vu()) {
			t.Errorf("Sort Not Working:%d:%d", it.I(), it.Vu())
		}
	}

}
