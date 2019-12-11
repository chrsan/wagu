package test

import (
	"encoding/binary"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestCopy(t *testing.T) {
	c := qt.New(t)
	ctx1 := NewContext()
	WasmCallCtors(ctx1)
	ptr1 := Int32New(ctx1)
	c.Assert(ptr1, qt.Not(qt.Equals), 0)
	Int32Set(ctx1, ptr1, 123)
	c.Assert(Int32Get(ctx1, ptr1), qt.Equals, int32(123))
	st := StackSave(ctx1)
	ptr2 := StackAlloc(ctx1, 4)
	ctx2 := ctx1.Copy()
	binary.LittleEndian.PutUint32(ctx1.Mem[ptr2:], uint32(123))
	binary.LittleEndian.PutUint32(ctx2.Mem[ptr2:], uint32(456))
	c.Assert(Int32Get(ctx1, ptr1), qt.Equals, int32(123))
	c.Assert(Int32Get(ctx2, ptr2), qt.Equals, int32(456))
	StackRestore(ctx1, st)
	StackRestore(ctx2, st)
	c.Assert(StackSave(ctx2), qt.Equals, StackSave(ctx1))
	Int32Set(ctx1, ptr1, 456)
	c.Assert(Int32Get(ctx1, ptr1), qt.Equals, int32(456))
	c.Assert(Int32Get(ctx2, ptr1), qt.Equals, int32(123))
	Free(ctx1, ptr1)
	Free(ctx2, ptr1)
}
