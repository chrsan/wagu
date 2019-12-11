package test

import (
	"encoding/binary"
	"math/bits"
)

func f2(ctx *Context, l0 int32) int32 {
	var l1 int32
	_ = l1
	var l2 int32
	_ = l2
	var l3 int32
	_ = l3
	var l4 int32
	_ = l4
	var l5 int32
	_ = l5
	var l6 int32
	_ = l6
	var l7 int32
	_ = l7
	var l8 int32
	_ = l8
	var l9 int32
	_ = l9
	var l10 int32
	_ = l10
	var l11 int32
	_ = l11
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	var s2i32 int32
	_ = s2i32
	var s3i32 int32
	_ = s3i32
	var s4i32 int32
	_ = s4i32
	var s5i32 int32
	_ = s5i32
	var s1i64 int64
	_ = s1i64
	var s2i64 int64
	_ = s2i64
	// get_global
	s0i32 = ctx.g0
	// const
	s1i32 = 16
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l11 = s0i32
	// set_global
	ctx.g0 = s0i32
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// block
	// get_local
	s0i32 = l0
	// const
	s1i32 = 244
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1024
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l6 = s0i32
		// const
		s1i32 = 16
		// get_local
		s2i32 = l0
		// const
		s3i32 = 11
		// binary: i32.add
		s2i32 = s2i32 + s3i32
		// const
		s3i32 = -8
		// binary: i32.and
		s2i32 = s2i32 & s3i32
		// get_local
		s3i32 = l0
		// const
		s4i32 = 11
		// binary: i32.lt_u
		if uint32(s3i32) < uint32(s4i32) {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		// select
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		// tee_local
		l5 = s1i32
		// const
		s2i32 = 3
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// tee_local
		l0 = s1i32
		// binary: i32.shr_u
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		// tee_local
		l1 = s0i32
		// const
		s1i32 = 3
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// if
		if s0i32 != 0 {
			// get_local
			s0i32 = l1
			// const
			s1i32 = -1
			// binary: i32.xor
			s0i32 = s0i32 ^ s1i32
			// const
			s1i32 = 1
			// binary: i32.and
			s0i32 = s0i32 & s1i32
			// get_local
			s1i32 = l0
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// tee_local
			l2 = s0i32
			// const
			s1i32 = 3
			// binary: i32.shl
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			// tee_local
			l4 = s0i32
			// const
			s1i32 = 1072
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// load: i32.load
			s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
			// tee_local
			l1 = s0i32
			// const
			s1i32 = 8
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// set_local
			l0 = s0i32
			// block
			// get_local
			s0i32 = l1
			// load: i32.load
			s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
			// tee_local
			l3 = s0i32
			// get_local
			s1i32 = l4
			// const
			s2i32 = 1064
			// binary: i32.add
			s1i32 = s1i32 + s2i32
			// tee_local
			l4 = s1i32
			// binary: i32.eq
			if s0i32 == s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			// if
			if s0i32 != 0 {
				// const
				s0i32 = 1024
				// get_local
				s1i32 = l6
				// const
				s2i32 = -2
				// get_local
				s3i32 = l2
				// binary: i32.rotl
				s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
				// binary: i32.and
				s1i32 = s1i32 & s2i32
				// store: i32.store
				binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
				// br
				goto lbl13
				// end_if
			}
			// const
			s0i32 = 1040
			// load: i32.load
			s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
			// get_local
			s0i32 = l3
			// get_local
			s1i32 = l4
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
			// get_local
			s0i32 = l4
			// get_local
			s1i32 = l3
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
			// end_block
		lbl13:
			// get_local
			s0i32 = l1
			// get_local
			s1i32 = l2
			// const
			s2i32 = 3
			// binary: i32.shl
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			// tee_local
			l2 = s1i32
			// const
			s2i32 = 3
			// binary: i32.or
			s1i32 = s1i32 | s2i32
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
			// get_local
			s0i32 = l1
			// get_local
			s1i32 = l2
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// tee_local
			l1 = s0i32
			// get_local
			s1i32 = l1
			// load: i32.load
			s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
			// const
			s2i32 = 1
			// binary: i32.or
			s1i32 = s1i32 | s2i32
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
			// br
			goto lbl0
			// end_if
		}
		// get_local
		s0i32 = l5
		// const
		s1i32 = 1032
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
		// tee_local
		l8 = s1i32
		// binary: i32.le_u
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl10
		}
		// get_local
		s0i32 = l1
		// if
		if s0i32 != 0 {
			// block
			// const
			s0i32 = 2
			// get_local
			s1i32 = l0
			// binary: i32.shl
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			// tee_local
			l2 = s0i32
			// const
			s1i32 = 0
			// get_local
			s2i32 = l2
			// binary: i32.sub
			s1i32 = s1i32 - s2i32
			// binary: i32.or
			s0i32 = s0i32 | s1i32
			// get_local
			s1i32 = l1
			// get_local
			s2i32 = l0
			// binary: i32.shl
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			// binary: i32.and
			s0i32 = s0i32 & s1i32
			// tee_local
			l0 = s0i32
			// const
			s1i32 = 0
			// get_local
			s2i32 = l0
			// binary: i32.sub
			s1i32 = s1i32 - s2i32
			// binary: i32.and
			s0i32 = s0i32 & s1i32
			// const
			s1i32 = -1
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// tee_local
			l0 = s0i32
			// get_local
			s1i32 = l0
			// const
			s2i32 = 12
			// binary: i32.shr_u
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			// const
			s2i32 = 16
			// binary: i32.and
			s1i32 = s1i32 & s2i32
			// tee_local
			l0 = s1i32
			// binary: i32.shr_u
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			// tee_local
			l1 = s0i32
			// const
			s1i32 = 5
			// binary: i32.shr_u
			s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
			// const
			s1i32 = 8
			// binary: i32.and
			s0i32 = s0i32 & s1i32
			// tee_local
			l2 = s0i32
			// get_local
			s1i32 = l0
			// binary: i32.or
			s0i32 = s0i32 | s1i32
			// get_local
			s1i32 = l1
			// get_local
			s2i32 = l2
			// binary: i32.shr_u
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			// tee_local
			l0 = s1i32
			// const
			s2i32 = 2
			// binary: i32.shr_u
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			// const
			s2i32 = 4
			// binary: i32.and
			s1i32 = s1i32 & s2i32
			// tee_local
			l1 = s1i32
			// binary: i32.or
			s0i32 = s0i32 | s1i32
			// get_local
			s1i32 = l0
			// get_local
			s2i32 = l1
			// binary: i32.shr_u
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			// tee_local
			l0 = s1i32
			// const
			s2i32 = 1
			// binary: i32.shr_u
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			// const
			s2i32 = 2
			// binary: i32.and
			s1i32 = s1i32 & s2i32
			// tee_local
			l1 = s1i32
			// binary: i32.or
			s0i32 = s0i32 | s1i32
			// get_local
			s1i32 = l0
			// get_local
			s2i32 = l1
			// binary: i32.shr_u
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			// tee_local
			l0 = s1i32
			// const
			s2i32 = 1
			// binary: i32.shr_u
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			// const
			s2i32 = 1
			// binary: i32.and
			s1i32 = s1i32 & s2i32
			// tee_local
			l1 = s1i32
			// binary: i32.or
			s0i32 = s0i32 | s1i32
			// get_local
			s1i32 = l0
			// get_local
			s2i32 = l1
			// binary: i32.shr_u
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// tee_local
			l2 = s0i32
			// const
			s1i32 = 3
			// binary: i32.shl
			s0i32 = s0i32 << (uint32(s1i32) & 31)
			// tee_local
			l3 = s0i32
			// const
			s1i32 = 1072
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// load: i32.load
			s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
			// tee_local
			l1 = s0i32
			// load: i32.load
			s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
			// tee_local
			l0 = s0i32
			// get_local
			s1i32 = l3
			// const
			s2i32 = 1064
			// binary: i32.add
			s1i32 = s1i32 + s2i32
			// tee_local
			l3 = s1i32
			// binary: i32.eq
			if s0i32 == s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			// if
			if s0i32 != 0 {
				// const
				s0i32 = 1024
				// get_local
				s1i32 = l6
				// const
				s2i32 = -2
				// get_local
				s3i32 = l2
				// binary: i32.rotl
				s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
				// binary: i32.and
				s1i32 = s1i32 & s2i32
				// tee_local
				l6 = s1i32
				// store: i32.store
				binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
				// br
				goto lbl16
				// end_if
			}
			// const
			s0i32 = 1040
			// load: i32.load
			s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
			// get_local
			s0i32 = l0
			// get_local
			s1i32 = l3
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
			// get_local
			s0i32 = l3
			// get_local
			s1i32 = l0
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
			// end_block
		lbl16:
			// get_local
			s0i32 = l1
			// const
			s1i32 = 8
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// set_local
			l0 = s0i32
			// get_local
			s0i32 = l1
			// get_local
			s1i32 = l5
			// const
			s2i32 = 3
			// binary: i32.or
			s1i32 = s1i32 | s2i32
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
			// get_local
			s0i32 = l1
			// get_local
			s1i32 = l5
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// tee_local
			l7 = s0i32
			// get_local
			s1i32 = l2
			// const
			s2i32 = 3
			// binary: i32.shl
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			// tee_local
			l2 = s1i32
			// get_local
			s2i32 = l5
			// binary: i32.sub
			s1i32 = s1i32 - s2i32
			// tee_local
			l3 = s1i32
			// const
			s2i32 = 1
			// binary: i32.or
			s1i32 = s1i32 | s2i32
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
			// get_local
			s0i32 = l1
			// get_local
			s1i32 = l2
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// get_local
			s1i32 = l3
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
			// get_local
			s0i32 = l8
			// if
			if s0i32 != 0 {
				// get_local
				s0i32 = l8
				// const
				s1i32 = 3
				// binary: i32.shr_u
				s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
				// tee_local
				l4 = s0i32
				// const
				s1i32 = 3
				// binary: i32.shl
				s0i32 = s0i32 << (uint32(s1i32) & 31)
				// const
				s1i32 = 1064
				// binary: i32.add
				s0i32 = s0i32 + s1i32
				// set_local
				l1 = s0i32
				// const
				s0i32 = 1044
				// load: i32.load
				s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
				// set_local
				l2 = s0i32
				// block
				// get_local
				s0i32 = l6
				// const
				s1i32 = 1
				// get_local
				s2i32 = l4
				// binary: i32.shl
				s1i32 = s1i32 << (uint32(s2i32) & 31)
				// tee_local
				l4 = s1i32
				// binary: i32.and
				s0i32 = s0i32 & s1i32
				// unary: i32.eqz
				if s0i32 == 0 {
					s0i32 = 1
				} else {
					s0i32 = 0
				}
				// if
				if s0i32 != 0 {
					// const
					s0i32 = 1024
					// get_local
					s1i32 = l4
					// get_local
					s2i32 = l6
					// binary: i32.or
					s1i32 = s1i32 | s2i32
					// store: i32.store
					binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
					// get_local
					s0i32 = l1
					// br
					goto lbl19
					// end_if
				}
				// get_local
				s0i32 = l1
				// load: i32.load
				s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
				// end_block
			lbl19:
				// set_local
				l4 = s0i32
				// get_local
				s0i32 = l1
				// get_local
				s1i32 = l2
				// store: i32.store
				binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
				// get_local
				s0i32 = l4
				// get_local
				s1i32 = l2
				// store: i32.store
				binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
				// get_local
				s0i32 = l2
				// get_local
				s1i32 = l1
				// store: i32.store
				binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
				// get_local
				s0i32 = l2
				// get_local
				s1i32 = l4
				// store: i32.store
				binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
				// end_if
			}
			// const
			s0i32 = 1044
			// get_local
			s1i32 = l7
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
			// const
			s0i32 = 1032
			// get_local
			s1i32 = l3
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
			// br
			goto lbl0
			// end_if
		}
		// const
		s0i32 = 1028
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l10 = s0i32
		// unary: i32.eqz
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl10
		}
		// get_local
		s0i32 = l10
		// const
		s1i32 = 0
		// get_local
		s2i32 = l10
		// binary: i32.sub
		s1i32 = s1i32 - s2i32
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// const
		s1i32 = -1
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l0 = s0i32
		// get_local
		s1i32 = l0
		// const
		s2i32 = 12
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// const
		s2i32 = 16
		// binary: i32.and
		s1i32 = s1i32 & s2i32
		// tee_local
		l0 = s1i32
		// binary: i32.shr_u
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		// tee_local
		l1 = s0i32
		// const
		s1i32 = 5
		// binary: i32.shr_u
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		// const
		s1i32 = 8
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// tee_local
		l2 = s0i32
		// get_local
		s1i32 = l0
		// binary: i32.or
		s0i32 = s0i32 | s1i32
		// get_local
		s1i32 = l1
		// get_local
		s2i32 = l2
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// tee_local
		l0 = s1i32
		// const
		s2i32 = 2
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// const
		s2i32 = 4
		// binary: i32.and
		s1i32 = s1i32 & s2i32
		// tee_local
		l1 = s1i32
		// binary: i32.or
		s0i32 = s0i32 | s1i32
		// get_local
		s1i32 = l0
		// get_local
		s2i32 = l1
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// tee_local
		l0 = s1i32
		// const
		s2i32 = 1
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// const
		s2i32 = 2
		// binary: i32.and
		s1i32 = s1i32 & s2i32
		// tee_local
		l1 = s1i32
		// binary: i32.or
		s0i32 = s0i32 | s1i32
		// get_local
		s1i32 = l0
		// get_local
		s2i32 = l1
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// tee_local
		l0 = s1i32
		// const
		s2i32 = 1
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// const
		s2i32 = 1
		// binary: i32.and
		s1i32 = s1i32 & s2i32
		// tee_local
		l1 = s1i32
		// binary: i32.or
		s0i32 = s0i32 | s1i32
		// get_local
		s1i32 = l0
		// get_local
		s2i32 = l1
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// const
		s1i32 = 2
		// binary: i32.shl
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		// const
		s1i32 = 1328
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l1 = s0i32
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
		// const
		s1i32 = -8
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// get_local
		s1i32 = l5
		// binary: i32.sub
		s0i32 = s0i32 - s1i32
		// set_local
		l3 = s0i32
		// get_local
		s0i32 = l1
		// set_local
		l2 = s0i32
		// loop
	lbl21:
		// block
		// get_local
		s0i32 = l2
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
		// tee_local
		l0 = s0i32
		// unary: i32.eqz
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// get_local
			s0i32 = l2
			// load: i32.load
			s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
			// tee_local
			l0 = s0i32
			// unary: i32.eqz
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			// br_if
			if s0i32 != 0 {
				goto lbl22
			}
			// end_if
		}
		// get_local
		s0i32 = l0
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
		// const
		s1i32 = -8
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// get_local
		s1i32 = l5
		// binary: i32.sub
		s0i32 = s0i32 - s1i32
		// tee_local
		l2 = s0i32
		// get_local
		s1i32 = l3
		// get_local
		s2i32 = l2
		// get_local
		s3i32 = l3
		// binary: i32.lt_u
		if uint32(s2i32) < uint32(s3i32) {
			s2i32 = 1
		} else {
			s2i32 = 0
		}
		// tee_local
		l2 = s2i32
		// select
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		// set_local
		l3 = s0i32
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l1
		// get_local
		s2i32 = l2
		// select
		if s2i32 != 0 {
			// s0i32 = s0i32
		} else {
			s0i32 = s1i32
		}
		// set_local
		l1 = s0i32
		// get_local
		s0i32 = l0
		// set_local
		l2 = s0i32
		// br
		goto lbl21
		// end_block
	lbl22:
		// end
		// get_local
		s0i32 = l1
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
		// set_local
		l9 = s0i32
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l1
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+12):]))
		// tee_local
		l4 = s1i32
		// binary: i32.ne
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// const
			s0i32 = 1040
			// load: i32.load
			s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
			// get_local
			s1i32 = l1
			// load: i32.load
			s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+8):]))
			// tee_local
			l0 = s1i32
			// binary: i32.le_u
			if uint32(s0i32) <= uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			// if
			if s0i32 != 0 {
				// get_local
				s0i32 = l0
				// load: i32.load
				s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
				// end_if
			}
			// get_local
			s0i32 = l0
			// get_local
			s1i32 = l4
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
			// get_local
			s0i32 = l4
			// get_local
			s1i32 = l0
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
			// br
			goto lbl1
			// end_if
		}
		// get_local
		s0i32 = l1
		// const
		s1i32 = 20
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l2 = s0i32
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l0 = s0i32
		// unary: i32.eqz
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// get_local
			s0i32 = l1
			// load: i32.load
			s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
			// tee_local
			l0 = s0i32
			// unary: i32.eqz
			if s0i32 == 0 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			// br_if
			if s0i32 != 0 {
				goto lbl9
			}
			// get_local
			s0i32 = l1
			// const
			s1i32 = 16
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// set_local
			l2 = s0i32
			// end_if
		}
		// loop
	lbl27:
		// get_local
		s0i32 = l2
		// set_local
		l7 = s0i32
		// get_local
		s0i32 = l0
		// tee_local
		l4 = s0i32
		// const
		s1i32 = 20
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l2 = s0i32
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l0 = s0i32
		// br_if
		if s0i32 != 0 {
			goto lbl27
		}
		// get_local
		s0i32 = l4
		// const
		s1i32 = 16
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l2 = s0i32
		// get_local
		s0i32 = l4
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
		// tee_local
		l0 = s0i32
		// br_if
		if s0i32 != 0 {
			goto lbl27
		}
		// end
		// get_local
		s0i32 = l7
		// const
		s1i32 = 0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// br
		goto lbl1
		// end_if
	}
	// const
	s0i32 = -1
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = -65
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl10
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 11
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l0 = s0i32
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// set_local
	l5 = s0i32
	// const
	s0i32 = 1028
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l7 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl10
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l5
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l2 = s0i32
	// block
	// block
	// block
	// block
	// const
	s0i32 = 0
	// get_local
	s1i32 = l0
	// const
	s2i32 = 8
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// tee_local
	l0 = s1i32
	// unary: i32.eqz
	if s1i32 == 0 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	// br_if
	if s1i32 != 0 {
		goto lbl31
	}
	// const
	s0i32 = 31
	// get_local
	s1i32 = l5
	// const
	s2i32 = 16777215
	// binary: i32.gt_u
	if uint32(s1i32) > uint32(s2i32) {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	// br_if
	if s1i32 != 0 {
		goto lbl31
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// const
	s2i32 = 1048320
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 8
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l0 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l1 = s0i32
	// get_local
	s1i32 = l1
	// const
	s2i32 = 520192
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l1 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 245760
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = 16
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// tee_local
	l3 = s1i32
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 15
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// get_local
	s2i32 = l3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// get_local
	s1i32 = l5
	// get_local
	s2i32 = l0
	// const
	s3i32 = 21
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// const
	s2i32 = 1
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// const
	s1i32 = 28
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// end_block
lbl31:
	// tee_local
	l8 = s0i32
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 1328
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 0
		// set_local
		l0 = s0i32
		// br
		goto lbl30
		// end_if
	}
	// get_local
	s0i32 = l5
	// const
	s1i32 = 0
	// const
	s2i32 = 25
	// get_local
	s3i32 = l8
	// const
	s4i32 = 1
	// binary: i32.shr_u
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// get_local
	s3i32 = l8
	// const
	s4i32 = 31
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l1 = s0i32
	// const
	s0i32 = 0
	// set_local
	l0 = s0i32
	// loop
lbl33:
	// block
	// get_local
	s0i32 = l3
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l5
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l6 = s0i32
	// get_local
	s1i32 = l2
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl34
	}
	// get_local
	s0i32 = l3
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l6
	// tee_local
	l2 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl34
	}
	// const
	s0i32 = 0
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l3
	// set_local
	l0 = s0i32
	// br
	goto lbl29
	// end_block
