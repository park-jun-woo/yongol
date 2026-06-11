//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-56 — STML 폼이 소비하는 PATCH op의 requestBody 필드가 전부 required (WARNING, 부분수정 모순)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm56PatchAllRequired warns when a PATCH data-action form consumes an
// operation whose requestBody fields are all required (plans/gen/frontend
// Phase035, BUG-124). PATCH means partial update, but all-required forces the
// user to resend every field — a contradiction. The codegen never relaxes zod
// on its own (required is the OpenAPI decision), so this points back to OpenAPI:
// mark the fields optional and the existing zod_chain .optional() path applies.
// It is scoped to operations actually consumed by an STML form (de-duplicated by
// operationId) to avoid noise on non-frontend PATCH APIs.
func tm56PatchAllRequired(page stml.PageSpec, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	seen := make(map[string]bool)
	var diags []diagnostic.Diagnostic
	for _, a := range stml.CollectChildActions(page.Children) {
		if len(a.Fields) == 0 || seen[a.OperationID] {
			continue
		}
		entry, ok := opMap[a.OperationID]
		if !ok || entry.method != "PATCH" {
			continue
		}
		if !requestBodyAllRequired(entry.op) {
			continue
		}
		seen[a.OperationID] = true
		diags = append(diags, diagnostic.Diagnostic{
			File:        page.FileName,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelWarning,
			Message:     fmt.Sprintf("[TM-56] PATCH operation %q has an all-required requestBody but is consumed by a partial-update form — every field must be resent", a.OperationID),
			Advice:      fmt.Sprintf("Mark the optional fields of %q's requestBody as not required in OpenAPI; the generated zod schema then relaxes them automatically", a.OperationID),
			OperationID: a.OperationID,
		})
	}
	return diags
}
