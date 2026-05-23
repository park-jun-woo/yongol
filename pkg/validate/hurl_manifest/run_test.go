//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what Run — Hurl↔Manifest 교차 검증 (XOH-06/07) 통합 검증

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	cases := []struct {
		name      string
		fs        *yongol.Fullstack
		wantEmpty bool
	}{
		{
			name:      "nil_manifest_no_diag",
			fs:        &yongol.Fullstack{},
			wantEmpty: true,
		},
		{
			name: "no_auth_config_no_diag",
			fs: &yongol.Fullstack{
				Manifest: &manifest.ProjectConfig{},
			},
			wantEmpty: true,
		},
		{
			name: "auth_present_no_hurl_entries_no_diag",
			fs: &yongol.Fullstack{
				Manifest: &manifest.ProjectConfig{
					Backend: manifest.Backend{Auth: &manifest.Auth{}},
				},
			},
			wantEmpty: true,
		},
		{
			name: "auth_present_protected_entry_produces_diag",
			fs: &yongol.Fullstack{
				Manifest: &manifest.ProjectConfig{
					Backend: manifest.Backend{Auth: &manifest.Auth{}},
				},
				HurlEntries: []hurl.HurlEntry{
					{Method: "GET", Path: "/api/users", StatusCode: "200", File: "t.hurl", Line: 1},
				},
			},
			wantEmpty: false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			diags := Run(c.fs)
			if c.wantEmpty && len(diags) != 0 {
				t.Errorf("expected no diags, got %d: %v", len(diags), diags)
			}
			if !c.wantEmpty && len(diags) == 0 {
				t.Errorf("expected diags, got none")
			}
		})
	}
}
