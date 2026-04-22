//ff:func feature=validate type=rule control=iteration dimension=2 topic=tsx-openapi
//ff:what XOT-3 — useForm().register('x') 필드가 페이지의 mutation operation 의 request body schema 에 존재하는지 검증
package tsx_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xot03FormField validates XOT-3: every PageSpec.FormFields entry must
// correspond to a property in the request body schema of at least one
// apiClient invocation on the same page.
//
// Scope rationale — a page may host multiple mutations (create + update);
// rather than ranking them by heuristic, we accept a match against any
// call's request body. This is intentionally permissive: TypeScript +
// openapi-typescript is the primary type-check (strict); XOT-3 is a
// secondary, advisory WARNING that only fires when the field is absent
// from **every** body schema reachable through the page.
func xot03FormField(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.TSXPages) == 0 {
		return nil
	}
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, page := range fs.TSXPages {
		if len(page.FormFields) == 0 {
			continue
		}
		// Union of body field names across all apiClient calls on this page.
		known := make(map[string]bool)
		anyBody := false
		for _, call := range page.Calls {
			req := g.Lookup["OpenAPI.request."+call.OperationID]
			if len(req) == 0 {
				continue
			}
			anyBody = true
			for k := range req {
				known[k] = true
			}
		}
		if !anyBody {
			// No mutation discoverable from AST — skip rather than flood
			// WARNINGs. AI may have wired forms to helpers that the visitor
			// doesn't recognize yet; TypeScript still catches real mismatches.
			continue
		}
		for _, field := range page.FormFields {
			if known[field.Name] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    page.File,
				Line:    field.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: "[XOT-3] useForm register('" + field.Name + "') is not in any matching OpenAPI request body schema on this page",
				Advice:  "openapi-typescript 로 생성된 타입을 확인하거나 openapi.yaml 의 requestBody schema 에 " + field.Name + " 를 추가하세요",
			})
		}
	}
	return diags
}
