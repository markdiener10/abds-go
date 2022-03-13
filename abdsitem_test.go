package abds

import (
	"errors"
	"testing"
)

//Test an externally define structure so we can integrate
//project specific structures into the abds framework

var gtestvals []interface{}

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
		return errors.New("Invalid Abds parameter passed")
	}
	//g.what = input.T("WHAT").V()
	return nil
}

func init() {

	gtestvals = []interface{}{
		nil,                    //0
		true,                   //1
		false,                  //2
		int(11),                //3
		int8(12),               //4
		int16(13),              //5
		int32(14),              //6
		int64(15),              //7
		uint(21),               //8
		uint8(22),              //9
		uint16(23),             //10
		uint32(24),             //11
		uint64(25),             //12
		float32(30.02),         //13
		float64(31.03),         //14
		complex(41, -43),       //15
		complex(42, -44),       //16
		[]byte{51, 52, 53, 54}, //17
		"string",               //18
		nil,                    //19 Ptr Bool true
		nil,                    //20 Ptr Bool false
		nil,                    //21 Ptr int
		nil,                    //22 Ptr int8
		nil,                    //23 Ptr int16
		nil,                    //24 Ptr int32
		nil,                    //25 Ptr int64
		nil,                    //26 Ptr uint
		nil,                    //27 Ptr uint8
		nil,                    //28 Ptr uint16
		nil,                    //29 Ptr uint32
		nil,                    //30 Ptr uint64
		nil,                    //31 Ptr float32
		nil,                    //32 Ptr float64
		nil,                    //33 Ptr complex64
		nil,                    //34 Ptr complex128
		nil,                    //35 Ptr []byte
		nil,                    //36 Ptr string
		nil,                    //37 Ptr Abds
		nil,                    //38 Ptr AbdsTransform
	}

	//Load up the pointers to values
	gtrue := true
	gtestvals[19] = &gtrue

	gfalse := false
	gtestvals[20] = &gfalse

	gint := int(60)
	gtestvals[21] = &gint

	gint8 := int8(61)
	gtestvals[22] = &gint8

	gint16 := int16(62)
	gtestvals[23] = &gint16

	gint32 := int32(63)
	gtestvals[24] = &gint32

	gint64 := int64(64)
	gtestvals[25] = &gint64

	guint := uint(70)
	gtestvals[26] = &guint

	guint8 := uint8(71)
	gtestvals[27] = &guint8

	guint16 := uint16(72)
	gtestvals[28] = &guint16

	guint32 := uint32(73)
	gtestvals[29] = &guint32

	guint64 := uint64(74)
	gtestvals[30] = &guint64

	gfloat32 := float32(80.1)
	gtestvals[31] = &gfloat32

	gfloat64 := float64(81.1)
	gtestvals[32] = &gfloat64

	gcomplex64 := complex64(complex(90, 91))
	gtestvals[33] = &gcomplex64

	gcomplex128 := complex128(complex(100, 101))
	gtestvals[34] = &gcomplex128

	gbyteptr := &[]byte{111, 112, 113, 114}
	gtestvals[35] = gbyteptr

	gstring := "stringptr"
	gtestvals[36] = &gstring

	gabds := New()
	gtestvals[37] = gabds

	gtrans := TestTransform{what: "string"}
	gtestvals[38] = &gtrans

}

func TestItemOperations(t *testing.T) {

	g := New()

	gitem := g.G("TAG")
	for _, git := range gtestvals {

		gitem.S(git)
		gitem.V()
		gitem.Vi()
		/*
			gitem.Vi8()
			gitem.Vi16()
			gitem.Vi32()
			gitem.Vi64()
			gitem.Vu()
			gitem.Vu8()
			gitem.Vu16()
			gitem.Vu32()
			gitem.Vu64()
			gitem.Vs()
		*/

	}

	//gitem.

}
