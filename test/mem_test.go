package test

func newMemory() []byte {
	mem := make([]byte, 16777216)
	copy(mem[1536:], []byte{0xa0, 0x6, 0x50})
	return mem
}
