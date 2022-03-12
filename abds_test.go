package abds

import (
	"testing"
)

func TestTopLevelOperations(t *testing.T) {

	g := New()
	if g == nil {
		t.Error("Unable to allocate abds struct")
	}
	if g.IsArray() {
		t.Error("Default Abds is an array when it should be tag (object) based")
	}
	if g.Is("GCMD", RECURSE) {
		t.Error("Default Abds has a GCMD tag when it should be empty")
	}
	if g.Is("GCMD") {
		t.Error("Default Abds has a GCMD tag when it should be empty")
	}

	if g.Len() > 0 {
		t.Error("Default Abds has elements when it should be empty")
	}

	//Tag syntax with direct item access
	for _, git := range gtestvals {
		g.P("TAG", git)
	}

	if g.T("CMD", RECURSE, NOAUTO) != nil {
		t.Error("Default Abds retrieves elements that should not exist")
	}

}

func TestTopLevelSort(t *testing.T) {

	g := NA(0)
	for idx := 0; idx < 10; idx++ {
		g.Add(idx + 1)
	}

	g.Sort(func(pa *AbdsItem, pb *AbdsItem) bool {
		if pa.Ri() < pb.Ri() {
			return true
		}
		return false
	})

	it := NewIter()
	for g.Iter(it) {
		//fmt.Println(it.index, it.pItem.Ri())
		if it.I() == 10 {
			if it.T().Ri() != 1 {
				t.Errorf("Sort Not Working:%d", it.T().Ri())
			}
		}
	}
}
