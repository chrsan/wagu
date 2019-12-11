package test

func f9(ctx *Context, l0 int32) int32 {
	var s0i32 int32
	_ = s0i32
	var s1i32 int32
	_ = s1i32
	// get_global
	s0i32 = ctx.g0
	// get_local
	s1i32 = l0
	// binary: i32.sub
	s0i32 = s0i32 - s1i32
	// const
	s1i32 = -16
	// binary: i32.and
	s0i32 = s0i32 & s1i32
	// tee_local
	l0 = s0i32
	// set_global
	ctx.g0 = s0i32
	// get_local
	s0i32 = l0
	// return
	return s0i32
}
