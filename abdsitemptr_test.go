package abds

import (
	"testing"
)

func TestItemPtrOperations(t *testing.T) {

	g := New()
	gitem := g.g("TAG")

	for _, git := range gatb {

		gitem.S(git)

		switch git.(type) {
		case *bool:
			gv := gitem.pB()
			*gv = !*git.(*bool)
			if gitem.vB() != *gv {
				t.Error("Ptr Bool not sticky")
			}
		case *int:
			gv := gitem.pi()
			*gv = *git.(*int) + 1
			if gitem.vi() != *gv {
				t.Error("Ptr int not sticky")
			}
		case *int8:
			gv := gitem.pi8()
			*gv = *git.(*int8) + 1
			if gitem.vi8() != *gv {
				t.Error("Ptr int8 not sticky")
			}
		case *int16:
			gv := gitem.pi16()
			*gv = *git.(*int16) + 1
			if gitem.vi16() != *gv {
				t.Error("Ptr int16 not sticky")
			}
		case *int32:
			gv := gitem.pi32()
			*gv = *git.(*int32) + 1
			if gitem.vi32() != *gv {
				t.Error("Ptr int32 not sticky")
			}
		case *int64:
			gv := gitem.pi64()
			*gv = *git.(*int64) + 1
			if gitem.vi64() != *gv {
				t.Error("Ptr int64 not sticky")
			}
		case *uint:
			gv := gitem.pu()
			*gv = *git.(*uint) + 1
			if gitem.vu() != *gv {
				t.Error("Ptr uint not sticky")
			}
		case *uint8:
			gv := gitem.pu8()
			*gv = *git.(*uint8) + 1
			if gitem.vu8() != *gv {
				t.Error("Ptr uint8 not sticky")
			}
		case *uint16:
			gv := gitem.pu16()
			*gv = *git.(*uint16) + 1
			if gitem.vu16() != *gv {
				t.Error("Ptr uint16 not sticky")
			}
		case *uint32:
			gv := gitem.pu32()
			*gv = *git.(*uint32) + 1
			if gitem.vu32() != *gv {
				t.Error("Ptr uint32 not sticky")
			}
		case *uint64:
			gv := gitem.pu64()
			*gv = *git.(*uint64) + 1
			if gitem.vu64() != *gv {
				t.Error("Ptr uint64 not sticky")
			}
		case *float32:
			gv := gitem.pf32()
			*gv = *git.(*float32) + 1
			if gitem.vf32() != *gv {
				t.Error("Ptr float32 not sticky")
			}
		case *float64:
			gv := gitem.pf64()
			*gv = *git.(*float64) + 1
			if gitem.vf64() != *gv {
				t.Error("Ptr float64 not sticky")
			}
		case *complex64:
			gv := gitem.pc64()
			*gv = *git.(*complex64) + 1
			if gitem.vc64() != *gv {
				t.Error("Ptr complex64 not sticky")
			}
		case *complex128:
			gv := gitem.pc128()
			*gv = *git.(*complex128) + 1
			if gitem.vc128() != *gv {
				t.Error("Ptr complex128 not sticky")
			}
		case *string:
			gv := gitem.ps()
			*gv = "Changed"
			if gitem.vs() != *gv {
				t.Error("Ptr string not sticky")
			}

		}

	}

}

func TestItemPtrAbdsOperations(t *testing.T) {

	//g := New()
	//gitem := g.G("TAG")

}
