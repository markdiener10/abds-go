package abds

import (
	"fmt"
	"strings"
)

//Agile binary data structure

//Bit Behavior Flags
const (
	ARRAY    uint64 = 1 << iota
	TAGMIXED        //Allow for mixed case tags (DEFAULT UPPER CASE)
	RECURSE
	NOERR //Suppress Error tracking
	CHILD //Indicate abds is a child (not a parent)
)

//TODO?:Maybe space for package code in high order bits
//Top 16 bits == 35,000 package identifiers

type Abds struct {
	flags uint64
	vals  []*AbdsItem
	errs  *AbdsErr //Nil, Errs
}

func (g *Abds) Len() uint {
	return uint(len(g.vals))
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
	if g.flags&TAGMIXED == 0 {
		tag = strings.ToUpper(tag)
	}
	pItem := g.find(tag, parmflag(RECURSE, parms...))
	if pItem == nil {
		return false
	}
	return true

}

func (g *Abds) IsErr() (result bool) {
	if g.errs == nil {
		return false
	}
	return g.errs.IsErr()
}

func (g *Abds) Errs() *AbdsErr {
	return g.errs
}

func (g *Abds) ErrorClear() bool {
	if g.errs == nil {
		return false
	}
	g.errs.Reset()
	return true
}

func (g *Abds) adderr(msg string, parms ...interface{}) {
	if g.errs == nil {
		return
	}
	g.errs.Log(fmt.Errorf(msg, parms...))
	return
}

func (g *Abds) Get(tag string, parms ...interface{}) *AbdsItem {
	return g.G(tag, parms)
}

func (g *Abds) G(tag string, parms ...interface{}) *AbdsItem {

	if g.IsArray() {
		g.adderr("ABDS Top Value called with empty tag")
		return nil
	}

	tag = strings.TrimSpace(tag)
	if tag == "" {
		g.adderr("ABDS Top Value called with empty tag")
		//But we now have a tagless item in our array
	}
	if g.flags&TAGMIXED == 0 {
		tag = strings.ToUpper(tag)
	}
	pItem := g.find(tag, parmflag(RECURSE, parms))
	if pItem != nil {
		return pItem
	}
	pItem = &AbdsItem{tag: tag, val: nil}
	g.vals = append(g.vals, pItem)
	return pItem
}

//Convenience function to add Abds as a child
func (g *Abds) New(tag string) *Abds {
	return g.N(tag)
}

func (g *Abds) N(tag string) *Abds {

	tag = strings.TrimSpace(tag)
	if tag == "" {
		g.adderr("ABDS Top Value called with empty tag")
		return nil
	}
	if g.flags&TAGMIXED == 0 {
		tag = strings.ToUpper(tag)
	}

	gabds := New()

	//Recycle the error pointer for children
	gabds.errs = g.errs
	gabds.flags |= CHILD
	if g.errs == nil {
		gabds.flags |= NOERR
	}

	pItem := g.find(tag, false)
	if pItem == nil {
		pItem = &AbdsItem{tag: tag, val: gabds}
		g.vals = append(g.vals, pItem)
		return gabds
	}
	pItem.S(gabds)
	return gabds
}

func (g *Abds) Ntran(tag string, tran *AbdsTransform) bool {

	tag = strings.TrimSpace(tag)
	if tag == "" {
		g.adderr("ABDS Ntran called with empty tag")
		return false
	}
	if tran == nil {
		g.adderr("ABDS Ntran called with null transform parameter")
		return false
	}
	if g.flags&TAGMIXED == 0 {
		tag = strings.ToUpper(tag)
	}

	pItem := g.find(tag, false)
	if pItem == nil {
		pItem = &AbdsItem{tag: tag, val: tran}
		g.vals = append(g.vals, pItem)
		return true
	}
	pItem.S(tran)
	return true
}

func (g *Abds) Copy(src *Abds) error {
	if src == nil {
		return fmt.Errorf("ABDS Invalid Abds parameter passed for copy")
	}
	//This will replace the current payload with a new one (and release memory)
	g.vals = make([]*AbdsItem, 0)
	//Almost like a C++11 friend function
	g.vals = append(g.vals, (*src).vals...)
	return nil
}

func New(parms ...interface{}) *Abds {
	gabds := &Abds{flags: parmflags(parms...), vals: make([]*AbdsItem, 0)}
	if gabds.flags&NOERR == 0 {
		gabds.errs = NewErrs()
	}
	return gabds
}
