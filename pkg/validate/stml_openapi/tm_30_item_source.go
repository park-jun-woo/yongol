//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what TM-30 — item.<Field> 파라미터 소스가 data-each 밖에서 쓰이거나 item 스키마에 없는 필드 참조 (ERROR)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm30ItemSource checks every item.<Field> data-param-* source on the page
// (page-flow Phase006, BUG-112 (3)): the source is only legal inside a
// data-each block, and the field must exist in the item schema of the
// enclosing data-each array (OpenAPI response, innermost each). The walk
// covers page.Children only — top-level actions appear there too, so a
// separate page.Actions pass would double-report. An unresolvable item
// schema (unknown operationId / data-each field not in the response) stays
// silent: TM-01/TM-07 own those diagnostics.
func tm30ItemSource(page stml.PageSpec, raif map[string]map[string]map[string]bool) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	tm30CheckChildren(page.Children, page.FileName, "", nil, false, raif, &diags)
	return diags
}
