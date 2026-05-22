//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-openapi
//ff:what XOS-66 — verifies that request fields referenced in SSaC are listed in the OpenAPI requestBody required array

package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xos66UsedFieldsRequired validates XOS-66: every SSaC-referenced
// `request.<field>` is declared in the OpenAPI requestBody's required list.
func xos66UsedFieldsRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		rs, ok := fs.RequestConstraints[fn.Name]
		if !ok {
			continue
		}
		for field := range collectRequestFields(fn) {
			fc, exists := rs[field]
			if !exists || fc.Required {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:        fn.FileName,
				Line:        fn.Line,
				Phase:       diagnostic.PhaseValidate,
				Level:       diagnostic.LevelError,
				Message:     "[XOS-66] field " + field + " is used in SSaC " + fn.Name + " but not marked required in OpenAPI requestBody",
				Advice:      "Add field " + field + " used by SSaC to the required array of the OpenAPI requestBody",
				OperationID: fn.Name,
			})
		}
	}
	return diags
}
