package ir

// LocalVar represents a WebAssembly local.
type LocalVar struct {
	ID  int
	Typ ValueType
}

// LocalVars returns the locals for f.
func (f *Function) LocalVars() []LocalVar {
	if len(f.Locals) == 0 {
		return nil
	}
	var locals []LocalVar
	p := len(f.Sig.ParamTypes)
	for i, l := range f.Locals {
		locals = append(locals, LocalVar{ID: p + i, Typ: l})
	}
	return locals
}

// HasReturnType returns whether f returns a value or not.
func (s *FunctionSig) HasReturnType() bool {
	return len(s.ReturnTypes) != 0
}

// ReturnType returns the return value type for s.
func (s *FunctionSig) ReturnType() ValueType {
	return s.ReturnTypes[0]
}
