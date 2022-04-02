package abds

import (
	_ "strconv"
	"testing"
)

func TestIterationObject(t *testing.T) {

	g := New()
	it := g.NewIter()
	for idx := 1; idx <= 10; idx++ {
		g.S(idx)
	}
	if g.IsErr() {
		t.Error("Abds triggering errors while adding in array mode")
	}
	if g.Len() != 10 {
		t.Errorf("Abds set to object mode but not adding 10 items:%d", g.Len())
	}
	for g.Iter(it) {
		if it.I() != it.Vu() {
			t.Errorf("Abds set to object mode but reports invalid Tag: %d:%d", it.I(), it.Vu())
		}

	}
}
