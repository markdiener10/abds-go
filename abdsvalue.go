package abds

func (g *Abds) Vb(tag string, parms ...interface{}) bool {
	return g.G(tag, parms...).Vb()
}

func (g *Abds) Vi(tag string, parms ...interface{}) int {
	return g.G(tag, parms...).Vi()
}

func (g *Abds) Vi8(tag string, parms ...interface{}) int8 {
	return g.G(tag, parms...).Vi8()
}

func (g *Abds) Vi16(tag string, parms ...interface{}) int16 {
	return g.G(tag, parms...).Vi16()
}

func (g *Abds) Vi32(tag string, parms ...interface{}) int32 {
	return g.G(tag, parms...).Vi32()
}

func (g *Abds) Vi64(tag string, parms ...interface{}) int64 {
	return g.G(tag, parms...).Vi64()
}

func (g *Abds) Vu(tag string, parms ...interface{}) uint {
	return g.G(tag, parms...).Vu()
}

func (g *Abds) Vu8(tag string, parms ...interface{}) uint8 {
	return g.G(tag, parms...).Vu8()
}

func (g *Abds) Vu16(tag string, parms ...interface{}) uint16 {
	return g.G(tag, parms...).Vu16()
}

func (g *Abds) Vu32(tag string, parms ...interface{}) uint32 {
	return g.G(tag, parms...).Vu32()
}

func (g *Abds) Vu64(tag string, parms ...interface{}) uint64 {
	return g.G(tag, parms...).Vu64()
}

func (g *Abds) Vf32(tag string, parms ...interface{}) float32 {
	return g.G(tag, parms...).Vf32()
}

func (g *Abds) Vf64(tag string, parms ...interface{}) float64 {
	return g.G(tag, parms...).Vf64()
}

func (g *Abds) Vc64(tag string, parms ...interface{}) complex64 {
	return g.G(tag, parms...).Vc64()
}

func (g *Abds) Vc128(tag string, parms ...interface{}) complex128 {
	return g.G(tag, parms...).Vc128()
}

func (g *Abds) Vbyte(tag string, parms ...interface{}) []byte {
	return g.G(tag, parms...).Vbyte()
}

func (g *Abds) Vrune(tag string, parms ...interface{}) []rune {
	return g.G(tag, parms...).Vrune()
}

func (g *Abds) Vs(tag string, parms ...interface{}) string {
	return g.G(tag, parms...).Vs()
}

func (g *Abds) Vch(tag string, parms ...interface{}) *Abds {
	return g.G(tag, parms...).Pabds()
}

func (g *Abds) Vadbs(tag string, parms ...interface{}) *Abds {
	return g.G(tag, parms...).Pabds()
}

func (g *Abds) Vtran(tag string, parms ...interface{}) interface{} {
	return g.G(tag, parms...).Ptran()
}
