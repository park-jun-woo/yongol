//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-03 — data-action이 GET 메서드 endpoint를 참조

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// tm03ActionGetMethod returns a TM-03 diagnostic when a data-action
// references an OpenAPI operation that uses the GET method.
func tm03ActionGetMethod(file, opID string) diagnostic.Diagnostic {
	return diagnostic.Diagnostic{
		File:    file,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[TM-03] data-action %q references a GET endpoint; actions require POST/PUT/DELETE", opID),
		Advice:  fmt.Sprintf("Change the HTTP method of %q in the OpenAPI spec to POST/PUT/DELETE, or use data-fetch instead", opID),
	}
}
