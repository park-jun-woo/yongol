//ff:func feature=crosscheck type=test-helper control=sequence topic=scenario-check
//ff:what assertEntry0 — ParseFile BodyFieldsCapturesAsserts 1번째 entry 검증

package hurl

import "testing"

// assertEntry0 verifies the body/headers/captures/asserts fields on the
// first entry parsed by the Phase002 extension fixture.
func assertEntry0(t *testing.T, e0 HurlEntry) {
	t.Helper()
	if e0.Method != "POST" || e0.Path != "/auth/login" || e0.StatusCode != "200" {
		t.Fatalf("entry[0] basics = %+v", e0)
	}
	if !containsAll(e0.BodyFields, []string{"email", "password"}) {
		t.Fatalf("entry[0] body fields = %v", e0.BodyFields)
	}
	if !hasHeader(e0.Headers, "Content-Type") || !hasHeader(e0.Headers, "X-CSRF-Token") {
		t.Fatalf("entry[0] headers = %+v", e0.Headers)
	}
	if len(e0.Captures) != 2 {
		t.Fatalf("entry[0] captures = %+v", e0.Captures)
	}
	if e0.Captures[0].Var != "token" || e0.Captures[0].JSONPath != "$.access_token" {
		t.Fatalf("capture[0] = %+v", e0.Captures[0])
	}
	if e0.Captures[1].Var != "csrf" || e0.Captures[1].Header != "X-CSRF-Token" {
		t.Fatalf("capture[1] = %+v", e0.Captures[1])
	}
	if len(e0.Asserts) != 2 {
		t.Fatalf("entry[0] asserts = %+v", e0.Asserts)
	}
	if e0.Asserts[0].JSONPath != "$.user.id" {
		t.Fatalf("assert[0] = %+v", e0.Asserts[0])
	}
}
