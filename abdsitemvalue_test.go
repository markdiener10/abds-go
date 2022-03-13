package abds

import (
	"testing"
)

func TestItemValueOperations(t *testing.T) {

	g := New()

	gitem := g.G("TAG")
	for _, git := range gtestvals {
		gitem.S(git)
		gitem.V()
		gitem.Vi()
		gitem.Vi8()
		gitem.Vi16()
		gitem.Vi32()
		gitem.Vi64()
		gitem.Vu()
		gitem.Vu8()
		gitem.Vu16()
		gitem.Vu32()
		gitem.Vu64()
		gitem.Vs()
		gitem.Vc64()
		gitem.Vc128()
	}

}
