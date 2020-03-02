package gen

import (
	"bytes"
	"text/template"

	"github.com/chrsan/wagu/ir"
)

// Context generates source for the context file.
func Context(pkg string, m *ir.Module, mmap bool, useUnsafe bool) ([]byte, error) {
	asset, err := lookupTemplate("context")
	if err != nil {
		return nil, err
	}
	fm := template.FuncMap{
		"type":  typ,
		"value": val,
		"sig": func(tableEntryFuncID int64) *ir.FunctionSig {
			idx := int(tableEntryFuncID)
			if idx < len(m.ImportedFunctions) {
				return m.ImportedFunctions[idx].Sig
			}
			return m.Functions[idx-len(m.ImportedFunctions)].Sig
		},
	}
	tmpl, err := template.New("context").Funcs(fm).Parse(string(asset))
	if err != nil {
		return nil, err
	}
	data := struct {
		Pkg       string
		M         *ir.Module
		MMap      bool
		UseUnsafe bool
	}{
		Pkg:       pkg,
		M:         m,
		MMap:      mmap,
		UseUnsafe: useUnsafe,
	}
	var b bytes.Buffer
	if err := tmpl.Execute(&b, data); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
