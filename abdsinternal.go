package abds

import (
	_ "fmt"
)

func (g *Abds) find(tag string, recurse bool) *AbdsItem {

	var index int
	var pItem *AbdsItem = nil

	for index = 0; index < len(g.Vals); index++ {
		pItem = g.Vals[index]
		if pItem == nil {
			continue
		}
		if pItem.Ts() != tag {
			continue
		}
		return pItem
	}
	if !recurse {
		return nil
	}

	var child *Abds
	var ok bool

	for index = 0; index < len(g.Vals); index++ {
		pItem = g.Vals[index]
		child, ok = pItem.val.(*Abds)
		if !ok {
			continue
		}
		if child == nil {
			continue
		}
		pItem = child.find(tag, true)
		if pItem == nil {
			continue
		}
		return pItem
	}
	return nil
}

func parmflag(flag uint64, parms ...interface{}) bool {
	for _, parm := range parms {
		gflag, ok := parm.(uint64)
		if !ok {
			continue
		}
		if gflag != flag {
			continue
		}
		return true
	}
	return false
}

func parmflags(parms ...interface{}) (flags uint64) {
	flags = 0
	for _, parm := range parms {
		flag, ok := parm.(uint64)
		if !ok {
			continue
		}
		flags |= flag
	}
	return flags
}

func parmerr(err error, parms ...interface{}) {
	for _, parm := range parms {
		gerr, ok := parm.(*AbdsErr)
		if !ok {
			continue
		}
		gerr.Log(err)
		break
	}
}
