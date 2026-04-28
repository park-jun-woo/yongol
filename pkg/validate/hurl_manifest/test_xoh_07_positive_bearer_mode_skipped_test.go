//ff:func feature=validate type=test control=sequence topic=hurl-manifest
//ff:what TestXoh07_Positive_BearerModeSkipped — bearer 모드는 CSRF 검사 skip

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh07_Positive_BearerModeSkipped(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Auth: &manifest.Auth{Type: "jwt", Mode: "bearer"}},
		},
		HurlEntries: []hurl.HurlEntry{
			{Method: "POST", Path: "/workflows", File: "t.hurl", Line: 1},
		},
	}
	if diags := xoh07CSRFOnMutation(fs); len(diags) != 0 {
		t.Fatalf("bearer mode should skip; got %+v", diags)
	}
}
