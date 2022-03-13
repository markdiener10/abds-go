package abds

import (
	_ "errors"
	_ "fmt"
	"testing"
)

type TestErrorTransform struct {
	what string
}

func TestErrorOperations(t *testing.T) {

	g := New()
	errs := NewErrs()

	//Direct item access with errors added
	gitem := g.G("TAG")
	for _, git := range gtestvals {
		gitem.S(git, errs)
	}

	if errs.Check() {
		t.Error("Default Abds has elements when it should be empty")
	}

	errs.Reset()

	if g.Is("CMD", errs) {
		t.Error("Default Abds retrieves elements that should not exist")
	}

	//Top Level access with errors added
	for _, git := range gtestvals {
		g.S("TAG", git, errs)
	}

	if errs.Check() {
		t.Error("Default Abds has elements when it should be empty")
	}

}

func TestForcedError(t *testing.T) {

	g := New()
	errs := NewErrs()
	gitem := g.G("TAG")

	if errs.Check() {
		t.Error("Default Abds has elements when it should be empty")
	}

	//Initialize structure that does not comform to AbdsTransform interface
	gtrans := &TestErrorTransform{what: "string"}

	gitem.S(gtrans, errs)

	if !errs.Check() {
		t.Error("Passing invalid structure to item put should trigger error")
	}

	//Example of using alternative member function to verify AbdsTransform Interface type
	//The line will fail in the linter if you uncomment it and then it will not compile
	//gitem.PStru(gtrans, errs)

}
