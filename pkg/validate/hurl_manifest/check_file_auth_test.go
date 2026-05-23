//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what checkFileAuth — 인증 선행 없는 보호 요청에 XOH-06 WARNING 생성 검증

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestCheckFileAuth(t *testing.T) {
	cases := []TestCheckFileAuthCase{
		{
			name:      "nil_entries_no_diag",
			entries:   nil,
			wantCount: 0,
		},
		{
			name: "auth_first_then_protected_no_diag",
			entries: []hurl.HurlEntry{
				{Method: "POST", Path: "/auth/login", StatusCode: "200", File: "t.hurl", Line: 1},
				{Method: "GET", Path: "/api/users", StatusCode: "200", File: "t.hurl", Line: 5},
			},
			wantCount: 0,
		},
		{
			name: "protected_before_auth_warning",
			entries: []hurl.HurlEntry{
				{Method: "GET", Path: "/api/users", StatusCode: "200", File: "t.hurl", Line: 1},
				{Method: "POST", Path: "/auth/login", StatusCode: "200", File: "t.hurl", Line: 5},
			},
			wantCount: 1,
		},
		{
			name: "protected_with_auth_header_no_diag",
			entries: []hurl.HurlEntry{
				{Method: "GET", Path: "/api/users", StatusCode: "200", File: "t.hurl", Line: 1,
					Headers: []hurl.HurlHeader{{Name: "Authorization", Value: "Bearer tok"}}},
			},
			wantCount: 0,
		},
		{
			name: "multiple_protected_before_auth",
			entries: []hurl.HurlEntry{
				{Method: "GET", Path: "/api/users", StatusCode: "200", File: "t.hurl", Line: 1},
				{Method: "GET", Path: "/api/orders", StatusCode: "200", File: "t.hurl", Line: 5},
				{Method: "POST", Path: "/auth/login", StatusCode: "200", File: "t.hurl", Line: 10},
			},
			wantCount: 2,
		},
		{
			name: "auth_with_non_2xx_does_not_count",
			entries: []hurl.HurlEntry{
				{Method: "POST", Path: "/auth/login", StatusCode: "401", File: "t.hurl", Line: 1},
				{Method: "GET", Path: "/api/users", StatusCode: "200", File: "t.hurl", Line: 5},
			},
			wantCount: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runCheckFileAuth(t, c)
		})
	}
}
