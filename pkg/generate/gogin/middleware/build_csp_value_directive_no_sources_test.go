//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what TestBuildCSPValue_DirectiveNoSources — 값이 빈 directive 도 키만 방출

package middleware

import "testing"

func TestBuildCSPValue_DirectiveNoSources(t *testing.T) {
	in := map[string][]string{
		"upgrade-insecure-requests": {},
		"default-src":               {"'self'"},
	}
	want := "default-src 'self'; upgrade-insecure-requests"
	if got := buildCSPValue(in); got != want {
		t.Fatalf("empty-source directive mismatch:\n want: %q\n got:  %q", want, got)
	}
}
