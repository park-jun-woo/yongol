//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-cors
//ff:what corsAllowOriginsHasWildcard — allow_origins 목록에 "*" 포함 여부 검증

package manifest

import "testing"

func TestCorsAllowOriginsHasWildcard(t *testing.T) {
	cases := []struct {
		name    string
		origins []string
		want    bool
	}{
		{name: "nil", origins: nil, want: false},
		{name: "empty", origins: []string{}, want: false},
		{name: "wildcard", origins: []string{"*"}, want: true},
		{name: "no_wildcard", origins: []string{"https://example.com"}, want: false},
		{name: "wildcard_among_others", origins: []string{"https://a.com", "*"}, want: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := corsAllowOriginsHasWildcard(c.origins)
			if got != c.want {
				t.Errorf("got %v, want %v", got, c.want)
			}
		})
	}
}
