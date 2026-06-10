//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what shouldCheckCSRF — mutating/non-auth/CSRF 헤더 조합별 판정 검증

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestShouldCheckCSRF(t *testing.T) {
	cases := []struct {
		name       string
		entry      hurl.HurlEntry
		headerName string
		want       bool
	}{
		{
			name:       "GET_not_mutating",
			entry:      hurl.HurlEntry{Method: "GET", Path: "/api/users"},
			headerName: "X-XSRF-TOKEN",
			want:       false,
		},
		{
			name:       "POST_auth_path_skip",
			entry:      hurl.HurlEntry{Method: "POST", Path: "/auth/login"},
			headerName: "X-XSRF-TOKEN",
			want:       false,
		},
		{
			name: "POST_with_csrf_header_skip",
			entry: hurl.HurlEntry{
				Method:  "POST",
				Path:    "/api/orders",
				Headers: []hurl.HurlHeader{{Name: "X-XSRF-TOKEN", Value: "tok"}},
			},
			headerName: "X-XSRF-TOKEN",
			want:       false,
		},
		{
			name: "POST_with_custom_csrf_header_skip",
			entry: hurl.HurlEntry{
				Method:  "POST",
				Path:    "/api/orders",
				Headers: []hurl.HurlHeader{{Name: "X-My-CSRF", Value: "tok"}},
			},
			headerName: "X-My-CSRF",
			want:       false,
		},
		{
			name: "POST_wrong_header_needs_check",
			entry: hurl.HurlEntry{
				Method:  "POST",
				Path:    "/api/orders",
				Headers: []hurl.HurlHeader{{Name: "X-CSRF-Token", Value: "tok"}},
			},
			headerName: "X-XSRF-TOKEN",
			want:       true,
		},
		{
			name:       "POST_no_csrf_needs_check",
			entry:      hurl.HurlEntry{Method: "POST", Path: "/api/orders"},
			headerName: "X-XSRF-TOKEN",
			want:       true,
		},
		{
			name:       "PUT_no_csrf_needs_check",
			entry:      hurl.HurlEntry{Method: "PUT", Path: "/api/orders/1"},
			headerName: "X-XSRF-TOKEN",
			want:       true,
		},
		{
			name:       "PATCH_no_csrf_needs_check",
			entry:      hurl.HurlEntry{Method: "PATCH", Path: "/api/orders/1"},
			headerName: "X-XSRF-TOKEN",
			want:       true,
		},
		{
			name:       "DELETE_no_csrf_needs_check",
			entry:      hurl.HurlEntry{Method: "DELETE", Path: "/api/orders/1"},
			headerName: "X-XSRF-TOKEN",
			want:       true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := shouldCheckCSRF(c.entry, c.headerName)
			if got != c.want {
				t.Errorf("shouldCheckCSRF(..., %q) = %v, want %v", c.headerName, got, c.want)
			}
		})
	}
}
