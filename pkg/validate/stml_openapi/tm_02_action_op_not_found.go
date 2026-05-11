//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-02 — data-action operationId가 OpenAPI에 없음

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// tm02ActionOpNotFound returns a TM-02 diagnostic when a data-action
// operationId does not exist in the OpenAPI spec.
func tm02ActionOpNotFound(file, opID string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    file,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[TM-02] data-action operationId %q is not defined in OpenAPI", opID),
		Advice:  fmt.Sprintf("Add operationId %q to the OpenAPI spec, or fix the data-action value in the STML file", opID),
	}
}
