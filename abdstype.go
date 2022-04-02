package abds

import (
	"reflect"
)

//Used to allow future structures to map into abds format
type AbdsTransform interface {
	Pack() *Abds
	UnPack(*Abds) error
}

//Have a global variable since this is the only struct we want in our nodes
var gAbdsTranType reflect.Type

func init() {
	gAbdsTranType = reflect.TypeOf((AbdsTransform)(nil))
}
