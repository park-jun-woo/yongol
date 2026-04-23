//ff:func feature=openapi-parse type=util control=iteration dimension=1
//ff:what DeriveSuccessStatus — HTTP method 관례에 따라 operation 의 성공 2xx 응답을 선택

package openapi

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// DeriveSuccessStatus returns the 2xx status code that the generated
// handler for op (served via method) should emit. Order of preference is
// method-conventional (see httpMethodSuccessOrder); if none of those are
// declared, fall back to the lowest declared 2xx code so unusual
// operations still produce runnable code. Returns 0 when op is nil or
// declares no 2xx response — callers should treat 0 as "generator
// cannot proceed" and surface a diagnostic (see XOS-80/81).
func DeriveSuccessStatus(op *openapi3.Operation, method string) int {
	if op == nil || op.Responses == nil {
		return 0
	}
	declared := declared2xx(op)
	if len(declared) == 0 {
		return 0
	}
	for _, candidate := range httpMethodSuccessOrder[strings.ToUpper(method)] {
		if declared[candidate] {
			return candidate
		}
	}
	// Fallback: pick the numerically-smallest declared 2xx. Deterministic
	// and keeps the generator moving when the author has used a non-standard
	// status (e.g. 207) or added a method we don't have a preference list
	// for. XOS-80 flags the mismatch separately.
	lowest := 0
	for code := range declared {
		if lowest == 0 || code < lowest {
			lowest = code
		}
	}
	return lowest
}