lbl34:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l3
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+20):]))
	// tee_local
	l6 = s1i32
	// get_local
	s2i32 = l6
	// get_local
	s3i32 = l3
	// get_local
	s4i32 = l1
	// const
	s5i32 = 29
	// binary: i32.shr_u
	s4i32 = int32(uint32(s4i32) >> (uint32(s5i32) & 31))
	// const
	s5i32 = 4
	// binary: i32.and
	s4i32 = s4i32 & s5i32
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+16):]))
	// tee_local
	l3 = s3i32
	// binary: i32.eq
	if s2i32 == s3i32 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l6
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l3
	// const
	s2i32 = 0
	// binary: i32.ne
	if s1i32 != s2i32 {
		s1i32 = 1
	} else {
		s1i32 = 0
	}
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l3
	// br_if
	if s0i32 != 0 {
		goto lbl33
	}
	// end
	// end_block
lbl30:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l4
	// binary: i32.or
	s0i32 = s0i32 | s1i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 2
		// get_local
		s1i32 = l8
		// binary: i32.shl
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		// tee_local
		l0 = s0i32
		// const
		s1i32 = 0
		// get_local
		s2i32 = l0
		// binary: i32.sub
		s1i32 = s1i32 - s2i32
		// binary: i32.or
		s0i32 = s0i32 | s1i32
		// get_local
		s1i32 = l7
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// tee_local
		l0 = s0i32
		// unary: i32.eqz
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl10
		}
		// get_local
		s0i32 = l0
		// const
		s1i32 = 0
		// get_local
		s2i32 = l0
		// binary: i32.sub
		s1i32 = s1i32 - s2i32
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// const
		s1i32 = -1
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l0 = s0i32
		// get_local
		s1i32 = l0
		// const
		s2i32 = 12
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// const
		s2i32 = 16
		// binary: i32.and
		s1i32 = s1i32 & s2i32
		// tee_local
		l0 = s1i32
		// binary: i32.shr_u
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		// tee_local
		l1 = s0i32
		// const
		s1i32 = 5
		// binary: i32.shr_u
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		// const
		s1i32 = 8
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// tee_local
		l3 = s0i32
		// get_local
		s1i32 = l0
		// binary: i32.or
		s0i32 = s0i32 | s1i32
		// get_local
		s1i32 = l1
		// get_local
		s2i32 = l3
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// tee_local
		l0 = s1i32
		// const
		s2i32 = 2
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// const
		s2i32 = 4
		// binary: i32.and
		s1i32 = s1i32 & s2i32
		// tee_local
		l1 = s1i32
		// binary: i32.or
		s0i32 = s0i32 | s1i32
		// get_local
		s1i32 = l0
		// get_local
		s2i32 = l1
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// tee_local
		l0 = s1i32
		// const
		s2i32 = 1
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// const
		s2i32 = 2
		// binary: i32.and
		s1i32 = s1i32 & s2i32
		// tee_local
		l1 = s1i32
		// binary: i32.or
		s0i32 = s0i32 | s1i32
		// get_local
		s1i32 = l0
		// get_local
		s2i32 = l1
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// tee_local
		l0 = s1i32
		// const
		s2i32 = 1
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// const
		s2i32 = 1
		// binary: i32.and
		s1i32 = s1i32 & s2i32
		// tee_local
		l1 = s1i32
		// binary: i32.or
		s0i32 = s0i32 | s1i32
		// get_local
		s1i32 = l0
		// get_local
		s2i32 = l1
		// binary: i32.shr_u
		s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// const
		s1i32 = 2
		// binary: i32.shl
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		// const
		s1i32 = 1328
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// set_local
		l0 = s0i32
		// end_if
	}
	// get_local
	s0i32 = l0
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl28
	}
	// end_block
