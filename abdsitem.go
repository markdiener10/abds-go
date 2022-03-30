package abds

import (
	_ "fmt"
)

type AbdsItem struct {
	tag interface{} //uint or String
	val interface{} //Range of values
}

func (g *AbdsItem) Clear() {
	//Remove the link between the value and the abds structure (pointers are garbage collected)
	//Do not touch the TAG value
	g.val = nil
}

//return Tag Values, Not Values themselves
func (g *AbdsItem) Ti() uint {
	if g.tag == nil {
		return 0
	}
	gt, ok := g.tag.(uint)
	if !ok {
		return 0
	}
	return gt
}

func (g *AbdsItem) Ts() string {
	if g.tag == nil {
		return ""
	}
	gt, ok := g.tag.(string)
	if !ok {
		return ""
	}
	return gt
}
