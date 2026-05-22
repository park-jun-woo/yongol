//ff:func feature=validate type=rule control=iteration dimension=2 topic=func-check
//ff:what XFS-42 — @call Inputs count → FuncRequest fields

package ssac_func

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfs42CallInputsCount validates XFS-42: @call Inputs count → FuncRequest fields
func xfs42CallInputsCount(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "call" || seq.Model == "" {
				continue
			}
			funcName := callFuncName(seq.Model)
			camelName := callFuncCamelName(seq.Model)
			if funcName == "" || camelName == "" {
				continue
			}
			reqFields, ok := g.Schemas["Func.request."+camelName]
			if !ok {
				continue
			}
			if len(seq.Inputs) != len(reqFields) {
				diags = append(diags, diagnostic.Diagnostic{
					File:  fn.FileName,
					Line:  seq.Line,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: "[XFS-42] @call " + seq.Model + " inputs " + itoa(len(seq.Inputs)) +
						" ≠ request fields " + itoa(len(reqFields)),
					Advice:      "Match the @call input count to the number of fields in the func Request struct",
					OperationID: fn.Name,
				})
			}
		}
	}
	return diags
}
