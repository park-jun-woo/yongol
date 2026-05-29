//ff:func feature=scenario type=parser control=sequence
//ff:what ParseFile이 .hurl에서 요청/응답 쌍을 올바르게 추출하는지 검증
package hurl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile(t *testing.T) {
	dir := t.TempDir()
	content := `# Login
POST {{host}}/auth/login
Content-Type: application/json
{
  "email": "test@test.com",
  "password": "pass"
}

HTTP 200
[Captures]
token: jsonpath "$.access_token"

# Create gig
POST {{host}}/gigs
Authorization: Bearer {{token}}
Content-Type: application/json
{
  "title": "Test"
}

HTTP 201
`
	path := filepath.Join(dir, "scenario-test.hurl")
	os.WriteFile(path, []byte(content), 0644)

	entries, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}

	if entries[0].Method != "POST" || entries[0].Path != "/auth/login" || entries[0].StatusCode != "200" {
		t.Errorf("entry[0] = %+v", entries[0])
	}
	if entries[1].Method != "POST" || entries[1].Path != "/gigs" || entries[1].StatusCode != "201" {
		t.Errorf("entry[1] = %+v", entries[1])
	}
}
