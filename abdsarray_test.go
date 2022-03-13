package abds

import (
	"testing"
)

func TestArrayOperations(t *testing.T) {

	//Array operations
	g := New(ARRAY)
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
		t.Errorf("Abds set to array mode but fails to add 10 items:%d", g.Len())
	}
	for idx = 1; idx <= 10; idx++ {
		pItem = g.AG(idx)
		if pItem.Ti() != idx {
			t.Errorf("Abds set to array mode but reports as incorrect tag number:%d : %d", pItem.Ti(), idx)
		}
	}

}
