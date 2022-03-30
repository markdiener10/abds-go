package abds

func (g *Abds) PB(tag string, parms ...interface{}) *bool {
	return g.G(tag, parms...).PB()
}

func (g *Abds) Pi(tag string, parms ...interface{}) *int {
	return g.G(tag, parms...).Pi()
}

func (g *Abds) Pi8(tag string, parms ...interface{}) *int8 {
	return g.G(tag, parms...).Pi8()
}

func (g *Abds) Pi16(tag string, parms ...interface{}) *int16 {
	return g.G(tag, parms...).Pi16()
}

func (g *Abds) Pi32(tag string, parms ...interface{}) *int32 {
	return g.G(tag, parms...).Pi32()
}

func (g *Abds) Pi64(tag string, parms ...interface{}) *int64 {
	return g.G(tag, parms...).Pi64()
}

func (g *Abds) Pu(tag string, parms ...interface{}) *uint {
	return g.G(tag, parms...).Pu()
}

func (g *Abds) Pu8(tag string, parms ...interface{}) *uint8 {
	return g.G(tag, parms...).Pu8()
}

func (g *Abds) Pu16(tag string, parms ...interface{}) *uint16 {
	return g.G(tag, parms...).Pu16()
}

func (g *Abds) Pu32(tag string, parms ...interface{}) *uint32 {
	return g.G(tag, parms...).Pu32()
}

func (g *Abds) Pu64(tag string, parms ...interface{}) *uint64 {
	return g.G(tag, parms...).Pu64()
}

func (g *Abds) Pf32(tag string, parms ...interface{}) *float32 {
	return g.G(tag, parms...).Pf32()
}

func (g *Abds) Pf64(tag string, parms ...interface{}) *float64 {
	return g.G(tag, parms...).Pf64()
}

func (g *Abds) Pc64(tag string, parms ...interface{}) *complex64 {
	return g.G(tag, parms...).Pc64()
}

func (g *Abds) Pc128(tag string, parms ...interface{}) *complex128 {
	return g.G(tag, parms...).Pc128()
}

func (g *Abds) Pb(tag string, parms ...interface{}) *[]byte {
	return g.G(tag, parms...).Pb()
}

func (g *Abds) Pbyte(tag string, parms ...interface{}) *[]byte {
	return g.G(tag, parms...).Pb()
}

func (g *Abds) Prune(tag string, parms ...interface{}) *[]rune {
	return g.G(tag, parms...).Prune()
}

func (g *Abds) Ps(tag string, parms ...interface{}) *string {
	return g.G(tag, parms...).Ps()
}

func (g *Abds) Padbs(tag string, parms ...interface{}) *Abds {
	return g.G(tag, parms...).Pabds()
}

func (g *Abds) Pch(tag string, parms ...interface{}) *Abds {
	return g.G(tag, parms...).Pabds()
}

func (g *Abds) Ptran(tag string, parms ...interface{}) *AbdsTransform {
	return g.G(tag, parms...).Ptran()
}
