//ff:func feature=crosscheck type=test control=sequence topic=scenario-check
//ff:what TestParseFile_BodyFieldsCapturesAsserts — body / captures / asserts / headers 파싱

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
	assertEntry0(t, entries[0])
	assertEntry1(t, entries[1])
}
