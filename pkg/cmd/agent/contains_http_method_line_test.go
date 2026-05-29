//ff:func feature=agent type=test control=iteration dimension=2
//ff:what TestContainsHTTPMethodLine — HTTP 메서드로 시작하는 라인 포함 여부 판별 검증

package agent

import "testing"

func TestContainsHTTPMethodLine(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"get line", "foo\nGET /users\nbar", true},
		{"indented post", "   POST /login", true},
		{"delete", "DELETE /x", true},
		{"no method", "foo\nbar\nbaz", false},
		{"empty", "", false},
		{"method substring not prefix", "MyGET something", false},
	}
	for _, c := range cases {
		if got := containsHTTPMethodLine(c.in); got != c.want {
			t.Errorf("%s: containsHTTPMethodLine(%q) = %v, want %v", c.name, c.in, got, c.want)
		}
	}
}
