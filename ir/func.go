package ir

import (
	"fmt"
	"sort"

	"github.com/go-interpreter/wagon/disasm"
	"github.com/go-interpreter/wagon/wasm"
	"github.com/go-interpreter/wagon/wasm/operators"
)

func readFunc(m *Module, wm *wasm.Module, f wasm.Function, idx int) (*Function, error) {
	var locals []ValueType
	for _, l := range f.Body.Locals {
		for i := 0; i < int(l.Count); i++ {
			locals = append(locals, valueType(l.Type))
		}
	}
	fn := &Function{
		Id:        uint32(idx + len(m.ImportedFunctions)),
		Sig:       funcSig(f.Sig),
		Locals:    locals,
		StackVars: new(StackVars),
	}
	c := fc{fn: fn, stackVars: map[ValueType]map[uint32]struct{}{}}
	instr, err := disasm.Disassemble(f.Body.Code)
	if err != nil {
		return nil, err
	}
	for _, ins := range instr {
		switch op := ins.Op; op.Code {
		case operators.Nop:
		case operators.Drop:
			c.stack = c.stack[0 : len(c.stack)-1]
		case operators.Unreachable:
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Unreachable{
					Unreachable: UnreachableExpr_UNREACHABLE,
				},
			})
		case operators.Select:
			p := c.popStack(3)
			r := c.pushStack(p[0].ValueType)
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Select{
					Select: &SelectExpr{
						Param1:  p[0],
						Param2:  p[1],
						Param3:  p[2],
						Returns: []*StackVar{r},
					},
				},
			})
		case operators.I32Const:
			r := c.pushStack(valueType(op.Returns))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Const{
					Const: &ConstExpr{
						Value: &Value{
							Kind: &Value_I32{
								I32: ins.Immediates[0].(int32),
							},
						},
						Returns: []*StackVar{r},
					},
				},
			})
		case operators.I64Const:
			r := c.pushStack(valueType(op.Returns))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Const{
					Const: &ConstExpr{
						Value: &Value{
							Kind: &Value_I64{
								I64: ins.Immediates[0].(int64),
							},
						},
						Returns: []*StackVar{r},
					},
				},
			})
		case operators.F32Const:
			r := c.pushStack(valueType(op.Returns))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Const{
					Const: &ConstExpr{
						Value: &Value{
							Kind: &Value_F32{
								F32: ins.Immediates[0].(float32),
							},
						},
						Returns: []*StackVar{r},
					},
				},
			})
		case operators.F64Const:
			r := c.pushStack(valueType(op.Returns))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Const{
					Const: &ConstExpr{
						Value: &Value{
							Kind: &Value_F64{
								F64: ins.Immediates[0].(float64),
							},
						},
						Returns: []*StackVar{r},
					},
				},
			})
		case operators.CurrentMemory:
			r := c.pushStack(valueType(op.Returns))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_CurrentMemory{
					CurrentMemory: &CurrentMemoryExpr{
						Returns: []*StackVar{r},
					},
				},
			})
		case operators.I32Add, operators.I32Sub, operators.I32Mul, operators.I32DivS, operators.I32DivU, operators.I32RemS,
			operators.I32RemU, operators.I32And, operators.I32Or, operators.I32Xor, operators.I32Shl, operators.I32ShrS,
			operators.I32ShrU, operators.I32Rotl, operators.I32Rotr, operators.I32Eq, operators.I32Ne, operators.I32LtS,
			operators.I32LtU, operators.I32LeS, operators.I32LeU, operators.I32GtS, operators.I32GtU, operators.I32GeS,
			operators.I32GeU, operators.I64Add, operators.I64Sub, operators.I64Mul, operators.I64DivS, operators.I64DivU,
			operators.I64RemS, operators.I64RemU, operators.I64And, operators.I64Or, operators.I64Xor, operators.I64Shl,
			operators.I64ShrS, operators.I64ShrU, operators.I64Rotl, operators.I64Rotr, operators.I64Eq, operators.I64Ne,
			operators.I64LtS, operators.I64LtU, operators.I64LeS, operators.I64LeU, operators.I64GtS, operators.I64GtU,
			operators.I64GeS, operators.I64GeU, operators.F32Add, operators.F32Sub, operators.F32Mul, operators.F32Div,
			operators.F32Min, operators.F32Max, operators.F32Copysign, operators.F32Eq, operators.F32Ne, operators.F32Lt,
			operators.F32Le, operators.F32Gt, operators.F32Ge, operators.F64Add, operators.F64Sub, operators.F64Mul,
			operators.F64Div, operators.F64Min, operators.F64Max, operators.F64Copysign, operators.F64Eq, operators.F64Ne,
			operators.F64Lt, operators.F64Le, operators.F64Gt, operators.F64Ge:
			p := c.popStack(2)
			r := c.pushStack(valueType(op.Returns))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Binary{
					Binary: &BinaryExpr{
						OpCode:  uint32(op.Code),
						Param1:  p[0],
						Param2:  p[1],
						Returns: []*StackVar{r},
					},
				},
			})
		case operators.I32Clz, operators.I32Ctz, operators.I32Popcnt, operators.I32Eqz, operators.I64Clz, operators.I64Ctz,
			operators.I64Popcnt, operators.I64Eqz, operators.F32Sqrt, operators.F32Ceil, operators.F32Floor, operators.F32Trunc,
			operators.F32Nearest, operators.F32Abs, operators.F32Neg, operators.F64Sqrt, operators.F64Ceil, operators.F64Floor,
			operators.F64Trunc, operators.F64Nearest, operators.F64Abs, operators.F64Neg, operators.I32WrapI64, operators.I64ExtendUI32,
			operators.I64ExtendSI32, operators.I32TruncUF32, operators.I32TruncUF64, operators.I64TruncUF32, operators.I64TruncUF64,
			operators.I32TruncSF32, operators.I32TruncSF64, operators.I64TruncSF32, operators.I64TruncSF64, operators.F32DemoteF64,
			operators.F64PromoteF32, operators.F32ConvertUI32, operators.F32ConvertUI64, operators.F64ConvertUI32, operators.F64ConvertUI64,
			operators.F32ConvertSI32, operators.F32ConvertSI64, operators.F64ConvertSI32, operators.F64ConvertSI64, operators.I32ReinterpretF32,
			operators.I64ReinterpretF64, operators.F32ReinterpretI32, operators.F64ReinterpretI64:
			p := c.popStack(1)
			r := c.pushStack(valueType(op.Returns))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Unary{
					Unary: &UnaryExpr{
						OpCode:  uint32(op.Code),
						Param:   p[0],
						Returns: []*StackVar{r},
					},
				},
			})
		case operators.GrowMemory:
			p := c.popStack(1)
			r := c.pushStack(valueType(op.Returns))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_GrowMemory{
					GrowMemory: &GrowMemoryExpr{
						Param:   p[0],
						Returns: []*StackVar{r},
					},
				},
			})
		case operators.I32Load, operators.I64Load,
			operators.I32Load8s, operators.I32Load16s, operators.I64Load8s, operators.I64Load16s, operators.I64Load32s, operators.I32Load8u,
			operators.I32Load16u, operators.I64Load8u, operators.I64Load16u, operators.I64Load32u, operators.F32Load, operators.F64Load:
			p := c.popStack(1)
			r := c.pushStack(valueType(op.Returns))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Load{
					Load: &LoadExpr{
						OpCode:  uint32(op.Code),
						Align:   ins.Immediates[0].(uint32),
						Offset:  ins.Immediates[1].(uint32),
						Param:   p[0],
						Returns: []*StackVar{r},
					},
				},
			})
		case operators.I32Store, operators.I32Store8, operators.I32Store16, operators.I64Store, operators.I64Store8, operators.I64Store16,
			operators.I64Store32, operators.F32Store, operators.F64Store:
			p := c.popStack(2)
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Store{
					Store: &StoreExpr{
						OpCode: uint32(op.Code),
						Align:  ins.Immediates[0].(uint32),
						Offset: ins.Immediates[1].(uint32),
						Param1: p[0],
						Param2: p[1],
					},
				},
			})
		case operators.GetGlobal:
			idx := ins.Immediates[0].(uint32)
			r := c.pushStack(typ(m.Globals[idx]))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_GetGlobal{
					GetGlobal: &GetGlobalExpr{
						Index:   idx,
						Returns: []*StackVar{r},
					},
				},
			})
		case operators.GetLocal:
			idx := ins.Immediates[0].(uint32)
			var vt ValueType
			if int(idx) < len(f.Sig.ParamTypes) {
				vt = valueType(f.Sig.ParamTypes[idx])
			} else {
				vt = locals[int(idx)-len(f.Sig.ParamTypes)]
			}
			r := c.pushStack(vt)
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_GetLocal{
					GetLocal: &GetLocalExpr{
						Index:   idx,
						Returns: []*StackVar{r},
					},
				},
			})
		case operators.SetGlobal:
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_SetGlobal{
					SetGlobal: &SetGlobalExpr{
						Index: ins.Immediates[0].(uint32),
						Param: c.popStack(1)[0],
					},
				},
			})
		case operators.SetLocal:
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_SetLocal{
					SetLocal: &SetLocalExpr{
						Index: ins.Immediates[0].(uint32),
						Param: c.popStack(1)[0],
					},
				},
			})
		case operators.TeeLocal:
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_TeeLocal{
					TeeLocal: &TeeLocalExpr{
						Index: ins.Immediates[0].(uint32),
						Param: c.stack[len(c.stack)-1],
					},
				},
			})
		case operators.Return:
			var r []*StackVar
			if len(f.Sig.ReturnTypes) != 0 {
				r = c.popStack(1)
			}
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Return{
					Return: &ReturnExpr{
						Returns: r,
					},
				},
			})
		case operators.Call:
			idx := ins.Immediates[0].(uint32)
			var sig *FunctionSig
			if int(idx) < len(m.ImportedFunctions) {
				sig = m.ImportedFunctions[idx].Sig
			} else {
				sig = funcSig(wm.FunctionIndexSpace[int(idx)-len(m.ImportedFunctions)].Sig)
			}
			p := c.popStack(len(sig.ParamTypes))
			var r []*StackVar
			if len(sig.ReturnTypes) != 0 {
				r = []*StackVar{c.pushStack(sig.ReturnTypes[0])}
			}
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Call{
					Call: &CallExpr{
						Index:   idx,
						Params:  p,
						Returns: r,
					},
				},
			})
		case operators.CallIndirect:
			idx := ins.Immediates[0].(uint32)
			sig := funcSig(&wm.Types.Entries[idx])
			p := c.popStack(len(sig.ParamTypes) + 1)
			var r []*StackVar
			if len(sig.ReturnTypes) != 0 {
				r = []*StackVar{c.pushStack(sig.ReturnTypes[0])}
			}
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_CallIndirect{
					CallIndirect: &CallIndirectExpr{
						Index:   idx,
						Params:  p,
						Returns: r,
					},
				},
			})
		case operators.Block:
			l := c.pushLabel()
			c.pushStackMark()
			c.pushBlockType(ins.Immediates[0].(wasm.BlockType))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Block{
					Block: &BlockExpr{
						LabelId: l,
					},
				},
			})
		case operators.Loop:
			l := c.pushLabel()
			c.pushStackMark()
			c.pushBlockType(ins.Immediates[0].(wasm.BlockType))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Loop{
					Loop: &LoopExpr{
						LabelId: l,
					},
				},
			})
		case operators.If:
			p := c.popStack(1)
			l := c.pushLabel()
			c.pushStackMark()
			c.pushBlockType(ins.Immediates[0].(wasm.BlockType))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_If{
					If: &IfExpr{
						LabelId: l,
						Param:   p[0],
					},
				},
			})
		case operators.Else:
			c.resetStack(c.stackMarks[len(c.stackMarks)-1])
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Else{
					Else: ElseExpr_ELSE,
				},
			})
		case operators.Br:
			l := c.findLabel(ins.Immediates[0].(uint32))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Br{
					Br: &BrExpr{
						LabelId: l,
					},
				},
			})
		case operators.BrIf:
			l := c.findLabel(ins.Immediates[0].(uint32))
			p := c.popStack(1)
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_BrIf{
					BrIf: &BrIfExpr{
						LabelId: l,
						Param:   p[0],
					},
				},
			})
		case operators.BrTable:
			p := c.popStack(1)
			var l []uint32
			for j := 1; j < len(ins.Immediates)-1; j++ {
				l = append(l, c.findLabel(ins.Immediates[j].(uint32)))
			}
			l = append(l, c.findLabel(ins.Immediates[len(ins.Immediates)-1].(uint32)))
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_BrTable{
					BrTable: &BrTableExpr{
						Labels: l,
						Param:  p[0],
					},
				},
			})
		case operators.End:
			c.popLabel()
			if m := c.popStackMark(); m != -1 {
				c.resetStack(m)
			}
			if t := c.popBlockType(); t != wasm.BlockTypeEmpty {
				c.pushStack(valueType(wasm.ValueType(t)))
			}
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_End{
					End: EndExpr_END,
				},
			})
		default:
			return nil, fmt.Errorf("unknown op: %s", op.Name)
		}
	}
	switch len(c.stack) {
	case 0:
	case 1:
		if len(f.Sig.ReturnTypes) != 0 {
			fn.Exprs = append(fn.Exprs, &Expr{
				Kind: &Expr_Return{
					Return: &ReturnExpr{
						Returns: c.stack,
					},
				},
			})
		}
	default:
		panic(fmt.Sprintf("%d values left in stack", len(c.stack)))
	}
	c.assignStackVars()
	return fn, nil
}

