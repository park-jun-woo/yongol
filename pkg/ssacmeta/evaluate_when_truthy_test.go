//ff:func feature=ssacmeta type=test control=iteration dimension=1
//ff:what TestevaluateWhenTruthy — evaluateWhenTruthy() bare `manifest.<path>` truthy 평가

package ssacmeta

import "testing"

func TestEvaluateWhenTruthy(t *testing.T) {
	manifest := map[string]any{
		"backend": map[string]any{
			"auth": map[string]any{
				"enabled":  true,
				"disabled": false,
				"name":     "jwt",
				"empty":    "",
			},
		},
	}
	cases := []struct {
		name string
		expr string
		want bool
	}{
		{"truthy-bool", "manifest.backend.auth.enabled", true},
		{"falsy-bool", "manifest.backend.auth.disabled", false},
		{"truthy-string", "manifest.backend.auth.name", true},
		{"empty-string", "manifest.backend.auth.empty", false},
		{"missing-path", "manifest.backend.auth.nope", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := evaluateWhenTruthy(c.expr, manifest); got != c.want {
				t.Errorf("evaluateWhenTruthy(%q) = %v, want %v", c.expr, got, c.want)
			}
		})
	}
}
