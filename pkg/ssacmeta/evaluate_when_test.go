//ff:func feature=ssacmeta type=test control=iteration dimension=1
//ff:what TestEvaluateWhen — table-driven EvaluateWhen cases

package ssacmeta

import "testing"

func TestEvaluateWhen(t *testing.T) {
	cases := []struct {
		expr string
		ctx  map[string]any
		want bool
	}{
		{"", nil, true},
		{"always", nil, true},
		{`manifest.cache.backend == "postgres"`, map[string]any{"cache": map[string]any{"backend": "postgres"}}, true},
		{`manifest.cache.backend == "postgres"`, map[string]any{"cache": map[string]any{"backend": "memory"}}, false},
		{`manifest.cache.backend == "postgres"`, map[string]any{}, false},
		{`manifest.backend.auth.refresh.enabled`, map[string]any{"backend": map[string]any{"auth": map[string]any{"refresh": map[string]any{"enabled": true}}}}, true},
		{`manifest.backend.auth.refresh.enabled`, map[string]any{"backend": map[string]any{"auth": map[string]any{"refresh": map[string]any{"enabled": false}}}}, false},
	}
	for _, c := range cases {
		got := EvaluateWhen(c.expr, c.ctx)
		if got != c.want {
			t.Errorf("EvaluateWhen(%q) = %v, want %v", c.expr, got, c.want)
		}
	}
}
