package abds

import (
	"testing"
)

func TestItemValueOperations(t *testing.T) {

	g := New()
	gitem := g.g("TAG")

	for _, git := range gatb {
		gitem.S(git)
		gitem.vi()
		gitem.vi8()
		gitem.vi16()
		gitem.vi32()
		gitem.vi64()
		gitem.vu()
		gitem.vu8()
		gitem.vu16()
		gitem.vu32()
		gitem.vu64()
		gitem.vs()
		gitem.vc64()
		gitem.vc128()
	}

}
