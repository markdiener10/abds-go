package abds

import (
	"fmt"
	_ "strings"
)

//Agile binary data structure

type AbdsIter struct {
	index uint
	pItem *AbdsItem
}

func NewIter() *AbdsIter {
	return &AbdsIter{index: 0, pItem: nil}
}

//Loop Iteration
func (g *Abds) Iter(piter *AbdsIter) bool {
	if piter == nil {
		return false
	}
	if piter.index >= g.Len() {
		return false
	}
	piter.pItem = g.Vals[piter.index]
	piter.index++
	return true
}

func (g *AbdsIter) Reset() {
	g.index = 0
	g.pItem = nil
}

func (g *AbdsIter) P(val interface{}, parms ...interface{}) {
	if g.pItem == nil {
		parmerr(fmt.Errorf("Iterator has no valid item:%d", g.index), parms...)
		return
	}
	g.pItem.S(val)
}

func (g *AbdsIter) T() *AbdsItem {
	return g.pItem
}

func (g *AbdsIter) I() uint {
	return g.index
}

func (g *AbdsIter) Ti() uint {
	if g.pItem == nil {
		return 0
	}
	return g.pItem.Ti()
}

func (g *AbdsIter) Ts() string {
	if g.pItem == nil {
		return ""
	}
	return g.pItem.Ts()
}
