//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what shouldCheckCSRF — mutating/non-auth/CSRF 헤더 조합별 판정 검증

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestShouldCheckCSRF(t *testing.T) {
	cases := []struct {
		name  string
		entry hurl.HurlEntry
		want  bool
	}{
		{
			name:  "GET_not_mutating",
			entry: hurl.HurlEntry{Method: "GET", Path: "/api/users"},
			want:  false,
		},
		{
			name:  "POST_auth_path_skip",
			entry: hurl.HurlEntry{Method: "POST", Path: "/auth/login"},
			want:  false,
		},
		{
			name: "POST_with_csrf_header_skip",
			entry: hurl.HurlEntry{
				Method:  "POST",
				Path:    "/api/orders",
				Headers: []hurl.HurlHeader{{Name: "X-CSRF-Token", Value: "tok"}},
			},
			want: false,
		},
		{
			name:  "POST_no_csrf_needs_check",
			entry: hurl.HurlEntry{Method: "POST", Path: "/api/orders"},
			want:  true,
		},
		{
			name:  "PUT_no_csrf_needs_check",
			entry: hurl.HurlEntry{Method: "PUT", Path: "/api/orders/1"},
			want:  true,
		},
		{
			name:  "PATCH_no_csrf_needs_check",
			entry: hurl.HurlEntry{Method: "PATCH", Path: "/api/orders/1"},
			want:  true,
		},
		{
			name:  "DELETE_no_csrf_needs_check",
			entry: hurl.HurlEntry{Method: "DELETE", Path: "/api/orders/1"},
			want:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := shouldCheckCSRF(c.entry)
			if got != c.want {
				t.Errorf("shouldCheckCSRF(...) = %v, want %v", got, c.want)
			}
		})
	}
}
