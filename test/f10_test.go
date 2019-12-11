package test

func f10(ctx *Context) int32 {
	var s0i32 int32
	_ = s0i32
	// const
	s0i32 = 4
	// call
	s0i32 = f2(ctx, s0i32)
	// return
	return s0i32
}
