//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-structural
//ff:what XOO-71 — password 계열 필드에 minLength 제약 누락

package openapi

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoo71PasswordNoMinLength flags request-body password fields that have no
// minLength constraint. Emitted as WARNING.
func xoo71PasswordNoMinLength(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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
			if !isPasswordField(name) {
				continue
			}
			if fc.MinLength != nil {
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
				Message: fmt.Sprintf("[XOO-71] password field %q in %s has no minLength constraint", name, opID),
				Advice:  "password 필드에 minLength: 8 (또는 정책 최솟값) 을 추가하세요",
			})
		}
	}
	return diags
}
