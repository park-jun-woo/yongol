//ff:func feature=gen-gogin type=util control=iteration dimension=1 topic=dos-guard
//ff:what resolveHTTPLimits — manifest.backend.http 에서 global + per-op limit 추출

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// resolveHTTPLimits computes effective global + per-op limits from
// manifest. Missing manifest → defaults only. Parse failures fall back to
// defaults (SEC validation rule catches bad values at validate-time — the
// generator stays forgiving so a single bad override does not halt codegen).
func resolveHTTPLimits(fs *yongol.Fullstack) (bodyLimit, multipartLimit int64, bodyOverrides, multipartOverrides map[string]int64) {
	bodyLimit = defaultBodyLimit
	multipartLimit = defaultMultipartLimit
	bodyOverrides = map[string]int64{}
	multipartOverrides = map[string]int64{}

	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.HTTP == nil {
		return
	}
	h := fs.Manifest.Backend.HTTP
	if n, ok := parseHTTPSize(h.BodyLimit); ok {
		bodyLimit = n
	}
	if n, ok := parseHTTPSize(h.MultipartLimit); ok {
		multipartLimit = n
	}

	// Build operationId → "METHOD PATH" index from OpenAPI doc.
	opToRoute := buildOperationRouteIndex(fs)
	for opID, ov := range h.Overrides {
		applyHTTPOverride(opToRoute, opID, ov, bodyOverrides, multipartOverrides)
	}
	return
}
