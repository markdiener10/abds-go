package abds

import (
	_ "strings"
)

type AbdsIter struct {
	index uint
	pItem *AbdsItem
}

func (g *AbdsIter) Reset() {
	g.index = 0
	g.pItem = nil
}

func (g *AbdsIter) I() uint {
	return g.index
}

func NewIter() *AbdsIter {
	return &AbdsIter{index: 0, pItem: nil}
}

//Loop Iteration
func (g *Abds) Iter(piter *AbdsIter) bool {
	if piter == nil {
		return false
	}
	if piter.index >= g.Len() {
		return false
	}
	piter.pItem = g.Vals[piter.index]
	piter.index++
	return true
}

func (g *AbdsIter) G() *AbdsItem {
	return g.pItem
}

func (g *AbdsIter) S(val interface{}, parms ...interface{}) {
	g.G().S(val, parms...)
}

// #############################################

func (g *AbdsIter) P() *AbdsItem {
	return g.G()
}

func (g *AbdsIter) Pb() *bool {
	return g.G().Pb()
}

func (g *AbdsIter) Pi() *int {
	return g.G().Pi()
}

func (g *AbdsIter) Pi8() *int8 {
	return g.G().Pi8()
}

func (g *AbdsIter) Pi16() *int16 {
	return g.G().Pi16()
}

func (g *AbdsIter) Pi32() *int32 {
	return g.G().Pi32()
}

func (g *AbdsIter) Pi64() *int64 {
	return g.G().Pi64()
}

func (g *AbdsIter) Pu() *uint {
	return g.G().Pu()
}

func (g *AbdsIter) Pu8() *uint8 {
	return g.G().Pu8()
}

func (g *AbdsIter) Pu16() *uint16 {
	return g.G().Pu16()
}

func (g *AbdsIter) Pu32() *uint32 {
	return g.G().Pu32()
}

func (g *AbdsIter) Pu64() *uint64 {
	return g.G().Pu64()
}

func (g *AbdsIter) Pf32() *float32 {
	return g.G().Pf32()
}

func (g *AbdsIter) Pf64() *float64 {
	return g.G().Pf64()
}

func (g *AbdsIter) Pc64() *complex64 {
	return g.G().Pc64()
}

func (g *AbdsIter) Pc128() *complex128 {
	return g.G().Pc128()
}

func (g *AbdsIter) Pbyte() *[]byte {
	return g.G().Pbyte()
}

func (g *AbdsIter) Prune() *[]rune {
	return g.G().Prune()
}

func (g *AbdsIter) Ps() *string {
	return g.G().Ps()
}

func (g *AbdsIter) Padbs() *Abds {
	return g.G().Pabds()
}

func (g *AbdsIter) Ptran() interface{} {
	return g.G().Ptran()
}

// #############################################

func (g *AbdsIter) Vb() bool {
	return g.G().Vb()
}

func (g *AbdsIter) Vi() int {
	return g.G().Vi()
}

func (g *AbdsIter) Vi8() int8 {
	return g.G().Vi8()
}

func (g *AbdsIter) Vi16() int16 {
	return g.G().Vi16()
}

func (g *AbdsIter) Vi32() int32 {
	return g.G().Vi32()
}

func (g *AbdsIter) Vi64() int64 {
	return g.G().Vi64()
}

func (g *AbdsIter) Vu() uint {
	return g.G().Vu()
}

func (g *AbdsIter) Vu8() uint8 {
	return g.G().Vu8()
}

func (g *AbdsIter) Vu16() uint16 {
	return g.G().Vu16()
}

func (g *AbdsIter) Vu32() uint32 {
	return g.G().Vu32()
}

func (g *AbdsIter) Vu64() uint64 {
	return g.G().Vu64()
}

func (g *AbdsIter) Vf32() float32 {
	return g.G().Vf32()
}

func (g *AbdsIter) Vf64() float64 {
	return g.G().Vf64()
}

func (g *AbdsIter) Vc64() complex64 {
	return g.G().Vc64()
}

func (g *AbdsIter) Vc128() complex128 {
	return g.G().Vc128()
}

func (g *AbdsIter) Vs() string {
	return g.G().Vs()
}
