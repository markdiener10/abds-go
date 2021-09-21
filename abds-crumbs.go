package abds

import (
	"errors"
)

//This is experimental for now
//These are optimized binary formats for transmission or storage

//We allow it to use a pointer since that would be the fastest
//way to fill a buffer is to let the library owner provide the buffer

//A binary map allows for accellerated parsing because
//we can build a binary breadcrumb trail
type Abdscrumb struct {
}

//Some way to represent a zero overhead parsing of the binary buffer
//That way we can just jam through and load each level
type Abdscrumbmap struct {
	Crumbs []Abdscrumb
}

func (g *Abds) GenCrumbmap(buf *[]byte) error {
	return errors.New("TDD Failure")
}

func (g *Abds) UseCrumbmap(*[]byte, *Abdscrumbmap) error {
	return errors.New("TDD Failure")
}
