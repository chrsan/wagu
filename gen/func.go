package gen

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/go-interpreter/wagon/wasm/operators"

	"github.com/chrsan/wagu/ir"
)

// Func generates source for a function.
func Func(pkg string, f *ir.Function, globals []*ir.Value, exported, exprComments bool) ([]byte, error) {
	asset, err := lookupTemplate("func")
	if err != nil {
		return nil, err
	}
	fm := template.FuncMap{
		"type":  typ,
		"value": val,
	}
	tmpl, err := template.New("func").Funcs(fm).Parse(string(asset))
	if err != nil {
		return nil, err
	}
	data := struct {
		Pkg          string
		F            *ir.Function
		StackVars    []sv
		Exprs        []expr
		Exported     bool
		ExprComments bool
	}{
		Pkg:          pkg,
		F:            f,
		StackVars:    stackVars(f),
		Exprs:        exprs(f, globals),
		Exported:     exported,
		ExprComments: exprComments,
	}
	var b bytes.Buffer
	if err := tmpl.Execute(&b, data); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

var svSuffixes = map[ir.ValueType]string{
	ir.ValueType_I32: "i32",
	ir.ValueType_I64: "i64",
	ir.ValueType_F32: "f32",
	ir.ValueType_F64: "f64",
}

type sv struct {
	ID     uint32
	Suffix string
	Type   ir.ValueType
}

func stackVars(f *ir.Function) []sv {
	var svs []sv
	it := func(ss []uint32, vt ir.ValueType) {
		for _, s := range ss {
			svs = append(svs, sv{ID: s, Suffix: svSuffixes[vt], Type: vt})
		}
	}
	it(f.StackVars.I32, ir.ValueType_I32)
	it(f.StackVars.I64, ir.ValueType_I64)
	it(f.StackVars.F32, ir.ValueType_F32)
	it(f.StackVars.F64, ir.ValueType_F64)
	return svs
}

func stackVar(v *ir.StackVar) string {
	return fmt.Sprintf("s%d%s", v.Id, svSuffixes[v.ValueType])
}

type expr struct {
	Name string
	Data map[string]interface{}
}

func exprs(f *ir.Function, globals []*ir.Value) []expr {
	var (
		blocks []*ir.Expr
		exprs  []expr
	)
	for _, x := range f.Exprs {
		e := expr{Data: map[string]interface{}{}}
		switch k := x.Kind.(type) {
		case *ir.Expr_Unreachable:
			e.Name = "unreachable"
		case *ir.Expr_Return:
			e.Name = "return"
			if len(k.Return.Returns) != 0 {
				e.Data["Var"] = stackVar(k.Return.Returns[0])
			}
		case *ir.Expr_GetLocal:
			e.Name = "get_local"
			e.Data["Var"] = stackVar(k.GetLocal.Returns[0])
			e.Data["Local"] = fmt.Sprintf("l%d", k.GetLocal.Index)
		case *ir.Expr_SetLocal:
			e.Name = "set_local"
			e.Data["Local"] = fmt.Sprintf("l%d", k.SetLocal.Index)
			e.Data["Var"] = stackVar(k.SetLocal.Param)
		case *ir.Expr_TeeLocal:
			e.Name = "tee_local"
			e.Data["Local"] = fmt.Sprintf("l%d", k.TeeLocal.Index)
			e.Data["Var"] = stackVar(k.TeeLocal.Param)
		case *ir.Expr_GetGlobal:
			if int(k.GetGlobal.Index) >= len(globals) {
				panic(fmt.Sprintf("global index out of bounds: %d", k.GetGlobal.Index))
			}
			e.Name = "get_global"
			e.Data["Var"] = stackVar(k.GetGlobal.Returns[0])
			e.Data["GlobalIdx"] = k.GetGlobal.Index
		case *ir.Expr_SetGlobal:
			if int(k.SetGlobal.Index) >= len(globals) {
				panic(fmt.Sprintf("global index out of bounds: %d", k.SetGlobal.Index))
			}
			e.Name = "set_global"
			e.Data["GlobalIdx"] = k.SetGlobal.Index
			e.Data["Var"] = stackVar(k.SetGlobal.Param)
		case *ir.Expr_Call:
			e.Name = "call"
			if len(k.Call.Returns) != 0 {
				e.Data["Var"] = stackVar(k.Call.Returns[0])
			}
			e.Data["FuncIdx"] = k.Call.Index
			var params []string
			for _, p := range k.Call.Params {
				params = append(params, stackVar(p))
			}
			e.Data["Params"] = params
		case *ir.Expr_CallIndirect:
			e.Name = "call_indirect"
			if len(k.CallIndirect.Returns) != 0 {
				e.Data["Var"] = stackVar(k.CallIndirect.Returns[0])
				e.Data["VarType"] = typ(k.CallIndirect.Returns[0].ValueType)
			}
			e.Data["TableVar"] = stackVar(k.CallIndirect.Params[len(k.CallIndirect.Params)-1])
			var params, paramTypes []string
			for _, v := range k.CallIndirect.Params[:len(k.CallIndirect.Params)-1] {
				params = append(params, stackVar(v))
				paramTypes = append(paramTypes, typ(v.ValueType))
			}
			e.Data["Params"] = params
			e.Data["ParamTypes"] = paramTypes
		case *ir.Expr_Select:
			e.Name = "select"
			e.Data["Var"] = stackVar(k.Select.Returns[0])
			e.Data["Param1"] = stackVar(k.Select.Param1)
			e.Data["Param2"] = stackVar(k.Select.Param2)
			e.Data["Param3"] = stackVar(k.Select.Param3)
		case *ir.Expr_Const:
			e.Name = "const"
			e.Data["Var"] = stackVar(k.Const.Returns[0])
			e.Data["Value"] = val(k.Const.Value)
		case *ir.Expr_Binary:
			e.Name = "binary"
			e.Data["Var"] = stackVar(k.Binary.Returns[0])
			e.Data["Param1"] = stackVar(k.Binary.Param1)
			e.Data["Param2"] = stackVar(k.Binary.Param2)
			op, err := operators.New(byte(k.Binary.OpCode))
			if err != nil {
				panic(err)
			}
			e.Data["Op"] = op.Name
		case *ir.Expr_Unary:
			e.Name = "unary"
			e.Data["Var"] = stackVar(k.Unary.Returns[0])
			e.Data["Param"] = stackVar(k.Unary.Param)
			op, err := operators.New(byte(k.Unary.OpCode))
			if err != nil {
				panic(err)
			}
			e.Data["Op"] = op.Name
		case *ir.Expr_Load:
			e.Name = "load"
			e.Data["Var"] = stackVar(k.Load.Returns[0])
			e.Data["Param"] = stackVar(k.Load.Param)
			e.Data["Offset"] = k.Load.Offset
			op, err := operators.New(byte(k.Load.OpCode))
			if err != nil {
				panic(err)
			}
			e.Data["Op"] = op.Name
		case *ir.Expr_Store:
			e.Name = "store"
			e.Data["Param1"] = stackVar(k.Store.Param1)
			e.Data["Param2"] = stackVar(k.Store.Param2)
			e.Data["Offset"] = k.Store.Offset
			op, err := operators.New(byte(k.Store.OpCode))
			if err != nil {
				panic(err)
			}
			e.Data["Op"] = op.Name
		case *ir.Expr_CurrentMemory:
			e.Name = "current_mem"
			e.Data["Var"] = stackVar(k.CurrentMemory.Returns[0])
		case *ir.Expr_GrowMemory:
			e.Name = "grow_mem"
			e.Data["Var"] = stackVar(k.GrowMemory.Returns[0])
			e.Data["Param"] = stackVar(k.GrowMemory.Param)
		case *ir.Expr_Loop:
			e.Name = "loop"
			blocks = append(blocks, x)
			e.Data["HasLabel"] = false
			if k.Loop.LabelUsed {
				e.Data["HasLabel"] = true
				e.Data["LabelID"] = k.Loop.LabelId
			}
		case *ir.Expr_Block:
			e.Name = "block"
			blocks = append(blocks, x)
		case *ir.Expr_If:
			e.Name = "if"
			blocks = append(blocks, x)
			e.Data["Param"] = stackVar(k.If.Param)
		case *ir.Expr_Else:
			e.Name = "else"
		case *ir.Expr_Br:
			e.Name = "br"
			e.Data["LabelID"] = k.Br.LabelId
		case *ir.Expr_BrIf:
			e.Name = "br_if"
			e.Data["Param"] = stackVar(k.BrIf.Param)
			e.Data["LabelID"] = k.BrIf.LabelId
		case *ir.Expr_BrTable:
			e.Name = "br_table"
			e.Data["Param"] = stackVar(k.BrTable.Param)
			e.Data["DefaultLabelID"] = k.BrTable.Labels[len(k.BrTable.Labels)-1]
			type c struct {
				Idx     int
				LabelID uint32
			}
			var cs []c
			for i := 0; i < len(k.BrTable.Labels)-1; i++ {
				cs = append(cs, c{Idx: i, LabelID: k.BrTable.Labels[i]})
			}
			e.Data["Cases"] = cs
		case *ir.Expr_End:
			e.Name = "end"
			e.Data["HasLabel"] = false
			if len(blocks) != 0 {
				b := blocks[len(blocks)-1]
				blocks = blocks[:len(blocks)-1]
				switch bk := b.Kind.(type) {
				case *ir.Expr_If:
					e.Name = "end_if"
					if bk.If.LabelUsed {
						e.Data["HasLabel"] = true
						e.Data["LabelID"] = bk.If.LabelId
					}
				case *ir.Expr_Block:
					e.Name = "end_block"
					if bk.Block.LabelUsed {
						e.Data["HasLabel"] = true
						e.Data["LabelID"] = bk.Block.LabelId
					}
				}
			}
		default:
			panic(fmt.Sprintf("unhandled expr: %T", k))
		}
		exprs = append(exprs, e)
	}
	return exprs
}
