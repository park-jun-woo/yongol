//ff:func feature=cli-init type=test control=sequence
//ff:what TestExtractDomain — path→domain: 정상/복수형/하이픈/빈세그/필드부족 분기 검증

package cliinit

import "testing"

func TestExtractDomain(t *testing.T) {
	cases := []struct{ in, want string }{
		{"POST /workflows/{id}", "workflow"},
		{"GET /tasks", "task"},
		{"GET /audit-logs", "audit"},
		{"GET /user", "user"},
		{"badpath", "unknown"},
		{"GET /", "unknown"},
		{"GET /s", "unknown"},
	}
	for _, c := range cases {
		if got := extractDomain(c.in); got != c.want {
			t.Errorf("extractDomain(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
