package abds

import (
	"reflect"
)

type AbdsItem struct {
	tag uint
	val interface{} //Range of values
}

func (g *AbdsItem) Clear() {
	//Remove the link between the value and the abds structure (pointers are garbage collected)
	//Do not touch the TAG value
	g.val = nil
}

func (g *AbdsItem) Ti() uint {
	return g.tag
}

func (g *AbdsItem) V() interface{} {
	return g.val
}

func (g *AbdsItem) cv(val interface{}) interface{} {

	if val == nil {
		return nil
	}

	var gPointer bool = false
	baseType := reflect.TypeOf(val)
	baseVal := reflect.ValueOf(val)
	if baseType.Kind() == reflect.Pointer {
		baseVal = reflect.Indirect(baseVal)
		baseType = baseVal.Type()
		gPointer = true
	}

	switch baseType.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		if !gPointer {
			gp := reflect.New(baseType)
			gp.Elem().Set(baseVal)
			val = gp.Interface()
		}
		break
	default:
		return nil
	}
	return val
}
