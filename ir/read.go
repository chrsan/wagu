package ir

import (
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/go-interpreter/wagon/wasm"
)

// ImportedGlobals maps a module/field name pair to a global value.
type ImportedGlobals map[string]map[string]*Value

// Read reads the WebAssembly in r and returns a new IR Module.
func Read(r io.Reader, ig ImportedGlobals) (*Module, error) {
	m, err := wasm.ReadModule(r, nil)
	if err != nil {
		return nil, err
	}
	return ReadModule(m, ig)
}

// ReadModule reads the WebAssembly module wm and returns a new IR Module.
func ReadModule(wm *wasm.Module, ig ImportedGlobals) (*Module, error) {
	ie := groupImports(wm)
	m := new(Module)
	if len(ie) != 0 {
		for _, e := range ie[wasm.ExternalFunction] {
			f := e.Type.(wasm.FuncImport)
			m.ImportedFunctions = append(m.ImportedFunctions, &ImportedFunction{
				ModuleName: e.ModuleName,
				FieldName:  e.FieldName,
				Sig:        funcSig(&wm.Types.Entries[f.Type]),
				TypeId:     f.Type,
			})
		}
	}
	if err := readGlobals(m, wm, ie, ig); err != nil {
		return nil, err
	}
	if err := readMemory(m, wm, ie); err != nil {
		return nil, err
	}
	if err := readTable(m, wm, ie); err != nil {
		return nil, err
	}
	if wm.Export != nil && wm.Export.Entries != nil {
		for name, e := range wm.Export.Entries {
			if e.Kind != wasm.ExternalFunction {
				continue
			}
			m.ExportedFunctions = append(m.ExportedFunctions, &ExportedFunction{
				Id:   e.Index,
				Name: name,
			})
		}
	}
	for i, f := range wm.FunctionIndexSpace {
		fn, err := readFunc(m, wm, f, i)
		if err != nil {
			return nil, err
		}
		m.Functions = append(m.Functions, fn)
	}
	return m, nil
}

func groupImports(m *wasm.Module) [][]wasm.ImportEntry {
	if m.Import == nil {
		return nil
	}
	ie := make([][]wasm.ImportEntry, 4)
	for _, e := range m.Import.Entries {
		switch e.Type.Kind() {
		case wasm.ExternalGlobal:
			ie[wasm.ExternalGlobal] = append(ie[wasm.ExternalGlobal], e)
		case wasm.ExternalFunction:
			ie[wasm.ExternalFunction] = append(ie[wasm.ExternalFunction], e)
		case wasm.ExternalMemory:
			ie[wasm.ExternalMemory] = append(ie[wasm.ExternalMemory], e)
		case wasm.ExternalTable:
			ie[wasm.ExternalTable] = append(ie[wasm.ExternalTable], e)
		}
	}
	return ie
}

func readGlobals(m *Module, wm *wasm.Module, ie [][]wasm.ImportEntry, ig ImportedGlobals) error {
	if len(ie) != 0 {
		for _, e := range ie[wasm.ExternalGlobal] {
			if e.ModuleName == "global" {
				switch e.FieldName {
				case "NaN":
					m.Globals = append(m.Globals, &Value{Kind: &Value_F64{F64: math.NaN()}})
					continue
				case "Infinity":
					m.Globals = append(m.Globals, &Value{Kind: &Value_F64{F64: math.Inf(0)}})
					continue
				}
			}
			g := e.Type.(wasm.GlobalVarImport).Type
			mm, ok := ig[e.ModuleName]
			if !ok {
				return fmt.Errorf("missing global module %s", e.ModuleName)
			}
			v, ok := mm[e.FieldName]
			if !ok {
				return fmt.Errorf("missing global field %s in module %s", e.FieldName, e.ModuleName)
			}
			vt := wasmValueType(v)
			if vt != g.Type {
				return fmt.Errorf("got type %v for global %s/%s, want %v", vt, e.ModuleName, e.FieldName, g.Type)
			}
			m.Globals = append(m.Globals, v)
		}
	}
	for i, e := range wm.GlobalIndexSpace {
		v, err := readInitExpr(wm, e.Init, m.Globals)
		if err != nil {
			return err
		}
		vt := wasmValueType(v)
		if vt != e.Type.Type {
			return fmt.Errorf("type mismatch for global %d: %v != %v", i, vt, e.Type.Type)
		}
		m.Globals = append(m.Globals, v)
	}
	return nil
}

