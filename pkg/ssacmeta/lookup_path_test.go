//ff:func feature=ssacmeta type=test control=iteration dimension=1
//ff:what TestlookupPath — lookupPath() dot-path 탐색: 성공/누락/비-map 중단 케이스

package ssacmeta

import "testing"

func TestLookupPath(t *testing.T) {
	m := map[string]any{
		"cache": map[string]any{
			"backend": "postgres",
		},
		"backend": map[string]any{
			"auth": map[string]any{
				"enabled": true,
			},
		},
		"scalar": "leaf",
	}
	cases := []struct {
		name   string
		path   string
		want   any
		wantOk bool
	}{
		{"single-level", "scalar", "leaf", true},
		{"two-levels", "cache.backend", "postgres", true},
		{"three-levels", "backend.auth.enabled", true, true},
		{"intermediate-map", "cache", m["cache"], true},
		{"missing-top", "nope", nil, false},
		{"missing-nested", "cache.missing", nil, false},
		{"descend-into-non-map", "scalar.deeper", nil, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, ok := lookupPath(m, c.path)
			if ok != c.wantOk {
				t.Fatalf("lookupPath(%q) ok = %v, want %v", c.path, ok, c.wantOk)
			}
			if !ok {
				return
			}
			// Compare scalars directly; for the map case just confirm non-nil.
			switch want := c.want.(type) {
			case string:
				if got != want {
					t.Errorf("lookupPath(%q) = %v, want %v", c.path, got, want)
				}
			case bool:
				if got != want {
					t.Errorf("lookupPath(%q) = %v, want %v", c.path, got, want)
				}
			default:
				if got == nil {
					t.Errorf("lookupPath(%q) = nil, want non-nil map", c.path)
				}
			}
		})
	}
}
