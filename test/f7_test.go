package test

import (
	"encoding/binary"
)

func f7(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// return
	return s0i32
}
