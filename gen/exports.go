package gen

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"

	"github.com/chrsan/wagu/ir"
)

// Exports generates source for the exports file.
func Exports(pkg string, m *ir.Module, camelCase, exported bool) ([]byte, error) {
	asset, err := lookupTemplate("exports")
	if err != nil {
		return nil, err
	}
	fm := template.FuncMap{
		"name": func(s string) string {
			if camelCase {
				s = cc(s, exported)
			} else if exported {
				s = strings.Title(s)
			}
			return s
		},
		"params": func(id uint32, emitTyp bool) string {
			var sb strings.Builder
			for i, p := range m.Functions[int(id)-len(m.ImportedFunctions)].Sig.ParamTypes {
				if i != 0 {
					sb.WriteString(", ")
				}
				sb.WriteString(fmt.Sprintf("p%d", i))
				if emitTyp {
					sb.WriteByte(' ')
					sb.WriteString(typ(p))
				}
			}
			return sb.String()
		},
		"returnType": func(id uint32) string {
			rt := m.Functions[int(id)-len(m.ImportedFunctions)].Sig.ReturnTypes
			if len(rt) == 0 {
				return ""
			}
			return typ(rt[0])
		},
	}
	tmpl, err := template.New("exports").Funcs(fm).Parse(string(asset))
	if err != nil {
		return nil, err
	}
	data := struct {
		Pkg string
		E   []*ir.ExportedFunction
	}{
		Pkg: pkg,
		E:   m.ExportedFunctions,
	}
	var b bytes.Buffer
	if err := tmpl.Execute(&b, data); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

var ccRE = regexp.MustCompile("(^[A-Za-z0-9])|_+([A-Za-z0-9])")

func cc(name string, titleCase bool) string {
	return ccRE.ReplaceAllStringFunc(name, func(s string) string {
		s = strings.ReplaceAll(s, "_", "")
		if !titleCase {
			return s
		}
		return strings.ToTitle(s)
	})
}
