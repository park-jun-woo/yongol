//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what xoh06AuthPrecondition — auth 미들웨어 선언 시 보호 endpoint 선행 인증 검증

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh06AuthPrecondition(t *testing.T) {
	authManifest := &manifest.ProjectConfig{
		Backend: manifest.Backend{Auth: &manifest.Auth{}},
	}

	cases := []TestXoh06AuthPreconditionCase{
		{
			name:      "nil_fullstack",
			fs:        nil,
			wantCount: 0,
		},
		{
			name:      "nil_manifest",
			fs:        &yongol.Fullstack{},
			wantCount: 0,
		},
		{
			name:      "no_auth_in_manifest",
			fs:        &yongol.Fullstack{Manifest: &manifest.ProjectConfig{}},
			wantCount: 0,
		},
		{
			name: "auth_present_no_entries",
			fs: &yongol.Fullstack{
				Manifest: authManifest,
			},
			wantCount: 0,
		},
		{
			name: "auth_present_protected_without_prior_auth",
			fs: &yongol.Fullstack{
				Manifest: authManifest,
				HurlEntries: []hurl.HurlEntry{
					{Method: "GET", Path: "/api/users", StatusCode: "200", File: "a.hurl", Line: 1},
				},
			},
			wantCount: 1,
		},
		{
			name: "auth_present_auth_first_then_protected",
			fs: &yongol.Fullstack{
				Manifest: authManifest,
				HurlEntries: []hurl.HurlEntry{
					{Method: "POST", Path: "/auth/login", StatusCode: "200", File: "a.hurl", Line: 1},
					{Method: "GET", Path: "/api/users", StatusCode: "200", File: "a.hurl", Line: 5},
				},
			},
			wantCount: 0,
		},
		{
			name: "separate_files_independent",
			fs: &yongol.Fullstack{
				Manifest: authManifest,
				HurlEntries: []hurl.HurlEntry{
					{Method: "POST", Path: "/auth/login", StatusCode: "200", File: "a.hurl", Line: 1},
					{Method: "GET", Path: "/api/users", StatusCode: "200", File: "b.hurl", Line: 1},
				},
			},
			wantCount: 1, // b.hurl has no auth
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runXoh06AuthPrecondition(t, c)
		})
	}
}
