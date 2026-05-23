//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what carriesAuthHeader — Authorization/Cookie 헤더 탐지 검증

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestCarriesAuthHeader(t *testing.T) {
	cases := []struct {
		name    string
		headers []hurl.HurlHeader
		want    bool
	}{
		{
			name:    "nil_headers_false",
			headers: nil,
			want:    false,
		},
		{
			name:    "empty_headers_false",
			headers: []hurl.HurlHeader{},
			want:    false,
		},
		{
			name:    "authorization_header",
			headers: []hurl.HurlHeader{{Name: "Authorization", Value: "Bearer tok"}},
			want:    true,
		},
		{
			name:    "authorization_lowercase",
			headers: []hurl.HurlHeader{{Name: "authorization", Value: "Bearer tok"}},
			want:    true,
		},
		{
			name:    "cookie_header",
			headers: []hurl.HurlHeader{{Name: "Cookie", Value: "session=abc"}},
			want:    true,
		},
		{
			name:    "cookie_lowercase",
			headers: []hurl.HurlHeader{{Name: "cookie", Value: "session=abc"}},
			want:    true,
		},
		{
			name:    "content_type_no_auth",
			headers: []hurl.HurlHeader{{Name: "Content-Type", Value: "application/json"}},
			want:    false,
		},
		{
			name: "auth_among_other_headers",
			headers: []hurl.HurlHeader{
				{Name: "Content-Type", Value: "application/json"},
				{Name: "Authorization", Value: "Bearer tok"},
				{Name: "Accept", Value: "*/*"},
			},
			want: true,
		},
		{
			name:    "mixed_case_AUTHORIZATION",
			headers: []hurl.HurlHeader{{Name: "AUTHORIZATION", Value: "Basic xyz"}},
			want:    true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := carriesAuthHeader(c.headers)
			if got != c.want {
				t.Errorf("carriesAuthHeader(...) = %v, want %v", got, c.want)
			}
		})
	}
}
