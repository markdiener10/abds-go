package abds

import (
	"fmt"
	"strings"
)

//Agile binary data structure

//Bit Behavior Flags
const (
	ARRAY   uint64 = 1 << iota
	TAGCASE        //Allow for mixed case tags
	RECURSE
)

//TODO?:Maybe space for package code in high order bits
//Top 16 bits == 35,000 package identifiers

type Abds struct {
	flags uint64
	Vals  []*AbdsItem
}

func (g *Abds) Len() uint {
	return uint(len(g.Vals))
}

func (g *Abds) Flags(flags uint64) *Abds {
	g.flags = flags
	return g
}

func (g *Abds) Is(tag string, parms ...interface{}) bool {

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

func (g *Abds) G(tag string, parms ...interface{}) *AbdsItem {

	tag = strings.TrimSpace(tag)
	if tag == "" {
		parmerr(fmt.Errorf("Top Value called with empty tag"), parms...)
		//But we now have a tagless item in our array
	}
	if g.flags&TAGCASE == 0 {
		tag = strings.ToUpper(tag)
	}
	pItem := g.find(tag, false)
	if pItem != nil {
		return pItem
	}
	pItem = &AbdsItem{tag: tag, val: nil}
	g.Vals = append(g.Vals, pItem)
	return pItem
}

//Convenience function to add
func (g *Abds) N(tag string, parms ...interface{}) *Abds {

	tag = strings.TrimSpace(tag)
	if tag == "" {
		parmerr(fmt.Errorf("Top Value called with empty tag"), parms...)
		return nil
	}
	if g.flags&TAGCASE == 0 {
		tag = strings.ToUpper(tag)
	}

	gabds := New(parms...)

	pItem := g.find(tag, false)
	if pItem == nil {
		pItem = &AbdsItem{tag: tag, val: gabds}
		g.Vals = append(g.Vals, pItem)
		return gabds
	}
	pItem.S(gabds)
	return gabds
}

func (g *Abds) Copy(src *Abds) error {
	if src == nil {
		return fmt.Errorf("Invalid Abds parameter passed for copy")
	}
	g.Vals = make([]*AbdsItem, 0)
	g.Vals = append(g.Vals, (*src).Vals...)
	return nil
}

func New(parms ...interface{}) *Abds {
	return &Abds{flags: parmflags(parms...), Vals: make([]*AbdsItem, 0)}
}
