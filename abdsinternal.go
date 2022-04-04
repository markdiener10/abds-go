package abds

func (g *Abds) dummy() *AbdsItem {
	return &AbdsItem{}
}

func (g *Abds) find(ntag uint, recurse bool) *AbdsItem {

	var index int
	var pItem *AbdsItem = nil

	for index = 0; index < len(g.vals); index++ {
		pItem = g.vals[index]
		if pItem == nil {
			continue
		}
		if pItem.tag != ntag {
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
		pItem = child.find(ntag, true)
		if pItem == nil {
			continue
		}
		return pItem
	}
	return nil
}

func parmbool(parms ...bool) bool {
	for _, parm := range parms {
		if parm {
			return true
		}
	}
	return false
}

func parmu(flag uint, parms ...uint) bool {
	for _, parm := range parms {
		if parm != flag {
			continue
		}
		return true
	}
	return false
}

func parmflags(parms ...uint) (flags uint) {
	flags = 0
	for _, parm := range parms {
		flags |= parm
	}
	return
}

func parmfilter(flag uint, parms ...uint) (flags uint) {
	flags = 0
	for _, parm := range parms {
		if flag == parm {
			continue
		}
		flags |= parm
	}
	return
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
