//ff:func feature=validate type=test control=sequence topic=hurl-manifest
//ff:what TestXoh06_Negative_NoPriorAuth — auth 선행 없으면 WARNING

package hurl_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh06_Negative_NoPriorAuth(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{Auth: &manifest.Auth{Type: "jwt", Mode: "bearer"}},
		},
		HurlEntries: []hurl.HurlEntry{
			{Method: "GET", Path: "/workflows", File: "t.hurl", Line: 1},
		},
	}
	diags := xoh06AuthPrecondition(fs)
	if len(diags) != 1 || !strings.Contains(diags[0].Message, "[XOH-06]") {
		t.Fatalf("want 1 XOH-06 diag, got %+v", diags)
	}
}
