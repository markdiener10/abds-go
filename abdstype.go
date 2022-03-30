package abds

import (
	_ "fmt"
	_ "math/cmplx"
	"reflect"
	_ "strconv"
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

func checkType(val interface{}, limit bool) bool {

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
	case *string:
	case *[]byte:
	case *[]rune:
	case Abds:
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
	case []byte:
	case []rune:
	case *Abds: //We also only want to have Abds pointers, not stack allocations
		if limit == true {
			return false
		}
		break
	default:

		//Only allow pointers
		if reflect.TypeOf(val).Kind() != reflect.Ptr {
			return false
		}

		//Only allow structure pointers that implement interface AbdsTransform
		//Strange requirement that we have a variable store the type instead of the type itself
		if !reflect.TypeOf(val).Implements(gAbdsTranType) {
			return false
		}
		if limit == true {
			return false
		}
	}
	return true
}
