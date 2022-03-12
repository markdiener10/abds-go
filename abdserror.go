package abds

type AbdsErr struct {
	count  uint
	errors []error
}

func (g *AbdsErr) Log(err error) {
	g.count++
	g.errors = append(g.errors, err)
}

func (g *AbdsErr) Check() bool {
	if g.count == 0 {
		return false
	}
	return true
}

func (g *AbdsErr) Content() []error {
	return g.errors
}

func (g *AbdsErr) Reset() {
	g.count = 0
	g.errors = make([]error, 0)
}

func NewErrs() *AbdsErr {
	return &AbdsErr{count: 0, errors: make([]error, 0)}
}
