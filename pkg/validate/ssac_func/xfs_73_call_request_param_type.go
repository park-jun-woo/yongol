//ff:func feature=validate type=rule control=iteration dimension=3 topic=func-check
//ff:what XFS-73 — @call input request.* OpenAPI param type must match Func Request field type

package ssac_func

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfs73CallRequestParamType validates XFS-73: when a @call input references
// request.<field>, the OpenAPI parameter's Go type (registered in
// Ground.Types["OpenAPI.paramType.<opID>.<field>"]) must be compatible with
// the target Func Request field type. Body fields are not registered as
// paramType entries and are skipped.
func xfs73CallRequestParamType(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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
			if funcName == "" {
				continue
			}
			for inputKey, inputValue := range seq.Inputs {
				if !strings.HasPrefix(inputValue, "request.") {
					continue
				}
				field := strings.TrimPrefix(inputValue, "request.")
				// OpenAPI param Go type lookup (path/query params only).
				paramType := g.Types["OpenAPI.paramType."+fn.Name+"."+field]
				if paramType == "" {
					continue
				}
				// Func Request field type lookup.
				reqType := g.Types["Func.request."+toCamelKey(funcName)+"."+inputKey]
				if reqType == "" {
					continue
				}
				if !TypesCompatible(paramType, reqType) {
					diags = append(diags, diagnostic.Diagnostic{
						File:  fn.FileName,
						Line:  seq.Line,
						Phase: diagnostic.PhaseValidate,
						Level: diagnostic.LevelError,
						Message: "[XFS-73] @call " + seq.Model + " input " + inputKey +
							": request." + field + " type " + paramType +
							" ≠ " + funcName + "Request." + inputKey + " type " + reqType,
						Advice: "Make the Func Request field " + inputKey +
							" type match the OpenAPI param type (" + paramType + ")",
					})
				}
			}
		}
	}
	return diags
}
