package gen

import (
	"fmt"
	"math"

	"github.com/chrsan/wagu/gen/internal"
	"github.com/chrsan/wagu/ir"
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o=internal/bindata.go -pkg=internal -modtime=1 ./template/...

func lookupTemplate(name string) ([]byte, error) {
	return internal.Asset(fmt.Sprintf("template/%s.go.tmpl", name))
}

func typ(arg interface{}) string {
	var t ir.ValueType
	switch x := arg.(type) {
	case ir.ValueType:
		t = x
	case *ir.Value:
		switch x.Kind.(type) {
		case *ir.Value_I32:
			t = ir.ValueType_I32
		case *ir.Value_I64:
			t = ir.ValueType_I64
		case *ir.Value_F32:
			t = ir.ValueType_F32
		case *ir.Value_F64:
			t = ir.ValueType_F64
		}
	default:
		panic(fmt.Sprintf("unexpected arg: %T", arg))
	}
	switch t {
	case ir.ValueType_I32:
		return "int32"
	case ir.ValueType_I64:
		return "int64"
	case ir.ValueType_F32:
		return "float32"
	default:
		return "float64"
	}
}

func val(v *ir.Value) interface{} {
	var x interface{}
	switch vk := v.Kind.(type) {
	case *ir.Value_I32:
		x = vk.I32
	case *ir.Value_I64:
		x = vk.I64
	case *ir.Value_F32:
		switch {
		case math.IsNaN(float64(vk.F32)):
			x = "float32(math.NaN())"
		case math.IsInf(float64(vk.F32), 0):
			x = "float32(math.Inf(0))"
		default:
			x = vk.F32
		}
	case *ir.Value_F64:
		switch {
		case math.IsNaN(vk.F64):
			x = "math.NaN()"
		case math.IsInf(vk.F64, 0):
			x = "math.Inf(0)"
		default:
			return vk.F64
		}
	}
	return x
}
