package abds

import (
	"strings"
)

//Agile binary data structure
const VERSION = 2

type Abds struct {
	root bool
	tags *AbdsTags
	vals []*AbdsItem
	link interface{} //*Abds or *Abdserrors
}

func (g *Abds) Version() uint {
	return VERSION
}

func (g *Abds) Len() uint {
	return uint(len(g.vals))
}

func New() *Abds {
	return &Abds{root: true, vals: make([]*AbdsItem, 0), tags: nil, link: NewErrs()}
}

func (g *Abds) Nchild(parms ...interface{}) *Abds {
	return g.Nch(parms...)
}

func (g *Abds) Nch(parms ...interface{}) *Abds {

	pAbds := &Abds{root: false, vals: make([]*AbdsItem, 0), tags: nil, link: g}
	pItem := &AbdsItem{tag: 0, val: pAbds}
	g.vals = append(g.vals, pItem)

	if len(parms) < 1 {
		return pAbds
	}
	stag, ok := parms[0].(string)
	if !ok {
		return pAbds
	}
	stag = strings.TrimSpace(stag)
	if stag == "" {
		return pAbds
	}
	pItem.tag = g.GetSetTag(stag)
	return pAbds
}

func (g *Abds) g(parm interface{}, parms ...bool) *AbdsItem {

	nidx := uint(0)
	var ntag uint
	var pItem *AbdsItem

	switch parm.(type) {
	case string:
		stag := strings.TrimSpace(parm.(string))
		if stag == "" {
			break
		}
		recurse := parmbool(parms...)
		ntag = g.GetTagi(stag)
		if ntag > 0 {
			pItem = g.find(ntag, recurse)
			if pItem != nil {
				return pItem
			}
		} else {
			ntag = g.GetSetTag(stag)
		}
		pItem = &AbdsItem{tag: ntag, val: nil}
		g.vals = append(g.vals, pItem)
		return pItem
	case int, int8, int16, int32, int64:
		nidx = uint(parm.(int))
	case uint, uint8, uint16, uint32, uint64:
		nidx = parm.(uint)
	case nil:
	default:
	}

	if nidx == 0 || uint(len(g.vals)) < nidx {
		pItem = &AbdsItem{tag: 0, val: nil}
		g.vals = append(g.vals, pItem)
	} else {
		pItem = g.vals[nidx-1]
	}
	return pItem
}
