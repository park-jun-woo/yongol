//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-05 — data-field가 OpenAPI request body schema에 없음

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// tm05FieldNotFound returns a TM-05 diagnostic when a data-field name
// is not found in the request body schema of the referenced operation.
func tm05FieldNotFound(file, opID, field string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    file,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[TM-05] data-field %q is not in the request body schema of operationId %q", field, opID),
		Advice:  fmt.Sprintf("Add field %q to the requestBody schema of %q in the OpenAPI spec, or remove it from the STML file", field, opID),
	}
}
