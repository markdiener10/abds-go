package abds

import (
	"fmt"
	_ "strings"
)

func (g *Abds) IsArray() bool {
	if g.flags&ARRAY == 0 {
		return false
	}
	return true
}

//We cannot prevent major circular references (developer caution)
func (g *Abds) Add(val interface{}) *AbdsItem {

	if !checkType(val, true) {
		parmerr(fmt.Errorf("ABDS Invalid Type passed for array add:%T", val), g.errs)
		val = nil
	}

	pItem := &AbdsItem{tag: g.Len() + 1, val: nil}
	pItem.S(val, g.errs)
	g.vals = append(g.vals, pItem)
	return pItem
}

//Convenience function to add *Abds to array
func (g *Abds) ANew() *Abds {
	return g.AN()
}

func (g *Abds) AN() *Abds {
	gabds := New(ARRAY)
	gabds.errs = g.errs
	gabds.flags |= CHILD
	if g.errs == nil {
		gabds.flags |= NOERR
	}
	pItem := &AbdsItem{tag: g.Len() + 1, val: gabds}
	g.vals = append(g.vals, pItem)
	return gabds
}

func (g *Abds) AInsert(idx uint, val interface{}) *AbdsItem {
	return g.AI(idx, val)
}

func (g *Abds) AIns(idx uint, val interface{}) *AbdsItem {
	return g.AI(idx, val)
}

func (g *Abds) AI(idx uint, val interface{}) *AbdsItem {

	if idx == 0 {
		idx = 1
	}
	if !checkType(val, true) {
		parmerr(fmt.Errorf("ABDS Invalid Type passed for array insert:%T idx:%d", val, idx), g.errs)
		val = nil
	}
	if g.Len() < idx {
		return g.Add(val)
	}
	pItemNew := &AbdsItem{tag: 0, val: nil}

	var pSrc *AbdsItem
	var pDst *AbdsItem

	g.vals = append(g.vals, pItemNew)

	//Shift the values up
	for slot := g.Len() - 2; slot >= idx-1; slot-- {
		pSrc = g.vals[slot]
		pDst = g.vals[slot+1]
		pSrc.val = pDst.val
	}

	//Relabel the trags
	for slot := idx - 1; idx < g.Len(); slot++ {
		pDst = g.vals[slot]
		pDst.tag = slot + 1
	}
	pItemNew = g.vals[idx-1]
	pItemNew.S(val)
	return pItemNew
}

func (g *Abds) AG(idx uint) *AbdsItem {
	if idx < 1 {
		idx = 1
	}
	if idx > g.Len() {
		return nil
	}
	return g.vals[idx-1]
}

func (g *Abds) AS(idx uint, val interface{}) {

	if !checkType(val,true) {
		g.adderr("ABDS Invalid Type passed for AS:%T : Idx.%d", val, idx)
		val = nil
	}

	pItem := g.AG(idx)
	if pItem == nil {
		g.adderr("ABDS Item for index not found:%d", idx)
	}
	pItem.S(val, g.errs)
}

// #############################################

func (g *Abds) APb(idx uint) *bool {
	return g.AG(idx).PB()
}

func (g *Abds) APi(idx uint) *int {
	return g.AG(idx).Pi()
}

func (g *Abds) APi8(idx uint) *int8 {
	return g.AG(idx).Pi8()
}

func (g *Abds) APi16(idx uint) *int16 {
	return g.AG(idx).Pi16()
}

func (g *Abds) APi32(idx uint) *int32 {
	return g.AG(idx).Pi32()
}

func (g *Abds) APi64(idx uint) *int64 {
	return g.AG(idx).Pi64()
}

func (g *Abds) APu(idx uint) *uint {
	return g.AG(idx).Pu()
}

func (g *Abds) APu8(idx uint) *uint8 {
	return g.AG(idx).Pu8()
}

func (g *Abds) APu16(idx uint) *uint16 {
	return g.AG(idx).Pu16()
}

func (g *Abds) APu32(idx uint) *uint32 {
	return g.AG(idx).Pu32()
}

func (g *Abds) APu64(idx uint) *uint64 {
	return g.AG(idx).Pu64()
}

func (g *Abds) APf32(idx uint) *float32 {
	return g.AG(idx).Pf32()
}

func (g *Abds) APf64(idx uint) *float64 {
	return g.AG(idx).Pf64()
}

func (g *Abds) APc64(idx uint) *complex64 {
	return g.AG(idx).Pc64()
}

func (g *Abds) APc128(idx uint) *complex128 {
	return g.AG(idx).Pc128()
}

func (g *Abds) APbyte(idx uint) *[]byte {
	return g.AG(idx).Pbyte()
}

func (g *Abds) APrune(idx uint) *[]rune {
	return g.AG(idx).Prune()
}

func (g *Abds) APs(idx uint) *string {
	return g.AG(idx).Ps()
}

func (g *Abds) APadbs(idx uint) *Abds {
	return g.AG(idx).Pabds()
}

func (g *Abds) APtran(idx uint) interface{} {
	return g.AG(idx).Ptran()
}

// #############################################

func (g *Abds) AVb(idx uint) bool {
	return g.AG(idx).Vb()
}

func (g *Abds) AVi(idx uint) int {
	return g.AG(idx).Vi()
}

func (g *Abds) AVi8(idx uint) int8 {
	return g.AG(idx).Vi8()
}

func (g *Abds) AVi16(idx uint) int16 {
	return g.AG(idx).Vi16()
}

func (g *Abds) AVi32(idx uint) int32 {
	return g.AG(idx).Vi32()
}

func (g *Abds) AVi64(idx uint) int64 {
	return g.AG(idx).Vi64()
}

func (g *Abds) AVu(idx uint) uint {
	return g.AG(idx).Vu()
}

func (g *Abds) AVu8(idx uint) uint8 {
	return g.AG(idx).Vu8()
}

func (g *Abds) AVu16(idx uint) uint16 {
	return g.AG(idx).Vu16()
}

func (g *Abds) AVu32(idx uint) uint32 {
	return g.AG(idx).Vu32()
}

func (g *Abds) AVu64(idx uint) uint64 {
	return g.AG(idx).Vu64()
}

func (g *Abds) AVf32(idx uint) float32 {
	return g.AG(idx).Vf32()
}

func (g *Abds) AVf64(idx uint) float64 {
	return g.AG(idx).Vf64()
}

func (g *Abds) AVc64(idx uint) complex64 {
	return g.AG(idx).Vc64()
}

func (g *Abds) AVc128(idx uint) complex128 {
	return g.AG(idx).Vc128()
}

func (g *Abds) AVs(idx uint) string {
	return g.AG(idx).Vs()
}

//Do NOT get []byte and []rune by value, only pointer access to them
//To avoid the big stack copy on return of a value
