package abds

import (
	"fmt"
	"strings"
)

//Allow for daisy-chaining value set by returned self (g *Abds)
func (g *Abds) S(parms ...interface{}) *Abds {

	if g == nil {
		return nil
	}

	var err error = nil
	var idx uint = 0
	var pItem *AbdsItem = nil

	switch len(parms) {
	case 0:
		return g
	case 1:
		if parms[0], err = checkval(parms[0]); err != nil {
			g.Log(err)
			return g
		}
		pItem = &AbdsItem{tag: 0, val: parms[0]}
		g.vals = append(g.vals, pItem)
		return g
	}

	if parms[1], err = checkval(parms[1]); err != nil {
		g.Log(err)
		return g
	}

	switch parms[0].(type) {
	case string:
		stag := strings.TrimSpace(parms[0].(string))
		if stag == "" {
			g.Log(fmt.Errorf("ABDSitem S() received empty tag parm"))
			return g
		}
		ntag := g.GetTagi(stag)
		if ntag > 0 {
			pItem := g.find(ntag, false)
			if pItem != nil {
				return g
			}
		}
	case int, int8, int16, int32, int64:
		idx = uint(parms[0].(int))
		break
	case uint, uint8, uint16, uint32, uint64:
		idx = parms[0].(uint)
		break
	default:
		g.Log(fmt.Errorf("ABDSitem S() invalid tag/idx parameter"))
		return g
	}

	if idx == 0 || g.Len() < idx {
		pItem = &AbdsItem{tag: 0, val: parms[1]}
		g.vals = append(g.vals, pItem)
		return g
	}
	pItem = g.vals[idx]
	pItem.val = parms[1]
	return g
}

//Array insert capability
func (g *Abds) Sinsert(idx uint, val interface{}) bool {

	if g == nil {
		return false
	}
	var err error
	if val, err = checkval(val); err != nil {
		return false
	}
	pItem := &AbdsItem{tag: 0, val: nil}
	pItem.val = val
	nLen := g.Len()

	//Always add one to the end
	g.vals = append(g.vals, nil)
	if nLen <= idx {
		return false
	}
	var pSrc *AbdsItem
	var pDst *AbdsItem

	//Shift the values up
	for slot := g.Len() - 2; slot >= idx-1; slot-- {
		pSrc = g.vals[slot]
		pDst = g.vals[slot+1]
		pDst.val = pSrc.val
		pDst.tag = pSrc.tag
		pSrc.val = nil
		pSrc.tag = 0
	}
	g.vals[idx-1] = pItem
	return true
}
