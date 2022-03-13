package abds

import (
	"testing"
)

func TestIteration(t *testing.T) {
	it := NewIter()
	errs := NewErrs()
	g := New(ARRAY)
	for idx := 0; idx < 10; idx++ {
		g.Add(idx+1, errs)
	}
	if errs.Check() {
		t.Error("Abds triggering errors while adding in array mode")
	}
	if g.Len() != 10 {
		t.Errorf("Abds set to array mode but not adding 10 items:%d", g.Len())
	}
	for g.Iter(it) {
		if it.I() != it.Vu() {
			t.Errorf("Abds set to array mode but reports invalid Tag: %d:%d", it.I(), it.Vu())
		}
	}
}
