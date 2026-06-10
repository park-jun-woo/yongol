//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what tm36NavItemDiags — 단일 data-nav 항목 검사: 정적 경로 매칭 / 페이지명 존재 / 필수 세그먼트 부재

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm36NavItemDiags validates a single data-nav entry of a layout. Three
// error branches: (1) a "/"-prefixed static path that matches no page's
// resolved route pattern ("/" itself is exempt — the index route is
// always emitted); (2) a page-name reference naming no STML page; (3) a
// page-name reference whose resolved route (stml.RoutePaths first
// pattern) carries a *required* parameter segment — the layout has no
// value to fill it. Optional segments (":Name?") are legal: the emitter
// strips them, the base path matches with them omitted.
func tm36NavItemDiags(l stml.LayoutSpec, item stml.NavItem, pages []stml.PageSpec) []diagnostic.Diagnostic {
	if strings.HasPrefix(item.Path, "/") {
		if item.Path == "/" {
			return nil // the index route is always emitted
		}
		if tm36StaticNavMatches(item.Path, pages) {
			return nil
		}
		return []diagnostic.Diagnostic{{
			File:    l.File,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-36] data-nav %q in layout %q does not resolve to any STML page route", item.Path, l.Name),
			Advice:  "Point data-nav at an existing page route, or use a page-name reference (the target page's STML filename without .html)",
		}}
	}

	var target *stml.PageSpec
	for i := range pages {
		if pages[i].Name == item.Path {
			target = &pages[i]
			break
		}
	}
	if target == nil {
		return []diagnostic.Diagnostic{{
			File:    l.File,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-36] data-nav %q in layout %q does not name any STML page", item.Path, l.Name),
			Advice:  "Use the target page's STML filename without .html (a page-name reference), or a \"/\"-prefixed static path",
		}}
	}
	patterns := stml.RoutePaths(*target)
	if len(patterns) == 0 {
		return nil
	}
	for _, seg := range strings.Split(patterns[0], "/") {
		if strings.HasPrefix(seg, ":") && !strings.HasSuffix(seg, "?") {
			return []diagnostic.Diagnostic{{
				File:    l.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-36] data-nav %q in layout %q resolves to route %q with required segment %s — a static menu link has no value to fill it", item.Path, l.Name, patterns[0], seg),
				Advice:  "Point data-nav at a page without required route params; navigation into a parameterized page belongs to data-link with data-link-params",
			}}
		}
	}
	return nil
}