type fc struct {
	fn         *Function
	stack      []*StackVar
	labelID    uint32
	labels     []label
	stackMarks []int
	blockTypes []wasm.BlockType
	stackVars  map[ValueType]map[uint32]struct{}
}

type label struct {
	id   uint32
	expr int
}

func (c *fc) pushStack(vt ValueType) *StackVar {
	sv := &StackVar{Id: uint32(len(c.stack)), ValueType: vt}
	m, ok := c.stackVars[vt]
	if !ok {
		m = map[uint32]struct{}{}
		c.stackVars[vt] = m
	}
	m[sv.Id] = struct{}{}
	c.stack = append(c.stack, sv)
	if len(c.stack) > int(c.fn.MaxStackSize) {
		c.fn.MaxStackSize = uint32(len(c.stack))
	}
	return sv
}

func (c *fc) popStack(n int) []*StackVar {
	if len(c.stack) == 0 {
		return nil
	}
	off := len(c.stack) - n
	s := make([]*StackVar, n)
	copy(s, c.stack[off:])
	c.stack = c.stack[:off]
	return s
}

func (c *fc) resetStack(mark int) {
	c.stack = c.stack[:mark]
}

func (c *fc) pushLabel() uint32 {
	l := c.labelID
	c.labelID++
	c.labels = append(c.labels, label{id: l, expr: len(c.fn.Exprs)})
	return l
}

