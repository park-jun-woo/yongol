//ff:func feature=crosscheck type=test control=sequence topic=scenario-check
//ff:what TestParseFile_Extensions — body fields / captures / asserts / headers 파싱 검증

package hurl

import (
	"os"
	"path/filepath"
	"testing"
)

// TestParseFile_BodyFieldsCapturesAsserts verifies the Phase002 parser
// extensions. A single entry exposes top-level JSON body fields,
// jsonpath asserts, captures, and request headers so XOH-03/04/07/08/09
// can interrogate them.
func TestParseFile_BodyFieldsCapturesAsserts(t *testing.T) {
	dir := t.TempDir()
	content := `POST {{host}}/auth/login
Content-Type: application/json
X-CSRF-Token: abc
{
  "email": "u@example.com",
  "password": "pw"
}

HTTP 200
[Captures]
token: jsonpath "$.access_token"
csrf: header "X-CSRF-Token"
[Asserts]
jsonpath "$.user.id" isInteger
jsonpath "$.user.email" == "u@example.com"

GET {{host}}/workflows/{{id}}
Authorization: Bearer {{token}}
HTTP 200
`
	path := filepath.Join(dir, "scenario.hurl")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	entries, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(entries) != 2 {
		t.Fatalf("want 2 entries, got %d", len(entries))
	}

	e0 := entries[0]
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

	e1 := entries[1]
	if e1.Method != "GET" || e1.Path != "/workflows/{{id}}" {
		t.Fatalf("entry[1] basics = %+v", e1)
	}
	if !hasHeader(e1.Headers, "Authorization") {
		t.Fatalf("entry[1] headers = %+v", e1.Headers)
	}
}

func containsAll(got, want []string) bool {
	set := map[string]struct{}{}
	for _, v := range got {
		set[v] = struct{}{}
	}
	for _, w := range want {
		if _, ok := set[w]; !ok {
			return false
		}
	}
	return true
}

func hasHeader(hs []HurlHeader, name string) bool {
	for _, h := range hs {
		if h.Name == name {
			return true
		}
	}
	return false
}
