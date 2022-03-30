package abds

type AbdsSortFunc func(p1 *AbdsItem, p2 *AbdsItem) bool

//Sorting functions (Depending on Size O(n))
func (g *Abds) Sort(fn AbdsSortFunc) {

	if fn == nil {
		return
	}
	var swapped bool
	var pitema *AbdsItem
	var pitemb *AbdsItem

	n := len(g.vals)

	for i := 0; i < n-1; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			pitema = g.vals[j]
			pitemb = g.vals[j+1]
			if fn(pitema, pitemb) == true {
				g.vals[j] = pitemb
				g.vals[j+1] = pitema
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}
