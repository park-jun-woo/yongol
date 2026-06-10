//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what TM-34 — manifest.frontend.index 가 없는 페이지명 / 필수 세그먼트 라우트 / data-route="/" 동시 선언 (ERROR)

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm34IndexTarget validates manifest.frontend.index (page-flow Phase009,
// BUG-114 (3)) — a page-name reference deciding what "/" redirects to.
// Three error branches: (1) a page already mounts "/" via data-route while
// frontend.index declares a redirect — the same decision stated twice, in
// contradiction; (2) the name matches no STML page; (3) the target page's
// resolved route (stml.RoutePaths first pattern) carries a *required*
// parameter segment — a redirect has no value to fill it. Optional-only
// segments (":Name?") are legal: the base path matches with them omitted,
// so the emitter strips them. A protected target is legal too —
// <ProtectedRoute> bounces unauthenticated visits to /login.
func tm34IndexTarget(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	index := fs.Manifest.Frontend.Index
	if index == "" {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, p := range fs.STMLPages {
		mounted := false
		for _, pattern := range stml.RoutePaths(p) {
			if pattern == "/" {
				mounted = true
				break
			}
		}
		if mounted {
			diags = append(diags, diagnostic.Diagnostic{
				File:    "manifest.yaml",
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-34] manifest.frontend.index %q contradicts page %q which already mounts \"/\" (data-route) — \"/\" cannot both render a page and redirect", index, p.FileName),
				Advice:  fmt.Sprintf("Keep one declaration: remove frontend.index from manifest.yaml, or drop data-route=\"/\" from %s", p.FileName),
			})
		}
	}

	var target *stml.PageSpec
	for i := range fs.STMLPages {
		if fs.STMLPages[i].Name == index {
			target = &fs.STMLPages[i]
			break
		}
	}
	if target == nil {
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-34] manifest.frontend.index %q does not name any STML page", index),
			Advice:  "Use the index page's STML filename without .html (a page-name reference, not a path)",
		})
		return diags
	}

	patterns := stml.RoutePaths(*target)
	if len(patterns) > 0 {
		for _, seg := range strings.Split(patterns[0], "/") {
			if strings.HasPrefix(seg, ":") && !strings.HasSuffix(seg, "?") {
				diags = append(diags, diagnostic.Diagnostic{
					File:    "manifest.yaml",
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[TM-34] manifest.frontend.index %q resolves to route %q with required segment %s — a redirect has no value to fill it", index, patterns[0], seg),
					Advice:  "Point frontend.index at a page without required route params (optional :Name? segments are fine — they are stripped)",
				})
				break
			}
		}
	}
	return diags
}