lbl29:
	// loop
lbl36:
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l5
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l2
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// get_local
	s2i32 = l1
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l2 = s0i32
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l1
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// set_local
	l4 = s0i32
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l1 = s0i32
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l1
		// else
	} else {
		// get_local
		s0i32 = l0
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
		// end_if
	}
	// tee_local
	l0 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl36
	}
	// end
	// end_block
lbl28:
	// get_local
	s0i32 = l4
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl10
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1032
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i32 = l5
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl10
	}
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
	// set_local
	l8 = s0i32
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l4
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+12):]))
	// tee_local
	l1 = s1i32
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1040
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// get_local
		s1i32 = l4
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+8):]))
		// tee_local
		l0 = s1i32
		// binary: i32.le_u
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// get_local
			s0i32 = l0
			// load: i32.load
			s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
			// end_if
		}
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l1
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
		// br
		goto lbl2
		// end_if
	}
	// get_local
	s0i32 = l4
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l0 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l4
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
		// tee_local
		l0 = s0i32
		// unary: i32.eqz
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl8
		}
		// get_local
		s0i32 = l4
		// const
		s1i32 = 16
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l3 = s0i32
		// end_if
	}
	// loop
lbl41:
	// get_local
	s0i32 = l3
	// set_local
	l6 = s0i32
	// get_local
	s0i32 = l0
	// tee_local
	l1 = s0i32
	// const
	s1i32 = 20
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l0 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl41
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l1
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l0 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl41
	}
	// end
	// get_local
	s0i32 = l6
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl2
	// end_block
