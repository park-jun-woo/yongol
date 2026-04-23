//ff:func feature=gen-gogin type=util control=sequence topic=dos-guard
//ff:what applyHTTPOverride — opID/route 가 매칭될 때 body/multipart override 두 맵에 기록

package boot

import pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"

// applyHTTPOverride writes body/multipart overrides into the provided
// maps for a single operationId/route pair. Kept separate so
// resolveHTTPLimits keeps control=iteration at depth 1 and no nested if
// branches pile up inside the outer range loop.
func applyHTTPOverride(opToRoute map[string]string, opID string, ov pmanifest.HTTPOverride, bodyOverrides, multipartOverrides map[string]int64) {
	route, ok := opToRoute[opID]
	if !ok {
		return
	}
	if n, ok := parseHTTPSize(ov.BodyLimit); ok {
		bodyOverrides[route] = n
	}
	if n, ok := parseHTTPSize(ov.MultipartLimit); ok {
		multipartOverrides[route] = n
	}
}
