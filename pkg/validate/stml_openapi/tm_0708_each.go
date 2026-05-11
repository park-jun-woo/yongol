//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-07/TM-08 — data-each 필드가 response에 없거나 array 타입이 아님

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm0708Each checks that each data-each field exists in the response schema
// (TM-07) and is of array type (TM-08).
func tm0708Each(eaches []stml.EachBlock, opID, file string, entry operationEntry) []diagnostic.Diagnostic {
	respFields := responseFields(entry.op)
	var diags []diagnostic.Diagnostic
	for _, e := range eaches {
		info, ok := respFields[e.Field]
		if !ok {
			diags = append(diags, diagnostic.Diagnostic{
				File:    file,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-07] data-each field %q is not in the response schema of operationId %q", e.Field, opID),
				Advice:  fmt.Sprintf("Add field %q to the response schema of %q in the OpenAPI spec, or remove the data-each from the STML file", e.Field, opID),
			})
			continue
		}
		if info.typ != "array" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    file,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-08] data-each field %q in operationId %q is not an array type (got %q)", e.Field, opID, info.typ),
				Advice:  fmt.Sprintf("Change the type of %q in the response schema to array, or use data-bind instead", e.Field),
			})
		}
	}
	return diags
}
