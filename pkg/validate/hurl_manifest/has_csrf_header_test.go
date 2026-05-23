//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what hasCSRFHeader — X-CSRF-Token 헤더 존재 여부 검증

package hurl_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

func TestHasCSRFHeader(t *testing.T) {
	cases := []struct {
		name    string
		headers []hurl.HurlHeader
		want    bool
	}{
		{name: "nil_headers", headers: nil, want: false},
		{name: "empty_headers", headers: []hurl.HurlHeader{}, want: false},
		{name: "exact_match", headers: []hurl.HurlHeader{{Name: "X-CSRF-Token", Value: "abc"}}, want: true},
		{name: "lowercase", headers: []hurl.HurlHeader{{Name: "x-csrf-token", Value: "abc"}}, want: true},
		{name: "uppercase", headers: []hurl.HurlHeader{{Name: "X-CSRF-TOKEN", Value: "abc"}}, want: true},
		{name: "unrelated_header", headers: []hurl.HurlHeader{{Name: "Content-Type", Value: "json"}}, want: false},
		{
			name: "csrf_among_others",
			headers: []hurl.HurlHeader{
				{Name: "Content-Type", Value: "json"},
				{Name: "X-CSRF-Token", Value: "tok"},
				{Name: "Accept", Value: "*/*"},
			},
			want: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := hasCSRFHeader(c.headers)
			if got != c.want {
				t.Errorf("hasCSRFHeader(...) = %v, want %v", got, c.want)
			}
		})
	}
}
