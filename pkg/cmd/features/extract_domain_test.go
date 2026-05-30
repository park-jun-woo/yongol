//ff:func feature=features type=test control=iteration dimension=1
//ff:what TestExtractDomain — path→domain: 정상/복수형 절단/하이픈 절단/빈세그/필드부족 분기 검증

package features

import "testing"

func TestExtractDomain(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"POST /workflows/{id}", "workflow"}, // trailing-s trimmed
		{"GET /tasks", "task"},
		{"POST /work-flows/{id}", "work"}, // hyphen cut (before plural trim)
		{"GET /user", "user"},             // no trailing s
		{"badpath", "unknown"},            // <2 fields
		{"GET /", "unknown"},              // empty seg after trim prefix
		{"GET /s", "unknown"},             // becomes empty after trailing-s trim
	}
	for _, c := range cases {
		if got := extractDomain(c.in); got != c.want {
			t.Errorf("extractDomain(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
