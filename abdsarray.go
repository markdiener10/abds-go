package abds

import (
	"fmt"
	_ "strings"
)

func (g *Abds) IsArray() bool {
	if g.flags&ISARRAY == 0 {
		return false
	}
	return true
}

func (g *Abds) Add(val interface{}, parms ...interface{}) bool {

	if !g.IsArray() {
		if parms == nil {
			parms = make([]interface{},0)
		}
		parms = append(parms,g)
		errset(fmt.Errorf("Adding array element to tag based abds"), parms...)
		return false
	}
	if !checkType(val) {
		if parms == nil {
			parms = make([]interface{},0)
		}
		parms = append(parms,g)
		errset(fmt.Errorf("Invalid Type passed for add:%x", val), parms...)
		return false
	}
	idx := g.Len()
	pItem := &AbdsItem{tag: idx + 1, Val: val}
	g.Vals = append(g.Vals, pItem)
	return true
}

func (g *Abds) I(idx uint) *AbdsItem {
	if idx < 1 {
		return nil
	}
	if idx > g.Len() {
		return nil
	}
	return g.Vals[idx-1]
}

func (g *Abds) IP(idx uint, val interface{}, parms ...interface{}) {
	if !checkType(val) {
		parms = append(parms, g)
		errset(fmt.Errorf("Invalid Type passed for add:%x", val), parms...)
		return
	}
	pItem := g.I(idx)
	if pItem == nil {
		parms = append(parms, g)		
		errset(fmt.Errorf("Item for index not found:%d", idx), parms...)
	}
	pItem.P(val, parms...)
}
