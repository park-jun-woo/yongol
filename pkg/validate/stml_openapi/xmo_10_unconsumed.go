//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what XMO-10 — Frontend ON에서 OpenAPI operationId가 STML/컴포넌트에 미소비이며 no-front도 아님 (ERROR)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmo10Unconsumed detects OpenAPI operationIds that are never referenced from
// any STML data-fetch, data-action, layout data-logout, or component
// api.<Op>( call while the frontend is ON. Operations tagged "no-front" are
// explicit backend-only decisions and are skipped. Frontend OFF skips the rule
// entirely. STML 0 pages is yielded to XMO-11 (single ERROR) rather than
// flooding one ERROR per op.
func xmo10Unconsumed(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if !frontendEnabled(fs) {
		return nil
	}
	if len(fs.STMLPages) == 0 {
		return nil
	}
	doc := fs.OpenAPIDoc
	if doc == nil || doc.Paths == nil {
		return nil
	}

	ops := collectOpIDs(doc)
	consumed := collectConsumedOps(fs.STMLPages, fs.Layouts, fs.SpecsDir, ops)

	var diags []diagnostic.Diagnostic
	for _, item := range doc.Paths.Map() {
		diags = append(diags, xmo10ItemDiags(item, consumed)...)
	}
	return diags
}
