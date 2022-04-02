package abds

type AbdsScanFunc func(level uint, pIter *AbdsIter, pCap *Abds) bool
type AbdsSelectFunc func(level uint, pIter *AbdsIter)

//Sorting functions (Depending on Size O(n))
func (g *Abds) Scan(fn AbdsScanFunc, parms ...bool) *Abds {
	pCap := New()
	if fn == nil {
		return pCap
	}
	norecurse := parmbool(parms...)
	g.scanlevel(1, norecurse, fn, pCap)
	return pCap
}

func (g *Abds) scanlevel(level uint, norecurse bool, fn AbdsScanFunc, pCap *Abds) bool {

	var iter AbdsIter
	var pItem *AbdsItem
	var pAbds *Abds
	var idx int
	var terminate bool
	var resb bool

	iter.pAbds = g
	iter.reset()

	for idx = 0; idx < len(g.vals); idx++ {
		iter.index = uint(idx + 1)
		iter.pItem = g.vals[idx]
		iter.flag = false
		terminate = fn(level, &iter, pCap)
		if iter.flag {
			g.delete(&iter)
			//Lower index so we re-process the same idx on the next loop
			idx--
		}
		if terminate {
			return terminate
		}
	}
	if norecurse {
		return false
	}
	for idx = 0; idx < len(g.vals); idx++ {
		pItem = g.vals[idx]
		//Only recurse the *Abds children
		pAbds, resb = pItem.val.(*Abds)
		if !resb {
			continue
		}
		terminate = pAbds.scanlevel(level+1, norecurse, fn, pCap)
		if terminate {
			return terminate
		}
	}
	return false
}
