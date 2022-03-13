package abds

func (g *AbdsItem) Pb() *bool {
	switch g.val.(type) {
	case *bool:
		break
	default:
		g.S(g.Vb())
	}
	return g.val.(*bool)
}

func (g *AbdsItem) Pi() *int {
	switch g.val.(type) {
	case *int:
		break
	default:
		g.S(g.Vi())
	}
	return (g.val.(*int))
}

func (g *AbdsItem) Pi8() *int8 {
	switch g.val.(type) {
	case *int8:
		break
	default:
		g.S(g.Vi8())
	}
	return (g.val.(*int8))
}

func (g *AbdsItem) Pi16() *int16 {
	switch g.val.(type) {
	case *int:
		break
	default:
		g.S(g.Vi16())
	}
	return (g.val.(*int16))
}

func (g *AbdsItem) Pi32() *int32 {
	switch g.val.(type) {
	case *int32:
		break
	default:
		g.S(g.Vi32())
	}
	return (g.val.(*int32))
}

func (g *AbdsItem) Pi64() *int64 {
	switch g.val.(type) {
	case *int64:
		break
	default:
		g.S(g.Vi64())
	}
	return (g.val.(*int64))
}

func (g *AbdsItem) Pu() *uint {
	switch g.val.(type) {
	case *uint:
		break
	default:
		g.S(g.Vu())
	}
	return (g.val.(*uint))
}

func (g *AbdsItem) Pu8() *uint8 {
	switch g.val.(type) {
	case *uint8:
		break
	default:
		g.S(g.Vu8())
	}
	return (g.val.(*uint8))
}

func (g *AbdsItem) Pu16() *uint16 {
	switch g.val.(type) {
	case *uint16:
		break
	default:
		g.S(g.Vu16())
	}
	return (g.val.(*uint16))
}

func (g *AbdsItem) Pu32() *uint32 {
	switch g.val.(type) {
	case *uint32:
		break
	default:
		g.S(g.Vu32())
	}
	return (g.val.(*uint32))
}

func (g *AbdsItem) Pu64() *uint64 {
	switch g.val.(type) {
	case *uint64:
		break
	default:
		g.S(g.Vu64())
	}
	return (g.val.(*uint64))
}

func (g *AbdsItem) Pf32() *float32 {
	switch g.val.(type) {
	case *float32:
		break
	default:
		g.S(g.Vf32())
	}
	return (g.val.(*float32))
}

func (g *AbdsItem) Pf64() *float64 {
	switch g.val.(type) {
	case *float64:
		break
	default:
		g.S(g.Vf64())
	}
	return (g.val.(*float64))
}

func (g *AbdsItem) Pc64() *complex64 {
	switch g.val.(type) {
	case *complex64:
		break
	default:
		g.S(g.Vc64())
	}
	return (g.val.(*complex64))
}

func (g *AbdsItem) Pc128() *complex128 {
	switch g.val.(type) {
	case *complex128:
		break
	default:
		g.S(g.Vc128())
	}
	return (g.val.(*complex128))
}

func (g *AbdsItem) Ps() *string {
	switch g.val.(type) {
	case *string:
		break
	default:
		g.S(g.Vs())
	}
	return g.val.(*string)
}

func (g *AbdsItem) Pbyte() *[]byte {
	switch g.val.(type) {
	case *[]byte:
		break
	default:
		g.S(g.Vb())
	}
	return g.val.(*[]byte)
}

func (g *AbdsItem) Prune() *[]rune {
	switch g.val.(type) {
	case *[]rune:
		break
	default:
		g.S(g.Vrune())
	}
	return g.val.(*[]rune)
}

func (g *AbdsItem) Pabds() *Abds {
	switch g.val.(type) {
	case *Abds:
		return g.val.(*Abds)
	}
	return nil
}

func (g *AbdsItem) Ptran() interface{} {
	switch g.val.(type) {
	case *AbdsTransform:
		return g.val
	}
	return nil
}
