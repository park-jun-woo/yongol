//ff:func feature=validate type=test control=sequence topic=hurl-manifest
//ff:what XOH-07 positive/negative — cookie 모드 mutation 에 X-CSRF-Token 헤더 포함

package hurl_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh07_Negative_MissingCSRFOnCookieMode(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Auth: &manifest.Auth{Type: "jwt", Mode: "cookie"}},
		},
		HurlEntries: []hurl.HurlEntry{
			{Method: "POST", Path: "/workflows", File: "t.hurl", Line: 1},
		},
	}
	diags := xoh07CSRFOnMutation(fs)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[XOH-07]") {
		t.Fatalf("want 1 XOH-07 diag, got %+v", diags)
	}
}

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
