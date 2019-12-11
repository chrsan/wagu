package test

import (
	"errors"
	"math"
)

type Context struct {
	Mem []byte
	g0  int32
}

func NewContext() *Context {
	c := &Context{Mem: newMemory()}
	c.g0 = 5244416
	return c
}

func (c *Context) Copy() *Context {
	d := *c
	d.Mem = make([]byte, len(c.Mem))
	copy(d.Mem, c.Mem)
	return &d
}

func (c *Context) growMem(size int) {
	c.Mem = append(c.Mem, make([]byte, size)...)
}

var (
	ErrDivByZero   = errors.New("div by zero")
	ErrIntOverflow = errors.New("int overflow")
)

func i32DivS(x, y int32) int32 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	if x == math.MinInt32 && y == -1 {
		panic(ErrIntOverflow)
	}
	return x / y
}

func i64DivS(x, y int64) int64 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	if x == math.MinInt64 && y == -1 {
		panic(ErrIntOverflow)
	}
	return x / y
}

func i32RemS(x, y int32) int32 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	if x == math.MinInt32 && y == -1 {
		return 0
	}
	return x % y
}

func i64RemS(x, y int64) int64 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	if x == math.MinInt64 && y == -1 {
		return 0
	}
	return x % y
}

func i32DivU(x, y int32) int32 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	return int32(uint32(x) / uint32(y))
}

func i64DivU(x, y int64) int64 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	return int64(uint64(x) / uint64(y))
}

func i32RemU(x, y int32) int32 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	return int32(uint32(x) % uint32(y))
}

func i64RemU(x, y int64) int64 {
	if y == 0 {
		panic(ErrDivByZero)
	}
	return int64(uint64(x) % uint64(y))
}
