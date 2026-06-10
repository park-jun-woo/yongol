//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-30 보조 — 액션 파라미터의 item.* 소스를 each 컨텍스트와 item 스키마로 검사

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm30CheckActionParams validates the item.<Field> param sources of one
// action. Outside a data-each every item.* source is an error (there is no
// row in scope). Inside, the field must exist in the enclosing data-each
// item schema; a nil itemFields (unknown op / unresolved each field —
// TM-01/TM-07 territory) stays silent.
func tm30CheckActionParams(a stml.ActionBlock, file string, itemFields map[string]bool, inEach bool) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, p := range a.Params {
		field, ok := itemParamField(p.Source)
		if !ok {
			continue
		}
		if !inEach {
			diags = append(diags, tm30OutsideEachDiag(file, a.OperationID, p.Source))
			continue
		}
		if itemFields == nil || itemFields[field] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        file,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[TM-30] data-param source %q references field %q which is not in the item schema of the enclosing data-each array", p.Source, field),
			Advice:      fmt.Sprintf("Add %q to the array item schema in the OpenAPI response, or fix the item.<Field> source", field),
			OperationID: a.OperationID,
		})
	}
	return diags
}
