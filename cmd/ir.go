package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/chrsan/wagu/ir"
	"github.com/go-interpreter/wagon/wasm"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
)

var (
	irOut     string
	irGlobals string
	irCmd     = &cobra.Command{
		Use:   "ir FILE.wasm",
		Short: "Convert a WebAssembly file to IR",
		Args:  cobra.ExactArgs(1),
		RunE:  genIR,
	}
)

func init() {
	irCmd.Flags().StringVarP(&irOut, "output", "o", "wagu.ir", "output file")
	irCmd.Flags().StringVarP(&irGlobals, "globals", "g", "", "globals JSON file")
}

func genIR(cmd *cobra.Command, args []string) error {
	f, err := os.Open(args[0])
	if err != nil {
		return err
	}
	defer f.Close()
	wm, err := wasm.ReadModule(f, nil)
	if err != nil {
		return err
	}
	ig, err := resolveGlobals(wm)
	if err != nil {
		return err
	}
	m, err := ir.ReadModule(wm, ig)
	if err != nil {
		return err
	}
	b, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(irOut, b, 0644)
}

func resolveGlobals(wm *wasm.Module) (ir.ImportedGlobals, error) {
	if irGlobals == "" {
		return nil, nil
	}
	data, err := ioutil.ReadFile(irGlobals)
	if err != nil {
		return nil, err
	}
	gm := map[string]map[string]json.Number{}
	if err := json.Unmarshal(data, &gm); err != nil {
		return nil, err
	}
	if wm.Import == nil {
		return nil, nil
	}
	var ig ir.ImportedGlobals
	for mn, fm := range gm {
		for fn, n := range fm {
			for _, e := range wm.Import.Entries {
				if e.ModuleName != mn || e.FieldName != fn || e.Type.Kind() != wasm.ExternalGlobal {
					continue
				}
				if ig == nil {
					ig = ir.ImportedGlobals{}
				}
				g := e.Type.(wasm.GlobalVarImport)
				if _, ok := ig[mn]; !ok {
					ig[mn] = map[string]*ir.Value{}
				}
				var v *ir.Value
				switch g.Type.Type {
				case wasm.ValueTypeI32:
					i, err := n.Int64()
					if err != nil {
						return nil, err
					}
					v = &ir.Value{
						Kind: &ir.Value_I32{
							I32: int32(i),
						},
					}
				case wasm.ValueTypeI64:
					i, err := n.Int64()
					if err != nil {
						return nil, err
					}
					v = &ir.Value{
						Kind: &ir.Value_I64{
							I64: i,
						},
					}
				case wasm.ValueTypeF32:
					f, err := n.Float64()
					if err != nil {
						return nil, err
					}
					v = &ir.Value{
						Kind: &ir.Value_F32{
							F32: float32(f),
						},
					}
				case wasm.ValueTypeF64:
					f, err := n.Float64()
					if err != nil {
						return nil, err
					}
					v = &ir.Value{
						Kind: &ir.Value_F64{
							F64: f,
						},
					}
				}
				ig[mn][fn] = v
			}
		}
	}
	return ig, nil
}
