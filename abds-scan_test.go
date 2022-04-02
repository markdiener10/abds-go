package abds

import (
	"testing"
)

func TestTopLevelScan(t *testing.T) {

	g := New()
	for idx := 0; idx < 10; idx++ {
		g.S(idx + 1)
	}

	fnScan := func(level uint, piter *AbdsIter, pCap *Abds) (terminate bool) {
		*pCap.Pu("COUNT")++
		return
	}

	pCap := g.Scan(fnScan, false)
	if pCap.Vu("COUNT") != 10 {
		t.Errorf("Abds scan test failed:%d", pCap.Vu("COUNT"))
	}

}
