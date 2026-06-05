//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-16 — data-invalidates operationId가 OpenAPI에 없거나 GET이 아님 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm16InvalidatesOpNotFound checks that every operationId listed in an action's
// data-invalidates is defined in the OpenAPI spec and served under GET (only
// queries can be invalidated). A missing operationId or a non-GET method yields
// an ERROR. An empty Invalidates list yields no diagnostics.
func tm16InvalidatesOpNotFound(a stml.ActionBlock, file string, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, opID := range a.Invalidates {
		entry, ok := opMap[opID]
		if !ok {
			diags = append(diags, diagnostic.Diagnostic{
				File:    file,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-16] data-invalidates operationId %q on action %q is not defined in OpenAPI", opID, a.OperationID),
				Advice:  fmt.Sprintf("Add operationId %q to the OpenAPI spec, or fix the data-invalidates value in the STML file", opID),
			})
			continue
		}
		if entry.method != "GET" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    file,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-16] data-invalidates operationId %q on action %q is a %s endpoint; only GET queries can be invalidated", opID, a.OperationID, entry.method),
				Advice:  fmt.Sprintf("Change data-invalidates to reference a GET operationId, since %q uses %s", opID, entry.method),
			})
		}
	}
	return diags
}
