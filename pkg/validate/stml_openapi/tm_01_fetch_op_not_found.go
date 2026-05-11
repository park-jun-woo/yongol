//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-01 — data-fetch operationId가 OpenAPI에 없음

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// tm01FetchOpNotFound returns a TM-01 diagnostic when a data-fetch
// operationId does not exist in the OpenAPI spec.
func tm01FetchOpNotFound(file, opID string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    file,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[TM-01] data-fetch operationId %q is not defined in OpenAPI", opID),
		Advice:  fmt.Sprintf("Add operationId %q to the OpenAPI spec, or fix the data-fetch value in the STML file", opID),
	}
}
