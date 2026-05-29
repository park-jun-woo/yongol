//ff:func feature=rule type=test control=sequence
//ff:what populateResponseSchema 2xx 케이스 — 첫 2xx가 back-compat 키에 등록되는지 검증

package ground

import (
	"testing"
)

// TestPopulateResponseSchema_2xx verifies the first 2xx response populates both
// status-code-specific (OpenAPI.response.<code>.<opID>) and back-compat
// (OpenAPI.response.<opID>) keys.
func TestPopulateResponseSchema_2xx(t *testing.T) {
	doc := buildDocWithOp("/w", "GET", "GetW", []string{"id", "name"})
	g := newGround()

	op := doc.Paths.Value("/w").Get
	populateResponseSchema(g, "GetW", op)

	// status-code-specific key must exist
	got200, ok := g.Schemas["OpenAPI.response.200.GetW"]
	if !ok {
		t.Fatalf("Schemas[OpenAPI.response.200.GetW] missing")
	}
	if len(got200) != 2 {
		t.Fatalf("200.GetW fields = %v, want 2", got200)
	}

	// back-compat primary 2xx key
	gotPrimary, ok := g.Schemas["OpenAPI.response.GetW"]
	if !ok {
		t.Fatalf("Schemas[OpenAPI.response.GetW] missing (back-compat)")
	}
	if len(gotPrimary) != 2 {
		t.Fatalf("back-compat primary fields = %v, want 2", gotPrimary)
	}
}
