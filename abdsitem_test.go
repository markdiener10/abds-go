package abds

import (
	"fmt"
	"testing"
)

//Test an externally define structure so we can integrate
//project specific structures into the abds framework

var gatb []interface{}
var gats []interface{}
var gate []interface{} //Error cases

type TestNotTransform struct {
	what string
}

type TestTransform struct {
	what string
}

func (g *TestTransform) Pack() *Abds {
	abds := New()
	abds.S("WHAT", g.what)
	return abds
}

func (g *TestTransform) UnPack(input *Abds) error {
	if input == nil {
		return fmt.Errorf("ABDS Invalid Abds parameter passed")
	}
	//g.what = input.T("WHAT").V()
	return nil
}

func init() {

	var gv interface{} = nil

	gatb = make([]interface{}, 0)
	gats = make([]interface{}, 0)
	gate = make([]interface{}, 0)

	gatb = append(gatb, nil)
	gatb = append(gatb, true)
	gatb = append(gatb, false)
	gatb = append(gatb, int(11))
	gatb = append(gatb, int8(12))
	gatb = append(gatb, int16(13))
	gatb = append(gatb, int32(14))
	gatb = append(gatb, int64(15))
	gatb = append(gatb, uint(21))
	gatb = append(gatb, uint8(22))
	gatb = append(gatb, uint16(23))
	gatb = append(gatb, uint32(24))
	gatb = append(gatb, uint64(25))
	gatb = append(gatb, float32(30.02))
	gatb = append(gatb, float64(31.03))
	gatb = append(gatb, complex(41, -42))
	gatb = append(gatb, complex(43, -44))
	gatb = append(gatb, "stringtest")

	gv = true
	gatb = append(gatb, &gv)
	gv = false
	gatb = append(gatb, &gv)
	gv = int(60)
	gatb = append(gatb, &gv)
	gv = int8(61)
	gatb = append(gatb, &gv)
	gv = int16(62)
	gatb = append(gatb, &gv)
	gv = int32(63)
	gatb = append(gatb, &gv)
	gv = int64(64)
	gatb = append(gatb, &gv)
	gv = uint(70)
	gatb = append(gatb, &gv)
	gv = uint8(71)
	gatb = append(gatb, &gv)
	gv = uint16(72)
	gatb = append(gatb, &gv)
	gv = uint32(73)
	gatb = append(gatb, &gv)
	gv = uint64(74)
	gatb = append(gatb, &gv)
	gv = float32(80.1)
	gatb = append(gatb, &gv)
	gv = float64(81.1)
	gatb = append(gatb, &gv)
	gv = complex64(complex(90, 91))
	gatb = append(gatb, &gv)
	gv = complex128(complex(100, 101))
	gatb = append(gatb, &gv)
	gv = "stringptr"
	gatb = append(gatb, &gv)
	gv = []byte{111, 112, 113, 114}
	gatb = append(gatb, &gv)
	gv = []float32{300.1, 300.2, 300.3, 300.4}
	gatb = append(gatb, &gv)
	gv = map[uint]int{
		400: 401,
		402: 403,
		404: 405,
		406: 407,
	}
	gatb = append(gatb, &gv)

	//Structural test cases
	gats = append(gats, New())
	gats = append(gats, &TestTransform{what: "transform"})
	gats = append(gats, &TestNotTransform{what: "nottransform"})

	//Several illegal entries
	gate = append(gate, Abds{})
	gate = append(gate, TestTransform{what: "transform"})
	gate = append(gate, TestNotTransform{what: "nottransform"})
	gate = append(gate, []byte{111, 112, 113, 114})
	gate = append(gate, []float32{300.1, 300.2, 300.3, 300.4})
	gate = append(gate, map[uint]int{
		400: 401,
		402: 403,
		404: 405,
		406: 407,
	})

}

func TestItemOperations(t *testing.T) {

	var err error
	var idx int

	_ = idx

	g := New()
	gitem := g.g("TAG")
	var gtest interface{}
	for idx, gtest = range gatb {
		err = gitem.S(gtest)
		if err != nil {
			g.Log(err)
		}
	}

	if g.IsErr() {
		t.Error("Test Item Operations errors A")
	}

	g.ErrorClear()

	for idx, gtest = range gats {
		err = gitem.S(gtest)
		if err != nil {
			g.Log(err)
		}

	}

	if g.IsErr() {
		t.Error("Test Item Operations errors B")
	}

	g.ErrorClear()
	for idx, gtest = range gate {
		err = gitem.S(gtest)
		if err != nil {
			g.Log(err)
		}
		if !g.IsErr() {
			t.Error("Test Item Operations should have errors C")
		}
		g.ErrorClear()
	}

}

func TestTopLevelArrayOperations(t *testing.T) {

	//Array operations
	g := New()
	if g.Len() != 0 {
		t.Errorf("Abds set to array mode but reports as non-empty:%d", g.Len())
	}
	var idx uint
	for idx = 1; idx <= 10; idx++ {
		g.S(idx)
	}
	if g.Len() != 10 {
		t.Errorf("Abds set to array mode but fails to add 10 items:%d", g.Len())
	}
	for idx = 1; idx <= 10; idx++ {
		if g.Vu(idx) != idx {
			t.Errorf("Abds set to array mode but reports as incorrect value:%d : %d", g.Vu(idx), idx)
		}
	}

	g = New()
	for idx = 1; idx <= 10; idx++ {
		g.S(idx, idx)
	}
	if g.Len() != 10 {
		t.Errorf("Abds set to array mode but fails to add 10 items:%d", g.Len())
	}
	for idx = 1; idx <= 10; idx++ {
		if g.Vu(idx) != idx {
			t.Errorf("Abds set to array mode but reports as incorrect value:%d : %d", g.Vu(idx), idx)
		}
	}

}

func TestTopLevelUninitialized(t *testing.T) {

	g := New()
	*g.Pu(1) += 2
	*g.Pu("TAG")++

	if g.Vu(1) != 2 {
		t.Errorf("Abds uinitialized array access not correct (2):%d", g.Vu(1))
	}
	if g.Vu("TAG") != 1 {
		t.Errorf("Abds uinitialized tag access not correct (1):%d", g.Vu("TAG"))
	}

}
