package abds

import (
	_ "errors"
)

const (
	NIL = iota
	BOOL
	INT
	INT8
	INT16
	INT32
	INT64
	UINT
	UINT8
	UINT16
	UINT32
	UINT64
	FLOAT32
	FLOAT64
	COMPLEX64
	COMPLEX128
	STRING
	STRINGPTR
	BYTEARR
	BYTEPTRd
	RUNEARR
	ABDS
	ABDSTRAN
)

type AbdsStream struct {
	len  uint64
	data []byte
}

func (g *Abds) GenStream(buf *[]byte, parms ...interface{}) {
	return
}

func (g *Abds) LoadStream(buf *[]byte, parms ...interface{}) error {
	return nil
}
