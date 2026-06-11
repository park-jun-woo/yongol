//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what tm42IndexEntryDiags — data-index 항목 하나의 정합 진단: page 부재 / 필수 세그먼트 라우트 / manifest.frontend.index 모순

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm42IndexEntryDiags validates one data-index entry: it must sit on a
// data-page node, the page's resolved route (stml.RoutePaths first pattern)
// must carry no *required* parameter segment — the TM-34 judgment, a
// redirect has no value to fill it — and when manifest.frontend.index is
// also declared the two must name the same page (the same decision stated
// twice must not contradict). A nonexistent data-page is TM-39's finding
// and stays silent here.
func tm42IndexEntryDiags(fs *yongol.Fullstack, e sitemapEntry) []diagnostic.Diagnostic {
	if e.Node.Page == "" {
		return []diagnostic.Diagnostic{{
			File:    fs.Sitemap.FileName,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-42] data-index at %s sits on an entry without data-page — \"/\" has no page to redirect to", e.Path),
			Advice:  "Move data-index onto an <li data-page=\"...\"> entry",
		}}
	}
	target := findPageByName(fs.STMLPages, e.Node.Page)
	if target == nil {
		return nil // TM-39 already reports the nonexistent page
	}

	var diags []diagnostic.Diagnostic
	patterns := stml.RoutePaths(*target)
	if len(patterns) > 0 {
		if seg := firstRequiredSegment(patterns[0]); seg != "" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.Sitemap.FileName,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-42] data-index page %q (at %s) resolves to route %q with required segment %s — a redirect has no value to fill it", e.Node.Page, e.Path, patterns[0], seg),
				Advice:  "Mark a page without required route params as the index (optional :Name? segments are fine — they are stripped)",
			})
		}
	}
	if fs.Manifest != nil && fs.Manifest.Frontend.Index != "" && fs.Manifest.Frontend.Index != e.Node.Page {
		diags = append(diags, diagnostic.Diagnostic{
			File:    fs.Sitemap.FileName,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-42] sitemap data-index names %q but manifest.frontend.index names %q — the same decision is declared twice, in contradiction", e.Node.Page, fs.Manifest.Frontend.Index),
			Advice:  "Make both declarations name the same page, or keep only one of them",
		})
	}
	return diags
}
