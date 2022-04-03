package abds

import (
	_ "fmt"
)

type AbdsItem struct {
	tag uint
	val interface{} //Range of values
}

func (g *AbdsItem) Clear() {
	//Remove the link between the value and the abds structure (pointers are garbage collected)
	//Do not touch the TAG value
	g.val = nil
}

func (g *AbdsItem) Ti() uint {
	return g.tag
}

func (g *AbdsItem) V() interface{} {
	return g.val
}
