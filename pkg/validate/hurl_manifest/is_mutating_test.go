//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what isMutating — HTTP method가 POST/PUT/PATCH/DELETE인지 판정 검증

package hurl_manifest

import "testing"

func TestIsMutating(t *testing.T) {
	cases := []struct {
		name   string
		method string
		want   bool
	}{
		{name: "POST", method: "POST", want: true},
		{name: "PUT", method: "PUT", want: true},
		{name: "PATCH", method: "PATCH", want: true},
		{name: "DELETE", method: "DELETE", want: true},
		{name: "GET", method: "GET", want: false},
		{name: "HEAD", method: "HEAD", want: false},
		{name: "OPTIONS", method: "OPTIONS", want: false},
		{name: "lowercase_post", method: "post", want: true},
		{name: "mixed_case_Put", method: "Put", want: true},
		{name: "lowercase_get", method: "get", want: false},
		{name: "empty", method: "", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := isMutating(c.method)
			if got != c.want {
				t.Errorf("isMutating(%q) = %v, want %v", c.method, got, c.want)
			}
		})
	}
}
