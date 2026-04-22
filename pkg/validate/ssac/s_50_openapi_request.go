//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-50 — every request.field reference exists in the OpenAPI request schema

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s50OpenAPIRequest validates S-50: every request.<field> referenced in Args
// must be a declared field on the OpenAPI request schema for the operation.
//
// Relationship: XOS-66 (openapi_ssac/) is a non-symmetric stricter constraint — it also
// requires that referenced fields appear in the requestBody.required array. S-50 only checks
// for existence in schema properties; XOS-66 handles the required-list check separately.
// When both violations apply (the field is absent from the schema itself), S-50 raises the
// ERROR first and XOS-66 catches the required-only gap as a separate signal.
func s50OpenAPIRequest(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			continue
		}
		fields, ok := g.Lookup["OpenAPI.request."+fn.Name]
		if !ok {
			continue
		}
		for _, seq := range fn.Sequences {
			for _, arg := range seq.Args {
				if arg.Source != "request" || arg.Field == "" {
					continue
				}
				if fields[arg.Field] {
					continue
				}
				diags = append(diags, diagnostic.Diagnostic{
					File:    fn.FileName,
					Line:    seq.Line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[S-50] request.%s not in OpenAPI request schema", arg.Field),
					Advice:  fmt.Sprintf("Add field %q to the requestBody of OpenAPI operationId %q", arg.Field, fn.Name),
				})
			}
		}
	}
	return diags
}
