//ff:func feature=validate type=rule control=iteration dimension=3 topic=func-check
//ff:what XFS-44 — @call Input type ↔ Request field type (literal and variable inference)

package ssac_func

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfs44CallInputType validates XFS-44: @call Input type ↔ Request field type.
// Each @call input "key: expression" is compared against
// Types["Func.request.<funcName>.<key>"]. Two resolution paths:
//   - literal (quoted string, numeric, bool, nil) → inferLiteralType
//   - bare variable → Types["SSaC.var.<funcName>.<var>"]
// Field access (var.Field) remains deferred until a variable-symbol populator
// registers field types.
func xfs44CallInputType(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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
			lookupKey := toCamelKey(funcName) // Ground.Types keys @func annotation (camelCase)
			for inputKey, inputValue := range seq.Inputs {
				reqType, ok := g.Types["Func.request."+lookupKey+"."+inputKey]
				if !ok {
					continue
				}
				sourceType := resolveInputType(g, fn.Name, inputValue)
				if sourceType == "" {
					continue
				}
				if !TypesCompatible(sourceType, reqType) {
					diags = append(diags, diagnostic.Diagnostic{
						File:  fn.FileName,
						Line:  seq.Line,
						Phase: diagnostic.PhaseValidate,
						Level: diagnostic.LevelError,
						Message: "[XFS-44] @call " + seq.Model + " input " + inputKey +
							" type " + sourceType + " ≠ " + funcName + "Request." + inputKey + " type " + reqType,
						Advice:      "Make the @call input " + inputKey + " type match the " + inputKey + " field type in the func Request",
						OperationID: fn.Name,
					})
				}
			}
		}
	}
	return diags
}
