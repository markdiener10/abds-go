package abds

func (g *Abds) find(tag string, recurse bool) *AbdsItem {

	var index int
	var pItem *AbdsItem = nil

	for index = 0; index < len(g.vals); index++ {
		pItem = g.vals[index]
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

	for index = 0; index < len(g.vals); index++ {
		pItem = g.vals[index]
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

		gbool, ok := parm.(bool)
		if ok {
			if gbool {
				flags |= ARRAY
			}
			continue
		}
		flag, ok := parm.(uint64)
		if !ok {
			continue
		}
		flags |= flag
	}
	return flags

}

func parmerr(gerr error, parms ...*AbdsErr) {

	for _, parm := range parms {
		if parm == nil {
			continue
		}
		parm.Log(gerr)
		return
	}
}
