package abds

func (g *Abds) VB(parm interface{}, parms ...bool) bool {
	return g.g(parm, parms...).vB()
}

func (g *Abds) Vi(parm interface{}, parms ...bool) int {
	return g.g(parm, parms...).vi()
}

func (g *Abds) Vi8(parm interface{}, parms ...bool) int8 {
	return g.g(parm, parms...).vi8()
}

func (g *Abds) Vi16(parm interface{}, parms ...bool) int16 {
	return g.g(parm, parms...).vi16()
}

func (g *Abds) Vi32(parm interface{}, parms ...bool) int32 {
	return g.g(parm, parms...).vi32()
}

func (g *Abds) Vi64(parm interface{}, parms ...bool) int64 {
	return g.g(parm, parms...).vi64()
}

func (g *Abds) Vu(parm interface{}, parms ...bool) uint {
	return g.g(parm, parms...).vu()
}

func (g *Abds) Vu8(parm interface{}, parms ...bool) uint8 {
	return g.g(parm, parms...).vu8()
}

func (g *Abds) Vu16(parm interface{}, parms ...bool) uint16 {
	return g.g(parm, parms...).vu16()
}

func (g *Abds) Vu32(parm interface{}, parms ...bool) uint32 {
	return g.g(parm, parms...).vu32()
}

func (g *Abds) Vu64(parm interface{}, parms ...bool) uint64 {
	return g.g(parm, parms...).vu64()
}

func (g *Abds) Vf32(parm interface{}, parms ...bool) float32 {
	return g.g(parm, parms...).vf32()
}

func (g *Abds) Vf64(parm interface{}, parms ...bool) float64 {
	return g.g(parm, parms...).vf64()
}

func (g *Abds) Vc64(parm interface{}, parms ...bool) complex64 {
	return g.g(parm, parms...).vc64()
}

func (g *Abds) Vc128(parm interface{}, parms ...bool) complex128 {
	return g.g(parm, parms...).vc128()
}

func (g *Abds) Vs(parm interface{}, parms ...bool) string {
	return g.g(parm, parms...).vs()
}
