package abds

import (
	"testing"
)

func TestItemPtrOperations(t *testing.T) {

	g := New()
	gitem := g.G("TAG")

	for _, git := range gtestvals {

		gitem.S(git)

		switch git.(type) {
		case *bool:
			gv := gitem.Pb()
			*gv = !*git.(*bool)
			if gitem.Vb() != *gv {
				t.Error("Ptr Bool not sticky")
			}
		case *int:
			gv := gitem.Pi()
			*gv = *git.(*int) + 1
			if gitem.Vi() != *gv {
				t.Error("Ptr int not sticky")
			}
		case *int8:
			gv := gitem.Pi8()
			*gv = *git.(*int8) + 1
			if gitem.Vi8() != *gv {
				t.Error("Ptr int8 not sticky")
			}
		case *int16:
			gv := gitem.Pi16()
			*gv = *git.(*int16) + 1
			if gitem.Vi16() != *gv {
				t.Error("Ptr int16 not sticky")
			}
		case *int32:
			gv := gitem.Pi32()
			*gv = *git.(*int32) + 1
			if gitem.Vi32() != *gv {
				t.Error("Ptr int32 not sticky")
			}
		case *int64:
			gv := gitem.Pi64()
			*gv = *git.(*int64) + 1
			if gitem.Vi64() != *gv {
				t.Error("Ptr int64 not sticky")
			}
		case *uint:
			gv := gitem.Pu()
			*gv = *git.(*uint) + 1
			if gitem.Vu() != *gv {
				t.Error("Ptr uint not sticky")
			}
		case *uint8:
			gv := gitem.Pu8()
			*gv = *git.(*uint8) + 1
			if gitem.Vu8() != *gv {
				t.Error("Ptr uint8 not sticky")
			}
		case *uint16:
			gv := gitem.Pu16()
			*gv = *git.(*uint16) + 1
			if gitem.Vu16() != *gv {
				t.Error("Ptr uint16 not sticky")
			}
		case *uint32:
			gv := gitem.Pu32()
			*gv = *git.(*uint32) + 1
			if gitem.Vu32() != *gv {
				t.Error("Ptr uint32 not sticky")
			}
		case *uint64:
			gv := gitem.Pu64()
			*gv = *git.(*uint64) + 1
			if gitem.Vu64() != *gv {
				t.Error("Ptr uint64 not sticky")
			}
		case *float32:
			gv := gitem.Pf32()
			*gv = *git.(*float32) + 1
			if gitem.Vf32() != *gv {
				t.Error("Ptr float32 not sticky")
			}
		case *float64:
			gv := gitem.Pf64()
			*gv = *git.(*float64) + 1
			if gitem.Vf64() != *gv {
				t.Error("Ptr float64 not sticky")
			}
		case *complex64:
			gv := gitem.Pc64()
			*gv = *git.(*complex64) + 1
			if gitem.Vc64() != *gv {
				t.Error("Ptr complex64 not sticky")
			}
		case *complex128:
			gv := gitem.Pc128()
			*gv = *git.(*complex128) + 1
			if gitem.Vc128() != *gv {
				t.Error("Ptr complex128 not sticky")
			}
		case *string:
			gv := gitem.Ps()
			*gv = "Changed"
			if gitem.Vs() != *gv {
				t.Error("Ptr string not sticky")
			}

		}

	}

}

func TestItemPtrAbdsOperations(t *testing.T) {

	//g := New()
	//gitem := g.G("TAG")
	

}
