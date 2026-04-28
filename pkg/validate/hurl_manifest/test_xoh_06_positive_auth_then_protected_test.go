//ff:func feature=validate type=test control=sequence topic=hurl-manifest
//ff:what TestXoh06_Positive_AuthThenProtected — auth 후 protected 요청 시 진단 없음

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh06_Positive_AuthThenProtected(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Auth: &manifest.Auth{Type: "jwt", Mode: "bearer"}},
		},
		HurlEntries: []hurl.HurlEntry{
			{Method: "POST", Path: "/auth/login", StatusCode: "200", File: "t.hurl", Line: 1},
			{Method: "GET", Path: "/workflows", File: "t.hurl", Line: 5,
				Headers: []hurl.HurlHeader{{Name: "Authorization", Value: "Bearer {{token}}"}}},
		},
	}
	if diags := xoh06AuthPrecondition(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %+v", diags)
	}
}
