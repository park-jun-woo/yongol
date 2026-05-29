//ff:func feature=crosscheck type=test control=sequence topic=scenario-check
//ff:what TestParseFile_VariableURL — 임의 {{var}} URL 엔트리 인식 + URLVar 채움 (BUG-092)

package hurl

import (
	"os"
	"path/filepath"
	"testing"
)

// TestParseFile_VariableURL verifies BUG-092: a request line whose URL
// uses any {{var}} prefix (not just {{host}}) is recognized as a new
// entry, so the following [Captures] line binds to THAT entry rather than
// the previous one. It also asserts URLVar is populated as the variable
// name, or "" for absolute http(s):// URLs.
func TestParseFile_VariableURL(t *testing.T) {
	dir := t.TempDir()
	content := `GET {{host}}/workflows
HTTP 200

POST {{authurl}}/auth/v1/token
HTTP 200
[Captures]
token: jsonpath "$.access_token"

GET http://localhost:8080/health
HTTP 200
`
	path := filepath.Join(dir, "scenario-var.hurl")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	entries, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(entries) != 3 {
		t.Fatalf("want 3 entries, got %d: %+v", len(entries), entries)
	}

	// entry[0]: {{host}} → URLVar "host"
	if entries[0].Method != "GET" || entries[0].Path != "/workflows" {
		t.Errorf("entry[0] = %+v", entries[0])
	}
	if entries[0].URLVar != "host" {
		t.Errorf("entry[0].URLVar = %q, want %q", entries[0].URLVar, "host")
	}

	// entry[1]: {{authurl}} recognized as its own entry, URLVar "authurl".
	if entries[1].Method != "POST" || entries[1].Path != "/auth/v1/token" {
		t.Errorf("entry[1] = %+v", entries[1])
	}
	if entries[1].URLVar != "authurl" {
		t.Errorf("entry[1].URLVar = %q, want %q", entries[1].URLVar, "authurl")
	}
	// The token capture must bind to entry[1] (the {{authurl}} request),
	// NOT the previous {{host}} entry.
	if len(entries[1].Captures) != 1 || entries[1].Captures[0].Var != "token" {
		t.Errorf("entry[1].Captures = %+v, want token capture", entries[1].Captures)
	}
	if len(entries[0].Captures) != 0 {
		t.Errorf("entry[0].Captures = %+v, want none (no misattribution)", entries[0].Captures)
	}

	// entry[2]: absolute http:// URL → URLVar "" (still recognized).
	if entries[2].Method != "GET" || entries[2].Path != "/health" {
		t.Errorf("entry[2] = %+v", entries[2])
	}
	if entries[2].URLVar != "" {
		t.Errorf("entry[2].URLVar = %q, want %q", entries[2].URLVar, "")
	}
}