lbl10:
	// const
	s0i32 = 1032
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l1 = s0i32
	// get_local
	s1i32 = l5
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1044
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// set_local
		l0 = s0i32
		// block
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l5
		// binary: i32.sub
		s0i32 = s0i32 - s1i32
		// tee_local
		l2 = s0i32
		// const
		s1i32 = 16
		// binary: i32.ge_u
		if uint32(s0i32) >= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// const
			s0i32 = 1032
			// get_local
			s1i32 = l2
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
			// const
			s0i32 = 1044
			// get_local
			s1i32 = l0
			// get_local
			s2i32 = l5
			// binary: i32.add
			s1i32 = s1i32 + s2i32
			// tee_local
			l3 = s1i32
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
			// get_local
			s0i32 = l3
			// get_local
			s1i32 = l2
			// const
			s2i32 = 1
			// binary: i32.or
			s1i32 = s1i32 | s2i32
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
			// get_local
			s0i32 = l0
			// get_local
			s1i32 = l1
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// get_local
			s1i32 = l2
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
			// get_local
			s0i32 = l0
			// get_local
			s1i32 = l5
			// const
			s2i32 = 3
			// binary: i32.or
			s1i32 = s1i32 | s2i32
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
			// br
			goto lbl43
			// end_if
		}
		// const
		s0i32 = 1044
		// const
		s1i32 = 0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// const
		s0i32 = 1032
		// const
		s1i32 = 0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l1
		// const
		s2i32 = 3
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l1
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l1 = s0i32
		// get_local
		s1i32 = l1
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
		// const
		s2i32 = 1
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
		// end_block
	lbl43:
		// get_local
		s0i32 = l0
		// const
		s1i32 = 8
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l0 = s0i32
		// br
		goto lbl0
		// end_if
	}
	// const
	s0i32 = 1036
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l1 = s0i32
	// get_local
	s1i32 = l5
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1036
		// get_local
		s1i32 = l1
		// get_local
		s2i32 = l5
		// binary: i32.sub
		s1i32 = s1i32 - s2i32
		// tee_local
		l1 = s1i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// const
		s0i32 = 1048
		// const
		s1i32 = 1048
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
		// tee_local
		l0 = s1i32
		// get_local
		s2i32 = l5
		// binary: i32.add
		s1i32 = s1i32 + s2i32
		// tee_local
		l2 = s1i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// get_local
		s0i32 = l2
		// get_local
		s1i32 = l1
		// const
		s2i32 = 1
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l5
		// const
		s2i32 = 3
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
		// get_local
		s0i32 = l0
		// const
		s1i32 = 8
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l0 = s0i32
		// br
		goto lbl0
		// end_if
	}
	// const
	s0i32 = 0
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l5
	// const
	s1i32 = 47
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// block
	// const
	s1i32 = 1496
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// if
	if s1i32 != 0 {
		// const
		s1i32 = 1504
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
		// br
		goto lbl46
		// end_if
	}
	// const
	s1i32 = 1508
	// const
	s2i64 = -1
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s1i32+0):], uint64(s2i64))
	// const
	s1i32 = 1500
	// const
	s2i64 = 17592186048512
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s1i32+0):], uint64(s2i64))
	// const
	s1i32 = 1496
	// get_local
	s2i32 = l11
	// const
	s3i32 = 12
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = -16
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// const
	s3i32 = 1431655768
	// binary: i32.xor
	s2i32 = s2i32 ^ s3i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s1i32+0):], uint32(s2i32))
	// const
	s1i32 = 1516
	// const
	s2i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s1i32+0):], uint32(s2i32))
	// const
	s1i32 = 1468
	// const
	s2i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s1i32+0):], uint32(s2i32))
	// const
	s1i32 = 4096
	// end_block
lbl46:
	// tee_local
	l2 = s1i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l6 = s0i32
	// const
	s1i32 = 0
	// get_local
	s2i32 = l2
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l7 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l2 = s0i32
	// get_local
	s1i32 = l5
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl0
	}
	// const
	s0i32 = 1464
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1456
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l8 = s0i32
		// get_local
		s1i32 = l2
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l9 = s0i32
		// get_local
		s1i32 = l8
		// binary: i32.le_u
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl0
		}
		// get_local
		s0i32 = l9
		// get_local
		s1i32 = l3
		// binary: i32.gt_u
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl0
		}
		// end_if
	}
	// const
	s0i32 = 1468
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+0)])
	// const
	s1i32 = 4
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl5
	}
	// block
	// block
	// const
	s0i32 = 1048
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1472
		// set_local
		l0 = s0i32
		// loop
	lbl52:
		// get_local
		s0i32 = l0
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l8 = s0i32
		// get_local
		s1i32 = l3
		// binary: i32.le_u
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// get_local
			s0i32 = l8
			// get_local
			s1i32 = l0
			// load: i32.load
			s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
			// binary: i32.add
			s0i32 = s0i32 + s1i32
			// get_local
			s1i32 = l3
			// binary: i32.gt_u
			if uint32(s0i32) > uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			// br_if
			if s0i32 != 0 {
				goto lbl50
			}
			// end_if
		}
		// get_local
		s0i32 = l0
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
		// tee_local
		l0 = s0i32
		// br_if
		if s0i32 != 0 {
			goto lbl52
		}
		// end
		// end_if
	}
	// const
	s0i32 = 0
	// call
	s0i32 = f0(ctx, s0i32)
	// tee_local
	l1 = s0i32
	// const
	s1i32 = -1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0i32 = l2
	// set_local
	l6 = s0i32
	// const
	s0i32 = 1500
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l0 = s0i32
	// const
	s1i32 = -1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l1
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l2
		// get_local
		s1i32 = l1
		// binary: i32.sub
		s0i32 = s0i32 - s1i32
		// get_local
		s1i32 = l1
		// get_local
		s2i32 = l3
		// binary: i32.add
		s1i32 = s1i32 + s2i32
		// const
		s2i32 = 0
		// get_local
		s3i32 = l0
		// binary: i32.sub
		s2i32 = s2i32 - s3i32
		// binary: i32.and
		s1i32 = s1i32 & s2i32
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l6 = s0i32
		// end_if
	}
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l5
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0i32 = l6
	// const
	s1i32 = 2147483646
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// const
	s0i32 = 1464
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l0 = s0i32
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1456
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l3 = s0i32
		// get_local
		s1i32 = l6
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l7 = s0i32
		// get_local
		s1i32 = l3
		// binary: i32.le_u
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl6
		}
		// get_local
		s0i32 = l7
		// get_local
		s1i32 = l0
		// binary: i32.gt_u
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl6
		}
		// end_if
	}
	// get_local
	s0i32 = l6
	// call
	s0i32 = f0(ctx, s0i32)
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl49
	}
	// br
	goto lbl4
	// end_block
lbl50:
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l1
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// get_local
	s1i32 = l7
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l6 = s0i32
	// const
	s1i32 = 2147483646
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl6
	}
	// get_local
	s0i32 = l6
	// call
	s0i32 = f0(ctx, s0i32)
	// tee_local
	l1 = s0i32
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i32 = l0
	// load: i32.load
	s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+4):]))
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl7
	}
	// get_local
	s0i32 = l1
	// set_local
	l0 = s0i32
	// end_block
lbl49:
	// get_local
	s0i32 = l0
	// set_local
	l1 = s0i32
	// block
	// get_local
	s0i32 = l5
	// const
	s1i32 = 48
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l6
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl56
	}
	// get_local
	s0i32 = l6
	// const
	s1i32 = 2147483646
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl56
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = -1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl56
	}
	// const
	s0i32 = 1504
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l6
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 0
	// get_local
	s2i32 = l0
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 2147483646
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// get_local
	s0i32 = l0
	// call
	s0i32 = f0(ctx, s0i32)
	// const
	s1i32 = -1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l6
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l6 = s0i32
		// br
		goto lbl4
		// end_if
	}
	// const
	s0i32 = 0
	// get_local
	s1i32 = l6
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// call
	s0i32 = f0(ctx, s0i32)
	// br
	goto lbl6
	// end_block
lbl56:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// br
	goto lbl6
	// end_block
lbl9:
	// const
	s0i32 = 0
	// set_local
	l4 = s0i32
	// br
	goto lbl1
	// end_block
lbl8:
	// const
	s0i32 = 0
	// set_local
	l1 = s0i32
	// br
	goto lbl2
	// end_block
lbl7:
	// get_local
	s0i32 = l1
	// const
	s1i32 = -1
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl4
	}
	// end_block
lbl6:
	// const
	s0i32 = 1468
	// const
	s1i32 = 1468
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// const
	s2i32 = 4
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl5:
	// get_local
	s0i32 = l2
	// const
	s1i32 = 2147483646
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l2
	// call
	s0i32 = f0(ctx, s0i32)
	// tee_local
	l1 = s0i32
	// const
	s1i32 = 0
	// call
	s1i32 = f0(ctx, s1i32)
	// tee_local
	l0 = s1i32
	// binary: i32.ge_u
	if uint32(s0i32) >= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = -1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = -1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// tee_local
	l6 = s0i32
	// get_local
	s1i32 = l5
	// const
	s2i32 = 40
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// end_block
lbl4:
	// const
	s0i32 = 1456
	// const
	s1i32 = 1456
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i32 = l6
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l0 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1460
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1460
		// get_local
		s1i32 = l0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// end_if
	}
	// block
	// block
	// block
	// const
	s0i32 = 1048
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1472
		// set_local
		l0 = s0i32
		// loop
	lbl63:
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l0
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
		// tee_local
		l2 = s1i32
		// get_local
		s2i32 = l0
		// load: i32.load
		s2i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s2i32+4):]))
		// tee_local
		l4 = s2i32
		// binary: i32.add
		s1i32 = s1i32 + s2i32
		// binary: i32.eq
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl61
		}
		// get_local
		s0i32 = l0
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
		// tee_local
		l0 = s0i32
		// br_if
		if s0i32 != 0 {
			goto lbl63
		}
		// end
		// br
		goto lbl60
		// end_if
	}
	// const
	s0i32 = 1040
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 0
	// get_local
	s2i32 = l1
	// get_local
	s3i32 = l0
	// binary: i32.ge_u
	if uint32(s2i32) >= uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1040
		// get_local
		s1i32 = l1
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// end_if
	}
	// const
	s0i32 = 0
	// set_local
	l0 = s0i32
	// const
	s0i32 = 1476
	// get_local
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1472
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1056
	// const
	s1i32 = -1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1060
	// const
	s1i32 = 1496
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1484
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// loop
lbl65:
	// get_local
	s0i32 = l0
	// const
	s1i32 = 3
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 1072
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 1064
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l3 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1076
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 32
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl65
	}
	// end
	// const
	s0i32 = 1036
	// get_local
	s1i32 = l6
	// const
	s2i32 = -40
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l0 = s1i32
	// const
	s2i32 = -8
	// get_local
	s3i32 = l1
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// const
	s3i32 = 7
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// const
	s3i32 = 0
	// get_local
	s4i32 = l1
	// const
	s5i32 = 8
	// binary: i32.add
	s4i32 = s4i32 + s5i32
	// const
	s5i32 = 7
	// binary: i32.and
	s4i32 = s4i32 & s5i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// tee_local
	l2 = s2i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l3 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1048
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l2 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 40
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 1052
	// const
	s1i32 = 1512
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl59
	// end_block
