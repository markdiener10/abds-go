package abds

import (
	_ "fmt"
)

func parmerr(parms ...interface{}) *AbdsErr {
	for _, parm := range parms {
		gerr, ok := parm.(*AbdsErr)
		if !ok {
			continue
		}
		return gerr
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

func (g *Abds) find(tag string, recurse bool) *AbdsItem {

	var index int
	var pItem *AbdsItem = nil

	if g.IsArray() {
		if !recurse {
			return nil
		}
		//
		goto CHILDSCAN
	}

	for index = 0; index < len(g.Vals); index++ {
		pItem = g.Vals[index]
		if pItem == nil {
			continue
		}
		if pItem.tag != tag {
			continue
		}
		return pItem
	}

	if !recurse {
		return nil
	}

	//Only scan the children AFTER scanning the parent
CHILDSCAN:

	var pchild *Abds
	var child Abds
	var ok bool

	for index = 0; index < len(g.Vals); index++ {
		pItem = g.Vals[index]
		child, ok = pItem.Val.(Abds)
		if ok {
			pchild = &child
		} else {
			pchild, ok = pItem.Val.(*Abds)
		}
		if !ok {
			continue
		}
		if pchild == nil {
			continue
		}
		pItem = pchild.find(tag, true)
		if pItem == nil {
			continue
		}
		return pItem
	}
	return nil
}

func errclear(g *Abds, errs *AbdsErr) {
	if g != nil {
		g.flags &= ^(ERRORSET)
	}
	if errs != nil {
		errs.Reset()
	}
}

func errset(err error, parms ...interface{}) {
	for _, parm := range parms {
		gabds, ok := parm.(*Abds)
		if ok {
			gabds.flags |= ERRORSET
			continue
		}
		gerr, ok := parm.(*AbdsErr)
		if !ok {
			continue
		}
		gerr.Log(err)
		break
	}
}
