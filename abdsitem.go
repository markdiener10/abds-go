package abds

import (
	"fmt"
	"math/cmplx"
	"reflect"
	"strconv"
)

//Used to allow future structures to map into abds format
type AbdsTransform interface {
	Pack() *Abds
	UnPack(*Abds) error
}

//This is for verifying with the interface AbdsTransform interface
var gAbdsTranType reflect.Type

func init() {
	gAbdsTranType = reflect.TypeOf((*AbdsTransform)(nil)).Elem()
}

type AbdsItem struct {
	tag interface{} //uint64 or String
	Val interface{} //Range of values
}

func (g *AbdsItem) Clear() {
	//Remove the link between the value and the abds structure (pointers are garbage collected)
	//Do not touch the TAG value (like ABDSERR)
	g.Val = nil
}

func (g *AbdsItem) Ti() uint {
	if g.tag == nil {
		return 0
	}
	gt, ok := g.tag.(uint)
	if !ok {
		return 0
	}
	return gt
}

func (g *AbdsItem) Ts() string {
	if g.tag == nil {
		return ""
	}
	gt, ok := g.tag.(string)
	if !ok {
		return ""
	}
	return gt
}

func (g *AbdsItem) G() interface{} {
	return g.Val
}

func (g *AbdsItem) P(val interface{}, parms ...interface{}) {

	if !checkType(val) {
		errset(fmt.Errorf("Tag:%s Invalid Type passed", g.tag), parms...)
		return
	}

	//Map pointers to locals and locals to pointers depending on type
	switch val.(type) {
	case *bool:
		gb, _ := val.(*bool)
		val = bool(*gb)
	case *int:
		gb, _ := val.(*int)
		val = int(*gb)
	case *int8:
		gb, _ := val.(*int8)
		val = int8(*gb)
	case *int16:
		gb, _ := val.(*int16)
		val = int16(*gb)
	case *int32: //*rune
		gb, _ := val.(*int32)
		val = int32(*gb)
	case *int64:
		gb, _ := val.(*int64)
		val = int64(*gb)
	case *uint:
		gb, _ := val.(*uint)
		val = uint(*gb)
	case *uint8: //*byte
		gb, _ := val.(*uint8)
		val = uint8(*gb)
	case *uint16:
		gb, _ := val.(*uint16)
		val = uint16(*gb)
	case *uint32:
		gb, _ := val.(*uint32)
		val = uint32(*gb)
	case *uint64:
		gb, _ := val.(*uint64)
		val = uint64(*gb)
	case *float32:
		gb, _ := val.(*float32)
		val = float32(*gb)
	case *float64:
		gb, _ := val.(*float64)
		val = float64(*gb)
	case *complex64:
		gb, _ := val.(*complex64)
		val = complex64(*gb)
	case *complex128:
		gb, _ := val.(*complex128)
		val = complex128(*gb)
	}
	g.Val = val
}

func checkType(val interface{}) bool {

	switch val.(type) {
	case *bool:
	case *int:
	case *int8:
	case *int16:
	case *int32: //*rune
	case *int64:
	case *uint:
	case *uint8: //*byte
	case *uint16:
	case *uint32:
	case *uint64:
	case *float32:
	case *float64:
	case *complex64:
	case *complex128:
	case nil:
	case bool:
	case int:
	case int8:
	case int16:
	case int32: //rune
	case int64:
	case uint:
	case uint8: //byte
	case uint16:
	case uint32:
	case uint64:
	case float32:
	case float64:
	case complex64:
	case complex128:
	case uintptr:
	case string:
	case *string:
	case []byte:
	case *[]byte:
	case []rune:
	case *[]rune:
	case *Abds: //We also only want to have Abds pointers, not stack allocations
		break
	default:

		//Only allow pointers
		if reflect.TypeOf(val).Kind() != reflect.Ptr {
			return false
		}

		//Only allow structure pointers that implement interface AbdsTransform
		if !reflect.TypeOf(val).Implements(gAbdsTranType) {
			return false
		}
	}
	return true
}