lbl61:
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+12)])
	// const
	s1i32 = 8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// br_if
	if s0i32 != 0 {
		goto lbl60
	}
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l3
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl60
	}
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// binary: i32.gt_u
	if uint32(s0i32) > uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl60
	}
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l4
	// get_local
	s2i32 = l6
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 1048
	// get_local
	s1i32 = l3
	// const
	s2i32 = -8
	// get_local
	s3i32 = l3
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// const
	s3i32 = 7
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// const
	s3i32 = 0
	// get_local
	s4i32 = l3
	// const
	s5i32 = 8
	// binary: i32.add
	s4i32 = s4i32 + s5i32
	// const
	s5i32 = 7
	// binary: i32.and
	s4i32 = s4i32 & s5i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// tee_local
	l0 = s2i32
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l1 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1036
	// const
	s1i32 = 1036
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// get_local
	s2i32 = l6
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l2 = s1i32
	// get_local
	s2i32 = l0
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l0 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l0
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 40
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 1052
	// const
	s1i32 = 1512
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// br
	goto lbl59
	// end_block
lbl60:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1040
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l4 = s1i32
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1040
		// get_local
		s1i32 = l1
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// get_local
		s0i32 = l1
		// set_local
		l4 = s0i32
		// end_if
	}
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l6
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// const
	s0i32 = 1472
	// set_local
	l0 = s0i32
	// block
	// block
	// block
	// block
	// block
	// block
	// loop
lbl73:
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// binary: i32.ne
	if s0i32 != s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l0
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
		// tee_local
		l0 = s0i32
		// br_if
		if s0i32 != 0 {
			goto lbl73
		}
		// br
		goto lbl72
		// end_if
	}
	// end
	// get_local
	s0i32 = l0
	// load: i32.load8_u
	s0i32 = int32(ctx.Mem[int(s0i32+12)])
	// const
	s1i32 = 8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl71
	}
	// end_block
lbl72:
	// const
	s0i32 = 1472
	// set_local
	l0 = s0i32
	// loop
lbl75:
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l2 = s0i32
	// get_local
	s1i32 = l3
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l2
		// get_local
		s1i32 = l0
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l4 = s0i32
		// get_local
		s1i32 = l3
		// binary: i32.gt_u
		if uint32(s0i32) > uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl70
		}
		// end_if
	}
	// get_local
	s0i32 = l0
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// set_local
	l0 = s0i32
	// br
	goto lbl75
	// unreachable
	panic("unreachable executed")
	// end
	// unreachable
	panic("unreachable executed")
	// end_block
lbl71:
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l0
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// get_local
	s2i32 = l6
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// const
	s1i32 = -8
	// get_local
	s2i32 = l1
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 7
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 0
	// get_local
	s3i32 = l1
	// const
	s4i32 = 8
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// const
	s4i32 = 7
	// binary: i32.and
	s3i32 = s3i32 & s4i32
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l9 = s0i32
	// get_local
	s1i32 = l5
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// const
	s1i32 = -8
	// get_local
	s2i32 = l2
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// const
	s2i32 = 7
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// const
	s2i32 = 0
	// get_local
	s3i32 = l2
	// const
	s4i32 = 8
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// const
	s4i32 = 7
	// binary: i32.and
	s3i32 = s3i32 & s4i32
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l1 = s0i32
	// get_local
	s1i32 = l9
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// get_local
	s1i32 = l5
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l5
	// get_local
	s1i32 = l9
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l7 = s0i32
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l3
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1048
		// get_local
		s1i32 = l7
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// const
		s0i32 = 1036
		// const
		s1i32 = 1036
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
		// get_local
		s2i32 = l0
		// binary: i32.add
		s1i32 = s1i32 + s2i32
		// tee_local
		l0 = s1i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// get_local
		s0i32 = l7
		// get_local
		s1i32 = l0
		// const
		s2i32 = 1
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
		// br
		goto lbl68
		// end_if
	}
	// get_local
	s0i32 = l1
	// const
	s1i32 = 1044
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1044
		// get_local
		s1i32 = l7
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// const
		s0i32 = 1032
		// const
		s1i32 = 1032
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
		// get_local
		s2i32 = l0
		// binary: i32.add
		s1i32 = s1i32 + s2i32
		// tee_local
		l0 = s1i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// get_local
		s0i32 = l7
		// get_local
		s1i32 = l0
		// const
		s2i32 = 1
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l7
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// get_local
		s1i32 = l0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// br
		goto lbl68
		// end_if
	}
	// get_local
	s0i32 = l1
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 3
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// const
	s1i32 = 1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l2
		// const
		s1i32 = -8
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// set_local
		l10 = s0i32
		// block
		// get_local
		s0i32 = l2
		// const
		s1i32 = 255
		// binary: i32.le_u
		if uint32(s0i32) <= uint32(s1i32) {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// get_local
			s0i32 = l1
			// load: i32.load
			s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
			// tee_local
			l3 = s0i32
			// get_local
			s1i32 = l2
			// const
			s2i32 = 3
			// binary: i32.shr_u
			s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
			// tee_local
			l4 = s1i32
			// const
			s2i32 = 3
			// binary: i32.shl
			s1i32 = s1i32 << (uint32(s2i32) & 31)
			// const
			s2i32 = 1064
			// binary: i32.add
			s1i32 = s1i32 + s2i32
			// binary: i32.ne
			if s0i32 != s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			// get_local
			s0i32 = l3
			// get_local
			s1i32 = l1
			// load: i32.load
			s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+12):]))
			// tee_local
			l2 = s1i32
			// binary: i32.eq
			if s0i32 == s1i32 {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			// if
			if s0i32 != 0 {
				// const
				s0i32 = 1024
				// const
				s1i32 = 1024
				// load: i32.load
				s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
				// const
				s2i32 = -2
				// get_local
				s3i32 = l4
				// binary: i32.rotl
				s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
				// binary: i32.and
				s1i32 = s1i32 & s2i32
				// store: i32.store
				binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
				// br
				goto lbl80
				// end_if
			}
			// get_local
			s0i32 = l3
			// get_local
			s1i32 = l2
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
			// get_local
			s0i32 = l2
			// get_local
			s1i32 = l3
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
			// br
			goto lbl80
			// end_if
		}
		// get_local
		s0i32 = l1
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+24):]))
		// set_local
		l8 = s0i32
		// block
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l1
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+12):]))
		// tee_local
		l6 = s1i32
		// binary: i32.ne
		if s0i32 != s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// get_local
			s0i32 = l4
			// get_local
			s1i32 = l1
			// load: i32.load
			s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+8):]))
			// tee_local
			l2 = s1i32
			// binary: i32.le_u
			if uint32(s0i32) <= uint32(s1i32) {
				s0i32 = 1
			} else {
				s0i32 = 0
			}
			// if
			if s0i32 != 0 {
				// get_local
				s0i32 = l2
				// load: i32.load
				s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+12):]))
				// end_if
			}
			// get_local
			s0i32 = l2
			// get_local
			s1i32 = l6
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
			// get_local
			s0i32 = l6
			// get_local
			s1i32 = l2
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
			// br
			goto lbl83
			// end_if
		}
		// block
		// get_local
		s0i32 = l1
		// const
		s1i32 = 20
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l3 = s0i32
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l5 = s0i32
		// br_if
		if s0i32 != 0 {
			goto lbl86
		}
		// get_local
		s0i32 = l1
		// const
		s1i32 = 16
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l3 = s0i32
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l5 = s0i32
		// br_if
		if s0i32 != 0 {
			goto lbl86
		}
		// const
		s0i32 = 0
		// set_local
		l6 = s0i32
		// br
		goto lbl83
		// end_block
	lbl86:
		// loop
	lbl87:
		// get_local
		s0i32 = l3
		// set_local
		l2 = s0i32
		// get_local
		s0i32 = l5
		// tee_local
		l6 = s0i32
		// const
		s1i32 = 20
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l3 = s0i32
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l5 = s0i32
		// br_if
		if s0i32 != 0 {
			goto lbl87
		}
		// get_local
		s0i32 = l6
		// const
		s1i32 = 16
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l3 = s0i32
		// get_local
		s0i32 = l6
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
		// tee_local
		l5 = s0i32
		// br_if
		if s0i32 != 0 {
			goto lbl87
		}
		// end
		// get_local
		s0i32 = l2
		// const
		s1i32 = 0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// end_block
	lbl83:
		// get_local
		s0i32 = l8
		// unary: i32.eqz
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl80
		}
		// block
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l1
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+28):]))
		// tee_local
		l2 = s1i32
		// const
		s2i32 = 2
		// binary: i32.shl
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		// const
		s2i32 = 1328
		// binary: i32.add
		s1i32 = s1i32 + s2i32
		// tee_local
		l3 = s1i32
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
		// binary: i32.eq
		if s0i32 == s1i32 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// get_local
			s0i32 = l3
			// get_local
			s1i32 = l6
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
			// get_local
			s0i32 = l6
			// br_if
			if s0i32 != 0 {
				goto lbl88
			}
			// const
			s0i32 = 1028
			// const
			s1i32 = 1028
			// load: i32.load
			s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
			// const
			s2i32 = -2
			// get_local
			s3i32 = l2
			// binary: i32.rotl
			s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
			// binary: i32.and
			s1i32 = s1i32 & s2i32
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
			// br
			goto lbl80
			// end_if
		}
		// get_local
		s0i32 = l8
		// const
		s1i32 = 16
		// const
		s2i32 = 20
		// get_local
		s3i32 = l8
		// load: i32.load
		s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+16):]))
		// get_local
		s4i32 = l1
		// binary: i32.eq
		if s3i32 == s4i32 {
			s3i32 = 1
		} else {
			s3i32 = 0
		}
		// select
		if s3i32 != 0 {
			// s1i32 = s1i32
		} else {
			s1i32 = s2i32
		}
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// get_local
		s1i32 = l6
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// get_local
		s0i32 = l6
		// unary: i32.eqz
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl80
		}
		// end_block
	lbl88:
		// get_local
		s0i32 = l6
		// get_local
		s1i32 = l8
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
		// get_local
		s0i32 = l1
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
		// tee_local
		l2 = s0i32
		// if
		if s0i32 != 0 {
			// get_local
			s0i32 = l6
			// get_local
			s1i32 = l2
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
			// get_local
			s0i32 = l2
			// get_local
			s1i32 = l6
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
			// end_if
		}
		// get_local
		s0i32 = l1
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
		// tee_local
		l2 = s0i32
		// unary: i32.eqz
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// br_if
		if s0i32 != 0 {
			goto lbl80
		}
		// get_local
		s0i32 = l6
		// get_local
		s1i32 = l2
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
		// get_local
		s0i32 = l2
		// get_local
		s1i32 = l6
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
		// end_block
	lbl80:
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l10
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l1 = s0i32
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l10
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l0 = s0i32
		// end_if
	}
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l1
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// const
	s2i32 = -2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l0
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l7
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 255
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l0
		// const
		s1i32 = 3
		// binary: i32.shr_u
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		// tee_local
		l1 = s0i32
		// const
		s1i32 = 3
		// binary: i32.shl
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		// const
		s1i32 = 1064
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l0 = s0i32
		// block
		// const
		s0i32 = 1024
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l2 = s0i32
		// const
		s1i32 = 1
		// get_local
		s2i32 = l1
		// binary: i32.shl
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		// tee_local
		l1 = s1i32
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// unary: i32.eqz
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// const
			s0i32 = 1024
			// get_local
			s1i32 = l1
			// get_local
			s2i32 = l2
			// binary: i32.or
			s1i32 = s1i32 | s2i32
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
			// get_local
			s0i32 = l0
			// br
			goto lbl92
			// end_if
		}
		// get_local
		s0i32 = l0
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
		// end_block
	lbl92:
		// set_local
		l1 = s0i32
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l7
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l7
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
		// get_local
		s0i32 = l7
		// get_local
		s1i32 = l0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
		// get_local
		s0i32 = l7
		// get_local
		s1i32 = l1
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
		// br
		goto lbl68
		// end_if
	}
	// get_local
	s0i32 = l7
	// block
	// const
	s1i32 = 0
	// get_local
	s2i32 = l0
	// const
	s3i32 = 8
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// tee_local
	l1 = s2i32
	// unary: i32.eqz
	if s2i32 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// br_if
	if s2i32 != 0 {
		goto lbl94
	}
	// const
	s1i32 = 31
	// get_local
	s2i32 = l0
	// const
	s3i32 = 16777215
	// binary: i32.gt_u
	if uint32(s2i32) > uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// br_if
	if s2i32 != 0 {
		goto lbl94
	}
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l1
	// const
	s3i32 = 1048320
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = 8
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l1 = s2i32
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l2 = s1i32
	// get_local
	s2i32 = l2
	// const
	s3i32 = 520192
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = 4
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l2 = s2i32
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l3 = s1i32
	// get_local
	s2i32 = l3
	// const
	s3i32 = 245760
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = 2
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l3 = s2i32
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 15
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// get_local
	s2i32 = l1
	// get_local
	s3i32 = l2
	// binary: i32.or
	s2i32 = s2i32 | s3i32
	// get_local
	s3i32 = l3
	// binary: i32.or
	s2i32 = s2i32 | s3i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l1 = s1i32
	// const
	s2i32 = 1
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l0
	// get_local
	s3i32 = l1
	// const
	s4i32 = 21
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = 1
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// const
	s2i32 = 28
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// end_block
lbl94:
	// tee_local
	l1 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l1
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 1328
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l2 = s0i32
	// block
	// const
	s0i32 = 1028
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l3 = s0i32
	// const
	s1i32 = 1
	// get_local
	s2i32 = l1
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l4 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1028
		// get_local
		s1i32 = l3
		// get_local
		s2i32 = l4
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// get_local
		s0i32 = l2
		// get_local
		s1i32 = l7
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// br
		goto lbl95
		// end_if
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 0
	// const
	s2i32 = 25
	// get_local
	s3i32 = l1
	// const
	s4i32 = 1
	// binary: i32.shr_u
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// get_local
	s3i32 = l1
	// const
	s4i32 = 31
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l1 = s0i32
	// loop
