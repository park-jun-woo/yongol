//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what TestXoh09_Positive_Used — capture 후 `{{var}}` 참조 존재 시 진단 없음

package hurl_openapi

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh09_Positive_Used(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "t.hurl")
	if err := os.WriteFile(path, []byte(`POST {{host}}/auth/login
HTTP 200
[Captures]
token: jsonpath "$.access_token"

GET {{host}}/me
Authorization: Bearer {{token}}
HTTP 200
`), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	fs := &yongol.Fullstack{
		HurlEntries: []hurl.HurlEntry{
			{
				Method: "POST", Path: "/auth/login", File: path, Line: 1,
				Captures: []hurl.HurlCapture{
					{Var: "token", Source: "jsonpath", JSONPath: "$.access_token", Line: 4},
				},
			},
			{Method: "GET", Path: "/me", File: path, Line: 6},
		},
	}
	if diags := xoh09UnusedCapture(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %+v", diags)
	}
}
