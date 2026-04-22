//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-57 — @call input types must match the FuncRequest field types

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s57FuncRequestType validates S-57: each @call argument's source-variable type
// must match the corresponding FuncRequest field type registered in Ground.
func s57FuncRequestType(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "call" {
				continue
			}
			method := extractMethod(seq)
			if method == "" {
				continue
			}
			lookupKey := "Func.request." + method
			if _, ok := g.Schemas[lookupKey]; !ok {
				continue
			}
			for _, arg := range seq.Args {
				if arg.Source == "" || arg.Field == "" {
					continue
				}
				sourceType := g.Types["SSaC.var."+fn.Name+"."+arg.Source]
				if sourceType == "" {
					continue
				}
				expected := g.Types[lookupKey+"."+arg.Field]
				if expected == "" || expected == sourceType {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[S-57] @call %s input %s type mismatch: %s vs %s", method, arg.Field, sourceType, expected),
					Advice:  "Align the Request type in the func spec with the SSaC input type",
				})
			}
		}
	}
	return diags
}
