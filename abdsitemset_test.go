package abds

import (
	"testing"
)

func TestItemSetOperations(t *testing.T) {

	g := New()

	gitem := g.G("TAG")
	for _, git := range gtestvals {
		gitem.S(git)
	}

}