func (c *fc) popLabel() {
	if len(c.labels) != 0 {
		c.labels = c.labels[:len(c.labels)-1]
	}
}

func (c *fc) findLabel(i uint32) uint32 {
	l := c.labels[len(c.labels)-1-int(i)]
	e := c.fn.Exprs[l.expr]
	switch k := e.Kind.(type) {
	case *Expr_Block:
		k.Block.LabelUsed = true
	case *Expr_If:
		k.If.LabelUsed = true
	case *Expr_Loop:
		k.Loop.LabelUsed = true
	default:
		panic(k)
	}
	c.fn.HasLabels = true
	return l.id
}

func (c *fc) pushStackMark() {
	c.stackMarks = append(c.stackMarks, len(c.stack))
}

func (c *fc) popStackMark() int {
	if len(c.stackMarks) == 0 {
		return -1
	}
	m := c.stackMarks[len(c.stackMarks)-1]
	c.stackMarks = c.stackMarks[:len(c.stackMarks)-1]
	return m
}

func (c *fc) pushBlockType(t wasm.BlockType) {
	c.blockTypes = append(c.blockTypes, t)
}

func (c *fc) popBlockType() wasm.BlockType {
	if len(c.blockTypes) == 0 {
		return wasm.BlockTypeEmpty
	}
	t := c.blockTypes[len(c.blockTypes)-1]
	c.blockTypes = c.blockTypes[:len(c.blockTypes)-1]
	return t
}

func (c *fc) assignStackVars() {
	a := func(s *[]uint32, vt ValueType) {
		m := c.stackVars[vt]
		for r := range m {
			*s = append(*s, r)
		}
		sort.Slice(*s, func(i, j int) bool {
			return (*s)[i] < (*s)[j]
		})
	}
	a(&c.fn.StackVars.I32, ValueType_I32)
	a(&c.fn.StackVars.I64, ValueType_I64)
	a(&c.fn.StackVars.F32, ValueType_F32)
	a(&c.fn.StackVars.F64, ValueType_F64)
}

func typ(v *Value) ValueType {
	switch v.Kind.(type) {
	case *Value_I32:
		return ValueType_I32
	case *Value_I64:
		return ValueType_I64
	case *Value_F32:
		return ValueType_F32
	case *Value_F64:
		return ValueType_F64
	}
	panic(v.Kind)
}
