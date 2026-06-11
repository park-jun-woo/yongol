//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-54 — 폼 필드가 prefill 2xx 응답 top-level 에 없으면 WARNING (빈 채로 시작)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm54FieldCoverage warns for each form field absent from the prefill response.
// The codegen still emits a compilable empty literal for it, so this is advisory
// (the input opens blank instead of prefilled), not a blocker.
func tm54FieldCoverage(a stml.ActionBlock, file string, resp map[string]responseFieldInfo) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, f := range a.Fields {
		if f.Name == "" {
			continue
		}
		if _, ok := resp[f.Name]; ok {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        file,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelWarning,
			Message:     fmt.Sprintf("[TM-54] form field %q of action %q is not in the data-prefill 2xx response of %q — it opens blank instead of prefilled", f.Name, a.OperationID, a.Prefill),
			Advice:      fmt.Sprintf("Add %q to the 2xx response schema of %q, or accept that this input starts empty", f.Name, a.Prefill),
			OperationID: a.OperationID,
		})
	}
	return diags
}
