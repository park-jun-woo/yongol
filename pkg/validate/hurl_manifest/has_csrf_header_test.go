//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what hasCSRFHeader — headerName 헤더 존재 여부 검증 (기본명/커스텀명/대소문자)

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestHasCSRFHeader(t *testing.T) {
	cases := []struct {
		name       string
		headers    []hurl.HurlHeader
		headerName string
		want       bool
	}{
		{name: "nil_headers", headers: nil, headerName: "X-XSRF-TOKEN", want: false},
		{name: "empty_headers", headers: []hurl.HurlHeader{}, headerName: "X-XSRF-TOKEN", want: false},
		{name: "exact_match", headers: []hurl.HurlHeader{{Name: "X-XSRF-TOKEN", Value: "abc"}}, headerName: "X-XSRF-TOKEN", want: true},
		{name: "lowercase", headers: []hurl.HurlHeader{{Name: "x-xsrf-token", Value: "abc"}}, headerName: "X-XSRF-TOKEN", want: true},
		{name: "mixed_case", headers: []hurl.HurlHeader{{Name: "X-Xsrf-Token", Value: "abc"}}, headerName: "X-XSRF-TOKEN", want: true},
		{name: "unrelated_header", headers: []hurl.HurlHeader{{Name: "Content-Type", Value: "json"}}, headerName: "X-XSRF-TOKEN", want: false},
		{
			name:       "legacy_name_not_default",
			headers:    []hurl.HurlHeader{{Name: "X-CSRF-Token", Value: "tok"}},
			headerName: "X-XSRF-TOKEN",
			want:       false,
		},
		{
			name:       "custom_name_match",
			headers:    []hurl.HurlHeader{{Name: "X-My-CSRF", Value: "tok"}},
			headerName: "X-My-CSRF",
			want:       true,
		},
		{
			name:       "custom_name_default_header_rejected",
			headers:    []hurl.HurlHeader{{Name: "X-XSRF-TOKEN", Value: "tok"}},
			headerName: "X-My-CSRF",
			want:       false,
		},
		{
			name: "csrf_among_others",
			headers: []hurl.HurlHeader{
				{Name: "Content-Type", Value: "json"},
				{Name: "X-XSRF-TOKEN", Value: "tok"},
				{Name: "Accept", Value: "*/*"},
			},
			headerName: "X-XSRF-TOKEN",
			want:       true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := hasCSRFHeader(c.headers, c.headerName)
			if got != c.want {
				t.Errorf("hasCSRFHeader(..., %q) = %v, want %v", c.headerName, got, c.want)
			}
		})
	}
}
