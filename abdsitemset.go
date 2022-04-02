package abds

import (
	"fmt"
)

func (g *AbdsItem) S(val interface{}) error {
	if g == nil {
		return fmt.Errorf("ABDS item received no parameters")
	}
	var err error
	if val, err = checkval(val); err != nil {
		return err
	}
	g.val = val
	return nil
}
