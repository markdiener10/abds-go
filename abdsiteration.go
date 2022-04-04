package abds

import (
	_ "strings"
)

type AbdsIter struct {
	index uint
	pItem *AbdsItem
	pAbds *Abds
	//stack []*Abds //Trace our history from root down tree
	flag bool
}

func (g *AbdsIter) reset() {
	g.index = 0
	g.pItem = nil
	g.flag = false
	//Keep the reference to *Abds unchanged
}

func (g *AbdsIter) I() uint {
	return g.index
}

func (g *AbdsIter) Delete() {
	if g == nil {
		return
	}
	g.flag = true
}

func (g *AbdsIter) Flag() {
	if g == nil {
		return
	}
	g.flag = true
}

func (g *Abds) NewIter() *AbdsIter {
	if g == nil {
		return nil
	}
	return &AbdsIter{index: 0, pItem: nil, pAbds: g}
}

func (g *Abds) Iter(piter *AbdsIter) bool {

	if g == nil {
		return false
	}

	if piter == nil {
		return false
	}
	if piter.flag {
		g.delete(piter)
		//Drop the index so we can the same process the same index on the next loop
		piter.index--
	}
	if piter.index >= g.Len() {
		return false
	}
	piter.flag = false
	piter.pItem = g.vals[piter.index]
	piter.index++
	return true
}

func (g *Abds) delete(pIter *AbdsIter) {
	if g == nil {
		return
	}
	pIter.flag = false
	//Zero Based offset from 1 based offset
	idx := pIter.index - 1
	g.vals[idx] = nil //Force garbage collection
	g.vals = append(g.vals[:idx], g.vals[idx+1:]...)
}

// #############################################
func (g *AbdsIter) IsTag() bool {

	if g == nil {
		return false
	}
	if g.pAbds == nil {
		return false
	}
	if g.pItem == nil {
		return false
	}
	if g.pItem.tag == 0 {
		return false
	}
	return true
}

func (g *AbdsIter) Tag() string {

	if g == nil {
		return ""
	}
	if g.pAbds == nil {
		return ""
	}
	if g.pItem == nil {
		return ""
	}
	if g.pItem.tag == 0 {
		return ""
	}
	return g.pAbds.GetTag(g.pItem.tag)
}

func (g *AbdsIter) Tagi() uint {

	if g == nil {
		return 0
	}
	if g.pAbds == nil {
		return 0
	}
	if g.pItem == nil {
		return 0
	}
	return g.pItem.tag
}

//Access and control
// ####################################

func (g *AbdsIter) g() *AbdsItem {
	return g.pItem
}

func (g *AbdsIter) S(val interface{}) {
	if g == nil {
		return
	}
	if g.pAbds == nil {
		return
	}
	g.pAbds.S(g.pItem, val)
}

func (g *AbdsIter) PB() *bool {
	return g.pItem.pB()
}

func (g *AbdsIter) Pi() *int {
	return g.pItem.pi()
}

func (g *AbdsIter) Pi8() *int8 {
	return g.pItem.pi8()
}

func (g *AbdsIter) Pi16() *int16 {
	return g.pItem.pi16()
}

func (g *AbdsIter) Pi32() *int32 {
	return g.pItem.pi32()
}

func (g *AbdsIter) Pi64() *int64 {
	return g.pItem.pi64()
}

func (g *AbdsIter) Pu() *uint {
	return g.pItem.pu()
}

func (g *AbdsIter) Pu8() *uint8 {
	return g.pItem.pu8()
}

func (g *AbdsIter) Pu16() *uint16 {
	return g.pItem.pu16()
}

func (g *AbdsIter) Pu32() *uint32 {
	return g.pItem.pu32()
}

func (g *AbdsIter) Pu64() *uint64 {
	return g.pItem.pu64()
}

func (g *AbdsIter) Pf32() *float32 {
	return g.pItem.pf32()
}

func (g *AbdsIter) Pf64() *float64 {
	return g.pItem.pf64()
}

func (g *AbdsIter) Pc64() *complex64 {
	return g.pItem.pc64()
}

func (g *AbdsIter) Pc128() *complex128 {
	return g.pItem.pc128()
}

func (g *AbdsIter) Pbyte() *[]byte {
	return g.pItem.pbyte()
}

func (g *AbdsIter) Ps() *string {
	return g.pItem.ps()
}

func (g *AbdsIter) Padbs() *Abds {
	return g.pItem.pch()
}

func (g *AbdsIter) Pch() *Abds {
	return g.pItem.pch()
}

func (g *AbdsIter) Praw() interface{} {
	return g.pItem.V()
}

func (g *AbdsIter) Pitem() *AbdsItem {
	return g.pItem
}

// #############################################

func (g *AbdsIter) V() interface{} {
	return g.pItem.V()
}

func (g *AbdsIter) Vb() bool {
	return g.pItem.vB()
}

func (g *AbdsIter) Vi() int {
	return g.pItem.vi()
}

func (g *AbdsIter) Vi8() int8 {
	return g.pItem.vi8()
}

func (g *AbdsIter) Vi16() int16 {
	return g.pItem.vi16()
}

func (g *AbdsIter) Vi32() int32 {
	return g.pItem.vi32()
}

func (g *AbdsIter) Vi64() int64 {
	return g.pItem.vi64()
}

func (g *AbdsIter) Vu() uint {
	return g.pItem.vu()
}

func (g *AbdsIter) Vu8() uint8 {
	return g.pItem.vu8()
}

func (g *AbdsIter) Vu16() uint16 {
	return g.pItem.vu16()
}

func (g *AbdsIter) Vu32() uint32 {
	return g.pItem.vu32()
}

func (g *AbdsIter) Vu64() uint64 {
	return g.pItem.vu64()
}

func (g *AbdsIter) Vf32() float32 {
	return g.pItem.vf32()
}

func (g *AbdsIter) Vf64() float64 {
	return g.pItem.vf64()
}

func (g *AbdsIter) Vc64() complex64 {
	return g.pItem.vc64()
}

func (g *AbdsIter) Vc128() complex128 {
	return g.pItem.vc128()
}

func (g *AbdsIter) Vs() string {
	//Comment More
	return g.pItem.vs()
}
