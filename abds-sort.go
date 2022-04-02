package abds

type AbdsSortFunc func(p1 *AbdsIter, p2 *AbdsIter) bool

//Sorting functions (Depending on Size O(n))
func (g *Abds) Sort(fn AbdsSortFunc) {

	if fn == nil {
		return
	}
	var swapped bool
	var itera AbdsIter
	var iterb AbdsIter

	itera.pAbds = g
	iterb.pAbds = g

	itera.reset()
	itera.reset()

	n := len(g.vals)

	for i := 0; i < n-1; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			itera.pItem = g.vals[j]
			iterb.pItem = g.vals[j+1]
			if fn(&itera, &iterb) == true {
				g.vals[j] = iterb.pItem
				g.vals[j+1] = itera.pItem
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}
