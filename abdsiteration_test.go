package abds

import (
	"testing"
)

func TestIteration(t *testing.T) {

	it := NewIter()
	errs := NewErrs()
	g := NA(0)
	for idx := 0; idx < 10; idx++ {
		g.Add(idx+2, errs)
	}
	if g.Len() != 10 {
		t.Error("Abds set to array mode but reports as tag (object) based")
	}

	for g.Iter(it) {
		if it.I() != it.T().Ti() {
			t.Errorf("Abds set to array mode but reports invalid Tag: %d:%d", it.I(), it.T().Ti())
		}
	}

}
