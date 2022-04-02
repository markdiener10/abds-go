package abds

import (
	"fmt"
)

func (g *Abds) Merge(from *Abds) error {
	return nil
}

func (g *Abds) Copy(src *Abds) error {
	if src == nil {
		return fmt.Errorf("ABDS Invalid Abds parameter passed for copy")
	}
	//This will replace the current payload with a new one (and release memory)
	g.vals = make([]*AbdsItem, 0)
	//Almost like a C++11 friend function
	g.vals = append(g.vals, (*src).vals...)
	return nil
}
