//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what buildCSPValue 단위 테스트 — directives map → 문자열 조합

package middleware

import "testing"

func TestBuildCSPValue_Empty(t *testing.T) {
	if v := buildCSPValue(nil); v != "" {
		t.Fatalf("nil map should yield empty string, got %q", v)
	}
	if v := buildCSPValue(map[string][]string{}); v != "" {
		t.Fatalf("empty map should yield empty string, got %q", v)
	}
}

func TestBuildCSPValue_Deterministic(t *testing.T) {
	in := map[string][]string{
		"script-src":      {"'self'", "cdn.example.com"},
		"default-src":     {"'self'"},
		"frame-ancestors": {"'none'"},
	}
	// Keys must be emitted in sorted order. "default-src" < "frame-ancestors" < "script-src"
	want := "default-src 'self'; frame-ancestors 'none'; script-src 'self' cdn.example.com"
	got := buildCSPValue(in)
	if got != want {
		t.Fatalf("buildCSPValue mismatch:\n want: %q\n got:  %q", want, got)
	}
}

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
