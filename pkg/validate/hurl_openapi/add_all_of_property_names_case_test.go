//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-openapi
//ff:what runAddAllOfPropertyNames — TestAddAllOfPropertyNames table-driven 개별 케이스 검증

package hurl_openapi

import "testing"

func runAddAllOfPropertyNames(t *testing.T, c TestAddAllOfPropertyNamesCase) {
	t.Helper()
	out := make(map[string]struct{})
	for k, v := range c.existing {
		out[k] = v
	}
	addAllOfPropertyNames(out, c.allOf)
	if len(out) != len(c.want) {
		t.Fatalf("got %d keys, want %d; out=%v", len(out), len(c.want), out)
	}
	for k := range c.want {
		if _, ok := out[k]; !ok {
			t.Errorf("missing key %q in output", k)
		}
	}
}
