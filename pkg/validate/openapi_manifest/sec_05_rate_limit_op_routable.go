//ff:func feature=validate type=rule control=iteration dimension=1 topic=config-check
//ff:what SEC-05 — every backend.rate_limit operationId must map to an OpenAPI route

package openapi_manifest

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// sec05RateLimitOpRoutable validates SEC-05: every key under
// backend.rate_limit (an operationId) must correspond to an operationId
// declared in the OpenAPI document, so codegen can resolve it to a gin
// route ("METHOD /path"). A rate_limit entry whose operationId has no
// route is silently dropped by the generator (blockRateLimit), producing a
// backend with no rate limiter despite a populated manifest (BUG-115).
//
// Routability criterion: an operationId is routable iff some OpenAPI
// operation declares it (non-empty). This is identical to the generator's
// buildOperationRouteIndex membership, which keys the same set under the
// same traversal/skip conditions. operationIDSet collects exactly that set,
// so this rule predicts generator routability without drift. SEC-05 is the
// direct successor of the retired SEC-03 (application-layer rate_limit
// operationId check), redefined against the current RouteRateLimit wiring,
// and a sibling of the active SEC-04 (http.overrides operationId check).
func sec05RateLimitOpRoutable(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || len(fs.Manifest.Backend.RateLimit) == 0 {
		return nil
	}
	opIDs := operationIDSet(fs)

	// Sort keys for deterministic diagnostic order.
	keys := make([]string, 0, len(fs.Manifest.Backend.RateLimit))
	for k := range fs.Manifest.Backend.RateLimit {
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
			Message: "[SEC-05] backend.rate_limit.\"" + k + "\" does not map to an OpenAPI route (operationId not found)",
			Advice:  "Use an existing OpenAPI operationId as the rate_limit key (case-sensitive); otherwise the rate limiter is silently omitted at codegen",
		})
	}
	return diags
}
