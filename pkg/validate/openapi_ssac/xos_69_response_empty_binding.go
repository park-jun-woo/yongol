//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ssac
//ff:what XOS-69 — @response binds 0 fields but OpenAPI 200 response schema has properties

package openapi_ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos69ResponseEmptyBinding warns when an SSaC `@response` (or `@response!`)
// binds zero fields while the OpenAPI 200 response schema declares properties.
// Empty binding causes codegen to return an empty struct, yielding `{}`.
func xos69ResponseEmptyBinding(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	prefix := "OpenAPI.response."
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			continue
		}
		fnPrefix := prefix + fn.Name + "."
		for _, seq := range fn.Sequences {
			if seq.Type != "response" {
				continue
			}
			if len(seq.Fields) != 0 {
				continue
			}
			// Check whether OpenAPI 200 response schema has any properties.
			hasProps := false
			for k := range g.Types {
				if strings.HasPrefix(k, fnPrefix) {
					hasProps = true
					break
				}
			}
			if !hasProps {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XOS-69] %s — @response binds 0 fields but OpenAPI 200 schema has properties", fn.Name),
				Advice:  "Bind response fields or remove properties from the OpenAPI 200 response schema",
			})
		}
	}
	return diags
}
