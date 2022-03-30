package abds

import (
	"strconv"
	"testing"
)

func TestIterationObject(t *testing.T) {
	it := NewIter()
	g := New()
	gs := ""
	for idx := 0; idx < 10; idx++ {
		gs = "Tag" + strconv.FormatUint(uint64(idx+1), 10)
		g.S(gs, idx+1)
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

func TestIterationArray(t *testing.T) {
	it := NewIter()
	g := New(ARRAY)
	for idx := 0; idx < 10; idx++ {
		g.Add(idx + 1)
	}
	if g.IsErr() {
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
