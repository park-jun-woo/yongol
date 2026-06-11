//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what TM-35 — frontend ON 인데 인덱스 미선언("/" 미점유 + frontend.index 없음 + sitemap data-index 없음) — 파일명 정렬 폴백 가시화 (WARNING)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm35IndexFallback surfaces the undeclared-index state (page-flow
// Phase009, BUG-114 (3)): the frontend is ON, no page mounts "/"
// (data-route), manifest.frontend.index is absent and no sitemap entry
// carries data-index (plans/stml/sitemap Phase001), so the emitter's
// fallback — the first public page in file-name sort order — decides the
// app's first screen by accident, not by declaration (Gozhip's
// forgot-password index). The advice names the page the fallback picks
// and both declaration vehicles. Zero STML pages stay silent — XMO-11
// already errors there and no fallback page exists to name.
func tm35IndexFallback(fs *yongol.Fullstack, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	if !frontendEnabled(fs) || len(fs.STMLPages) == 0 {
		return nil
	}
	if fs.Manifest.Frontend.Index != "" {
		return nil
	}
	// A sitemap data-index is an index declaration too (plans/stml/sitemap
	// Phase001) — its consistency is TM-42's job, not a fallback situation.
	if sitemapDeclaresIndex(fs.Sitemap) {
		return nil
	}
	for _, p := range fs.STMLPages {
		for _, pattern := range stml.RoutePaths(p) {
			if pattern == "/" {
				return nil
			}
		}
	}

	fallbackFile, fallbackPath := indexFallbackPage(fs, opMap)
	picked := fmt.Sprintf("page %q (%s)", fallbackFile, fallbackPath)
	if fallbackFile == "" {
		picked = fmt.Sprintf("%s (every candidate page is protected or parameterized)", fallbackPath)
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: fmt.Sprintf("[TM-35] no index is declared — the \"/\" route falls back to the first public page in file-name sort order, currently %s", picked),
		Advice:  "Declare the index: set manifest frontend.index to an STML page name (\"/\" redirects there), mount a page at \"/\" with data-route=\"/\", or mark a sitemap entry with data-index",
	}}
}
