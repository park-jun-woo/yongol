//ff:func feature=openapi-parse type=test control=sequence
//ff:what DeriveSuccessStatus 단위 테스트 — HTTP method 관례 우선순위 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func opWith(codes ...string) *openapi3.Operation {
	resp := openapi3.NewResponses()
	for _, c := range codes {
		resp.Set(c, &openapi3.ResponseRef{Value: openapi3.NewResponse()})
	}
	return &openapi3.Operation{Responses: resp}
}

// TestDeriveSuccessStatus_PostPrefers201 — POST with both 201 and 200
// picks 201 per RFC 9110 resource-creation convention.
func TestDeriveSuccessStatus_PostPrefers201(t *testing.T) {
	got := DeriveSuccessStatus(opWith("200", "201"), "POST")
	if got != 201 {
		t.Fatalf("POST 200+201 → %d, want 201", got)
	}
}

// TestDeriveSuccessStatus_PostFallsBackTo200 — POST operation that only
// declares 200 still gets a valid status. Common for idempotent POSTs.
func TestDeriveSuccessStatus_PostFallsBackTo200(t *testing.T) {
	got := DeriveSuccessStatus(opWith("200"), "POST")
	if got != 200 {
		t.Fatalf("POST 200 → %d, want 200", got)
	}
}

// TestDeriveSuccessStatus_DeletePrefers204 — DELETE with 204 + 200
// picks 204 (No Content is the conventional success for DELETE).
func TestDeriveSuccessStatus_DeletePrefers204(t *testing.T) {
	got := DeriveSuccessStatus(opWith("200", "204"), "DELETE")
	if got != 204 {
		t.Fatalf("DELETE 200+204 → %d, want 204", got)
	}
}

// TestDeriveSuccessStatus_GetReturns200 — GET operation with only the
// default 200 declaration.
func TestDeriveSuccessStatus_GetReturns200(t *testing.T) {
	got := DeriveSuccessStatus(opWith("200"), "GET")
	if got != 200 {
		t.Fatalf("GET 200 → %d, want 200", got)
	}
}

// TestDeriveSuccessStatus_NoTwoXX — operation declares only 4xx responses
// → function returns 0 so callers surface a XOS-22/XOS-81 diagnostic.
func TestDeriveSuccessStatus_NoTwoXX(t *testing.T) {
	got := DeriveSuccessStatus(opWith("400"), "POST")
	if got != 0 {
		t.Fatalf("POST 400 → %d, want 0", got)
	}
}

// TestDeriveSuccessStatus_LowestFallback — odd case: declared 2xx is not
// in the method's preference list (e.g. 207 Multi-Status). Fallback
// picks the lowest-numbered declared 2xx deterministically.
func TestDeriveSuccessStatus_LowestFallback(t *testing.T) {
	got := DeriveSuccessStatus(opWith("207", "208"), "POST")
	if got != 207 {
		t.Fatalf("POST 207+208 → %d, want 207", got)
	}
}

// TestDeriveSuccessStatus_NilOperation — defensive: nil op returns 0
// without panicking.
func TestDeriveSuccessStatus_NilOperation(t *testing.T) {
	if got := DeriveSuccessStatus(nil, "GET"); got != 0 {
		t.Fatalf("nil op → %d, want 0", got)
	}
}
