package abds

import (
	_ "errors"
	_ "fmt"
	"testing"
)

func TestErrorOperations(t *testing.T) {

	g := New()
	errs := NewErrs()

	//Direct item access with errors added
	gitem := g.T("TAG")
	for _, git := range gtestvals {
		gitem.P(git,errs)
	}

	if errs.Check() {
		t.Error("Default Abds has elements when it should be empty")
	}

	errs.Reset()

	if g.T("CMD", RECURSE, NOAUTO, errs) != nil {
		t.Error("Default Abds retrieves elements that should not exist")
	}

	//Top Level access with errors added
	for _, git := range gtestvals {
		g.P("TAG",git,errs)
	}

	if errs.Check() {
		t.Error("Default Abds has elements when it should be empty")
	}

}
