package abds

import (
	"fmt"
	"reflect"
)

func (g *Abds) dummy() *AbdsItem {
	return &AbdsItem{}
}

func (g *Abds) find(ntag uint, recurse bool) *AbdsItem {

	var index int
	var pItem *AbdsItem = nil

	for index = 0; index < len(g.vals); index++ {
		pItem = g.vals[index]
		if pItem == nil {
			continue
		}
		if pItem.tag != ntag {
			continue
		}
		return pItem
	}
	if !recurse {
		return nil
	}

	var child *Abds
	var ok bool

	for index = 0; index < len(g.vals); index++ {
		pItem = g.vals[index]
		child, ok = pItem.val.(*Abds)
		if !ok {
			continue
		}
		if child == nil {
			continue
		}
		pItem = child.find(ntag, true)
		if pItem == nil {
			continue
		}
		return pItem
	}
	return nil
}

func parmbool(parms ...bool) bool {
	for _, parm := range parms {
		if parm {
			return true
		}
	}
	return false
}

func parmu(flag uint, parms ...uint) bool {
	for _, parm := range parms {
		if parm != flag {
			continue
		}
		return true
	}
	return false
}

func parmflags(parms ...uint) (flags uint) {
	flags = 0
	for _, parm := range parms {
		flags |= parm
	}
	return
}

func parmfilter(flag uint, parms ...uint) (flags uint) {
	flags = 0
	for _, parm := range parms {
		if flag == parm {
			continue
		}
		flags |= parm
	}
	return
}

func parmerr(gerr error, parms ...*AbdsErr) {

	for _, parm := range parms {
		if parm == nil {
			continue
		}
		parm.Log(gerr)
		return
	}
}

func checkval(val interface{}) (interface{}, error) {

	if val == nil {
		return val, nil
	}

	var gPointer bool = false
	baseType := reflect.TypeOf(val)
	baseVal := reflect.ValueOf(val)
	if baseType.Kind() == reflect.Pointer {
		baseVal = reflect.Indirect(baseVal)
		baseType = baseVal.Type()
		gPointer = true
	}

	if baseType.Kind() == reflect.Interface {
		baseVal = baseVal.Elem()
		baseType = baseVal.Type()
	}

	switch baseType.Kind() {
	case reflect.Invalid, reflect.Chan, reflect.Func, reflect.Interface, reflect.Slice, reflect.UnsafePointer:
		return nil, fmt.Errorf("ABDSitem preset() invalid type:%s", baseType.String())
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:

		if !gPointer {
			gp := reflect.New(baseType)
			gp.Elem().Set(baseVal)
			val = gp.Interface()
		}
		break
	case reflect.Array:
		if gPointer == false {
			return nil, fmt.Errorf("ABDSitem S() only array pointers")
		}
	case reflect.Map:
		if gPointer == false {
			return nil, fmt.Errorf("ABDSitem S() only map pointers")
		}
	case reflect.Struct:
		if gPointer == false {
			return nil, fmt.Errorf("ABDSitem S() only struct pointers")
		}
	default:
		return nil, fmt.Errorf("ABDSitem S() unknown reflect type:%s", baseType.String())
	}
	return val, nil
}
