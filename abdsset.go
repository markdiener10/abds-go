package abds

func (g *Abds) S(tag string, val interface{}, parms ...interface{}) {
	g.G(tag, parms...).S(val, parms...)
}
