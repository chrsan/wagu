package ir

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/go-interpreter/wagon/wasm"
	"github.com/go-interpreter/wagon/wasm/leb128"
)

func readInitExpr(m *wasm.Module, b []byte, importedGlobals []*Value) (*Value, error) {
	var v *Value
	if len(b) <= 2 || b[len(b)-1] != 0x0b {
		return v, fmt.Errorf("invalid init expr: %#v", b)
	}
	t := b[0]
	b = b[1 : len(b)-1]
	switch t {
	case 0x41:
		r := bytes.NewReader(b)
		i, err := leb128.ReadVarint32(r)
		if err != nil {
			return v, err
		}
		v = &Value{
			Kind: &Value_I32{
				I32: i,
			},
		}
		return v, nil
	case 0x42:
		r := bytes.NewReader(b)
		i, err := leb128.ReadVarint64(r)
		if err != nil {
			return v, err
		}
		v = &Value{
			Kind: &Value_I64{
				I64: i,
			},
		}
		return v, nil
	case 0x43:
		if len(b) != 4 {
			return v, fmt.Errorf("invalid init expr: %#v", b)
		}
		v = &Value{
			Kind: &Value_F32{
				F32: math.Float32frombits(binary.LittleEndian.Uint32(b)),
			},
		}
		return v, nil
	case 0x44:
		if len(b) != 8 {
			return v, fmt.Errorf("invalid init expr: %#v", b)
		}
		v = &Value{
			Kind: &Value_F64{
				F64: math.Float64frombits(binary.LittleEndian.Uint64(b)),
			},
		}
		return v, nil
	case 0x23:
		r := bytes.NewReader(b)
		i, err := leb128.ReadVarUint32(r)
		if err != nil {
			return v, err
		}
		if int(i) >= len(importedGlobals) {
			return v, fmt.Errorf("invalid global index %d", i)
		}
		return importedGlobals[int(i)], nil
	}
	return v, fmt.Errorf("invalid init expr: %#v", b)
}
