//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what processAuthEntry — auth 경로 2xx/비2xx, auth헤더 유무, 보호 경로에 대한 상태 갱신/WARNING 생성 검증

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestProcessAuthEntry(t *testing.T) {
	cases := []TestProcessAuthEntryCase{
		{
			name:       "auth_path_2xx_sets_auth",
			entry:      hurl.HurlEntry{Method: "POST", Path: "/auth/login", StatusCode: "200"},
			authIssued: false,
			wantAuth:   true,
		},
		{
			name:       "auth_path_non_2xx_keeps_false",
			entry:      hurl.HurlEntry{Method: "POST", Path: "/auth/login", StatusCode: "401"},
			authIssued: false,
			wantAuth:   false,
		},
		{
			name:       "auth_path_non_2xx_keeps_true",
			entry:      hurl.HurlEntry{Method: "POST", Path: "/auth/login", StatusCode: "401"},
			authIssued: true,
			wantAuth:   true,
		},
		{
			name:       "protected_path_with_auth_issued_no_diag",
			entry:      hurl.HurlEntry{Method: "GET", Path: "/api/users", StatusCode: "200"},
			authIssued: true,
			wantAuth:   true,
		},
		{
			name: "protected_path_with_auth_header_no_diag",
			entry: hurl.HurlEntry{
				Method: "GET", Path: "/api/users", StatusCode: "200",
				Headers: []hurl.HurlHeader{{Name: "Authorization", Value: "Bearer tok"}},
			},
			authIssued:    false,
			wantAuth:      false,
			wantDiagCount: 0,
		},
		{
			name:          "protected_path_no_auth_produces_warning",
			entry:         hurl.HurlEntry{Method: "GET", Path: "/api/users", StatusCode: "200", File: "t.hurl", Line: 3},
			authIssued:    false,
			wantAuth:      false,
			wantDiagCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runProcessAuthEntry(t, c)
		})
	}
}
