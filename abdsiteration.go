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
	err := g.pItem.S(val)
	if err != nil {
		g.pAbds.Errs().Log(err)
	}
}

func (g *AbdsIter) PB() *bool {
	return g.g().pB()
}

func (g *AbdsIter) Pi() *int {
	return g.g().pi()
}

func (g *AbdsIter) Pi8() *int8 {
	return g.g().pi8()
}

func (g *AbdsIter) Pi16() *int16 {
	return g.g().pi16()
}

func (g *AbdsIter) Pi32() *int32 {
	return g.g().pi32()
}

func (g *AbdsIter) Pi64() *int64 {
	return g.g().pi64()
}

func (g *AbdsIter) Pu() *uint {
	return g.g().pu()
}

func (g *AbdsIter) Pu8() *uint8 {
	return g.g().pu8()
}

func (g *AbdsIter) Pu16() *uint16 {
	return g.g().pu16()
}

func (g *AbdsIter) Pu32() *uint32 {
	return g.g().pu32()
}

func (g *AbdsIter) Pu64() *uint64 {
	return g.g().pu64()
}

func (g *AbdsIter) Pf32() *float32 {
	return g.g().pf32()
}

func (g *AbdsIter) Pf64() *float64 {
	return g.g().pf64()
}

func (g *AbdsIter) Pc64() *complex64 {
	return g.g().pc64()
}

func (g *AbdsIter) Pc128() *complex128 {
	return g.g().pc128()
}

func (g *AbdsIter) Pbyte() *[]byte {
	return g.g().pbyte()
}

func (g *AbdsIter) Ps() *string {
	return g.g().ps()
}

func (g *AbdsIter) Padbs() *Abds {
	return g.g().pch()
}

func (g *AbdsIter) Pch() *Abds {
	return g.g().pch()
}

func (g *AbdsIter) Praw() interface{} {
	return g.g().v()
}

// #############################################

func (g *AbdsIter) V() interface{} {
	return g.g().v()
}

func (g *AbdsIter) Vb() bool {
	return g.g().vB()
}

func (g *AbdsIter) Vi() int {
	return g.g().vi()
}

func (g *AbdsIter) Vi8() int8 {
	return g.g().vi8()
}

func (g *AbdsIter) Vi16() int16 {
	return g.g().vi16()
}

func (g *AbdsIter) Vi32() int32 {
	return g.g().vi32()
}

func (g *AbdsIter) Vi64() int64 {
	return g.g().vi64()
}

func (g *AbdsIter) Vu() uint {
	return g.g().vu()
}

func (g *AbdsIter) Vu8() uint8 {
	return g.g().vu8()
}

func (g *AbdsIter) Vu16() uint16 {
	return g.g().vu16()
}

func (g *AbdsIter) Vu32() uint32 {
	return g.g().vu32()
}

func (g *AbdsIter) Vu64() uint64 {
	return g.g().vu64()
}

func (g *AbdsIter) Vf32() float32 {
	return g.g().vf32()
}

func (g *AbdsIter) Vf64() float64 {
	return g.g().vf64()
}

func (g *AbdsIter) Vc64() complex64 {
	return g.g().vc64()
}

func (g *AbdsIter) Vc128() complex128 {
	return g.g().vc128()
}

func (g *AbdsIter) Vs() string {
	return g.g().vs()
}
