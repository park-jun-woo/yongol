//ff:func feature=validate type=test control=sequence topic=hurl-manifest
//ff:what TestXoh07_Positive_CSRFPresent — X-CSRF-Token 존재 시 진단 없음

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh07_Positive_CSRFPresent(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Auth: &manifest.Auth{Type: "jwt", Mode: "cookie"}},
		},
		HurlEntries: []hurl.HurlEntry{
			{Method: "POST", Path: "/workflows", File: "t.hurl", Line: 1,
				Headers: []hurl.HurlHeader{{Name: "X-CSRF-Token", Value: "{{csrf}}"}}},
		},
	}
	if diags := xoh07CSRFOnMutation(fs); len(diags) != 0 {
		t.Fatalf("want 0 diags, got %+v", diags)
	}
}
