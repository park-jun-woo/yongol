//ff:func feature=validate type=rule control=iteration dimension=1 topic=config-check
//ff:what SEC-04 — backend.http.overrides.<key> 의 <key> 는 OpenAPI operationId 에 존재해야 함

package openapi_manifest

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// sec04HTTPOverridesOperationID validates SEC-04: every key under
// backend.http.overrides must correspond to an operationId declared in the
// OpenAPI document. Catches typos (CamelCase vs camelCase, pluralization)
// that would otherwise silently disable the override at runtime.
func sec04HTTPOverridesOperationID(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.HTTP == nil {
		return nil
	}
	overrides := fs.Manifest.Backend.HTTP.Overrides
	if len(overrides) == 0 {
		return nil
	}
	opIDs := operationIDSet(fs)

	// Sort keys for deterministic diagnostic order.
	keys := make([]string, 0, len(overrides))
	for k := range overrides {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var diags []diagnostic.Diagnostic
	for _, k := range keys {
		if !opIDs[k] {
			diags = append(diags, diagnostic.Diagnostic{
				File:    "manifest.yaml",
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[SEC-04] backend.http.overrides.\"" + k + "\" 가 OpenAPI operationId 에 존재하지 않습니다",
				Advice:  "OpenAPI operationId 의 대소문자·철자와 일치하는지 확인하세요",
			})
		}
	}
	return diags
}

// operationIDSet collects all operationIds declared in the OpenAPI doc.
func operationIDSet(fs *yongol.Fullstack) map[string]bool {
	s := map[string]bool{}
	if fs == nil || fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return s
	}
	for _, pi := range fs.OpenAPIDoc.Paths.Map() {
		for _, op := range pi.Operations() {
			if op == nil || op.OperationID == "" {
				continue
			}
			s[op.OperationID] = true
		}
	}
	return s
}
