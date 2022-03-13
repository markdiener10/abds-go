package abds

import (
	"fmt"
)

func (g *AbdsItem) S(val interface{}, parms ...interface{}) {

	if !checkType(val) {
		parmerr(fmt.Errorf("Tag:%s Invalid Type passed", g.tag), parms...)
		return
	}

	switch val.(type) {
	case bool:
		gp := new(bool)
		*gp = val.(bool)
		val = gp
	case int:
		gp := new(int)
		*gp = val.(int)
		val = gp
	case int8:
		gp := new(int8)
		*gp = val.(int8)
		val = gp
	case int16:
		gp := new(int16)
		*gp = val.(int16)
		val = gp
	case int32: //*rune
		gp := new(int32)
		*gp = val.(int32)
		val = gp
	case int64:
		gp := new(int64)
		*gp = val.(int64)
		val = gp
	case uint:
		gp := new(uint)
		*gp = val.(uint)
		val = gp
	case uint8: //*byte
		gp := new(uint8)
		*gp = val.(uint8)
		val = gp
	case uint16:
		gp := new(uint16)
		*gp = val.(uint16)
		val = gp
	case uint32:
		gp := new(uint32)
		*gp = val.(uint32)
		val = gp
	case uint64:
		gp := new(uint64)
		*gp = val.(uint64)
		val = gp
	case float32:
		gp := new(float32)
		*gp = val.(float32)
		val = gp
	case float64:
		gp := new(float64)
		*gp = val.(float64)
		val = gp
	case complex64:
		gp := new(complex64)
		*gp = val.(complex64)
		val = gp
	case complex128:
		gp := new(complex128)
		*gp = val.(complex128)
		val = gp
	case []byte:
		gb := val.([]byte)
		gp := new([]byte)
		*gp = append(*gp, gb...)
		val = gp
	case []rune:
		gb := val.([]rune)
		gp := new([]rune)
		*gp = append(*gp, gb...)
		val = gp
	case string:
		gp := new(string)
		*gp = val.(string)
		val = gp
	case Abds:
		gp := new(Abds)
		gs := val.(Abds)
		(*gp).Copy(&gs) //Could be an expensive copy but allow it
		val = gp
	}
	g.val = val
}

func (g *AbdsItem) SetAbds(val *Abds, parms ...interface{}) {
	g.S(val, parms...)
}


//Provide a member function that will enforce at compile time the AbdsTransform interface
func (g *AbdsItem) SetStru(val *AbdsTransform, parms ...interface{}) {
	g.S(val, parms...)
}
