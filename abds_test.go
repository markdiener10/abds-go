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

	//Tag syntax with top level access
	for _, git := range gtestvals {
		g.S("TAG", git)
	}

	if g.Is("CMD", RECURSE) {
		t.Error("Default Abds retrieves elements that should not exist")
	}

}
