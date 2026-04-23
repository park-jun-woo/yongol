//ff:func feature=openapi-parse type=accessor control=sequence
//ff:what Declared2xx — declared2xx 의 외부 노출 래퍼 (XOS-82 등 교차 검증에서 공용)

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// Declared2xx is the exported form of declared2xx for callers outside
// this package (e.g. the openapi_ssac validator that implements XOS-82).
// Returns a fresh map so callers may mutate it safely.
func Declared2xx(op *openapi3.Operation) map[int]bool {
	return declared2xx(op)
}
