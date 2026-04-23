//ff:func feature=openapi-parse type=util control=sequence
//ff:what DeriveSuccessStatus — HTTP method 관례에 따라 operation 의 성공 2xx 응답을 선택

package openapi

import (
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// httpMethodSuccessOrder lists the 2xx status codes yongol considers for a
// given HTTP method, in priority order. The generator emits a response for
// the first entry that the operation actually declares. Values follow RFC
// 9110 REST conventions:
//
//   - POST   : 201 (Created) → 200 (OK)
//   - PUT    : 200 (OK) → 204 (No Content)
//   - PATCH  : 200 (OK) → 204 (No Content)
//   - DELETE : 204 (No Content) → 200 (OK)
//   - GET    : 200 (OK)
//
// SSaC `@post/@put/...` directives describe the DB operation semantics and
// intentionally do not determine HTTP status — HTTP method is the source
// of truth for transport-layer status. Authors may declare multiple 2xx
// responses; yongol picks exactly one.
var httpMethodSuccessOrder = map[string][]int{
	"POST":   {201, 200},
	"PUT":    {200, 204},
	"PATCH":  {200, 204},
	"DELETE": {204, 200},
	"GET":    {200},
}

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

// Declared2xx is the exported form of declared2xx for callers outside
// this package (e.g. the openapi_ssac validator that implements XOS-82).
// Returns a fresh map so callers may mutate it safely.
func Declared2xx(op *openapi3.Operation) map[int]bool {
	return declared2xx(op)
}
