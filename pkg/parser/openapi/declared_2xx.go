//ff:func feature=openapi-parse type=util control=iteration dimension=1
//ff:what declared2xx — operation 의 2xx 응답 코드 집합 반환

package openapi

import (
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
)

// declared2xx returns the set of 2xx status codes declared on op.
// "default" and non-numeric keys are ignored. Used by DeriveSuccessStatus
// and by the XOS-80/81/82 validators that reason about the same set.
func declared2xx(op *openapi3.Operation) map[int]bool {
	out := map[int]bool{}
	if op == nil || op.Responses == nil {
		return out
	}
	for key := range op.Responses.Map() {
		code, err := strconv.Atoi(key)
		if err != nil {
			continue
		}
		if code >= 200 && code < 300 {
			out[code] = true
		}
	}
	return out
}
