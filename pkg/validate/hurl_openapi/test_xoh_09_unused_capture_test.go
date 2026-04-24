//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what XOH-09 positive/negative — unused captured variable

package hurl_openapi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh09_Negative_Unused(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "t.hurl")
	if err := os.WriteFile(path, []byte(`POST {{host}}/auth/login
HTTP 200
[Captures]
token: jsonpath "$.access_token"
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
		},
	}
	diags := xoh09UnusedCapture(fs)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[XOH-09]") {
		t.Fatalf("want 1 XOH-09 diag, got %+v", diags)
	}
}

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
