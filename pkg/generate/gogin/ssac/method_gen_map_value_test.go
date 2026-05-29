//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what methodGen.mapValue 단위 테스트 (literal/non-request/request. 분기)

package ssac

import "testing"

func TestMethodGenMapValue(t *testing.T) {
	g := &methodGen{
		PathParams: map[string]bool{"id": true},
	}
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"quoted literal passthrough", `"hello"`, `"hello"`},
		{"no dot passthrough", "rawVar", "rawVar"},
		{"non-request source passthrough", "currentUser.ID", "currentUser.ID"},
		{"request path param", "request.id", "request.Id"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := g.mapValue(tc.in); got != tc.want {
				t.Errorf("mapValue(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
