//ff:func feature=validate type=rule control=iteration dimension=3 topic=openapi-ssac
//ff:what XOS-67 — the value type of each @response field must be compatible with the expected type in the OpenAPI response schema

package openapi_ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate/ssac_func"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos67ResponseFieldType validates that each `@response { key: value }` pair
// has a value whose type is compatible with the OpenAPI response schema's
// declared type for that key.
//
// Lookup (populators: Phase005/007):
//
//	expected = Types["OpenAPI.response.<funcName>.<key>"]
//	actual   = inferResponseValueType(g, funcName, value)
//
// Mismatch → ERROR. Unresolvable (expected=="" or actual=="") → skip.
func xos67ResponseFieldType(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			continue
		}
		for seqIdx, seq := range fn.Sequences {
			if seq.Type != "response" {
				continue
			}
			for key, value := range seq.Fields {
				expected := g.Types["OpenAPI.response."+fn.Name+"."+key]
				if expected == "" {
					continue
				}
				actual := inferResponseValueType(g, fn.Name, value)
				if actual == "" {
					continue
				}
				// expected is now the format-aware oapi-codegen type
				// (registerOpenAPIResponseProps): a { type: string,
				// format: date-time } field resolves to time.Time and
				// { format: uuid } to openapi_types.UUID, so a DDL
				// TIMESTAMPTZ (time.Time) or UUID column (openapi_types.UUID)
				// matches directly through TypesCompatible — no special-case
				// format marker allowance is needed.
				if ssac_func.TypesCompatible(actual, expected) {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:  fn.FileName,
					Line:  seq.Line,
					Phase: diagnostic.PhaseValidate,
					Level: diagnostic.LevelError,
					Message: fmt.Sprintf("[XOS-67] %s seq[%d] — @response field %q = %q: type %s ≠ OpenAPI response %q expected type %s",
						fn.Name, seqIdx, key, value, actual, key, expected),
					Advice: "Correct the value type so that it is compatible with the OpenAPI response schema's generated Go type. " +
						"A { type: string, format: uuid } field is generated as openapi_types.UUID (not string or pgtype.UUID); " +
						"{ format: date-time } as time.Time. Bind a value whose type matches, or fix the field's OpenAPI format.",
					OperationID: fn.Name,
				})
			}
		}
	}
	return diags
}
