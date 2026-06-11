//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what prefillRuleDiags — 페이지의 prefill 관련 규칙 TM-54/55/56 묶음 실행

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// prefillRuleDiags runs the data-prefill cross-validation rules of
// plans/gen/frontend Phase035 (BUG-124) over one page: TM-54 (prefill source /
// field coverage), TM-55 (edit form without prefill), TM-56 (PATCH all-required).
func prefillRuleDiags(page stml.PageSpec, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, tm54PrefillSource(page, opMap)...)
	diags = append(diags, tm55EditFormNoPrefill(page, opMap)...)
	diags = append(diags, tm56PatchAllRequired(page, opMap)...)
	return diags
}
