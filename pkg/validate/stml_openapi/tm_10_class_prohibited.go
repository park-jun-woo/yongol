//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what TM-10 — STML 요소에 class 속성 직접 사용 금지 (ERROR)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm10ClassProhibited scans all STML pages and emits TM-10 ERROR for any
// element that has a non-empty ClassName. Designer custom styles must use
// <!-- @override class="..." --> comments instead.
func tm10ClassProhibited(pages []stml.PageSpec) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, page := range pages {
		for _, f := range page.Fetches {
			diags = append(diags, checkFetchClass(f, page.FileName)...)
		}
		for _, a := range page.Actions {
			diags = append(diags, checkActionClass(a, page.FileName)...)
		}
		for _, c := range page.Children {
			diags = append(diags, checkChildClass(c, page.FileName)...)
		}
	}
	return diags
}
