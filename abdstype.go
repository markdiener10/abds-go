package abds

import (
	"reflect"
)

//Used to allow future structures to map a stream of bytes
type AbdsTransform interface {
	Unpack() (string, []byte)
}

//Have a global variable since this is the only struct we want in our nodes
var gAbdsTranType reflect.Type

func init() {
	gAbdsTranType = reflect.TypeOf((AbdsTransform)(nil))
}
