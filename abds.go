package abds

import (
	"errors"
)

//This is our agile binary data structure to replace a good number
//of things including protobuffers, json processing internally
//and brittle structural definitions (that constantly need to be updated)

//Object packaging
type Abdso struct {
	Vals map[string]interface{}
}

//Array packaging
type Abdsa struct {
	Vals []interface{}
}

type Abds struct {
	Vals interface{}
}

func (g *Abds) P(tag string, val interface{}) (*interface{}, error) {
	return nil, errors.New("TDD Failure")
}

func (g *Abds) G(tag string) (interface{}, error) {
	return nil, errors.New("TDD Failure")
}

func (g *Abds) Pi(index uint64, val interface{}) (*interface{}, error) {
	return nil, errors.New("TDD Failure")
}

func (g *Abds) Gi(index uint64) (interface{}, error) {
	return nil, errors.New("TDD Failure")
}

func (g *Abds) IsArray() bool {
	return false
}

func (g *Abds) Len() uint64 {
	return 0
}
