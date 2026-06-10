//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-31 — data-link 대상 페이지명이 STML 페이지 집합에 없음 (ERROR, 허공으로의 링크)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm31LinkTargetNotFound checks every data-link on the page (page-flow
// Phase007, BUG-114 (1)): the target must name an existing STML page
// (filename without .html). A typo'd target would emit a <Link> into the
// void — a route react-router never matches. The walk covers page.Children
// only, like TM-30.
func tm31LinkTargetNotFound(page stml.PageSpec, pages []stml.PageSpec) []diagnostic.Diagnostic {
	var refs []linkRefCtx
	collectLinkRefs(page.Children, "", nil, false, nil, &refs)
	if len(refs) == 0 {
		return nil
	}
	names := map[string]bool{}
	for _, p := range pages {
		names[p.Name] = true
	}
	var diags []diagnostic.Diagnostic
	for _, ref := range refs {
		if names[ref.Link.TargetPage] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    page.FileName,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-31] data-link target %q does not name any STML page", ref.Link.TargetPage),
			Advice:  "Use the target page's STML filename without .html (a page-name reference, not a path)",
		})
	}
	return diags
}
