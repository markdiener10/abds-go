package abds

func (g *AbdsItem) pB() *bool {

	switch g.val.(type) {
	case *bool:
		break
	default:
		g.S(g.vB())
	}
	return g.val.(*bool)
}

func (g *AbdsItem) pi() *int {

	switch g.val.(type) {
	case *int:
		break
	default:
		g.S(g.vi())
	}
	return (g.val.(*int))
}

func (g *AbdsItem) pi8() *int8 {

	switch g.val.(type) {
	case *int8:
		break
	default:
		g.S(g.vi8())
	}
	return (g.val.(*int8))
}

func (g *AbdsItem) pi16() *int16 {

	switch g.val.(type) {
	case *int:
		break
	default:
		g.S(g.vi16())
	}
	return (g.val.(*int16))
}

func (g *AbdsItem) pi32() *int32 {

	switch g.val.(type) {
	case *int32:
		break
	default:
		g.S(g.vi32())
	}
	return (g.val.(*int32))
}

func (g *AbdsItem) pi64() *int64 {

	switch g.val.(type) {
	case *int64:
		break
	default:
		g.S(g.vi64())
	}
	return (g.val.(*int64))
}

func (g *AbdsItem) pu() *uint {

	switch g.val.(type) {
	case *uint:
		break
	default:
		g.S(g.vu())
	}
	return (g.val.(*uint))
}

func (g *AbdsItem) pu8() *uint8 {

	switch g.val.(type) {
	case *uint8:
		break
	default:
		g.S(g.vu8())
	}
	return (g.val.(*uint8))
}

func (g *AbdsItem) pu16() *uint16 {

	switch g.val.(type) {
	case *uint16:
		break
	default:
		g.S(g.vu16())
	}
	return (g.val.(*uint16))
}

func (g *AbdsItem) pu32() *uint32 {

	switch g.val.(type) {
	case *uint32:
		break
	default:
		g.S(g.vu32())
	}
	return (g.val.(*uint32))
}

func (g *AbdsItem) pu64() *uint64 {

	switch g.val.(type) {
	case *uint64:
		break
	default:
		g.S(g.vu64())
	}
	return (g.val.(*uint64))
}

func (g *AbdsItem) pf32() *float32 {

	switch g.val.(type) {
	case *float32:
		break
	default:
		g.S(g.vf32())
	}
	return (g.val.(*float32))
}

func (g *AbdsItem) pf64() *float64 {

	switch g.val.(type) {
	case *float64:
		break
	default:
		g.S(g.vf64())
	}
	return (g.val.(*float64))
}

func (g *AbdsItem) pc64() *complex64 {

	switch g.val.(type) {
	case *complex64:
		break
	default:
		g.S(g.vc64())
	}
	return (g.val.(*complex64))
}

func (g *AbdsItem) pc128() *complex128 {

	switch g.val.(type) {
	case *complex128:
		break
	default:
		g.S(g.vc128())
	}
	return (g.val.(*complex128))
}

func (g *AbdsItem) ps() *string {

	switch g.val.(type) {
	case *string:
		break
	default:
		g.S(g.vs())
	}
	return g.val.(*string)
}

func (g *AbdsItem) pb() *[]byte {

	switch g.val.(type) {
	case *[]byte:
		break
	default:
		g.S(g.vB())
	}
	return g.val.(*[]byte)
}

func (g *AbdsItem) pbyte() *[]byte {
	return g.pb()
}

func (g *AbdsItem) pabds() *Abds {
	return g.pch()
}

func (g *AbdsItem) pchild() *Abds {
	return g.pch()
}

func (g *AbdsItem) pch() *Abds {
	switch g.val.(type) {
	case *Abds:
		return g.val.(*Abds)
	}
	return nil
}

func (g *AbdsItem) ptran() *AbdsTransform {
	switch g.val.(type) {
	case *AbdsTransform:
		return g.val.(*AbdsTransform)
	}
	return nil
}
