package abds

//Allow for daisy-chaining value set
func (g *Abds) S(tag string, val interface{}) *Abds {
	g.G(tag).S(val, g.errs)
	return g
}
