package abds

import (
	"testing"
)

func TestTopLevelObjectOperations(t *testing.T) {

	g := New()
	if g.Len() != 0 {
		t.Errorf("Abds set to object mode but reports as non-empty:%d", g.Len())
	}
	g.S("TAG", "VALUE")
	if g.Len() != 1 {
		t.Errorf("Abds set to object mode but reports as non-empty:%d", g.Len())
	}

	gchild := g.Nch("TAGCHILD").S("TAG1", "VALUE1").S("TAG2", "VALUE2")
	if gchild.Len() != 2 {
		t.Errorf("Abds child does not have 3 tags:%d", g.Len())
	}

	if g.IsErr() {
		t.Errorf("Abds detected errors:%d", g.Len())
	}

	//Need to test array

	//Need to test child (abds)

	//Need to test transform

	//garray := g.AN()
	//garray.AN()

}

//func (g *Abds) G(tag string, parms ...interface{}) *AbdsItem {

//Convenience function to add
//func (g *Abds) N(tag string, parms ...interface{}) *Abds {

//func (g *Abds) Copy(src *Abds) error {
