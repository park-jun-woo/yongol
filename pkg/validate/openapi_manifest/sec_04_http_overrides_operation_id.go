//ff:func feature=validate type=rule control=iteration dimension=1 topic=config-check
//ff:what SEC-04 — every <key> under backend.http.overrides must exist as an OpenAPI operationId

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
		if opIDs[k] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    "manifest.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[SEC-04] backend.http.overrides.\"" + k + "\" does not exist as an OpenAPI operationId",
			Advice:  "Check that the key matches the OpenAPI operationId exactly (case-sensitive)",
		})
	}
	return diags
}
