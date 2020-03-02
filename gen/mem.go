package gen

import (
	"bytes"
	"text/template"

	"github.com/chrsan/wagu/ir"
)

// Mem generates source for the mem file.
func Mem(pkg string, mem *ir.Memory, mmap bool) ([]byte, error) {
	asset, err := lookupTemplate("mem")
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New("mem").Parse(string(asset))
	if err != nil {
		return nil, err
	}
	var (
		maxOff       uint32
		sizeAtMaxOff int
		segmentBytes int
	)
	if len(mem.Segments) != 0 {
		first := true
		for off, b := range mem.Segments {
			segmentBytes += len(b)
			if first {
				maxOff = off
				sizeAtMaxOff = len(b)
				first = false
				continue
			}
			if off > maxOff {
				maxOff = off
				sizeAtMaxOff = len(b)
			}
		}
	}
	data := struct {
		Pkg          string
		Mem          *ir.Memory
		MMap         bool
		MaxOff       uint32
		SizeAtMaxOff int
		SegmentBytes int
	}{
		Pkg:          pkg,
		Mem:          mem,
		MMap:         mmap,
		MaxOff:       maxOff,
		SizeAtMaxOff: sizeAtMaxOff,
		SegmentBytes: segmentBytes,
	}
	var b bytes.Buffer
	if err := tmpl.Execute(&b, data); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
