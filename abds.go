package abds

import (
	"fmt"
	"strings"
)

//Agile binary data structure

//Bit Behavior Flags
const (
	ISARRAY  uint64 = 1 << iota
	TAGCASE         //Allow for mixed case tags
	ERRORSET        //Set a bit to indicate that we have errors
	RECURSE
	NOAUTO
	FORCEAUTO
)

//TODO?:Maybe space for package code in high order bits
//Top 16 bits == 35,000 package identifiers

type Abds struct {
	flags uint64
	Vals  []*AbdsItem
}

type AbdsSortFunc func(p1 *AbdsItem, p2 *AbdsItem) bool

func (g *Abds) IsErr() bool {
	if g.flags&ERRORSET == 0 {
		return false
	}
	return true
}

func (g *Abds) Is(tag string, parms ...interface{}) bool {

	if g.IsArray() {
		return false
	}
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return false
	}
	if g.flags&TAGCASE == 0 {
		tag = strings.ToUpper(tag)
	}

	pItem := g.find(tag, parmflag(RECURSE, parms...))
	if pItem == nil {
		return false
	}
	return true

}

func (g *Abds) T(tag string, parms ...interface{}) *AbdsItem {

	tag = strings.TrimSpace(tag)
	if tag == "" {
		if parms == nil {
			parms = make([]interface{},0)
		}	
		parms = append(parms,g)	
		errset(fmt.Errorf("Called Tag with empty tag"), parms...)
		return nil
	}

	if g.IsArray() {
		if parms == nil {
			parms = make([]interface{},0)
		}	
		parms = append(parms,g)	
		errset(fmt.Errorf("Called Tag in Array Mode:%s", tag), parms...)
		return nil
	}

	if g.flags&TAGCASE == 0 {
		tag = strings.ToUpper(tag)
	}
	pItem := g.find(tag, false)
	if pItem != nil {
		return pItem
	}

	if !parmflag(FORCEAUTO, parms...) {
		if g.flags&NOAUTO != 0 {
			return nil
		}
		if parmflag(NOAUTO, parms...) {
			return nil
		}
	}
	pItem = &AbdsItem{tag: tag, Val: nil}
	g.Vals = append(g.Vals, pItem)
	return pItem
}

func (g *Abds) P(tag string, val interface{}, parms ...interface{}) {
	pItem := g.T(tag, parms...)
	if pItem == nil {
		return
	}
	pItem.P(val, parms...)
}

func (g *Abds) Len() uint {
	return uint(len(g.Vals))
}

func (g *Abds) Flags(flags uint64) *Abds {
	//Default case: most flexible
	if flags == 0 {
		g.flags = 0
		return g
	}

	/*

		//We could add flag edge triggers and re-configure the abds structure
		var Item *AbdsItem
		var index int

		for index = 0; index < len(g.Vals); index++ {
			Item = &g.Vals[index]
			if flags&ABDSTAGCASE != 0 {
				Item.Tag = strings.toUpper(Item.Tag)
			}
		}
	*/
	g.flags = flags
	return g
}

//Sorting functions (Depending on Size O(n))
func (g *Abds) Sort(fn AbdsSortFunc) {

	if fn == nil {
		return
	}
	var swapped bool
	var pitema *AbdsItem
	var pitemb *AbdsItem

	n := len(g.Vals)

	for i := 0; i < n-1; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			pitema = g.Vals[j]
			pitemb = g.Vals[j+1]
			if fn(pitema, pitemb) == true {
				g.Vals[j] = pitemb
				g.Vals[j+1] = pitema
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}

func N(flags uint64) *Abds {
	return New().Flags(flags)
}

func NA(flags uint64) *Abds {
	return New().Flags(ISARRAY | flags)
}

func New() *Abds {
	return &Abds{flags: 0, Vals: make([]*AbdsItem, 0)}
}