lbl97:
	// get_local
	s0i32 = l1
	// tee_local
	l2 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l0
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl69
	}
	// get_local
	s0i32 = l3
	// const
	s1i32 = 29
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l3
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l3 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l1
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl97
	}
	// end
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l7
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// end_block
lbl95:
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l7
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l7
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// br
	goto lbl68
	// end_block
lbl70:
	// const
	s0i32 = 1036
	// get_local
	s1i32 = l6
	// const
	s2i32 = -40
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l0 = s1i32
	// const
	s2i32 = -8
	// get_local
	s3i32 = l1
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// const
	s3i32 = 7
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// const
	s3i32 = 0
	// get_local
	s4i32 = l1
	// const
	s5i32 = 8
	// binary: i32.add
	s4i32 = s4i32 + s5i32
	// const
	s5i32 = 7
	// binary: i32.and
	s4i32 = s4i32 & s5i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// tee_local
	l2 = s2i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l7 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1048
	// get_local
	s1i32 = l1
	// get_local
	s2i32 = l2
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l2 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l7
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// const
	s1i32 = 40
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// const
	s0i32 = 1052
	// const
	s1i32 = 1512
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l4
	// const
	s2i32 = 39
	// get_local
	s3i32 = l4
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// const
	s3i32 = 7
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// const
	s3i32 = 0
	// get_local
	s4i32 = l4
	// const
	s5i32 = -39
	// binary: i32.add
	s4i32 = s4i32 + s5i32
	// const
	s5i32 = 7
	// binary: i32.and
	s4i32 = s4i32 & s5i32
	// select
	if s4i32 != 0 {
		// s2i32 = s2i32
	} else {
		s2i32 = s3i32
	}
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// const
	s2i32 = -47
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l0 = s1i32
	// get_local
	s2i32 = l0
	// get_local
	s3i32 = l3
	// const
	s4i32 = 16
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// binary: i32.lt_u
	if uint32(s2i32) < uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// select
	if s2i32 != 0 {
		// s0i32 = s0i32
	} else {
		s0i32 = s1i32
	}
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 27
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1480
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l2
	// const
	s1i32 = 1472
	// load: i64.load
	s1i64 = int64(binary.LittleEndian.Uint64(ctx.Mem[int(s1i32+0):]))
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+8):], uint64(s1i64))
	// const
	s0i32 = 1480
	// get_local
	s1i32 = l2
	// const
	s2i32 = 8
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1476
	// get_local
	s1i32 = l6
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1472
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1484
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// const
	s1i32 = 24
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l0 = s0i32
	// loop
