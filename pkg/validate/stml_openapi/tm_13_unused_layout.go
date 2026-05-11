//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-13 — layouts/에 정의된 레이아웃이 어떤 페이지에서도 사용되지 않음 (WARNING)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm13UnusedLayout detects layouts that are neither referenced by any page's
// data-layout attribute nor set as the manifest's defaultLayout.
func tm13UnusedLayout(pages []stml.PageSpec, layouts []stml.LayoutSpec, defaultLayout string) []diagnostic.Diagnostic {
	used := make(map[string]struct{})
	if defaultLayout != "" {
		used[defaultLayout] = struct{}{}
	}
	for _, page := range pages {
		if page.Layout != "" {
			used[page.Layout] = struct{}{}
		}
	}

	var diags []diagnostic.Diagnostic
	for _, l := range layouts {
		if _, ok := used[l.Name]; !ok {
			diags = append(diags, diagnostic.Diagnostic{
				File:    l.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[TM-13] layout %q is defined but never used by any page or defaultLayout", l.Name),
				Advice:  fmt.Sprintf("Add data-layout=%q to a page, set it as defaultLayout in manifest.yaml, or remove layouts/%s.html", l.Name, l.Name),
			})
		}
	}
	return diags
}
