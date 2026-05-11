//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-11 — STML 페이지의 data-layout 값이 layouts/에 없음 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm11LayoutNotFound checks that every page's Layout value references an
// existing layout defined in Layouts. Pages with an empty Layout are skipped.
func tm11LayoutNotFound(pages []stml.PageSpec, layouts []stml.LayoutSpec) []diagnostic.Diagnostic {
	names := make(map[string]struct{}, len(layouts))
	for _, l := range layouts {
		names[l.Name] = struct{}{}
	}

	var diags []diagnostic.Diagnostic
	for _, page := range pages {
		if page.Layout == "" {
			continue
		}
		if _, ok := names[page.Layout]; !ok {
			diags = append(diags, diagnostic.Diagnostic{
				File:    page.FileName,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-11] data-layout %q on page %q not found in layouts/", page.Layout, page.Name),
				Advice:  fmt.Sprintf("Create layouts/%s.html or fix the data-layout attribute value", page.Layout),
			})
		}
	}
	return diags
}
