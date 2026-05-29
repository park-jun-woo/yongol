//ff:func feature=ssacmeta type=test control=iteration dimension=1
//ff:what TestevaluateWhenEquality — evaluateWhenEquality() `manifest.<path> == "<v>"` 비교

package ssacmeta

import (
	"strings"
	"testing"
)

func TestEvaluateWhenEquality(t *testing.T) {
	manifest := map[string]any{
		"cache": map[string]any{"backend": "postgres"},
	}
	cases := []struct {
		name string
		expr string
		want bool
	}{
		{"match", `manifest.cache.backend == "postgres"`, true},
		{"mismatch", `manifest.cache.backend == "memory"`, false},
		{"missing-path", `manifest.cache.missing == "postgres"`, false},
		{"non-string-rhs-empty", `manifest.cache.backend == ""`, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			i := strings.Index(c.expr, "==")
			if i <= 0 {
				t.Fatalf("test expr lacks ==: %q", c.expr)
			}
			if got := evaluateWhenEquality(c.expr, i, manifest); got != c.want {
				t.Errorf("evaluateWhenEquality(%q) = %v, want %v", c.expr, got, c.want)
			}
		})
	}
}
