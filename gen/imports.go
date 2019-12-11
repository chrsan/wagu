package gen

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/chrsan/wagu/ir"
)

// ImportsInterface generates source for the ImportedFuncs interface file.
func ImportsInterface(pkg string, m *ir.Module) ([]byte, error) {
	asset, err := lookupTemplate("imports")
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New("imports").Parse(string(asset))
	if err != nil {
		return nil, err
	}
	data := struct {
		Pkg           string
		ImportedFuncs []importedFunc
	}{
		Pkg: pkg,
	}
	for _, f := range m.ImportedFunctions {
		data.ImportedFuncs = append(data.ImportedFuncs, impFunc(f, true))
	}
	var b bytes.Buffer
	if err := tmpl.Execute(&b, data); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// ImportedFunc generates source for an imported function.
func ImportedFunc(pkg string, id int, f *ir.ImportedFunction) ([]byte, error) {
	asset, err := lookupTemplate("imp_func")
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New("imp_func").Parse(string(asset))
	if err != nil {
		return nil, err
	}
	data := struct {
		Pkg          string
		ID           int
		ImportedFunc importedFunc
	}{
		Pkg:          pkg,
		ID:           id,
		ImportedFunc: impFunc(f, false),
	}
	var b bytes.Buffer
	if err := tmpl.Execute(&b, data); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type importedFunc struct {
	Module, Field string
	ReturnType    string
	ParamTypes    string
	Params        string
}

func impFunc(f *ir.ImportedFunction, decl bool) importedFunc {
	var rt string
	if len(f.Sig.ReturnTypes) != 0 {
		rt = typ(f.Sig.ReturnTypes[0])
	}
	var sb1, sb2 strings.Builder
	for i, p := range f.Sig.ParamTypes {
		sb1.WriteString(", ")
		sb2.WriteString(", ")
		n := "p"
		if !decl {
			n = "l"
		}
		ps := fmt.Sprintf("%s%d", n, i)
		sb1.WriteString(ps)
		sb2.WriteString(ps)
		sb1.WriteByte(' ')
		sb1.WriteString(typ(p))
	}
	return importedFunc{
		Module:     f.ModuleName,
		Field:      f.FieldName,
		ReturnType: rt,
		ParamTypes: sb1.String(),
		Params:     sb2.String(),
	}
}
