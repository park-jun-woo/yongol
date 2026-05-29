//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what TestBuildCSPValue_Deterministic — 키 정렬 순서 보장

package middleware

import "testing"

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
