package abds

func (g *Abds) PB(parm interface{}) *bool {
	return g.g(parm).pB()
}

func (g *Abds) Pi(parm interface{}) *int {
	return g.g(parm).pi()
}

func (g *Abds) Pi8(parm interface{}) *int8 {
	return g.g(parm).pi8()
}

func (g *Abds) Pi16(parm interface{}) *int16 {
	return g.g(parm).pi16()
}

func (g *Abds) Pi32(parm interface{}) *int32 {
	return g.g(parm).pi32()
}

func (g *Abds) Pi64(parm interface{}) *int64 {
	return g.g(parm).pi64()
}

func (g *Abds) Pu(parm interface{}) *uint {
	return g.g(parm).pu()
}

func (g *Abds) Pu8(parm interface{}) *uint8 {
	return g.g(parm).pu8()
}

func (g *Abds) Pu16(parm interface{}) *uint16 {
	return g.g(parm).pu16()
}

func (g *Abds) Pu32(parm interface{}) *uint32 {
	return g.g(parm).pu32()
}

func (g *Abds) Pu64(parm interface{}) *uint64 {
	return g.g(parm).pu64()
}

func (g *Abds) Pf32(parm interface{}) *float32 {
	return g.g(parm).pf32()
}

func (g *Abds) Pf64(parm interface{}) *float64 {
	return g.g(parm).pf64()
}

func (g *Abds) Pc64(parm interface{}) *complex64 {
	return g.g(parm).pc64()
}

func (g *Abds) Pc128(parm interface{}) *complex128 {
	return g.g(parm).pc128()
}

func (g *Abds) Pb(parm interface{}) *[]byte {
	return g.g(parm).pb()
}

func (g *Abds) Pbyte(parm interface{}) *[]byte {
	return g.g(parm).pb()
}

func (g *Abds) Ps(parm interface{}) *string {
	return g.g(parm).ps()
}

func (g *Abds) Padbs(parm interface{}) *Abds {
	return g.g(parm).pabds()
}

func (g *Abds) Pch(parm interface{}) *Abds {
	return g.g(parm).pabds()
}

//Map, Struct, Array
func (g *Abds) Praw(parm interface{}) interface{} {
	return g.g(parm).v()
}