lbl98:
	// get_local
	s0i32 = l0
	// const
	s1i32 = 7
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l4
	// binary: i32.lt_u
	if uint32(s0i32) < uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl98
	}
	// end
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl59
	}
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l2
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
	// const
	s2i32 = -2
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// get_local
	s2i32 = l3
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l4 = s1i32
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// const
	s1i32 = 255
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l4
		// const
		s1i32 = 3
		// binary: i32.shr_u
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		// tee_local
		l1 = s0i32
		// const
		s1i32 = 3
		// binary: i32.shl
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		// const
		s1i32 = 1064
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l0 = s0i32
		// block
		// const
		s0i32 = 1024
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l2 = s0i32
		// const
		s1i32 = 1
		// get_local
		s2i32 = l1
		// binary: i32.shl
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		// tee_local
		l1 = s1i32
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// unary: i32.eqz
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// const
			s0i32 = 1024
			// get_local
			s1i32 = l1
			// get_local
			s2i32 = l2
			// binary: i32.or
			s1i32 = s1i32 | s2i32
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
			// get_local
			s0i32 = l0
			// br
			goto lbl100
			// end_if
		}
		// get_local
		s0i32 = l0
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
		// end_block
	lbl100:
		// set_local
		l1 = s0i32
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l3
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l3
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
		// get_local
		s0i32 = l3
		// get_local
		s1i32 = l0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
		// get_local
		s0i32 = l3
		// get_local
		s1i32 = l1
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
		// br
		goto lbl59
		// end_if
	}
	// get_local
	s0i32 = l3
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l3
	// block
	// const
	s1i32 = 0
	// get_local
	s2i32 = l4
	// const
	s3i32 = 8
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// tee_local
	l0 = s2i32
	// unary: i32.eqz
	if s2i32 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// br_if
	if s2i32 != 0 {
		goto lbl102
	}
	// const
	s1i32 = 31
	// get_local
	s2i32 = l4
	// const
	s3i32 = 16777215
	// binary: i32.gt_u
	if uint32(s2i32) > uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// br_if
	if s2i32 != 0 {
		goto lbl102
	}
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l0
	// const
	s3i32 = 1048320
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = 8
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l0 = s2i32
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l1 = s1i32
	// get_local
	s2i32 = l1
	// const
	s3i32 = 520192
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = 4
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l1 = s2i32
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l2 = s1i32
	// get_local
	s2i32 = l2
	// const
	s3i32 = 245760
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = 2
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l2 = s2i32
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 15
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// get_local
	s2i32 = l0
	// get_local
	s3i32 = l1
	// binary: i32.or
	s2i32 = s2i32 | s3i32
	// get_local
	s3i32 = l2
	// binary: i32.or
	s2i32 = s2i32 | s3i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l0 = s1i32
	// const
	s2i32 = 1
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l4
	// get_local
	s3i32 = l0
	// const
	s4i32 = 21
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = 1
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// const
	s2i32 = 28
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// end_block
lbl102:
	// tee_local
	l0 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 1328
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// block
	// const
	s0i32 = 1028
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l2 = s0i32
	// const
	s1i32 = 1
	// get_local
	s2i32 = l0
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l6 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1028
		// get_local
		s1i32 = l2
		// get_local
		s2i32 = l6
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l3
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// get_local
		s0i32 = l3
		// get_local
		s1i32 = l1
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
		// br
		goto lbl103
		// end_if
	}
	// get_local
	s0i32 = l4
	// const
	s1i32 = 0
	// const
	s2i32 = 25
	// get_local
	s3i32 = l0
	// const
	s4i32 = 1
	// binary: i32.shr_u
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// get_local
	s3i32 = l0
	// const
	s4i32 = 31
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l1
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l1 = s0i32
	// loop
lbl105:
	// get_local
	s0i32 = l1
	// tee_local
	l2 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l4
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl67
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 29
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// set_local
	l1 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l1
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l6 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l1 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl105
	}
	// end
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// end_block
lbl103:
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// br
	goto lbl59
	// end_block
lbl69:
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l7
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l7
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l7
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// end_block
lbl68:
	// get_local
	s0i32 = l9
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l0 = s0i32
	// br
	goto lbl0
	// end_block
lbl67:
	// get_local
	s0i32 = l2
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// end_block
lbl59:
	// const
	s0i32 = 1036
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l5
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl3
	}
	// const
	s0i32 = 1036
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l5
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l1 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1048
	// const
	s1i32 = 1048
	// load: i32.load
	s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+0):]))
	// tee_local
	l0 = s1i32
	// get_local
	s2i32 = l5
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// tee_local
	l2 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l1
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l5
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l0 = s0i32
	// br
	goto lbl0
	// end_block
lbl3:
	// const
	s0i32 = 1520
	// const
	s1i32 = 48
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 0
	// set_local
	l0 = s0i32
	// br
	goto lbl0
	// end_block
lbl2:
	// block
	// get_local
	s0i32 = l8
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl106
	}
	// block
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 1328
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l4
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l3
		// get_local
		s1i32 = l1
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// get_local
		s0i32 = l1
		// br_if
		if s0i32 != 0 {
			goto lbl107
		}
		// const
		s0i32 = 1028
		// get_local
		s1i32 = l7
		// const
		s2i32 = -2
		// get_local
		s3i32 = l0
		// binary: i32.rotl
		s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
		// binary: i32.and
		s1i32 = s1i32 & s2i32
		// tee_local
		l7 = s1i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// br
		goto lbl106
		// end_if
	}
	// get_local
	s0i32 = l8
	// const
	s1i32 = 16
	// const
	s2i32 = 20
	// get_local
	s3i32 = l8
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+16):]))
	// get_local
	s4i32 = l4
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl106
	}
	// end_block
lbl107:
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l8
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l0 = s0i32
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l1
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
		// end_if
	}
	// get_local
	s0i32 = l4
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// tee_local
	l0 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl106
	}
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// end_block
lbl106:
	// block
	// get_local
	s0i32 = l2
	// const
	s1i32 = 15
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l4
		// get_local
		s1i32 = l2
		// get_local
		s2i32 = l5
		// binary: i32.add
		s1i32 = s1i32 + s2i32
		// tee_local
		l0 = s1i32
		// const
		s2i32 = 3
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l4
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l0 = s0i32
		// get_local
		s1i32 = l0
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
		// const
		s2i32 = 1
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
		// br
		goto lbl110
		// end_if
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l5
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l3 = s0i32
	// get_local
	s1i32 = l2
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// get_local
	s1i32 = l3
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l2
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l2
	// const
	s1i32 = 255
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l2
		// const
		s1i32 = 3
		// binary: i32.shr_u
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		// tee_local
		l1 = s0i32
		// const
		s1i32 = 3
		// binary: i32.shl
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		// const
		s1i32 = 1064
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l0 = s0i32
		// block
		// const
		s0i32 = 1024
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// tee_local
		l2 = s0i32
		// const
		s1i32 = 1
		// get_local
		s2i32 = l1
		// binary: i32.shl
		s1i32 = s1i32 << (uint32(s2i32) & 31)
		// tee_local
		l1 = s1i32
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// unary: i32.eqz
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// const
			s0i32 = 1024
			// get_local
			s1i32 = l1
			// get_local
			s2i32 = l2
			// binary: i32.or
			s1i32 = s1i32 | s2i32
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
			// get_local
			s0i32 = l0
			// br
			goto lbl113
			// end_if
		}
		// get_local
		s0i32 = l0
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
		// end_block
	lbl113:
		// set_local
		l1 = s0i32
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l3
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l3
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
		// get_local
		s0i32 = l3
		// get_local
		s1i32 = l0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
		// get_local
		s0i32 = l3
		// get_local
		s1i32 = l1
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
		// br
		goto lbl110
		// end_if
	}
	// get_local
	s0i32 = l3
	// block
	// const
	s1i32 = 0
	// get_local
	s2i32 = l2
	// const
	s3i32 = 8
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// tee_local
	l0 = s2i32
	// unary: i32.eqz
	if s2i32 == 0 {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// br_if
	if s2i32 != 0 {
		goto lbl115
	}
	// const
	s1i32 = 31
	// get_local
	s2i32 = l2
	// const
	s3i32 = 16777215
	// binary: i32.gt_u
	if uint32(s2i32) > uint32(s3i32) {
		s2i32 = 1
	} else {
		s2i32 = 0
	}
	// br_if
	if s2i32 != 0 {
		goto lbl115
	}
	// get_local
	s1i32 = l0
	// get_local
	s2i32 = l0
	// const
	s3i32 = 1048320
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = 8
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l0 = s2i32
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l1 = s1i32
	// get_local
	s2i32 = l1
	// const
	s3i32 = 520192
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = 4
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l1 = s2i32
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l5 = s1i32
	// get_local
	s2i32 = l5
	// const
	s3i32 = 245760
	// binary: i32.add
	s2i32 = s2i32 + s3i32
	// const
	s3i32 = 16
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = 2
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// tee_local
	l5 = s2i32
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// const
	s2i32 = 15
	// binary: i32.shr_u
	s1i32 = int32(uint32(s1i32) >> (uint32(s2i32) & 31))
	// get_local
	s2i32 = l0
	// get_local
	s3i32 = l1
	// binary: i32.or
	s2i32 = s2i32 | s3i32
	// get_local
	s3i32 = l5
	// binary: i32.or
	s2i32 = s2i32 | s3i32
	// binary: i32.sub
	s1i32 = s1i32 - s2i32
	// tee_local
	l0 = s1i32
	// const
	s2i32 = 1
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// get_local
	s2i32 = l2
	// get_local
	s3i32 = l0
	// const
	s4i32 = 21
	// binary: i32.add
	s3i32 = s3i32 + s4i32
	// binary: i32.shr_u
	s2i32 = int32(uint32(s2i32) >> (uint32(s3i32) & 31))
	// const
	s3i32 = 1
	// binary: i32.and
	s2i32 = s2i32 & s3i32
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// const
	s2i32 = 28
	// binary: i32.add
	s1i32 = s1i32 + s2i32
	// end_block
lbl115:
	// tee_local
	l0 = s1i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+28):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i64 = 0
	// store: i64.store
	binary.LittleEndian.PutUint64(ctx.Mem[int(s0i32+16):], uint64(s1i64))
	// get_local
	s0i32 = l0
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 1328
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l1 = s0i32
	// block
	// block
	// get_local
	s0i32 = l7
	// const
	s1i32 = 1
	// get_local
	s2i32 = l0
	// binary: i32.shl
	s1i32 = s1i32 << (uint32(s2i32) & 31)
	// tee_local
	l5 = s1i32
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// const
		s0i32 = 1028
		// get_local
		s1i32 = l5
		// get_local
		s2i32 = l7
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l3
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// br
		goto lbl117
		// end_if
	}
	// get_local
	s0i32 = l2
	// const
	s1i32 = 0
	// const
	s2i32 = 25
	// get_local
	s3i32 = l0
	// const
	s4i32 = 1
	// binary: i32.shr_u
	s3i32 = int32(uint32(s3i32) >> (uint32(s4i32) & 31))
	// binary: i32.sub
	s2i32 = s2i32 - s3i32
	// get_local
	s3i32 = l0
	// const
	s4i32 = 31
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l1
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// set_local
	l5 = s0i32
	// loop
