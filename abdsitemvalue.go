package abds

import (
	"math/cmplx"
	"strconv"
)

func (g *AbdsItem) vB() bool {

	var val uint64 = 0

	switch g.val.(type) {
	case nil:
		return false
	case *bool:
		return bool(*g.val.(*bool))
	case *int:
		val = uint64(*g.val.(*int))
	case *int8:
		val = uint64(*g.val.(*int8))
	case *int16:
		val = uint64(*g.val.(*int16))
	case *int32:
		val = uint64(*g.val.(*int32))
	case *int64:
		val = uint64(*g.val.(*int64))
	case *uint:
		val = uint64(*g.val.(*uint))
	case *uint8:
		val = uint64(*g.val.(*uint8))
	case *uint16:
		val = uint64(*g.val.(*uint16))
	case *uint32:
		val = uint64(*g.val.(*uint32))
	case *uint64:
		val = uint64(*g.val.(*uint64))
	case *float32:
		val = uint64(*g.val.(*float32))
	case *float64:
		val = uint64(*g.val.(*float64))
	case *complex64:
		val = uint64(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		val = uint64(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseUint(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			gi = 0
		}
		val = gi
	case *string:
		gi, err := strconv.ParseUint(string(*g.val.(*string)), 10, 64)
		if err != nil {
			gi = 0
		}
		val = gi
	}
	if val == 0 {
		return false
	}
	return true
}

func (g *AbdsItem) vi() int {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return 1
		}
		break
	case *int:
		return int(*g.val.(*int))
	case *int8:
		return int(*g.val.(*int8))
	case *int16:
		return int(*g.val.(*int16))
	case *int32:
		return int(*g.val.(*int32))
	case *int64:
		return int(*g.val.(*int64))
	case *uint:
		return int(*g.val.(*uint))
	case *uint8:
		return int(*g.val.(*uint8))
	case *uint16:
		return int(*g.val.(*uint16))
	case *uint32:
		return int(*g.val.(*uint32))
	case *uint64:
		return int(*g.val.(*uint64))
	case *float32:
		return int(*g.val.(*float32))
	case *float64:
		return int(*g.val.(*float64))
	case *complex64:
		return int(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		return int(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return int(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return int(gi)
	}
	return 0
}

func (g *AbdsItem) vi8() int8 {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return 1
		}
		break
	case *int:
		return int8(*g.val.(*int))
	case *int8:
		return int8(*g.val.(*int8))
	case *int16:
		return int8(*g.val.(*int16))
	case *int32:
		return int8(*g.val.(*int32))
	case *int64:
		return int8(*g.val.(*int64))
	case *uint:
		return int8(*g.val.(*uint))
	case *uint8:
		return int8(*g.val.(*uint8))
	case *uint16:
		return int8(*g.val.(*uint16))
	case *uint32:
		return int8(*g.val.(*uint32))
	case *uint64:
		return int8(*g.val.(*uint64))
	case *float32:
		return int8(*g.val.(*float32))
	case *float64:
		return int8(*g.val.(*float64))
	case *complex64:
		return int8(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		return int8(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return int8(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return int8(gi)
	}
	return 0
}

func (g *AbdsItem) vi16() int16 {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return 1
		}
		break
	case *int:
		return int16(*g.val.(*int))
	case *int8:
		return int16(*g.val.(*int8))
	case *int16:
		return int16(*g.val.(*int16))
	case *int32:
		return int16(*g.val.(*int32))
	case *int64:
		return int16(*g.val.(*int64))
	case *uint:
		return int16(*g.val.(*uint))
	case *uint8:
		return int16(*g.val.(*uint8))
	case *uint16:
		return int16(*g.val.(*uint16))
	case *uint32:
		return int16(*g.val.(*uint32))
	case *uint64:
		return int16(*g.val.(*uint64))
	case *float32:
		return int16(*g.val.(*float32))
	case *float64:
		return int16(*g.val.(*float64))
	case *complex64:
		return int16(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		return int16(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return int16(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			return 0
		}
		return int16(gi)
	}
	return 0
}

func (g *AbdsItem) vi32() int32 {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return 1
		}
		break
	case *int:
		return int32(*g.val.(*int))
	case *int8:
		return int32(*g.val.(*int8))
	case *int16:
		return int32(*g.val.(*int16))
	case *int32:
		return int32(*g.val.(*int32))
	case *int64:
		return int32(*g.val.(*int64))
	case *uint:
		return int32(*g.val.(*uint))
	case *uint8:
		return int32(*g.val.(*uint8))
	case *uint16:
		return int32(*g.val.(*uint16))
	case *uint32:
		return int32(*g.val.(*uint32))
	case *uint64:
		return int32(*g.val.(*uint64))
	case *float32:
		return int32(*g.val.(*float32))
	case *float64:
		return int32(*g.val.(*float64))
	case *complex64:
		return int32(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		return int32(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return int32(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return int32(gi)
	}
	return 0
}

func (g *AbdsItem) vi64() int64 {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return 1
		}
		break
	case *int:
		return int64(*g.val.(*int))
	case *int8:
		return int64(*g.val.(*int8))
	case *int16:
		return int64(*g.val.(*int16))
	case *int32:
		return int64(*g.val.(*int32))
	case *int64:
		return int64(*g.val.(*int64))
	case *uint:
		return int64(*g.val.(*uint))
	case *uint8:
		return int64(*g.val.(*uint8))
	case *uint16:
		return int64(*g.val.(*uint16))
	case *uint32:
		return int64(*g.val.(*uint32))
	case *uint64:
		return int64(*g.val.(*uint64))
	case *float32:
		return int64(*g.val.(*float32))
	case *float64:
		return int64(*g.val.(*float64))
	case *complex64:
		return int64(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		return int64(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return int64(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return int64(gi)
	}
	return 0
}

func (g *AbdsItem) vu() uint {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return 1
		}
		break
	case *int:
		return uint(*g.val.(*int))
	case *int8:
		return uint(*g.val.(*int8))
	case *int16:
		return uint(*g.val.(*int16))
	case *int32:
		return uint(*g.val.(*int32))
	case *int64:
		return uint(*g.val.(*int64))
	case *uint:
		return uint(*g.val.(*uint))
	case *uint8:
		return uint(*g.val.(*uint8))
	case *uint16:
		return uint(*g.val.(*uint16))
	case *uint32:
		return uint(*g.val.(*uint32))
	case *uint64:
		return uint(*g.val.(*uint64))
	case *float32:
		return uint(*g.val.(*float32))
	case *float64:
		return uint(*g.val.(*float64))
	case *complex64:
		return uint(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		return uint(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return uint(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return uint(gi)
	}
	return 0
}

func (g *AbdsItem) vu8() uint8 {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return 1
		}
		break
	case *int:
		return uint8(*g.val.(*int))
	case *int8:
		return uint8(*g.val.(*int8))
	case *int16:
		return uint8(*g.val.(*int16))
	case *int32:
		return uint8(*g.val.(*int32))
	case *int64:
		return uint8(*g.val.(*int64))
	case *uint:
		return uint8(*g.val.(*uint))
	case *uint8:
		return uint8(*g.val.(*uint8))
	case *uint16:
		return uint8(*g.val.(*uint16))
	case *uint32:
		return uint8(*g.val.(*uint32))
	case *uint64:
		return uint8(*g.val.(*uint64))
	case *float32:
		return uint8(*g.val.(*float32))
	case *float64:
		return uint8(*g.val.(*float64))
	case *complex64:
		return uint8(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		return uint8(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return uint8(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return uint8(gi)
	}
	return 0
}

func (g *AbdsItem) vu16() uint16 {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return 1
		}
		break
	case *int:
		return uint16(*g.val.(*int))
	case *int8:
		return uint16(*g.val.(*int8))
	case *int16:
		return uint16(*g.val.(*int16))
	case *int32:
		return uint16(*g.val.(*int32))
	case *int64:
		return uint16(*g.val.(*int64))
	case *uint:
		return uint16(*g.val.(*uint))
	case *uint8:
		return uint16(*g.val.(*uint8))
	case *uint16:
		return uint16(*g.val.(*uint16))
	case *uint32:
		return uint16(*g.val.(*uint32))
	case *uint64:
		return uint16(*g.val.(*uint64))
	case *float32:
		return uint16(*g.val.(*float32))
	case *float64:
		return uint16(*g.val.(*float64))
	case *complex64:
		return uint16(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		return uint16(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return uint16(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return uint16(gi)
	}
	return 0
}

func (g *AbdsItem) vu32() uint32 {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return 1
		}
		break
	case *int:
		return uint32(*g.val.(*int))
	case *int8:
		return uint32(*g.val.(*int8))
	case *int16:
		return uint32(*g.val.(*int16))
	case *int32:
		return uint32(*g.val.(*int32))
	case *int64:
		return uint32(*g.val.(*int64))
	case *uint:
		return uint32(*g.val.(*uint))
	case *uint8:
		return uint32(*g.val.(*uint8))
	case *uint16:
		return uint32(*g.val.(*uint16))
	case *uint32:
		return uint32(*g.val.(*uint32))
	case *uint64:
		return uint32(*g.val.(*uint64))
	case *float32:
		return uint32(*g.val.(*float32))
	case *float64:
		return uint32(*g.val.(*float64))
	case *complex64:
		return uint32(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		return uint32(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return uint32(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return uint32(gi)
	}
	return 0
}

func (g *AbdsItem) vu64() uint64 {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return 1
		}
		break
	case *int:
		return uint64(*g.val.(*int))
	case *int8:
		return uint64(*g.val.(*int8))
	case *int16:
		return uint64(*g.val.(*int16))
	case *int32:
		return uint64(*g.val.(*int32))
	case *int64:
		return uint64(*g.val.(*int64))
	case *uint:
		return uint64(*g.val.(*uint))
	case *uint8:
		return uint64(*g.val.(*uint8))
	case *uint16:
		return uint64(*g.val.(*uint16))
	case *uint32:
		return uint64(*g.val.(*uint32))
	case *uint64:
		return uint64(*g.val.(*uint64))
	case *float32:
		return uint64(*g.val.(*float32))
	case *float64:
		return uint64(*g.val.(*float64))
	case *complex64:
		return uint64(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		return uint64(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return uint64(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return uint64(gi)
	}
	return 0
}

func (g *AbdsItem) vf32() float32 {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return 1.0
		}
		break
	case *int:
		return float32(*g.val.(*int))
	case *int8:
		return float32(*g.val.(*int8))
	case *int16:
		return float32(*g.val.(*int16))
	case *int32:
		return float32(*g.val.(*int32))
	case *int64:
		return float32(*g.val.(*int64))
	case *uint:
		return float32(*g.val.(*uint))
	case *uint8:
		return float32(*g.val.(*uint8))
	case *uint16:
		return float32(*g.val.(*uint16))
	case *uint32:
		return float32(*g.val.(*uint32))
	case *uint64:
		return float32(*g.val.(*uint64))
	case *float32:
		return float32(*g.val.(*float32))
	case *float64:
		return float32(*g.val.(*float64))
	case *complex64:
		return float32(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		return float32(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return float32(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return float32(gi)
	}
	return 0.0
}

func (g *AbdsItem) vf64() float64 {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return 1.0
		}
		break
	case *int:
		return float64(*g.val.(*int))
	case *int8:
		return float64(*g.val.(*int8))
	case *int16:
		return float64(*g.val.(*int16))
	case *int32:
		return float64(*g.val.(*int32))
	case *int64:
		return float64(*g.val.(*int64))
	case *uint:
		return float64(*g.val.(*uint))
	case *uint8:
		return float64(*g.val.(*uint8))
	case *uint16:
		return float64(*g.val.(*uint16))
	case *uint32:
		return float64(*g.val.(*uint32))
	case *uint64:
		return float64(*g.val.(*uint64))
	case *float32:
		return float64(*g.val.(*float32))
	case *float64:
		return float64(*g.val.(*float64))
	case *complex64:
		return float64(cmplx.Abs(complex128(*g.val.(*complex64))))
	case *complex128:
		return float64(cmplx.Abs(complex128(*g.val.(*complex128))))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			return 0
		}
		return float64(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return float64(gi)
	}
	return 0.0
}

func (g *AbdsItem) vc64() complex64 {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return complex(1, 0)
		}
		break
	case *int:
		return complex64(complex(float64(*g.val.(*int)), 0))
	case *int8:
		return complex64(complex(float64(*g.val.(*int8)), 0))
	case *int16:
		return complex64(complex(float64(*g.val.(*int16)), 0))
	case *int32:
		return complex64(complex(float64(*g.val.(*int32)), 0))
	case *int64:
		return complex64(complex(float64(*g.val.(*int64)), 0))
	case *uint:
		return complex64(complex(float64(*g.val.(*uint)), 0))
	case *uint8:
		return complex64(complex(float64(*g.val.(*uint8)), 0))
	case *uint16:
		return complex64(complex(float64(*g.val.(*uint16)), 0))
	case *uint32:
		return complex64(complex(float64(*g.val.(*uint32)), 0))
	case *uint64:
		return complex64(complex(float64(*g.val.(*uint64)), 0))
	case *float32:
		return complex64(complex(float64(*g.val.(*float32)), 0))
	case *float64:
		return complex64(complex(float64(*g.val.(*float64)), 0))
	case *complex64:
		return complex64(*g.val.(*complex64))
	case *complex128:
		return complex64(*g.val.(*complex128))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return complex64(complex(float64(gi), 0))
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return complex64(complex(float64(gi), 0))
	}
	return complex(0, 0)
}

func (g *AbdsItem) vc128() complex128 {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return complex(1, 0)
		}
		break
	case *int:
		return complex128(complex(float64(*g.val.(*int)), 0))
	case *int8:
		return complex128(complex(float64(*g.val.(*int8)), 0))
	case *int16:
		return complex128(complex(float64(*g.val.(*int16)), 0))
	case *int32:
		return complex128(complex(float64(*g.val.(*int32)), 0))
	case *int64:
		return complex128(complex(float64(*g.val.(*int64)), 0))
	case *uint:
		return complex128(complex(float64(*g.val.(*uint)), 0))
	case *uint8:
		return complex128(complex(float64(*g.val.(*uint8)), 0))
	case *uint16:
		return complex128(complex(float64(*g.val.(*uint16)), 0))
	case *uint32:
		return complex128(complex(float64(*g.val.(*uint32)), 0))
	case *uint64:
		return complex128(complex(float64(*g.val.(*uint64)), 0))
	case *float32:
		return complex128(complex(float64(*g.val.(*float32)), 0))
	case *float64:
		return complex128(complex(float64(*g.val.(*float64)), 0))
	case *complex64:
		return complex128(*g.val.(*complex64))
	case *complex128:
		return complex128(*g.val.(*complex128))
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.val.(*[]byte)), 10, 64)
		if err != nil {
			break
		}
		return complex128(complex(float64(gi), 0))
	case *string:
		gi, err := strconv.ParseInt(string(*g.val.(*string)), 10, 64)
		if err != nil {
			break
		}
		return complex128(complex(float64(gi), 0))
	}
	return complex(0, 0)
}

func (g *AbdsItem) vs() string {

	switch g.val.(type) {
	case nil:
		return "NULL"
	case *bool:
		if g.val == true {
			return "TRUE"
		}
		return "FALSE"
	case *int:
		return strconv.FormatInt(int64(*g.val.(*int)), 10)
	case *int8:
		return strconv.FormatInt(int64(*g.val.(*int8)), 10)
	case *int16:
		return strconv.FormatInt(int64(*g.val.(*int16)), 10)
	case *int32:
		return strconv.FormatInt(int64(*g.val.(*int32)), 10)
	case *int64:
		return strconv.FormatInt(int64(*g.val.(*int64)), 10)
	case *uint:
		return strconv.FormatUint(uint64(*g.val.(*uint)), 10)
	case *uint8:
		return strconv.FormatUint(uint64(*g.val.(*uint8)), 10)
	case *uint16:
		return strconv.FormatUint(uint64(*g.val.(*uint16)), 10)
	case *uint32:
		return strconv.FormatUint(uint64(*g.val.(*uint32)), 10)
	case *uint64:
		return strconv.FormatUint(uint64(*g.val.(*uint64)), 10)
	case *float32:
		return strconv.FormatFloat(float64(*g.val.(*float32)), 'f', -1, 32)
	case *float64:
		return strconv.FormatFloat(float64(*g.val.(*float64)), 'f', -1, 64)
	case *complex64:
		return strconv.FormatComplex(complex128(*g.val.(*complex64)), 'f', -1, 64)
	case *complex128:
		return strconv.FormatComplex(complex128(*g.val.(*complex128)), 'f', -1, 128)
	case *[]byte:
		return string(*g.val.(*[]byte))
	case *string:
		return string(*g.val.(*string))
	}
	return ""
}

func (g *AbdsItem) vbyte() []byte {

	switch g.val.(type) {
	case *bool:
		if g.val == true {
			return []byte{1}
		}
		return []byte{0}
	case *int:
		return []byte(strconv.FormatInt(int64(*g.val.(*int)), 10))
	case *int8:
		return []byte(strconv.FormatInt(int64(*g.val.(*int8)), 10))
	case *int16:
		return []byte(strconv.FormatInt(int64(*g.val.(*int16)), 10))
	case *int32:
		return []byte(strconv.FormatInt(int64(*g.val.(*int32)), 10))
	case *int64:
		return []byte(strconv.FormatInt(int64(*g.val.(*int64)), 10))
	case *uint:
		return []byte(strconv.FormatUint(uint64(*g.val.(*uint)), 10))
	case *uint8:
		return []byte(strconv.FormatUint(uint64(*g.val.(*uint8)), 10))
	case *uint16:
		return []byte(strconv.FormatUint(uint64(*g.val.(*uint16)), 10))
	case *uint32:
		return []byte(strconv.FormatUint(uint64(*g.val.(*uint32)), 10))
	case *uint64:
		return []byte(strconv.FormatUint(uint64(*g.val.(*uint64)), 10))
	case *float32:
		return []byte(strconv.FormatFloat(float64(*g.val.(*float32)), 'f', -1, 32))
	case *float64:
		return []byte(strconv.FormatFloat(float64(*g.val.(*float64)), 'f', -1, 64))
	case *complex64:
		return []byte(strconv.FormatComplex(complex128(*g.val.(*complex64)), 'f', -1, 64))
	case *complex128:
		return []byte(strconv.FormatComplex(complex128(*g.val.(*complex128)), 'f', -1, 128))
	case *[]byte:
		return *g.val.(*[]byte)
	case *string:
		return []byte(*g.val.(*string))
	}
	return []byte{}
}
