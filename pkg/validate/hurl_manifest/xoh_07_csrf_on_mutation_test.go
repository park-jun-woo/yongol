//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what xoh07CSRFOnMutation — cookie/hybrid 모드에서 mutating 요청의 CSRF 헤더 검사 검증

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoh07CSRFOnMutation(t *testing.T) {
	cookieAuth := &manifest.Auth{Mode: "cookie"}
	hybridAuth := &manifest.Auth{Mode: "hybrid"}
	bearerAuth := &manifest.Auth{Mode: "bearer"}
	defaultAuth := &manifest.Auth{} // defaults to "cookie"

	cases := []TestXoh07CSRFOnMutationCase{
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
			name: "bearer_mode_no_diag",
			fs: &yongol.Fullstack{
				Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{Auth: bearerAuth}},
				HurlEntries: []hurl.HurlEntry{
					{Method: "POST", Path: "/api/orders", File: "t.hurl", Line: 1},
				},
			},
			wantCount: 0,
		},
		{
			name: "cookie_mode_mutating_no_csrf_produces_warning",
			fs: &yongol.Fullstack{
				Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{Auth: cookieAuth}},
				HurlEntries: []hurl.HurlEntry{
					{Method: "POST", Path: "/api/orders", File: "t.hurl", Line: 1},
				},
			},
			wantCount: 1,
		},
		{
			name: "hybrid_mode_mutating_no_csrf_produces_warning",
			fs: &yongol.Fullstack{
				Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{Auth: hybridAuth}},
				HurlEntries: []hurl.HurlEntry{
					{Method: "DELETE", Path: "/api/orders/1", File: "t.hurl", Line: 1},
				},
			},
			wantCount: 1,
		},
		{
			name: "default_mode_mutating_no_csrf_produces_warning",
			fs: &yongol.Fullstack{
				Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{Auth: defaultAuth}},
				HurlEntries: []hurl.HurlEntry{
					{Method: "PUT", Path: "/api/orders/1", File: "t.hurl", Line: 1},
				},
			},
			wantCount: 1,
		},
		{
			name: "cookie_mode_get_no_diag",
			fs: &yongol.Fullstack{
				Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{Auth: cookieAuth}},
				HurlEntries: []hurl.HurlEntry{
					{Method: "GET", Path: "/api/orders", File: "t.hurl", Line: 1},
				},
			},
			wantCount: 0,
		},
		{
			name: "cookie_mode_with_csrf_header_no_diag",
			fs: &yongol.Fullstack{
				Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{Auth: cookieAuth}},
				HurlEntries: []hurl.HurlEntry{
					{Method: "POST", Path: "/api/orders", File: "t.hurl", Line: 1,
						Headers: []hurl.HurlHeader{{Name: "X-CSRF-Token", Value: "tok"}}},
				},
			},
			wantCount: 0,
		},
		{
			name: "cookie_mode_auth_path_exempt",
			fs: &yongol.Fullstack{
				Manifest: &manifest.ProjectConfig{Backend: manifest.Backend{Auth: cookieAuth}},
				HurlEntries: []hurl.HurlEntry{
					{Method: "POST", Path: "/auth/login", File: "t.hurl", Line: 1},
				},
			},
			wantCount: 0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runXoh07CSRFOnMutation(t, c)
		})
	}
}
