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

	//Direct item access with errors added
	gitem := g.G("TAG")
	for _, git := range gtestvals {
		gitem.S(git, g.errs)
	}

	if g.IsErr() {
		t.Error("Default Abds has elements when it should be empty")
	}

	if g.Is("CMD") {
		t.Error("Default Abds retrieves elements that should not exist")
	}

	//Top Level access with errors added
	for _, git := range gtestvals {
		g.S("TAG", git)
	}

	if g.IsErr() {
		t.Error("Loading values triggered errors")
	}

}

func TestForcedError(t *testing.T) {

	g := New()
	if g.IsErr() {
		t.Error("Default Abds has errors when it should not")
	}

	//Initialize structure that does not comform to AbdsTransform interface
	gtrans := &TestErrorTransform{what: "string"}

	g.S("TAG", gtrans)
	if !g.IsErr() {
		t.Error("Passing invalid structure should trigger error")
	}
	gitem := g.G("TAG")
	gitem.S(gtrans, g.Errs())
	if !g.IsErr() {
		t.Error("Passing invalid structure to item put should trigger error")
	}

	//Example of using alternative member function to verify AbdsTransform Interface type
	//The line will fail in the linter if you uncomment it and then it will not compile
	//gitem.PStru(gtrans, errs)

}
