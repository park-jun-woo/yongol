//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-structural
//ff:what XOO-72 — email-type fields are missing format: email

package openapi

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoo72EmailNoFormat flags request-body email fields that do not declare
// format: email. Emitted as WARNING.
func xoo72EmailNoFormat(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.RequestConstraints) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	opIDs := make([]string, 0, len(fs.RequestConstraints))
	for opID := range fs.RequestConstraints {
		opIDs = append(opIDs, opID)
	}
	sort.Strings(opIDs)
	for _, opID := range opIDs {
		fields := fs.RequestConstraints[opID]
		fieldNames := make([]string, 0, len(fields))
		for name := range fields {
			fieldNames = append(fieldNames, name)
		}
		sort.Strings(fieldNames)
		for _, name := range fieldNames {
			fc := fields[name]
			if !isEmailField(name) {
				continue
			}
			if fc.Format == "email" {
				continue
			}
			line := fs.OpenAPILines.RequestFieldLine(opID, name)
			if line == 0 {
				line = fc.Line
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    "api/openapi.yaml",
				Line:    line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XOO-72] email field %q in %s has no format: email constraint", name, opID),
				Advice:  "Add format: email to the email field",
			})
		}
	}
	return diags
}
