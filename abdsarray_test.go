package abds

import (
	"testing"
)

func TestArrayOperations(t *testing.T) {

	//Array operations
	g := NA(0)
	if !g.IsArray() {
		t.Error("Abds set to array mode but reports as tag (object) based")
	}
	if g.Len() != 0 {
		t.Errorf("Abds set to array mode but reports as non-empty:%d", g.Len())
	}

	var idx uint
	var pItem *AbdsItem

	for idx = 0; idx < 10; idx++ {
		g.Add(idx + 1)
	}
	if g.Len() != 10 {
		t.Error("Abds set to array mode but reports as tag (object) based")
	}
	for idx = 1; idx <= 10; idx++ {
		pItem = g.I(idx)
		if pItem.Ti() != idx {
			t.Errorf("Abds set to array mode but reports as incorrect tag number:%d : %d", pItem.Ti(), idx)
		}
	}

}