func (g *AbdsItem) Rs() string {

	switch g.Val.(type) {
	case nil:
		return "NULL"
	case bool:
		if g.Val == true {
			return "FALSE"
		} else {
			return "FALSE"
		}
	case int:
		return strconv.FormatInt(int64(g.Val.(int)), 10)
	case int8:
		return strconv.FormatInt(int64(g.Val.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(g.Val.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(g.Val.(int32)), 10)
	case int64:
		return strconv.FormatInt(int64(g.Val.(int64)), 10)
	case uint:
		return strconv.FormatInt(int64(g.Val.(uint)), 10)
	case uint8:
		return strconv.FormatInt(int64(g.Val.(uint)), 10)
	case uint16:
		return strconv.FormatInt(int64(g.Val.(uint)), 10)
	case uint32:
		return strconv.FormatInt(int64(g.Val.(uint)), 10)
	case uint64:
		return strconv.FormatInt(int64(g.Val.(uint)), 10)
	case float32:
		return strconv.FormatFloat(float64(g.Val.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(float64(g.Val.(float64)), 'f', -1, 64)
	case complex64:
		return strconv.FormatComplex(complex128(g.Val.(complex64)), 'f', -1, 64)
	case complex128:
		return strconv.FormatComplex(complex128(g.Val.(complex128)), 'f', -1, 128)
	case uintptr:
		return fmt.Sprintf("%d", g.Val.(uintptr))
	case []byte:
		return string(g.Val.([]byte))
	case string:
		return g.Val.(string)
	case *[]byte:
		return string(*g.Val.(*[]byte))
	case *string:
		return string(*g.Val.(*string))
	//case []rune, *[]rune:
	//return ""
	//WE may be able to return JSON format for the Abds and AbdsTransform stuff
	//case Abds, AbdsTransform, *Abds, *AbdsTransform:
	//	return ""
	default:
		return ""
	}
}

func (g *AbdsItem) Ri() int {

	switch g.Val.(type) {
	case nil:
		return 0
	case bool:
		if g.Val == true {
			return 1
		} else {
			return 0
		}
	case int:
		return g.Val.(int)
	case int8:
		return int(g.Val.(int8))
	case int16:
		return int(g.Val.(int16))
	case int32:
		return int(g.Val.(int32))
	case int64:
		return int(g.Val.(int64))
	case uint:
		return int(g.Val.(uint))
	case uint8:
		return int(g.Val.(uint8))
	case uint16:
		return int(g.Val.(uint16))
	case uint32:
		return int(g.Val.(uint32))
	case uint64:
		return int(g.Val.(uint64))
	case float32:
		return int(g.Val.(float32))
	case float64:
		return int(g.Val.(float64))
	case complex64:
		return int(cmplx.Abs(complex128(g.Val.(complex64))))
	case complex128:
		return int(cmplx.Abs(complex128(g.Val.(complex128))))
	case uintptr:
		return int(g.Val.(uintptr))
	case []byte:
		gi, err := strconv.ParseInt(string(g.Val.([]byte)), 10, 64)
		if err != nil {
			return 0
		}
		return int(gi)
	case *[]byte:
		gi, err := strconv.ParseInt(string(*g.Val.(*[]byte)), 10, 64)
		if err != nil {
			return 0
		}
		return int(gi)
	case string:
		gi, err := strconv.ParseInt(string(g.Val.(string)), 10, 64)
		if err != nil {
			return 0
		}
		return int(gi)
	case *string:
		gi, err := strconv.ParseInt(string(*g.Val.(*string)), 10, 64)
		if err != nil {
			return 0
		}
		return int(gi)
	//case []rune, *[]rune:
	//	return 0
	//case Abds, AbdsTransform, *Abds, *AbdsTransform:
	//	return 0
	default:
		return 0
	}
}
