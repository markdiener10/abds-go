package abds

import (
	"fmt"
	_ "strings"
)

func (g *Abds) IsArray() bool {
	if g.flags&ARRAY == 0 {
		return false
	}
	return true
}

func (g *Abds) Add(val interface{}, parms ...interface{}) bool {

	if !checkType(val) {
		parmerr(fmt.Errorf("Invalid Type passed for add:%x", val), parms...)
		return false
	}
	idx := g.Len()
	pItem := &AbdsItem{tag: idx + 1, val: val}
	g.Vals = append(g.Vals, pItem)
	return true
}

func (g *Abds) AG(idx uint) *AbdsItem {
	if idx < 1 {
		return nil
	}
	if idx > g.Len() {
		return nil
	}
	return g.Vals[idx-1]
}

func (g *Abds) AN(val interface{}, parms ...interface{}) (*Abds, bool) {
	return nil, false
}

func (g *Abds) AP(idx uint, val interface{}, parms ...interface{}) {
	if !checkType(val) {
		parmerr(fmt.Errorf("Invalid Type passed for add:%x", val), parms...)
		return
	}
	pItem := g.AG(idx)
	if pItem == nil {
		parmerr(fmt.Errorf("Item for index not found:%d", idx), parms...)
	}
	pItem.S(val, parms...)
}
