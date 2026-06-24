//ff:func feature=validate type=util control=selection
//ff:what step.run — 도메인 모드에서 step.Run 을 once/merged/per-view 로 디스패치 (BUG-141)
package validate

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// run invokes s.Run with the right Fullstack scoping for the project mode.
//
// Single-site projects (and every non per-domain-SSOT step) run once on fs,
// exactly as before. In domain mode the singular OpenAPIDoc/STMLPages are nil
// (the data lives under the Domain* maps), so a step that reads those singular
// fields must be scoped:
//
//   - DomainAware  → run ONCE on the full fs (body already iterates all docs /
//     the merged Ground, or guards on Domains): keeps cross-domain detection.
//   - DomainMerged → run ONCE on fs.MergedOpenAPIView() so reverse rules that
//     consult the full op/scheme set against global config see every domain.
//   - usesPerDomainDoc (OpenAPI/STML-gated, neither flag) → run once per
//     fs.DomainView(name), reusing the single-site codepath per domain
//     (Decision A) and aggregating diagnostics in sorted domain order.
//
// The shared Ground is built once on the full fs and carried by every view via
// the shallow copy, so per-view rule bodies consulting Ground see the merged set.
func (s step) run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if !fs.IsDomained() {
		return s.Run(fs)
	}
	switch {
	case s.DomainAware:
		return s.Run(fs)
	case s.DomainMerged:
		return s.Run(fs.MergedOpenAPIView())
	case s.usesPerDomainDoc():
		var diags []diagnostic.Diagnostic
		for _, name := range fs.DomainNames() {
			diags = append(diags, s.Run(fs.DomainView(name))...)
		}
		return diags
	default:
		return s.Run(fs)
	}
}
