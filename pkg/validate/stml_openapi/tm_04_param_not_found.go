//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-04 — data-param이 OpenAPI parameters에 없음

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm04Params checks that each data-param-* has a matching parameter in the
// OpenAPI operation.
func tm04Params(params []stml.ParamBind, opID, file string, entry operationEntry) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, p := range params {
		if !hasMatchingParam(entry, p.Name) {
			diags = append(diags, diagnostic.Diagnostic{
				File:    file,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-04] data-param %q is not declared in the parameters of operationId %q", p.Name, opID),
				Advice:  fmt.Sprintf("Add parameter %q to operationId %q in the OpenAPI spec, or remove it from the STML file", p.Name, opID),
			})
		}
	}
	return diags
}

// hasMatchingParam returns true if the operation has a parameter with the
// given name (case-insensitive).
func hasMatchingParam(entry operationEntry, name string) bool {
	if entry.op == nil {
		return false
	}
	for _, p := range entry.op.Parameters {
		if p == nil || p.Value == nil {
			continue
		}
		if strings.EqualFold(p.Value.Name, name) {
			return true
		}
	}
	return false
}
