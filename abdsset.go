package abds

import (
	"fmt"
	"reflect"
	"strings"
)

//Allow for daisy-chaining value set by returned self (g *Abds)
func (g *Abds) S(parms ...interface{}) *Abds {

	if g == nil {
		return nil
	}

	var idx uint = 0
	var pItem *AbdsItem = nil
	var gch *Abds
	var ok bool

	switch len(parms) {
	case 0:
		return g
	case 1:
		parms[0] = g.checkval(parms[0])
		gch, ok = parms[0].(*Abds)
		if ok {
			gch.root = false
			if g.root {
				gch.link = g
			} else {
				gch.link = g.link
			}
		}
		pItem = &AbdsItem{tag: 0, val: parms[0]}
		g.vals = append(g.vals, pItem)
		return g
	}

	parms[1] = g.checkval(parms[1])

	if gch, ok = parms[0].(*Abds); ok {
		gch.root = false
		if g.root {
			gch.link = g
		} else {
			gch.link = g.link
		}
	}

	switch parms[0].(type) {
	case *AbdsItem:
		pItem = parms[0].(*AbdsItem)
		pItem.val = parms[1]
		return g
	case string:
		stag := strings.TrimSpace(parms[0].(string))
		if stag == "" {
			g.Log(fmt.Errorf("ABDSitem S() received empty tag parm"))
			return g
		}
		ntag := g.GetTagi(stag)
		if ntag > 0 {
			pItem := g.find(ntag, false)
			if pItem != nil {
				return g
			}
		}
	case int, int8, int16, int32, int64:
		idx = uint(parms[0].(int))
		break
	case uint, uint8, uint16, uint32, uint64:
		idx = parms[0].(uint)
		break
	default:
		g.Log(fmt.Errorf("ABDSitem S() invalid tag/idx/AbdsItem parameter"))
		return g
	}

	if idx == 0 || g.Len() < idx {
		pItem = &AbdsItem{tag: 0, val: parms[1]}
		g.vals = append(g.vals, pItem)
		return g
	}
	pItem = g.vals[idx]
	pItem.val = parms[1]
	return g
}

func (g *Abds) checkval(val interface{}) interface{} {

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

	if baseType.Kind() == reflect.Interface {
		baseVal = baseVal.Elem()
		baseType = baseVal.Type()
	}

	switch baseType.Kind() {
	case reflect.Invalid, reflect.Chan, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		g.Log(fmt.Errorf("ABDSitem preset() invalid type:%s", baseType.String()))
		return nil
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		if !gPointer {
			gp := reflect.New(baseType)
			gp.Elem().Set(baseVal)
			val = gp.Interface()
		}
		break
	case reflect.Slice:
	case reflect.Array:
	case reflect.Map, reflect.Struct:
	default:
		g.Log(fmt.Errorf("ABDSitem S() unknown reflect type:%s", baseType.String()))
		return nil
	}
	return val
}

//Array insert capability
func (g *Abds) Sinsert(idx uint, val interface{}) bool {

	if g == nil {
		return false
	}
	val = g.checkval(val)
	pItem := &AbdsItem{tag: 0, val: val}
	g.vals = append(g.vals, pItem)
	if g.Len() <= idx {
		return true
	}
	var pSrc *AbdsItem
	var pDst *AbdsItem

	//Shift the values up
	for slot := g.Len() - 1; slot > idx-1; slot-- {
		pSrc = g.vals[slot-1]
		pDst = g.vals[slot]
		pDst.val = pSrc.val
		pDst.tag = pSrc.tag
		pSrc.val = nil
		pSrc.tag = 0
	}
	pItem = g.vals[idx-1]
	pItem.tag = 0
	pItem.val = val
	return true
}
