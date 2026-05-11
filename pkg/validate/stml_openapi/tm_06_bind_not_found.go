//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-06 — data-bind가 OpenAPI response schema에 없음

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm06Binds checks that each data-bind field exists in the response schema.
func tm06Binds(binds []stml.FieldBind, opID, file string, entry operationEntry) []diagnostic.Diagnostic {
	respFields := responseFields(entry.op)
	var diags []diagnostic.Diagnostic
	for _, b := range binds {
		fieldName := b.Name
		// For dotted paths like "User.Name", check the top-level key.
		if idx := strings.IndexByte(fieldName, '.'); idx >= 0 {
			fieldName = fieldName[:idx]
		}
		if _, ok := respFields[fieldName]; !ok {
			diags = append(diags, diagnostic.Diagnostic{
				File:    file,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-06] data-bind %q is not in the response schema of operationId %q", b.Name, opID),
				Advice:  fmt.Sprintf("Add field %q to the response schema of %q in the OpenAPI spec, or remove the data-bind from the STML file", fieldName, opID),
			})
		}
	}
	return diags
}