lbl119:
	// get_local
	s0i32 = l5
	// tee_local
	l1 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+4):]))
	// const
	s1i32 = -8
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// get_local
	s1i32 = l2
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl116
	}
	// get_local
	s0i32 = l0
	// const
	s1i32 = 29
	// binary: i32.shr_u
	s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
	// set_local
	l5 = s0i32
	// get_local
	s0i32 = l0
	// const
	s1i32 = 1
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// set_local
	l0 = s0i32
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// const
	s2i32 = 4
	// binary: i32.and
	s1i32 = s1i32 & s2i32
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l6 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l5 = s0i32
	// br_if
	if s0i32 != 0 {
		goto lbl119
	}
	// end
	// get_local
	s0i32 = l6
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
	// end_block
lbl117:
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// br
	goto lbl110
	// end_block
lbl116:
	// get_local
	s0i32 = l1
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
	// tee_local
	l0 = s0i32
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// const
	s1i32 = 0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l1
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
	// end_block
lbl110:
	// get_local
	s0i32 = l4
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l0 = s0i32
	// br
	goto lbl0
	// end_block
lbl1:
	// block
	// get_local
	s0i32 = l9
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl120
	}
	// block
	// get_local
	s0i32 = l1
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+28):]))
	// tee_local
	l0 = s0i32
	// const
	s1i32 = 2
	// binary: i32.shl
	s0i32 = s0i32 << (uint32(s1i32) & 31)
	// const
	s1i32 = 1328
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l2 = s0i32
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
	// get_local
	s1i32 = l1
	// binary: i32.eq
	if s0i32 == s1i32 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l2
		// get_local
		s1i32 = l4
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// get_local
		s0i32 = l4
		// br_if
		if s0i32 != 0 {
			goto lbl121
		}
		// const
		s0i32 = 1028
		// get_local
		s1i32 = l10
		// const
		s2i32 = -2
		// get_local
		s3i32 = l0
		// binary: i32.rotl
		s2i32 = int32(bits.RotateLeft32(uint32(s2i32), int(s3i32)))
		// binary: i32.and
		s1i32 = s1i32 & s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
		// br
		goto lbl120
		// end_if
	}
	// get_local
	s0i32 = l9
	// const
	s1i32 = 16
	// const
	s2i32 = 20
	// get_local
	s3i32 = l9
	// load: i32.load
	s3i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s3i32+16):]))
	// get_local
	s4i32 = l1
	// binary: i32.eq
	if s3i32 == s4i32 {
		s3i32 = 1
	} else {
		s3i32 = 0
	}
	// select
	if s3i32 != 0 {
		// s1i32 = s1i32
	} else {
		s1i32 = s2i32
	}
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l4
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl120
	}
	// end_block
lbl121:
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l9
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+16):]))
	// tee_local
	l0 = s0i32
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l4
		// get_local
		s1i32 = l0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+16):], uint32(s1i32))
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l4
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
		// end_if
	}
	// get_local
	s0i32 = l1
	// load: i32.load
	s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+20):]))
	// tee_local
	l0 = s0i32
	// unary: i32.eqz
	if s0i32 == 0 {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// br_if
	if s0i32 != 0 {
		goto lbl120
	}
	// get_local
	s0i32 = l4
	// get_local
	s1i32 = l0
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+20):], uint32(s1i32))
	// get_local
	s0i32 = l0
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+24):], uint32(s1i32))
	// end_block
lbl120:
	// block
	// get_local
	s0i32 = l3
	// const
	s1i32 = 15
	// binary: i32.le_u
	if uint32(s0i32) <= uint32(s1i32) {
		s0i32 = 1
	} else {
		s0i32 = 0
	}
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l1
		// get_local
		s1i32 = l3
		// get_local
		s2i32 = l5
		// binary: i32.add
		s1i32 = s1i32 + s2i32
		// tee_local
		l0 = s1i32
		// const
		s2i32 = 3
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l1
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// tee_local
		l0 = s0i32
		// get_local
		s1i32 = l0
		// load: i32.load
		s1i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s1i32+4):]))
		// const
		s2i32 = 1
		// binary: i32.or
		s1i32 = s1i32 | s2i32
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
		// br
		goto lbl124
		// end_if
	}
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// const
	s2i32 = 3
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l1
	// get_local
	s1i32 = l5
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// tee_local
	l4 = s0i32
	// get_local
	s1i32 = l3
	// const
	s2i32 = 1
	// binary: i32.or
	s1i32 = s1i32 | s2i32
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+4):], uint32(s1i32))
	// get_local
	s0i32 = l3
	// get_local
	s1i32 = l4
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// get_local
	s0i32 = l8
	// if
	if s0i32 != 0 {
		// get_local
		s0i32 = l8
		// const
		s1i32 = 3
		// binary: i32.shr_u
		s0i32 = int32(uint32(s0i32) >> (uint32(s1i32) & 31))
		// tee_local
		l5 = s0i32
		// const
		s1i32 = 3
		// binary: i32.shl
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		// const
		s1i32 = 1064
		// binary: i32.add
		s0i32 = s0i32 + s1i32
		// set_local
		l0 = s0i32
		// const
		s0i32 = 1044
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+0):]))
		// set_local
		l2 = s0i32
		// block
		// const
		s0i32 = 1
		// get_local
		s1i32 = l5
		// binary: i32.shl
		s0i32 = s0i32 << (uint32(s1i32) & 31)
		// tee_local
		l5 = s0i32
		// get_local
		s1i32 = l6
		// binary: i32.and
		s0i32 = s0i32 & s1i32
		// unary: i32.eqz
		if s0i32 == 0 {
			s0i32 = 1
		} else {
			s0i32 = 0
		}
		// if
		if s0i32 != 0 {
			// const
			s0i32 = 1024
			// get_local
			s1i32 = l5
			// get_local
			s2i32 = l6
			// binary: i32.or
			s1i32 = s1i32 | s2i32
			// store: i32.store
			binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
			// get_local
			s0i32 = l0
			// br
			goto lbl127
			// end_if
		}
		// get_local
		s0i32 = l0
		// load: i32.load
		s0i32 = int32(binary.LittleEndian.Uint32(ctx.Mem[int(s0i32+8):]))
		// end_block
	lbl127:
		// set_local
		l5 = s0i32
		// get_local
		s0i32 = l0
		// get_local
		s1i32 = l2
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
		// get_local
		s0i32 = l5
		// get_local
		s1i32 = l2
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
		// get_local
		s0i32 = l2
		// get_local
		s1i32 = l0
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+12):], uint32(s1i32))
		// get_local
		s0i32 = l2
		// get_local
		s1i32 = l5
		// store: i32.store
		binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+8):], uint32(s1i32))
		// end_if
	}
	// const
	s0i32 = 1044
	// get_local
	s1i32 = l4
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// const
	s0i32 = 1032
	// get_local
	s1i32 = l3
	// store: i32.store
	binary.LittleEndian.PutUint32(ctx.Mem[int(s0i32+0):], uint32(s1i32))
	// end_block
lbl124:
	// get_local
	s0i32 = l1
	// const
	s1i32 = 8
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_local
	l0 = s0i32
	// end_block
lbl0:
	// get_local
	s0i32 = l11
	// const
	s1i32 = 16
	// binary: i32.add
	s0i32 = s0i32 + s1i32
	// set_global
	ctx.g0 = s0i32
	// get_local
	s0i32 = l0
	// return
	return s0i32
}
