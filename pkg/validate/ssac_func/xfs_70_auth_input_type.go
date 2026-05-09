//ff:func feature=validate type=rule control=iteration dimension=2 topic=func-check
//ff:what XFS-70 — @auth input value type must be string-compatible (authz.CheckRequest fields are string)

package ssac_func

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfs70AuthInputType validates XFS-70: every @auth input value must resolve
// to a string-compatible Go type. authz.CheckRequest fields (ResourceID,
// etc.) are always string; passing a pgtype.UUID or int64 causes a build
// failure that validate should catch early.
func xfs70AuthInputType(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type != "auth" {
				continue
			}
			for inputKey, inputValue := range seq.Inputs {
				sourceType := resolveInputType(g, fn.Name, inputValue)
				if sourceType == "" {
					continue
				}
				if !TypesCompatible(sourceType, "string") {
					diags = append(diags, diagnostic.Diagnostic{
						File:  fn.FileName,
						Line:  seq.Line,
						Phase: diagnostic.PhaseValidate,
						Level: diagnostic.LevelError,
						Message: "[XFS-70] @auth input " + inputKey +
							" value type " + sourceType + " is not string-compatible" +
							" (authz.CheckRequest." + inputKey + " is string)",
						Advice: "Use a string-typed source (e.g. request.id for path params) " +
							"instead of a DB row UUID field",
					})
				}
			}
		}
	}
	return diags
}
