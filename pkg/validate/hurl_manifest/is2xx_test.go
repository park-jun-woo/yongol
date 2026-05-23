//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what is2xx — hurl status code 2xx 판정 검증

package hurl_manifest

import "testing"

func TestIs2xx(t *testing.T) {
	cases := []struct {
		name string
		code string
		want bool
	}{
		{name: "empty_treated_as_success", code: "", want: true},
		{name: "200", code: "200", want: true},
		{name: "201", code: "201", want: true},
		{name: "204", code: "204", want: true},
		{name: "299", code: "299", want: true},
		{name: "301_redirect", code: "301", want: false},
		{name: "400_bad_request", code: "400", want: false},
		{name: "401_unauthorized", code: "401", want: false},
		{name: "500_server_error", code: "500", want: false},
		{name: "short_string", code: "20", want: false},
		{name: "long_string", code: "2001", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := is2xx(c.code)
			if got != c.want {
				t.Errorf("is2xx(%q) = %v, want %v", c.code, got, c.want)
			}
		})
	}
}