func readMemory(m *Module, wm *wasm.Module, ie [][]wasm.ImportEntry) error {
	var mem []wasm.Memory
	if len(ie) != 0 {
		for _, e := range ie[wasm.ExternalMemory] {
			em := e.Type.(wasm.MemoryImport).Type
			mem = append(mem, em)
		}
	}
	if wm.Memory != nil {
		for _, e := range wm.Memory.Entries {
			mem = append(mem, e)
		}
	}
	switch len(mem) {
	case 0:
		return nil
	case 1:
		size := mem[0].Limits.Initial
		if size == 0 {
			return nil
		}
		m.Memory = &Memory{Size: size * 65536}
		if wm.Data == nil || len(wm.Data.Entries) == 0 {
			return nil
		}
		m.Memory.Segments = make(map[uint32][]byte, len(wm.Data.Entries))
		for _, e := range wm.Data.Entries {
			off, err := readInitExpr(wm, e.Offset, m.Globals)
			if err != nil {
				return err
			}
			vt := wasmValueType(off)
			if vt != wasm.ValueTypeI32 {
				return fmt.Errorf("invalid memory offset type: %v", vt)
			}
			o := off.Kind.(*Value_I32)
			m.Memory.Segments[uint32(o.I32)] = e.Data
		}
		return nil
	}
	return errors.New("multiple memories not supported")
}

func readTable(m *Module, wm *wasm.Module, ie [][]wasm.ImportEntry) error {
	var ts []wasm.Table
	if len(ie) != 0 {
		for _, e := range ie[wasm.ExternalTable] {
			t := e.Type.(wasm.TableImport).Type
			if t.ElementType != wasm.ElemTypeAnyFunc {
				return fmt.Errorf("unsupported table element type: %v", t.ElementType)
			}
			ts = append(ts, t)
		}
	}
	if wm.Table != nil {
		for _, e := range wm.Table.Entries {
			if e.ElementType != wasm.ElemTypeAnyFunc {
				return fmt.Errorf("unsupported table element type: %v", e.ElementType)
			}
			ts = append(ts, e)
		}
	}
	switch len(ts) {
	case 0:
		return nil
	case 1:
		size := int(ts[0].Limits.Initial)
		if size == 0 {
			return nil
		}
		table := make([]uint32, size)
		for i := 0; i < size; i++ {
			table[i] = math.MaxUint32
		}
		if wm.Elements == nil {
			return nil
		}
		for _, e := range wm.Elements.Entries {
			off, err := readInitExpr(wm, e.Offset, m.Globals)
			if err != nil {
				return err
			}
			vt := wasmValueType(off)
			if vt != wasm.ValueTypeI32 {
				return fmt.Errorf("invalid table offset type: %v", vt)
			}
			o := off.Kind.(*Value_I32)
			copy(table[int(o.I32):], e.Elems)
		}
		for _, e := range table {
			if e == math.MaxUint32 {
				m.Table = append(m.Table, &TableEntry{FunctionId: -1})
				continue
			}
			var n int
			if int(e) < len(m.ImportedFunctions) {
				n = len(m.ImportedFunctions[int(e)].Sig.ParamTypes)
			} else {
				n = len(wm.FunctionIndexSpace[int(e)-len(m.ImportedFunctions)].Sig.ParamTypes)
			}
			m.Table = append(m.Table, &TableEntry{FunctionId: int64(e), NumParams: uint32(n)})
		}
		return nil
	}
	return errors.New("multiple tables not supported")
}

func funcSig(sig *wasm.FunctionSig) *FunctionSig {
	s := new(FunctionSig)
	for _, p := range sig.ParamTypes {
		s.ParamTypes = append(s.ParamTypes, valueType(p))
	}
	for _, r := range sig.ReturnTypes {
		s.ReturnTypes = append(s.ReturnTypes, valueType(r))
	}
	return s
}

func wasmValueType(v *Value) wasm.ValueType {
	switch v.Kind.(type) {
	case *Value_I32:
		return wasm.ValueTypeI32
	case *Value_I64:
		return wasm.ValueTypeI64
	case *Value_F32:
		return wasm.ValueTypeF32
	case *Value_F64:
		return wasm.ValueTypeF64
	}
	panic(v.Kind)
}

func valueType(v wasm.ValueType) ValueType {
	switch v {
	case wasm.ValueTypeI32:
		return ValueType_I32
	case wasm.ValueTypeI64:
		return ValueType_I64
	case wasm.ValueTypeF32:
		return ValueType_F32
	case wasm.ValueTypeF64:
		return ValueType_F64
	}
	panic(v)
}
