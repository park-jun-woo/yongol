//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what isNoFront — OpenAPI 오퍼레이션이 "no-front" 태그를 가졌는지 판정

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// noFrontTag marks an OpenAPI operation as backend-only (not consumed by any
// STML page or component).
const noFrontTag = "no-front"

// isNoFront reports whether op carries the "no-front" tag.
func isNoFront(op *openapi3.Operation) bool {
	if op == nil {
		return false
	}
	for _, tag := range op.Tags {
		if tag == noFrontTag {
			return true
		}
	}
	return false
}
