package abds

import (
	"testing"
)

//Test an externally define structure so we can integrate
//project specific structures into the abds framework

type TestStruct struct {
	what string
}

var gatb []interface{}

func init() {

	var gv interface{} = nil

	gatb = make([]interface{}, 0)

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
	gatb = append(gatb, TestStruct{what: "whatnonpointer"})
	gatb = append(gatb, Abds{root: true, link: NewErrs(), vals: make([]*AbdsItem, 0), tags: NewTags()})
	gatb = append(gatb, map[uint]int{
		50: 51,
		52: 53,
		54: 55,
		56: 67,
	})

	ga := []int{61, 62, 63, 64}
	gatb = append(gatb, ga[1:2])

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
	gatb = append(gatb, &TestStruct{what: "whatpointer"})
	gv = map[uint]int{
		400: 401,
		402: 403,
		404: 405,
		406: 407,
	}
	gatb = append(gatb, &gv)

	gv = ga[1:2]
	gatb = append(gatb, &gv)

	gatb = append(gatb, New())

}

func TestItemOperations(t *testing.T) {

	//var idx int

	g := New()
	var val interface{}
	for _, val = range gatb {
		g.S(val)
	}

	if g.IsErr() {
		t.Errorf("Test Item Operations errors A:%s", g.Errs().Content())
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
