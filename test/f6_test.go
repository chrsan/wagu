package test

import (
	"encoding/binary"
)

func f6(ctx *Context, l0 int32, l1 int32) {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
}
