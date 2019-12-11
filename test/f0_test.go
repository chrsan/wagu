package test

import (
	"encoding/binary"
)

func f0(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	// current_mem
	s0i32 = int32(len(ctx.Mem) / 65536)
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1536
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l0 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l2 = s0i32
	// get_local
	s1i32 = l1
	// const
	s2i32 = 16
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1520
		// const
		s1i32 = 48
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// const
		s0i32 = -1
		// return
		return s0i32
		// end_if
	}
	// const
	s0i32 = 1536
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// return
	return s0i32
}
