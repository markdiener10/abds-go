package abds

import "fmt"

type AbdsErr struct {
	count  uint
	errors []error
}

func (g *AbdsErr) Log(err error) {
	g.count++
	g.errors = append(g.errors, err)
}

func (g *AbdsErr) Is() bool {
	if g.count == 0 {
		return false
	}
	return true
}

func (g *AbdsErr) IsErr() bool {
	return g.Check()
}

func (g *AbdsErr) Check() bool {
	if g.count == 0 {
		return false
	}
	return true
}

func (g *AbdsErr) Content() []error {
	gerrs := make([]error, 0)
	gerrs = append(gerrs, g.errors...)
	return gerrs
}

func (g *AbdsErr) Reset() {
	g.count = 0
	g.errors = make([]error, 0)
}

func NewErrs() *AbdsErr {
	return &AbdsErr{count: 0, errors: make([]error, 0)}
}

//ABDS error related functions

func (g *Abds) IsErr() bool {
	if g.link == nil {
		return false
	}
	if g.root {
		gerrs := g.link.(*AbdsErr)
		return gerrs.IsErr()
	}
	glink := g.link.(*Abds)
	return glink.IsErr()
}

func (g *Abds) Errs(parms ...bool) *AbdsErr {
	if g.link == nil {
		return nil
	}
	if g.root {
		if parmbool(parms...) {
			gret := &AbdsErr{count: 0, errors: make([]error, 0)}
			gerrs := g.link.(*AbdsErr).Content()
			for _, gerr := range gerrs {
				gret.Log(gerr)
			}
			g.link.(*AbdsErr).Reset()
			return gret
		}
		return g.link.(*AbdsErr)
	}
	return g.link.(*Abds).Errs()
}

func (g *Abds) Log(err error) {
	var pErrs *AbdsErr
	if g == nil {
		return
	}
	if g.link == nil {
		return
	}
	if g.root {
		pErrs = g.link.(*AbdsErr)
	} else {
		pErrs = g.link.(*Abds).Errs()
	}
	if pErrs == nil {
		return
	}
	pErrs.Log(err)
}

func (g *Abds) ErrorClear() {
	//Cannot remove if child
	if !g.root {
		return
	}
	gerrs := g.Errs()
	gerrs.Reset()
}

func (g *Abds) adderr(msg string, parms ...interface{}) {
	gerrs := g.Errs()
	gerrs.Log(fmt.Errorf(msg, parms...))
	return
}
