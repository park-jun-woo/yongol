//ff:func feature=gen-gogin type=test control=sequence topic=security-headers
//ff:what TestBuildCSPValue_Empty — nil / 빈 map 이면 빈 문자열 반환

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
