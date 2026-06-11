//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what XMO-12 — no-front 태그인데 STML/컴포넌트가 실제로 소비 중 (WARNING, 낡은 태그)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmo12NoFrontConsumed warns when an operation tagged "no-front" is actually
// consumed by an STML page, a layout data-logout, or a component. The tag
// claims the endpoint is backend-only, so a real consumption means the tag is
// stale or wrong.
func xmo12NoFrontConsumed(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if !frontendEnabled(fs) {
		return nil
	}
	doc := fs.OpenAPIDoc
	if doc == nil || doc.Paths == nil {
		return nil
	}

	ops := collectOpIDs(doc)
	consumed := collectConsumedOps(fs.STMLPages, fs.Layouts, fs.Sitemap, fs.SpecsDir, ops)

	var diags []diagnostic.Diagnostic
	for _, item := range doc.Paths.Map() {
		diags = append(diags, xmo12ItemDiags(item, consumed)...)
	}
	return diags
}
